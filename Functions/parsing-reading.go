package functions

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func ParseArgs(args []string) ([]string,map[string]bool) {
	paths, flags := []string{}, make(map[string]bool)
	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			arg = strings.TrimPrefix(arg, "--")
			switch arg {
				case "recursive": flags["Recursive"] = true
				case "reverse": flags["Reverse"] = true
				case "all": flags["All"] = true
				case "help": flags["Help"] = true
				default: fmt.Printf("myls: unrecognized option '--%v'\nTry 'myls --help' for more information\n", string(arg)); os.Exit(0)
			}
		} else if strings.HasPrefix(arg, "-") && len(arg) != 1 && arg[1] != '/' {
			arg = strings.TrimPrefix(arg, "-")
			for i := 0; i < len(arg); i++ {
				switch arg[i] {
					case 'R': flags["Recursive"] = true
					case 'r': flags["Reverse"] = true
					case 'a': flags["All"] = true
					case 't': flags["Time"] = true
					case 'l': flags["LongFormat"] = true
					default: fmt.Printf("./myls: invalid option '--%v'\nTry './myls --help' for more information\n", string(arg[i])) ;os.Exit(0)
				}
			}
		} else {
			paths = append(paths, arg)
		}
	}
	return paths, flags
}

func CheckPath(path string, flags map[string]bool) ([]fs.FileInfo, int) {
	var List []fs.FileInfo
	link := 0
	items, err := os.ReadDir(path)
	if err != nil {
		currentDir, err := os.Lstat(path)
		if err != nil {
			fmt.Printf("myls: cannot access '%v': %v\n", path, err)
			os.Exit(0)
		}
		List = append(List, currentDir)
	} else {
		currentDir, err := os.Lstat(path)
		if err == nil && !strings.HasSuffix(path, "/") && currentDir.Mode()&os.ModeSymlink != 0 {
			List = append(List, currentDir)
			link = -1
		} else {
			if flags["All"] {
				List, err = HidenDirectories(List)
				if err != nil {
					fmt.Printf("myls: cannot access '%v': %v\n", path, err)
					os.Exit(0)
				}
			}
			for _, item := range items {
				itemInfo, err := item.Info()
				if err != nil {
					return List, link
				}
				List = append(List, itemInfo)
			}
		}
	}
	return List, link
}

func HidenDirectories(List []os.FileInfo) ([]os.FileInfo, error) {
    currentDir, err := os.Stat(".")
    if err != nil {
        return List, err
    }
    parentDir, err := os.Stat("..")
    if err != nil {
        return List, err
    }
    return append(List, currentDir, parentDir), nil
}
