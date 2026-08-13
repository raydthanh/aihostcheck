package main

import (
	"testing"
	"time"
)

func TestDefaultProbeTimeoutCoversColdWindowsPowerShellStart(t *testing.T) {
	if defaultProbeTimeout < 15*time.Second {
		t.Fatalf("defaultProbeTimeout = %s, want at least 15s for the supported Windows CIM GPU probe", defaultProbeTimeout)
	}
}
