// +build linux,cgo,!static_build,journald,!journald_compat

package journald // import "github.com/ellcrys/docker/daemon/logger/journald"

// #cgo pkg-config: libsystemd
import "C"
