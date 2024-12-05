package functions

import (
	"strings"
)

func SortPath(slice []string) []string {
	for i := 0; i < len(slice)-1; i++ {
		for j := i+1; j < len(slice); j++ {
			if strings.ToLower(slice[i]) > strings.ToLower(slice[j]) {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}

func SortLs(slice []LongFormatInfo) {
	for i := 0; i < len(slice); i++ {
		for j := i+1; j < len(slice); j++ {
			if strings.ToLower(getKey(slice[i].FileName)) > strings.ToLower(getKey(slice[j].FileName)) {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}

func getKey(filename string) string {
	if strings.HasPrefix(filename, ".") {
		for i := 0; i < len(filename); i++ {
			if filename[i] != '.' {
				return filename[i:]
			}
		}
	}
	return filename
}

func SortByTime(slice []LongFormatInfo) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if !slice[j].Time.After(slice[j+1].Time) {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func ReverseSorting(slice []LongFormatInfo) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
