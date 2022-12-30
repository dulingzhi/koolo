package main

import (
	"encoding/binary"
	"fmt"
	zapLogger "github.com/hectorgimenez/koolo/cmd/koolo/log"
	"github.com/hectorgimenez/koolo/internal/config"
	"github.com/hectorgimenez/koolo/internal/helper/tcp"
	"github.com/hectorgimenez/koolo/internal/memory"
	"time"
	"unsafe"
)

func main() {
	logger, err := zapLogger.NewLogger(config.Config.Debug, config.Config.LogFilePath)
	process, err := memory.NewProcess(logger)
	if err != nil {
		panic(err)
	}

	gd := memory.NewGameReader(process)

	var buffSize tcp.DWORD
	_ = tcp.GetExtendedTCPTable(uintptr(0), &buffSize, true, 2, tcp.TCP_TABLE_OWNER_PID_ALL, 0)
	var buffTable = make([]byte, int(buffSize))
	err = tcp.GetExtendedTCPTable(uintptr(unsafe.Pointer(&buffTable[0])), &buffSize, true, 2, tcp.TCP_TABLE_OWNER_PID_ALL, 0)

	count := *(*uint32)(unsafe.Pointer(&buffTable[0]))
	const structLen = 24
	for n, pos := uint32(0), 4; n < count && pos+structLen <= len(buffTable); n, pos = n+1, pos+structLen {
		state := *(*uint32)(unsafe.Pointer(&buffTable[pos]))
		if state < 1 || state > 12 {
			panic(state)
		}
		laddr := binary.BigEndian.Uint32(buffTable[pos+4 : pos+8])
		lport := binary.BigEndian.Uint16(buffTable[pos+8 : pos+10])
		raddr := binary.BigEndian.Uint32(buffTable[pos+12 : pos+16])
		rport := binary.BigEndian.Uint16(buffTable[pos+16 : pos+18])
		pid := *(*uint32)(unsafe.Pointer(&buffTable[pos+20]))

		if uint(pid) == process.GetPID() && state == tcp.MIB_TCP_STATE_ESTAB && rport == 443 {
			buffTable[pos] = tcp.MIB_TCP_STATE_DELETE_TCB
			newState := *(*uint32)(unsafe.Pointer(&buffTable[pos]))
			fmt.Println(newState)
			fmt.Printf("%5d = %d %08x:%d %08x:%d pid:%d\n", n, state, laddr, lport, raddr, rport, pid)
			newSlice := buffTable[pos : pos+24]
			state := *(*uint32)(unsafe.Pointer(&newSlice[0]))
			fmt.Printf("%5d = %d %08x:%d %08x:%d pid:%d\n", n, state, laddr, lport, raddr, rport, pid)
			tcp.SetTCPEntry(uintptr(unsafe.Pointer(&newSlice[0])))
		}
	}
	fmt.Println(count)
	fmt.Println(err)

	start := time.Now()
	for true {
		d := gd.GetData(true)
		fmt.Println(d.MercHPPercent())
		time.Sleep(time.Millisecond * 500)
	}

	fmt.Printf("Read time: %dms\n", time.Since(start).Milliseconds())
}
