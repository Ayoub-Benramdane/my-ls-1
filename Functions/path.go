package functions

import (
	"fmt"
)

func Path(path []string) {
	if len(path[0]) > 10 && path[0][1:5] == "[0m." {
		path[0] = path[0][10:]
	}
	flags, paths := ParseArgs(path)
	if flags["Help"] {
		fmt.Println("Usage: myls [OPTION]... [FILE]...\nList information about the FILEs (the current directory by default).\nSort entries alphabetically if none of -cftuvSUX nor --sort is specified.\n\nMandatory arguments to long options are mandatory for short options too.\n  -R, --recursive     list subdirectories recursively\n  -r, --reverse      reverse order while sorting\n  -a, --all          do not ignore entries starting with .\n  -l                 use a long listing format\n  -t                 sort by time, newest first; see --time")
		return
	} else if len(paths) == 0 {
		paths = append(paths, ".")
	}
	dirSlice, fileSlice := SplitPath(paths)
	FileSlice(fileSlice, dirSlice, flags)
	DirSlice(fileSlice, dirSlice, flags)
}
