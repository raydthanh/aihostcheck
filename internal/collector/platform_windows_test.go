//go:build windows

package collector

import "testing"

func TestParseRegistryValue(t *testing.T) {
	output := "HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion\r\n    ProductName    REG_SZ    Microsoft Windows 11 Pro\r\n"
	if got := parseRegistryValue(output, "ProductName"); got != "Microsoft Windows 11 Pro" {
		t.Fatalf("value = %q", got)
	}
}
