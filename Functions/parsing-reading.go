package functions

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func ParseArgs(args []string, Flags *map[string]bool) ([]string) {
	paths := []string{}
	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			arg = strings.TrimPrefix(arg, "--")
			switch arg {
				case "recursive": (*Flags)["Recursive"] = true
				case "reverse": (*Flags)["Reverse"] = true
				case "all": (*Flags)["All"] = true
				case "help": (*Flags)["Help"] = true
				default: fmt.Printf("myls: unrecognized option '--%v'\nTry 'myls --help' for more information\n", string(arg)); os.Exit(0)
			}
		} else if strings.HasPrefix(arg, "-") && len(arg) != 1 && arg[1] != '/' {
			arg = strings.TrimPrefix(arg, "-")
			for i := 0; i < len(arg); i++ {
				switch arg[i] {
					case 'R': (*Flags)["Recursive"] = true
					case 'r': (*Flags)["Reverse"] = true
					case 'a': (*Flags)["All"] = true
					case 't': (*Flags)["Time"] = true
					case 'l': (*Flags)["LongFormat"] = true
					default: fmt.Printf("./myls: invalid option '--%v'\nTry './myls --help' for more information\n", string(arg[i])) ;os.Exit(0)
				}
			}
		} else {
			paths = append(paths, arg)
		}
	}
	return paths
}

func CheckPath(path string, flags map[string]bool) []fs.FileInfo {
	var List []fs.FileInfo
	if flags["All"] {
		currentDir, err := os.Stat(".")
		if err != nil {
			fmt.Printf("myls: cannot access '%v': %v\n", path, err)
			os.Exit(0)
		}
		parentDir, err := os.Stat("..")
		if err != nil {
			fmt.Printf("myls: cannot access '%v': %v\n", path, err)
			os.Exit(0)
		}
		List = append(List, currentDir, parentDir)
	}
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
