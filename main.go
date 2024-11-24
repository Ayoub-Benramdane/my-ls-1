package main

import (
	"fmt"
	myls "my-ls-1/Functions"
	"os"
	"strings"
)

func main() {
	flags := myls.ParseArgs(os.Args[1:])
	var paths []string
	for _, arg := range os.Args[1:] {
		if !strings.HasPrefix(arg, "-") {
			paths = append(paths, arg)
		}
	}
	if flags["Help"] {
		fmt.Println("Usage: myls [OPTION]... [FILE]...\nList information about the FILEs (the current directory by default).\nSort entries alphabetically if none of -cftuvSUX nor --sort is specified.\n\nMandatory arguments to long options are mandatory for short options too.\n  -R, --recursive     list subdirectories recursively\n  -r, --reverse      reverse order while sorting\n  -a, --all          do not ignore entries starting with .\n  -l                 use a long listing format\n  -t                 sort by time, newest first; see --time")
		return
	}
	if len(paths) == 0 {
		paths = append(paths, ".")
	}
	for _, path := range paths {
		myls.MyLs(path, flags)
	}
}
