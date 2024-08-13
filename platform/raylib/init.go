package rl

import (
	"runtime"
)

func init() {
	runtime.LockOSThread()
}
