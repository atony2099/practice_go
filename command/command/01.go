package main

import (
	"flag"
	"fmt"
	"os"
)

func cmd() {

	str := flag.String("s", "defaultValue", "Help text.")
	ints := flag.Int("a", 123, "int value")

	flag.Usage = func() {
		fmt.Println("hello world")
	}

	flag.Parse()

	fmt.Println(*str, *ints)

}

func main() {

	// fmt.Println(len(os.Args))

	cmd()
	// sub()

}

func sub() {

	// 1. foo cmd
	foocmd := flag.NewFlagSet("foo", flag.ExitOnError)

	// // 2.  enable
	foolEnable := foocmd.Bool("enable", false, "enable")

	if len(os.Args) < 3 {
		os.Exit(1)
	}

	foocmd.Parse(os.Args[2:])

	fmt.Println(*foolEnable)
}
