package report

import (
	"fmt"
	"io"
	"sort"

	"github.com/raydthanh/aihostcheck/internal/model"
)

func WriteText(w io.Writer, r model.Report) {
	fmt.Fprintf(w, "AIHostCheck %s — %s\n", r.ToolVersion, r.Platform)
	keys := make([]string, 0, len(r.Capabilities))
	for k := range r.Capabilities {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		c := r.Capabilities[k]
		value := c.Value
		if value == "" {
			value = "-"
		}
		fmt.Fprintf(w, "%-18s %-18s %s\n", k, c.Status, value)
	}
}
