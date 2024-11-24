package functions

import (
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type LongFormatInfo struct {
	Permissions fs.FileMode
	NumberLinks string
	User        string
	Group       string
	Size        string
	Time        time.Time
	FileName    string
}

func MyLs(path string, flags map[string]bool) {
	list := CheckPath(path)
	masterSlice := []LongFormatInfo{}
	var User, Group, NumberLinks string
	var total int
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
		list = append(list, currentDir, parentDir)
	}
	for _, item := range list {
		if !flags["All"] && item.Name()[0] == '.' {
			continue
		}
		if stat, ok := item.Sys().(*syscall.Stat_t); ok {
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
		total += int(item.Size()) / 1020
		size := strconv.Itoa(int(item.Size()))
		element := LongFormatInfo{item.Mode(), NumberLinks, User, Group, size, item.ModTime(), item.Name()}
		masterSlice = append(masterSlice, element)
	}
	SortAlphabetic(&masterSlice)
	if flags["Time"] {
		SortByTime(&masterSlice)
	}
	if flags["Reverse"] {
		ReverseSorting(&masterSlice)
	}
	if flags["LongFormat"] {
		fmt.Println("total", int(total))
		for _, item := range masterSlice {
			fmt.Printf("%v %3s %3s %3s %10s %3s %s\n",
				item.Permissions,
				item.NumberLinks,
				item.User,
				item.Group,
				item.Size,
				item.Time.Format("Jan 01 00:00"),
				item.FileName,
			)
		}
	} else {
		if AddSingleQuotes(path) {
			path = "'" + path + "'"
		}
		if flags["Recursive"] {
			fmt.Printf("%v:\n", path)
		}
		Displaying(masterSlice)
	}
	for _, item := range list {
		if flags["Recursive"] && item.IsDir() {
			if !flags["All"] && strings.HasPrefix(item.Name(), ".") || item.Name() == "." || item.Name() == ".." {
				continue
			}
			if !strings.HasSuffix(path, "/") {
				path += "/"
			}
			fmt.Println()
			MyLs(path+item.Name(), flags)
		}
	}
}

func AddSingleQuotes(s string) bool {
	for _, r := range s {
		if r < 44 {
			return true
		}
	}
	return false
}

func Displaying(masterSlice []LongFormatInfo) {
	if len(masterSlice) > 20 {
		BigSlice(masterSlice)
	} else if len(masterSlice) > 10 {
		MeduimSlice(masterSlice)
	} else {
		for _, item := range masterSlice {
			if AddSingleQuotes(item.FileName) {
				item.FileName = "'" + item.FileName + "'"
			}
			fmt.Printf("%v  ", item.FileName)
		}
		fmt.Println()
	}
}

func BigSlice(masterSlice []LongFormatInfo) {
	for i, item := range masterSlice {
		if AddSingleQuotes(item.FileName) {
			item.FileName = "'" + item.FileName + "'"
		}
		if i%3 == 0 {
			fmt.Printf("%v  ", item.FileName)
		}
	}
	fmt.Println()
	for i, item := range masterSlice {
		if AddSingleQuotes(item.FileName) {
			item.FileName = "'" + item.FileName + "'"
		}
		if i%2 == 0 && i%3 != 0 {
			fmt.Printf("%v  ", item.FileName)
		}
	}
	fmt.Println()
	for i, item := range masterSlice {
		if AddSingleQuotes(item.FileName) {
			item.FileName = "'" + item.FileName + "'"
		}
		if i%2 != 0 && i%3 != 0 {
			fmt.Printf("%v  ", item.FileName)
		}
	}
	fmt.Println()
}

func MeduimSlice(masterSlice []LongFormatInfo) {
	for i, item := range masterSlice {
		if AddSingleQuotes(item.FileName) {
			item.FileName = "'" + item.FileName + "'"
		}
		if i%2 == 0 {
			fmt.Printf("%v  ", item.FileName)
		}
	}
	fmt.Println()
	for i, item := range masterSlice {
		if AddSingleQuotes(item.FileName) {
			item.FileName = "'" + item.FileName + "'"
		}
		if i%2 != 0 {
			fmt.Printf("%v  ", item.FileName)
		}
	}
	fmt.Println()
}
