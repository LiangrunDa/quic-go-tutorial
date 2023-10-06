package utils

import "strconv"

func Humanize(dataSize string) int64 {
	unit := dataSize[len(dataSize)-2:]
	value, _ := strconv.ParseFloat(dataSize[:len(dataSize)-2], 64)
	var bytes int64
	switch unit {
	case "MB":
		bytes = int64(value * 1024 * 1024)
	case "GB":
		bytes = int64(value * 1024 * 1024 * 1024)
	case "KB":
		bytes = int64(value * 1024)
	default:
		bytes = int64(value)
	}
	return bytes
}
