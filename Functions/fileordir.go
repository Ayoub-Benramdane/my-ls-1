package functions

import (
	"fmt"
	"os"
)

func SplitPath(paths []string) ([]string, []string) {
	dirSlice, fileSlice := []string{}, []string{}
	for _, path := range paths {
		if _, err := os.ReadDir(path); err == nil {
			dirSlice = append(dirSlice, path)
		} else {
			fileSlice = append(fileSlice, path)
		}
	}
	return dirSlice, fileSlice
}

func FileSlice(fileSlice, dirSlice []string, flags map[string]bool) {
	for _, path := range fileSlice {
		MyLs(path, flags, 0)
	}
}

func DirSlice(fileSlice, dirSlice []string, flags map[string]bool, totalPath int) {
	if len(fileSlice) != 0 && len(dirSlice) != 0 {
		fmt.Println()
	}
	for i, path := range dirSlice {
		if len(dirSlice) != 1 {
			fmt.Printf("%v:\n", path)
		}
		MyLs(path, flags, totalPath)
		if i != len(dirSlice)-1 {
			fmt.Println()
		}
	}
}
