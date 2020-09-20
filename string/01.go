package main

import "fmt"

func snakeCasedName(name string) string {
	newstr := make([]rune, 0)
	for idx, chr := range name {
		if isUpper := 'A' <= chr && chr <= 'Z'; isUpper {
			if idx > 0 {
				newstr = append(newstr, '_')
			}
			chr -= ('A' - 'a')
		}
		newstr = append(newstr, chr)
	}

	return string(newstr)
}

func main() {

	var test map[string]string

	fmt.Println(len(test), test == nil)

	var c = snakeCasedName("Abc123")
	fmt.Println(c)

}
