//go:build tools

package tools

import (
	_ "github.com/boumenot/gocover-cobertura"
	_ "github.com/jstemmer/go-junit-report/v2"
	_ "golang.org/x/perf/cmd/benchstat"
	_ "golang.org/x/vuln/cmd/govulncheck"
)
