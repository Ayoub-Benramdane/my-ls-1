package functions

import (
	"fmt"
	"io/fs"
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

func IsDir(list []fs.FileInfo) bool {
	for _, item := range list {
		if item.IsDir() {
			return true
		}
	}
	return false
}

func DirSlice(fileSlice, dirSlice []string, flags map[string]bool) {
	if len(fileSlice) != 0 && len(dirSlice) != 0 {
		fmt.Println()
	}
	for i, path := range dirSlice {
		if len(dirSlice) != 1 {
			fmt.Printf("%v:\n", path)
		}
		Len := MyLs(path, flags)
		if Len != 0 {
			fmt.Println()
		}
		if i != len(dirSlice)-1 {
			fmt.Println()
		}
	}
}

func FileSlice(fileSlice, dirSlice []string, flags map[string]bool) {
	for _, path := range fileSlice {
		MyLs(path, flags)
	}
	if len(fileSlice) != 0 {
		fmt.Println()
	}
}
