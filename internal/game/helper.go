package game

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/hectorgimenez/koolo/internal/process"
	"unsafe"
)

type DataProvider struct {
	Context process.Context
}

type UnitHashTable struct {
	Units [128]uintptr
}

func (dp DataProvider) GetUnitHashTable(offset int) {
	unitTableOffset := dp.Context.GetUnitHashtableOffset()

	data := dp.Context.ReadBytesFromMemory(unitTableOffset, uint(unsafe.Sizeof(UnitHashTable{})))
	buf := &bytes.Buffer{}
	binary.Write(buf, binary.LittleEndian, data)

	buf.Write(data)
	ht := UnitHashTable{}
	binary.Read(buf, binary.LittleEndian, &ht)
	fmt.Println(unitTableOffset)
	size := unsafe.Sizeof(UnitHashTable{})
	fmt.Println(size)
}
