package rl

/*
#cgo darwin LDFLAGS: -framework Cocoa -framework CoreFoundation -framework CoreVideo -framework IOKit
#cgo linux LDFLAGS: -ldl -lm -pthread -lrt -lwayland-client -lwayland-cursor -lwayland-egl -lxkbcommon
#cgo windows LDFLAGS: -lgdi32 -lole32 -lwinmm
#cgo LDFLAGS: -Llib -lraylib
*/
import "C"
import "runtime"

func init() {
	runtime.LockOSThread()
}
