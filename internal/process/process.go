package process

import "unsafe"

type Process struct {
	PID                  uint
	ModuleBaseAddress    unsafe.Pointer
	ModuleBaseAddressPtr uintptr
	ModuleBaseSize       uint
}
