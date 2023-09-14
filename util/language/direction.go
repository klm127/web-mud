package language

import "github.com/pwsdc/web-mud/shared/enum"

var dirMap map[string]string
var dirOpposite map[string]string
var dirFull map[string]string

func init() {
	dirMap = make(map[string]string)
	dirOpposite = make(map[string]string)
	dirFull = make(map[string]string)
	for _, v := range enum.Dirs {
		dirMap[v] = v
	}
	dirMap["north"] = enum.North
	dirMap["south"] = enum.South
	dirMap["east"] = enum.East
	dirMap["west"] = enum.West

	dirMap["northeast"] = enum.NorthEast
	dirMap["neast"] = enum.NorthEast
	dirMap["northe"] = enum.NorthEast

	dirMap["southeast"] = enum.SouthEast
	dirMap["southe"] = enum.SouthEast
	dirMap["seast"] = enum.SouthEast

	dirMap["northwest"] = enum.NorthWest
	dirMap["northw"] = enum.NorthWest
	dirMap["nwest"] = enum.NorthWest

	dirMap["southwest"] = enum.SouthWest
	dirMap["swest"] = enum.SouthWest
	dirMap["southw"] = enum.SouthWest

	dirMap["in"] = enum.In
	dirMap["enter"] = enum.In

	dirMap["exit"] = enum.Out
	dirMap["out"] = enum.Out

	dirMap["up"] = enum.Up
	dirMap["down"] = enum.Down

	dirOpposite[enum.North] = enum.South
	dirOpposite[enum.South] = enum.North
	dirOpposite[enum.East] = enum.West
	dirOpposite[enum.West] = enum.East
	dirOpposite[enum.NorthEast] = enum.SouthWest
	dirOpposite[enum.SouthWest] = enum.NorthEast
	dirOpposite[enum.NorthWest] = enum.SouthEast
	dirOpposite[enum.SouthEast] = enum.NorthWest
	dirOpposite[enum.Up] = enum.Down
	dirOpposite[enum.Down] = enum.Up
	dirOpposite[enum.In] = enum.Out
	dirOpposite[enum.Out] = enum.In

	dirFull[enum.North] = "north"
	dirFull[enum.South] = "south"
	dirFull[enum.East] = "east"
	dirFull[enum.West] = "west"
	dirFull[enum.NorthEast] = "northeast"
	dirFull[enum.SouthWest] = "southwest"
	dirFull[enum.NorthWest] = "northwest"
	dirFull[enum.SouthEast] = "southeast"
	dirFull[enum.Down] = "down"
	dirFull[enum.In] = "in"
	dirFull[enum.Out] = "out"
}

func ParseDirection(in string) (string, bool) {
	v, ok := dirMap[in]
	return v, ok
}

// in must be an enumerated direction, eg, 'n'. Returns, eg, 's'.
func ParseOppositeDirection(in string) string {
	return dirOpposite[in]
}

// in must be an enumerated direction, eg, 'n'. Returns, eg, 'north'
func ParseDirectionFull(in string) string {
	return dirFull[in]
}
