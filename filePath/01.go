package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	names := "hello.abcdd.go"

	s := strings.TrimSuffix(names, filepath.Ext(names))

	fmt.Println(s)

	// extension := filepath.Ext(names)
	// base := path.Base(names)

	// fmt.Println("Extension 1:", base)

	// fmt.Println(extension)

	// extension = filepath.Ext("./hello.jpg")
	// fmt.Println("Extension 2:", extension)
}
