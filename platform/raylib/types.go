package rl

/*
#cgo CFLAGS: -I"/System/Volumes/Data/opt/homebrew/Cellar/raylib/5.0/include/"
#cgo LDFLAGS: -L"/System/Volumes/Data/opt/homebrew/Cellar/raylib/5.0/lib/" -lraylib

#include "raylib.h"
#include "raymath.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

type AudioStream struct {
	Buffer     *C.rAudioBuffer
	SampleRate uint32
	SampleSize uint32
	Channels   uint32
	_          [4]byte
}

type Camera struct {
	Offset   Vector
	Target   Vector
	Rotation float32
	Zoom     float32
}

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

type Music struct {
	Stream      AudioStream
	SampleCount uint32
	Looping     bool
	CtxType     int32
	CtxData     unsafe.Pointer
}

type Rectangle struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

type Sound struct {
	Stream      AudioStream
	SampleCount uint32
	_           [4]byte
}

type Texture struct {
	ID      uint32
	Width   int32
	Height  int32
	Mipmaps int32
	Format  int32
}

type Vector struct {
	X float32
	Y float32
}
