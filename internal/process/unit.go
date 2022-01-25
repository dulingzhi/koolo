package process

import (
	"fmt"
	"github.com/hectorgimenez/koolo/internal/game"
)

const (
	// Unit Type
	unitTypePlayer  UnitType = 0
	unitTypeNPC     UnitType = 1
	unitTypeObject  UnitType = 2
	unitTypeMissile UnitType = 3
	unitTypeItem    UnitType = 4
	unitTypeWarp    UnitType = 5

	// Player Class
	playerClassAmazon      PlayerClass = 0
	playerClassSorceress   PlayerClass = 1
	playerClassNecromancer PlayerClass = 2
	playerClassPaladin     PlayerClass = 3
	playerClassBarbarian   PlayerClass = 4
	playerClassDruid       PlayerClass = 5
	playerClassAssassin    PlayerClass = 6
	playerClassInvalid     PlayerClass = 7

	// Unit Stats
	life      Stat = 6
	maxLife   Stat = 7
	mana      Stat = 8
	maxMana   Stat = 9
	gold      Stat = 14
	stashGold Stat = 15
)

type UnitType uint
type PlayerClass uint
type Stat uint8

type UnitHashTable struct {
	Units [128]uint64
}

type UnitAny struct {
	UnitType     UnitType    `bin:"len:4, offsetStart:0"`
	TxtFileNo    uint        `bin:"len:4, offsetStart:4"`
	UnitId       uint        `bin:"len:4, offsetStart:8"`
	UnitDataPtr  uint        `bin:"len:8, offsetStart:16"`
	ActPtr       uint        `bin:"len:8, offsetStart:32"`
	PathPtr      uint        `bin:"len:8, offsetStart:56"`
	StatsListPtr uint        `bin:"len:8, offsetStart:136"`
	InventoryPtr uint        `bin:"len:8, offsetStart:144"`
	X            uint16      `bin:"len:2, offsetStart:196"`
	Y            uint16      `bin:"len:2, offsetStart:198"`
	PlayerClass  PlayerClass `bin:"len:4, offsetStart:372"`
	IsCorpse     bool        `bin:"len:1, offsetStart:422"`
}

type StatList struct {
	OwnerType uint      `bin:"len:4, offsetStart:8"`
	OwnerId   uint      `bin:"len:4, offsetStart:12"`
	Flags     uint      `bin:"len:4, offsetStart:28"`
	BaseStats StatArray `bin:"offsetStart:48"`
	Stats     StatArray `bin:"offsetStart:128"`
}

func (ua StatList) Size() uint {
	return 2000
}

type StatArray struct {
	FirstStatPtr uint `bin:"len:8"`
	Size         int  `bin:"len:8"`
	Capacity     int  `bin:"len:8"`
}

type StatValue struct {
	Layer uint16 `bin:"len:2"`
	Stat  int16  `bin:"len:2"`
	Value int    `bin:"len:4"`
}

func (ua UnitAny) Size() uint {
	return 423
}

func (ua UnitAny) ToUnit(c Context) game.UnitAny {
	statList := StatList{}
	data := c.ReadBytesFromMemory(uintptr(ua.StatsListPtr), statList.Size())
	err := c.bytesToStruct(data, &statList)
	if err != nil {
		fmt.Println(err)
	}

	type statValues struct {
		List []StatValue `bin:"len:8"`
	}
	statValuesList := statValues{}
	data = c.ReadBytesFromMemoryTimes(uintptr(statList.Stats.FirstStatPtr), 8, statList.Stats.Size)
	err = c.bytesToStruct(data, &statValuesList)
	if err != nil {
		fmt.Println(err)
	}

	name := string(c.ReadBytesFromMemory(uintptr(ua.UnitDataPtr), 16))
	fmt.Println(name)

	return game.UnitAny{
		Name: name,
	}
}
