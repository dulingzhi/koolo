package main

import (
	"github.com/hectorgimenez/koolo/internal/process"
)

func main() {
	c, _ := process.NewContext()

	c.GetPlayer()
}
