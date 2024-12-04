package main

import (
	"fmt"
	"os"

	myls "my-ls-1/Functions"
)

func main() {
	Flags := make(map[string]bool)
	paths := myls.ParseArgs(os.Args[1:], &Flags)
	if (Flags)["Help"] {
		fmt.Println("Usage: myls [OPTION]... [FILE]...\nList information about the FILEs (the current directory by default).\nSort entries alphabetically if none of -cftuvSUX nor --sort is specified.\n\nMandatory arguments to long options are mandatory for short options too.\n  -R, --recursive     list subdirectories recursively\n  -r, --reverse      reverse order while sorting\n  -a, --all          do not ignore entries starting with .\n  -l                 use a long listing format\n  -t                 sort by time, newest first; see --time")
		return
	} else if len(paths) == 0 {
		paths = append(paths, ".")
	}
	dirSlice, fileSlice := myls.SplitPath(paths)
	myls.FileSlice(fileSlice, dirSlice, &Flags)
	myls.DirSlice(fileSlice, dirSlice, &Flags)
}
