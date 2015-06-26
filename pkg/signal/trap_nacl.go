//+build nacl

package signal

import (
	"runtime"

	"github.com/Sirupsen/logrus"
)

// Trap sets up a simplified signal "trap", appropriate for common
// behavior expected from a vanilla unix command-line tool in general
// (and the Docker engine in particular).
//
// * If SIGINT or SIGTERM are received, `cleanup` is called, then the process is terminated.
// * If SIGINT or SIGTERM are received 3 times before cleanup is complete, then cleanup is
//   skipped and the process is terminated immediately (allows force quit of stuck daemon)
// * A SIGQUIT always causes an exit without cleanup, with a goroutine dump preceding exit.
//
func Trap(cleanup func()) {
}

func DumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	// Note that if the daemon is started with a less-verbose log-level than "info" (the default), the goroutine
	// traces won't show up in the log.
	logrus.Infof("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
}
