// +build !windows

package system // import "github.com/ellcrys/docker/pkg/system"

// LCOWSupported returns true if Linux containers on Windows are supported.
func LCOWSupported() bool {
	return false
}
