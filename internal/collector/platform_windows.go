//go:build windows

package collector

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/raydthanh/aihostcheck/internal/model"
	"github.com/raydthanh/aihostcheck/internal/probe"
)

func platformCapabilities(timeout time.Duration) map[string]model.Capability {
	c := map[string]model.Capability{}
	c["cpu"] = model.Capability{Status: model.Detected, Value: fmt.Sprintf("%d logical CPUs", runtime.NumCPU()), Evidence: []model.Evidence{{Source: "go_runtime", Detail: "runtime.NumCPU"}}}
	c["shell"] = windowsShellCapability()
	c["os_version"] = windowsVersionCapability(timeout)
	c["package_manager"] = allAvailable(timeout, []commandSpec{{name: "winget.exe", args: []string{"--version"}}, {name: "choco.exe", args: []string{"--version"}}, {name: "scoop.cmd", args: []string{"--version"}}})
	c["gpu"] = windowsGPUCapability(timeout)
	var mem struct {
		Length                                                                                               uint32
		MemoryLoad                                                                                           uint32
		TotalPhys, AvailPhys, TotalPageFile, AvailPageFile, TotalVirtual, AvailVirtual, AvailExtendedVirtual uint64
	}
	mem.Length = uint32(unsafe.Sizeof(mem))
	kernel := syscall.NewLazyDLL("kernel32.dll")
	if r, _, _ := kernel.NewProc("GlobalMemoryStatusEx").Call(uintptr(unsafe.Pointer(&mem))); r != 0 {
		c["memory"] = model.Capability{Status: model.Detected, Value: fmt.Sprintf("%d bytes", mem.TotalPhys), Evidence: []model.Evidence{{Source: "win32_api", Detail: "GlobalMemoryStatusEx"}}}
	} else {
		c["memory"] = model.Capability{Status: model.Error, Evidence: []model.Evidence{{Source: "win32_api", Detail: "GlobalMemoryStatusEx failed"}}}
	}

	root := windowsSystemRoot()
	if root == "" {
		c["root_storage"] = model.Capability{Status: model.Unknown, Evidence: []model.Evidence{{Source: "environment", Detail: "system volume could not be identified"}}}
		return c
	}
	var free, total, totalFree uint64
	rootPointer, err := syscall.UTF16PtrFromString(root)
	if err != nil {
		c["root_storage"] = model.Capability{Status: model.Error, Evidence: []model.Evidence{{Source: "win32_api", Detail: "system volume path could not be encoded"}}}
		return c
	}
	if r, _, _ := kernel.NewProc("GetDiskFreeSpaceExW").Call(uintptr(unsafe.Pointer(rootPointer)), uintptr(unsafe.Pointer(&free)), uintptr(unsafe.Pointer(&total)), uintptr(unsafe.Pointer(&totalFree))); r != 0 {
		c["root_storage"] = model.Capability{Status: model.Detected, Value: fmt.Sprintf("%d bytes total, %d bytes available", total, free), Details: map[string]string{"volume": root}, Evidence: []model.Evidence{{Source: "win32_api", Detail: "GetDiskFreeSpaceExW on the system volume"}}}
	} else {
		c["root_storage"] = model.Capability{Status: model.Error, Evidence: []model.Evidence{{Source: "win32_api", Detail: "GetDiskFreeSpaceExW failed on the system volume"}}}
	}
	return c
}

func windowsShellCapability() model.Capability {
	result := model.Capability{Status: model.Unknown, Details: map[string]string{"kind": "active_shell_not_inferred"}, Evidence: []model.Evidence{{Source: "environment", Detail: "the active parent shell cannot be established safely from platform defaults"}}}
	if commandProcessor := filepath.Base(os.Getenv("ComSpec")); commandProcessor != "." && commandProcessor != "" {
		result.Details["default_command_processor"] = commandProcessor
		result.Evidence = append(result.Evidence, model.Evidence{Source: "environment", Detail: "basename of ComSpec identifies the default command processor only"})
	}
	return result
}

func windowsVersionCapability(timeout time.Duration) model.Capability {
	product, productEvidence, status := windowsRegistryValue(timeout, "ProductName")
	if status != model.Detected {
		return model.Capability{Status: status, Evidence: productEvidence}
	}
	value := product
	evidence := productEvidence
	if displayVersion, displayEvidence, displayStatus := windowsRegistryValue(timeout, "DisplayVersion"); displayStatus == model.Detected {
		value += " " + displayVersion
		evidence = append(evidence, displayEvidence...)
	}
	if build, buildEvidence, buildStatus := windowsRegistryValue(timeout, "CurrentBuildNumber"); buildStatus == model.Detected {
		value += " (build " + build + ")"
		evidence = append(evidence, buildEvidence...)
	}
	return model.Capability{Status: model.Detected, Value: value, Evidence: evidence}
}

func windowsRegistryValue(timeout time.Duration, name string) (string, []model.Evidence, model.Status) {
	r := probe.Run(timeout, "reg.exe", "query", `HKLM\SOFTWARE\Microsoft\Windows NT\CurrentVersion`, "/v", name)
	if r.Status != model.Detected {
		return "", r.Evidence, r.Status
	}
	if value := parseRegistryValue(r.Output, name); value != "" {
		return value, r.Evidence, model.Detected
	}
	return "", []model.Evidence{{Source: "command", Detail: "reg.exe returned no parseable " + name + " value"}}, model.Unknown
}

func parseRegistryValue(output, name string) string {
	for _, line := range strings.Split(output, "\n") {
		fields := strings.Fields(line)
		if len(fields) >= 3 && strings.EqualFold(fields[0], name) && strings.HasPrefix(strings.ToUpper(fields[1]), "REG_") {
			return strings.Join(fields[2:], " ")
		}
	}
	return ""
}

func windowsGPUCapability(timeout time.Duration) model.Capability {
	const query = "Get-CimInstance -ClassName Win32_VideoController | Select-Object -ExpandProperty Name"
	for _, shell := range []string{"pwsh.exe", "powershell.exe"} {
		r := probe.Run(timeout, shell, "-NoLogo", "-NoProfile", "-NonInteractive", "-Command", query)
		if r.Status == model.NotDetected {
			continue
		}
		if r.Status != model.Detected {
			return model.Capability{Status: r.Status, Evidence: r.Evidence}
		}
		lines := nonEmptyLines(r.Output)
		if len(lines) == 0 {
			return model.Capability{Status: model.Unknown, Evidence: []model.Evidence{{Source: "command", Detail: "CIM query completed without a video-controller result"}}}
		}
		return model.Capability{Status: model.Detected, Value: strings.Join(lines, "; "), Evidence: r.Evidence, Details: map[string]string{"inventory": "Win32_VideoController via CIM"}}
	}
	return model.Capability{Status: model.Unknown, Evidence: []model.Evidence{{Source: "executable_lookup", Detail: "PowerShell is unavailable, so CIM GPU inventory was not run"}}}
}

func windowsSystemRoot() string {
	if volume := filepath.VolumeName(os.Getenv("SystemRoot")); volume != "" {
		return volume + `\`
	}
	if volume := strings.TrimSpace(os.Getenv("SystemDrive")); volume != "" {
		return strings.TrimRight(volume, `\/`) + `\`
	}
	return ""
}
