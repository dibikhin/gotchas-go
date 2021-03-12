package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/rivo/uniseg"
)

func main() {
	b := "Käse"

	len1 := len([]rune(b))
	len2 := bytes.Count([]byte(b), nil) - 1
	len3 := strings.Count(b, "") - 1
	len4 := utf8.RuneCountInString(b)

	len5 := uniseg.GraphemeClusterCount(b)

	fmt.Println(len1) // 5
	fmt.Println(len2) // 5
	fmt.Println(len3) // 5
	fmt.Println(len4) // 5

	fmt.Println(len5) // 4 (!)
}
