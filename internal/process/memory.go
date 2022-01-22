package process

import (
	"encoding/binary"
)

func (c Context) GetUnitHashtableOffset() uintptr {
	memory := c.GetProcessMemory()
	addr := c.FindPatternEx(memory, "\x48\x8d\x00\x00\x00\x00\x00\x8b\xd1", "xx?????xx")

	relativeAddress := addr + 3

	bytes := c.ReadBytesFromMemory(relativeAddress, 4)
	offsetInt := uintptr(binary.LittleEndian.Uint32(bytes))
	delta := offsetInt - c.process.ModuleBaseAddressPtr

	return c.process.ModuleBaseAddressPtr + addr + 7 + delta
}

func (c Context) findPattern(memory []byte, pattern, mask string) int {
	patternLength := len(pattern)
	for i := 0; i < int(c.process.ModuleBaseSize)-patternLength; i++ {
		found := true
		for j := 0; j < patternLength; j++ {
			if string(mask[j]) != "?" && string(pattern[j]) != string(memory[i+j]) {
				found = false
				break
			}
		}

		if found {
			return i
		}
	}

	return 0
}

func (c Context) FindPatternEx(memory []byte, pattern, mask string) uintptr {
	offset := c.findPattern(memory, pattern, mask)

	if offset != 0 {
		return c.process.ModuleBaseAddressPtr + uintptr(offset)
	}

	return 0
}
