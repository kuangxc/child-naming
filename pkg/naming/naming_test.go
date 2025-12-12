package naming

import "testing"

func TestRun(t *testing.T) {
	// Run triggers output.SaveExcel which calls os.Exit(0) by design.
	// Skip to avoid terminating `go test`.
	t.Skip("Run exits the process via output.SaveExcel(os.Exit), not suitable for unit tests")
}
