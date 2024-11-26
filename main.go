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


	// dirSlice, fileSlice := []LongFormatInfo{},[]LongFormatInfo{}

	// for _, path := range paths {
	// 	if path.IsDir() {
	// 		dirSlice = append(dirSlice, item)
	// 	} else {
	// 		fileSlice = append(fileSlice, item)
	// 	}
	// }

	// for _, item := range fileSlice {
	// 	fmt.Printf("%v  ", item.FileName)
	// }
	// fmt.Println("\n")
	// if len(dirSlice) != 0 {
	// 	for _, item := range dirSlice {
	// 		fmt.Printf("%v:\n", item.FileName)
	// 		for _, item := range dirSlice {
	// 			fmt.Printf("%v  ", item.FileName)
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	for _, path := range paths {
		myls.MyLs(path, flags)
	}
}
