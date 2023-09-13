package language

import "github.com/pwsdc/web-mud/shared/enum"

var dirMap map[string]string

func init() {
	dirMap = make(map[string]string)
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
}

func ParseDirection(in string) (string, bool) {
	v, ok := dirMap[in]
	return v, ok
}
