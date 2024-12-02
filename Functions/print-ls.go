package functions

import (
	"fmt"
	"io/fs"
	"strconv"
)

func AddSingleQuotes(name string, permission fs.FileMode) string {
	green, blue, reset, color := "\033[32m", "\033[34m", "\033[0m", "\033[0m"
	if fmt.Sprintf("%s", permission)[0] == 'd' {
		color = blue
	} else if fmt.Sprintf("%s", permission)[1:4] == "rwx" {
		color = green
	}
	for _, r := range name {
		if r < 44 {
			return "'" + color + name + reset + "'"
		}
	}
	return color + name + reset
}

func LongFormat(slice []LongFormatInfo) {
	for i, item := range slice {
		item.FileName = AddSingleQuotes(item.FileName, item.Permissions)
		fmt.Printf("%v %"+strconv.Itoa(len(item.NumberLinks))+"s %-5s %-5s %7d %-10s %s",
			item.Permissions,
			item.NumberLinks,
			item.User,
			item.Group,
			item.Size,
			item.Time.Format("Jan 02 15:04"),
			item.FileName,
		)
		if i != len(slice)-1 {
			fmt.Println()
		}
	}
}

func ShortFormat(masterSlice []LongFormatInfo) {
	if len(masterSlice) > 20 {
		BigSlice(masterSlice)
	} else if len(masterSlice) > 10 {
		MeduimSlice(masterSlice)
	} else {
		for _, item := range masterSlice {
			item.FileName = AddSingleQuotes(item.FileName, item.Permissions)
			fmt.Printf("%v  ", item.FileName)
		}
	}
}

func BigSlice(masterSlice []LongFormatInfo) {
	for i, item := range masterSlice {
		item.FileName = AddSingleQuotes(item.FileName, item.Permissions)
		if i%3 == 0 {
			fmt.Printf("%v  ", item.FileName)
		}
	}
	fmt.Println()
	for i, item := range masterSlice {
		item.FileName = AddSingleQuotes(item.FileName, item.Permissions)
		if i%2 == 0 && i%3 != 0 {
			fmt.Printf("%v  ", item.FileName)
		}
	}
	fmt.Println()
	for i, item := range masterSlice {
		item.FileName = AddSingleQuotes(item.FileName, item.Permissions)
		if i%2 != 0 && i%3 != 0 {
			fmt.Printf("%v  ", item.FileName)
		}
	}
	fmt.Println()
}

func MeduimSlice(masterSlice []LongFormatInfo) {
	for i, item := range masterSlice {
		item.FileName = AddSingleQuotes(item.FileName, item.Permissions)
		if i%2 == 0 {
			fmt.Printf("%v  ", item.FileName)
		}
	}
	fmt.Println()
	for i, item := range masterSlice {
		item.FileName = AddSingleQuotes(item.FileName, item.Permissions)
		if i%2 != 0 {
			fmt.Printf("%v  ", item.FileName)
		}
	}
	fmt.Println()
}
