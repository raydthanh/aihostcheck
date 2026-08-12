//go:build linux || darwin

package collector

import "testing"

func TestParseMacGPUs(t *testing.T) {
	input := `{"SPDisplaysDataType":[{"sppci_model":"Apple M3 Pro"},{"_name":"External GPU"},{"sppci_model":"Apple M3 Pro"}]}`
	got, err := parseMacGPUs(input)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0] != "Apple M3 Pro" || got[1] != "External GPU" {
		t.Fatalf("models = %#v", got)
	}
}
