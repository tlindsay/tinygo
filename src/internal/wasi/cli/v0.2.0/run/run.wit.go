// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package run represents the exported interface "wasi:cli/run@0.2.0".
package run

import (
	"internal/cm"
)

//go:wasmexport wasi:cli/run@0.2.0#run
//export wasi:cli/run@0.2.0#run
func wasmexport_Run() (result0 uint32) {
	result := Exports.Run()
	result0 = cm.BoolToU32(result)
	return
}