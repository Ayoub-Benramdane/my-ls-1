package functions

import (
	"fmt"
	"io/fs"
	"os"
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
		if !flags["All"] && strings.HasPrefix(item.Name(), ".") {
			continue
		}
		if stat, ok := item.Sys().(*syscall.Stat_t); ok {
			if flags["All"] || !strings.HasPrefix(item.Name(), ".") {
				*total += int(stat.Blocks)
			}
			User = fmt.Sprintf("%d", stat.Uid)
			Group = fmt.Sprintf("%d", stat.Gid)
			NumberLinks = fmt.Sprintf("%d", stat.Nlink)
		}
		if user, err := user.LookupId(User); err == nil {
			User = user.Username
		}
		if group, err := user.LookupGroupId(Group); err == nil {
			Group = group.Name
		}
		masterSlice = append(masterSlice, LongFormatInfo{item.Mode(), NumberLinks, User, Group, item.Size(), item.ModTime(), item.Name()})
	}
	return masterSlice
}

func MyLs(path string, flags map[string]bool, totalPath int) {
	list, total := CheckPath(path, flags)
	masterSlice := MasterSlice(list, flags, &total)
	SortLs(masterSlice)
	if flags["Time"] {
		SortByTime(masterSlice)
	}
	if flags["Reverse"] {
		ReverseSorting(masterSlice)
	}
	_, err := os.ReadDir(path)
	if err == nil && (flags["Recursive"] || totalPath > 1) {
		fmt.Printf("%v:\n", AddSingleQuotes(path))
	}
	if flags["LongFormat"] {
		_, err := os.ReadDir(path)
		if total != -1 && err == nil {
			fmt.Println("total", total/2)
		}
		LongFormat(masterSlice)
	} else {
		ShortFormat(masterSlice, totalPath)
	}
	if flags["Recursive"] {
		for _, item := range list {
			if item.IsDir() {
				if flags["All"] || !strings.HasPrefix(item.Name(), ".") {
					fmt.Println()
				}
				Recursive(item, path, flags, totalPath)
			}
		}
	}
}

func Recursive(item fs.FileInfo, path string, flags map[string]bool, totalPath int) {
	if !flags["All"] && strings.HasPrefix(item.Name(), ".") || item.Name() == "." || item.Name() == ".." {
		return
	} else if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	MyLs(path+item.Name(), flags, totalPath)
}
