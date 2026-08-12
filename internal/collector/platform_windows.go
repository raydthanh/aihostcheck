//go:build windows

package collector

import (
	"fmt"
	"runtime"
	"syscall"
	"time"
	"unsafe"

	"github.com/raydthanh/aihostcheck/internal/model"
	"github.com/raydthanh/aihostcheck/internal/probe"
)

func platformCapabilities(timeout time.Duration) map[string]model.Capability {
	c := map[string]model.Capability{}
	c["cpu"] = model.Capability{Status: model.Detected, Value: fmt.Sprintf("%d logical CPUs", runtime.NumCPU()), Evidence: []model.Evidence{{Source: "go_runtime", Detail: "runtime.NumCPU"}}}
	c["shell"] = model.Capability{Status: model.Detected, Value: "cmd.exe / PowerShell", Evidence: []model.Evidence{{Source: "platform_default", Detail: "Windows native shells"}}}
	c["os_version"] = probe.Command(timeout, "reg.exe", "query", `HKLM\SOFTWARE\Microsoft\Windows NT\CurrentVersion`, "/v", "ProductName")
	c["package_manager"] = firstWindows(timeout, []string{"winget", "choco", "scoop"})
	c["gpu"] = probe.Command(timeout, "wmic.exe", "path", "win32_VideoController", "get", "name")
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
		c["memory"] = model.Capability{Status: model.Error}
	}
	var free, total, totalFree uint64
	root, _ := syscall.UTF16PtrFromString(`C:\`)
	if r, _, _ := kernel.NewProc("GetDiskFreeSpaceExW").Call(uintptr(unsafe.Pointer(root)), uintptr(unsafe.Pointer(&free)), uintptr(unsafe.Pointer(&total)), uintptr(unsafe.Pointer(&totalFree))); r != 0 {
		c["root_storage"] = model.Capability{Status: model.Detected, Value: fmt.Sprintf("%d bytes total, %d bytes available", total, free), Evidence: []model.Evidence{{Source: "win32_api", Detail: "GetDiskFreeSpaceExW C drive"}}}
	} else {
		c["root_storage"] = model.Capability{Status: model.Error}
	}
	return c
}
func firstWindows(t time.Duration, names []string) model.Capability {
	for _, n := range names {
		r := probe.Command(t, n, "--version")
		if r.Status == model.Detected {
			r.Details = map[string]string{"manager": n}
			return r
		}
	}
	return model.Capability{Status: model.NotDetected}
}
