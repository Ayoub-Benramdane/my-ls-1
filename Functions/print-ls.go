package functions

import "fmt"

func AddSingleQuotes(s string) string {
	for _, r := range s {
		if r < 44 {
			return "'" + s + "'"
		}
	}
	return s
}

func LongFormat(slice []LongFormatInfo) {
	for i, item := range slice {
		item.FileName = AddSingleQuotes(item.FileName)
		fmt.Printf("%v %3s %s %s %4d %s %s",
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
			item.FileName = AddSingleQuotes(item.FileName)
			fmt.Printf("%v  ", item.FileName)
		}
	}
}

func BigSlice(masterSlice []LongFormatInfo) {
	for i, item := range masterSlice {
		item.FileName = AddSingleQuotes(item.FileName)
		if i%3 == 0 {
			fmt.Printf("%v  ", item.FileName)
		}
	}
	fmt.Println()
	for i, item := range masterSlice {
		item.FileName = AddSingleQuotes(item.FileName)
		if i%2 == 0 && i%3 != 0 {
			fmt.Printf("%v  ", item.FileName)
		}
	}
	fmt.Println()
	for i, item := range masterSlice {
		item.FileName = AddSingleQuotes(item.FileName)
		if i%2 != 0 && i%3 != 0 {
			fmt.Printf("%v  ", item.FileName)
		}
	}
	fmt.Println()
}

func MeduimSlice(masterSlice []LongFormatInfo) {
	for i, item := range masterSlice {
		item.FileName = AddSingleQuotes(item.FileName)
		if i%2 == 0 {
			fmt.Printf("%v  ", item.FileName)
		}
	}
	fmt.Println()
	for i, item := range masterSlice {
		item.FileName = AddSingleQuotes(item.FileName)
		if i%2 != 0 {
			fmt.Printf("%v  ", item.FileName)
		}
	}
	fmt.Println()
}
