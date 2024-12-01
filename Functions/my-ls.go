package functions

import (
	"fmt"
	"io/fs"
	"os/user"
	"strings"
	"syscall"
	"time"
)

type LongFormatInfo struct {
	Permissions fs.FileMode
	NumberLinks string
	User        string
	Group       string
	Size        int64
	Time        time.Time
	FileName    string
}

func MasterSlice(list []fs.FileInfo, flags map[string]bool, total *int) []LongFormatInfo {
	masterSlice := []LongFormatInfo{}
	var User, Group, NumberLinks string
	for _, item := range list {
		if !flags["All"] && item.Name()[0] == '.' {
			continue
		}
		if stat, ok := item.Sys().(*syscall.Stat_t); ok {
			User = fmt.Sprintf("%d", stat.Uid)
			Group = fmt.Sprintf("%d", stat.Gid)
			if item.Name()[0] == '.' {
				NumberLinks = fmt.Sprintf("%d", stat.Nlink-1)
			} else {
				NumberLinks = fmt.Sprintf("%d", stat.Nlink)
			}
		}
		if user, err := user.LookupId(User); err == nil {
			User = user.Username
		}
		if group, err := user.LookupGroupId(Group); err == nil {
			Group = group.Name
		}
		*total += int(item.Size()) / 1020
		element := LongFormatInfo{item.Mode(), NumberLinks, User, Group, item.Size(), item.ModTime(), item.Name()}
		masterSlice = append(masterSlice, element)
	}
	return masterSlice
}

func MyLs(path string, flags map[string]bool) int {
	list := CheckPath(path, flags)
	total := 0
	masterSlice := MasterSlice(list, flags, &total)
	SortLs(masterSlice)
	if flags["Time"] {
		SortByTime(masterSlice)
	}
	if flags["Reverse"] {
		ReverseSorting(masterSlice)
	}
	if flags["LongFormat"] {
		fmt.Println("total", int(total))
		LongFormat(masterSlice)
	} else {
		path = AddSingleQuotes(path)
		if flags["Recursive"] {
			fmt.Printf("%v:\n", path)
		}
		ShortFormat(masterSlice)
		if flags["Recursive"] && path != "." && len(masterSlice) != 0 && IsDir(list) {
			fmt.Println()
		}
	}
	for _, item := range list {
		if flags["Recursive"] && item.IsDir() {
			fmt.Println()
			Recursive(item, path, flags)
		}
	}
	return len(masterSlice)
}

func Recursive(item fs.FileInfo, path string, flags map[string]bool) {
	if !flags["All"] && strings.HasPrefix(item.Name(), ".") || item.Name() == "." || item.Name() == ".." {
		return
	}
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	MyLs(path+item.Name(), flags)
}
