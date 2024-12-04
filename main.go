package main

import (
	"os"

	myls "my-ls-1/Functions"
)

func main() {
	myls.Path(os.Args[1:])
}
