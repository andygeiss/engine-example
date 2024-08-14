package rl

/*
#cgo CFLAGS: -Iinclude
#include "raylib.h"
#include "raymath.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

func DrawTexture(texture Texture, posX, posY, tX, tY, w, h, rot float32) {
	cCol := ptr[Color, C.Color](White)
	cDst := ptr[Rectangle, C.Rectangle](Rectangle{X: posX, Y: posY, Width: w, Height: h})
	cOri := ptr[Vector, C.Vector2](Vector{X: 0, Y: 0})
	cRec := ptr[Rectangle, C.Rectangle](Rectangle{X: tX, Y: tY, Width: w, Height: h})
	cRot := (C.float)(rot)
	cTex := ptr[Texture, C.Texture2D](texture)
	C.DrawTexturePro(*cTex, *cRec, *cDst, *cOri, cRot, *cCol)
}

func LoadTexture(path string) Texture {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	ctex := C.LoadTexture(cpath)
	tex := *ptr[C.Texture, Texture](ctex)
	return tex
}

func UnloadTexture(texture Texture) {
	C.UnloadTexture(*ptr[Texture, C.Texture](texture))
}
