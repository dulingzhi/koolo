package process

import (
	"fmt"
	"github.com/winlabs/gowin32"
	"golang.org/x/sys/windows"
	"time"
	"unsafe"
)

const moduleName = "D2R.exe"

type Context struct {
	process Process
	handler windows.Handle
}

func NewContext() (Context, error) {
	module, err := getGameModule()
	if err != nil {
		return Context{}, err
	}

	h, err := windows.OpenProcess(0x0010, false, uint32(module.ProcessID))
	if err != nil {
		return Context{}, err
	}

	return Context{
		handler: h,
		process: Process{
			PID:                  module.ProcessID,
			ModuleBaseAddress:    unsafe.Pointer(module.ModuleBaseAddress),
			ModuleBaseAddressPtr: uintptr(unsafe.Pointer(module.ModuleBaseAddress)),
			ModuleBaseSize:       module.ModuleBaseSize,
		}}, nil
}

func getGameModule() (gowin32.ModuleInfo, error) {
	processes := make([]uint32, 2048)
	length := uint32(0)
	err := windows.EnumProcesses(processes, &length)
	if err != nil {
		return gowin32.ModuleInfo{}, err
	}

	for _, process := range processes {
		module, _ := getMainModule(process)

		if module.ExePath == "C:\\Program Files (x86)\\Diablo II Resurrected\\D2R.exe" {
			return module, nil
		}
	}

	return gowin32.ModuleInfo{}, err
}

func getMainModule(pid uint32) (gowin32.ModuleInfo, error) {
	mi, err := gowin32.GetProcessModules(pid)
	if err != nil {
		return gowin32.ModuleInfo{}, err
	}
	for _, m := range mi {
		if m.ModuleName == moduleName {
			return m, nil
		}
	}

	return gowin32.ModuleInfo{}, err
}

func (c Context) GetProcessMemory() []byte {
	start := time.Now()
	var data = make([]byte, c.process.ModuleBaseSize)
	err := windows.ReadProcessMemory(c.handler, uintptr(c.process.ModuleBaseAddress), &data[0], uintptr(c.process.ModuleBaseSize), nil)
	if err != nil {
		fmt.Printf("err: %s", err.Error())
	}
	fmt.Printf("Main module memory loaded in: %dms\n", time.Since(start).Milliseconds())

	return data
}

func (c Context) ReadBytesFromMemory(baseAddress uintptr, size uint) []byte {
	var data = make([]byte, size)
	err := windows.ReadProcessMemory(c.handler, baseAddress, &data[0], uintptr(size), nil)
	if err != nil {
		fmt.Printf("err: %s", err.Error())
	}

	return data
}
