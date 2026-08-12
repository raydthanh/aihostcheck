package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/raydthanh/aihostcheck/internal/collector"
	"github.com/raydthanh/aihostcheck/internal/report"
)

var version = "dev"

func main() {
	jsonOutput := flag.Bool("json", false, "emit the machine-readable JSON report")
	showVersion := flag.Bool("version", false, "print the AIHostCheck version and exit")
	timeout := flag.Duration("timeout", 3*time.Second, "maximum duration of each command probe")
	flag.Parse()
	if *showVersion {
		fmt.Fprintln(os.Stdout, version)
		return
	}
	if *timeout <= 0 || *timeout > 30*time.Second {
		fmt.Fprintln(os.Stderr, "timeout must be greater than zero and no more than 30s")
		os.Exit(2)
	}
	r := collector.Collect(*timeout, version)
	if *jsonOutput {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(r); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}
	report.WriteText(os.Stdout, r)
}
