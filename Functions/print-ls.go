package functions

import (
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

func AddSingleQuotes(name string) string {
	runes := []rune{' ', '*', '?', '(', ')', '$', '\\', '\'', '&', '|', '<', '>', '~'}
	for _, r := range runes {
		if strings.ContainsRune(name, r) {
			return "'" + name + "'"
		}
	}
	return name
}

func Color(name string, permission fs.FileMode) string {
	green, blue, reset, yellow := "\033[32m", "\033[34m", "\033[0m", "\033[33m"
	if fmt.Sprintf("%s", permission)[0] == 'd' {
		return blue + name + reset
	} else if fmt.Sprintf("%s", permission)[0] == '-' && fmt.Sprintf("%s", permission)[3:4] != "x" {
		return name
	} else if fmt.Sprintf("%s", permission)[0] == '-' {
		return green + name + reset
	}
	return yellow + name + reset
}

func LongFormat(slice []LongFormatInfo) {
	for _, item := range slice {
		item.FileName = AddSingleQuotes(item.FileName)
		fmt.Printf("%v %"+strconv.Itoa(len(item.NumberLinks))+"s %s %s %7d %s ",
			item.Permissions,
			item.NumberLinks,
			item.User,
			item.Group,
			item.Size,
			item.Time.Format("Jan 2 15:04"),
		)
		target, err := os.Readlink(item.FileName)
		if err != nil {
			fmt.Printf("%s\n", Color(item.FileName, item.Permissions))
		} else {
			fmt.Printf("\033[36m%s\033[0m -> %s\n", item.FileName, target)
		}
	}
}

func ShortFormat(masterSlice []LongFormatInfo) {
	for _, item := range masterSlice {
		item.FileName = AddSingleQuotes(item.FileName)
		_, err := os.Readlink(item.FileName)
		if err != nil {
			fmt.Printf("%v  ", Color(item.FileName, item.Permissions))
		} else {
			fmt.Printf("\033[31;40m%s\033[0m  ", item.FileName)
		}
	}
}
