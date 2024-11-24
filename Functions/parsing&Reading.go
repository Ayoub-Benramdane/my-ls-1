package functions

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func ParseArgs(args []string) map[string]bool {
	Flags := make(map[string]bool)
	Flags["LongFormat"] = false
	Flags["Recursive"] = false
	Flags["Reverse"] = false
	Flags["Time"] = false
	Flags["Help"] = false
	Flags["All"] = false
	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			arg = strings.TrimPrefix(arg, "--")
			if arg == "recursive" {
				Flags["Recursive"] = true
			} else if arg == "reverse" {
				Flags["Reverse"] = true
			} else if arg == "all" {
				Flags["All"] = true
			} else if arg == "help" {
				Flags["Help"] = true
			} else {
				fmt.Printf("myls: unrecognized option -- '%v'\nTry 'myls --help' for more information\n", string(arg))
				os.Exit(0)
			}
		} else if strings.HasPrefix(arg, "-") {
			arg = strings.TrimPrefix(arg, "-")
			for i := 0; i < len(arg); i++ {
				if arg[i] == 'R' {
					Flags["Recursive"] = true
				} else if arg[i] == 'r' {
					Flags["Reverse"] = true
				} else if arg[i] == 'a' {
					Flags["All"] = true
				} else if arg[i] == 't' {
					Flags["Time"] = true
				} else if arg[i] == 'l' {
					Flags["LongFormat"] = true
				} else {
					fmt.Printf("myls: unrecognized option -- '%v'\nTry 'myls --help' for more information\n", string(arg[i]))
					os.Exit(0)
				}
			}
		}
	}
	return Flags
}

func CheckPath(path string) []fs.FileInfo {
	var List []fs.FileInfo
	items, err := os.ReadDir(path)
	if err != nil {
		currentDir, err := os.Stat(path)
		if err != nil {
			fmt.Printf("myls: cannot access '%v': %v\n", path, err)
			os.Exit(0)
		}
		List = append(List, currentDir)
	}
	for _, item := range items {
		itemInfo, err := item.Info()
		if err != nil {
			return List
		}
		List = append(List, itemInfo)
	}
	return List
}
