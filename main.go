package main

import (
	"fmt"
	"os"

	myls "my-ls-1/Functions"
)

func main() {
	flags, paths := myls.ParseArgs(os.Args[1:])
	if flags["Help"] {
		fmt.Println("Usage: myls [OPTION]... [FILE]...\nList information about the FILEs (the current directory by default).\nSort entries alphabetically if none of -cftuvSUX nor --sort is specified.\n\nMandatory arguments to long options are mandatory for short options too.\n  -R, --recursive     list subdirectories recursively\n  -r, --reverse      reverse order while sorting\n  -a, --all          do not ignore entries starting with .\n  -l                 use a long listing format\n  -t                 sort by time, newest first; see --time")
		return
	}
	if len(paths) == 0 {
		paths = append(paths, ".")
	}
	dirSlice, fileSlice := SplitPath(paths)
	for _, path := range fileSlice {
		myls.MyLs(path, flags)
	}
	if len(fileSlice) != 0 {
		fmt.Println()
		if len(dirSlice) != 0 {
			fmt.Println()
		}
	}
	for i, path := range dirSlice {
		fmt.Printf("%v:\n", path)
		myls.MyLs(path, flags)
		fmt.Println()
		if i != len(dirSlice)-1 {
			fmt.Println()
		}
	}
}

func SplitPath(paths []string) ([]string, []string) {
	dirSlice, fileSlice := []string{}, []string{}
	for _, path := range paths {
		if myls.Dir(path) {
			dirSlice = append(dirSlice, path)
		} else {
			fileSlice = append(fileSlice, path)
		}
	}
	return dirSlice, fileSlice
}
