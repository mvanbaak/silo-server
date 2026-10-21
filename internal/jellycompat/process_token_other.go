//go:build !linux && !darwin

package jellycompat

// processToken returns an identity token for a running process so callers can
// tell whether a recorded PID still belongs to the process they started. Only
// linux and darwin implement it; elsewhere the empty string makes the caller
// fall back to processIsRunning.
func processToken(int) string { return "" }
