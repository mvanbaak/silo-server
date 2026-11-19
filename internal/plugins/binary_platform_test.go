package plugins

import (
	"os"
	"runtime"
	"testing"
)

// TestBinaryPlatformReadsThisHostsBinaries asserts the reader reports the
// platform this test binary was built for. It fails on FreeBSD before the
// build-info lookup, because every FreeBSD ELF was reported as linux.
func TestBinaryPlatformReadsThisHostsBinaries(t *testing.T) {
	self, err := os.Executable()
	if err != nil {
		t.Fatalf("locate the test binary: %v", err)
	}

	goos, goarch, ok := binaryPlatform(self)
	if !ok {
		t.Fatal("binaryPlatform did not recognize this host's own binary")
	}
	if goos != runtime.GOOS || goarch != runtime.GOARCH {
		t.Fatalf("binaryPlatform(%s) = %s/%s, want %s/%s", self, goos, goarch, runtime.GOOS, runtime.GOARCH)
	}
	if err := checkBinaryPlatform(self); err != nil {
		t.Fatalf("checkBinaryPlatform(%s) = %v, want nil", self, err)
	}
}
