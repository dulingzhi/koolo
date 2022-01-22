package main

import (
	"github.com/hectorgimenez/koolo/internal/game"
	"github.com/hectorgimenez/koolo/internal/process"
)

func main() {
	c, _ := process.NewContext()
	dp := game.DataProvider{Context: c}

	dp.GetUnitHashTable(0)
}
