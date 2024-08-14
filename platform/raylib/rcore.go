package rl

/*
#cgo CFLAGS: -Iinclude
#include "raylib.h"
#include "raymath.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

func ptr[A, B any](input A) *B {
	return (*B)(unsafe.Pointer(&input))
}

func BeginDrawing() {
	C.BeginDrawing()
}

func BeginMode2D(camera Camera) {
	C.BeginMode2D(*ptr[Camera, C.Camera2D](camera))
}

func ClearBackground(color Color) {
	C.ClearBackground(*ptr[Color, C.Color](color))
}

func CloseWindow() {
	C.CloseWindow()
}

func EndDrawing() {
	C.EndDrawing()
}

func EndMode2D() {
	C.EndMode2D()
}

func GetFrameTime() float32 {
	return float32(C.GetFrameTime())
}

func GetFPS() int32 {
	return int32(C.GetFPS())
}

func InitWindow(width, height int32, title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.InitWindow(C.int(width), C.int(height), cTitle)
}

func IsKeyDown(key int32) bool {
	ckey := (C.int)(key)
	return bool(C.IsKeyDown(ckey))
}

func IsWindowReady() bool {
	return bool(C.IsWindowReady())
}

func SetTargetFPS(fps int) {
	C.SetTargetFPS(C.int(fps))
}

func WindowShouldClose() bool {
	return bool(C.WindowShouldClose())
}
