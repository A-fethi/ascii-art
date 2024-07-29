package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	chars := ""
	args := os.Args[1:]
	a:= 0
	if len(args) == 1 {
		// MapFiller('s')

		for _, v := range args[0] {
			// MapFiller(v)
			if !strings.ContainsRune(chars, v) {
				chars += string(v)
			}
		}
		r := []rune(chars)
		for i := 0; i < len(r)-1; i++ {
			for j := 0; j < len(r)-1-i; j++ {
				if r[i] > r[i+1] {
					r[i], r[i+1] = r[i+1], r[i]
				}
			}
		}
		for _, vv := range r {
			MapFiller(vv)
		}
	}
	for a < 8 {
		for _, char := range args[0] {
			fmt.Print(Map[string(char)][a])
		}
		fmt.Println()
		a++
	}
	fmt.Println(Map)
}

var Map = make(map[string][]string)

func MapFiller(r rune) map[string][]string {
	counter := 0
	c := 0
	content, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		defer content.Close()

		scaner := bufio.NewScanner(content)
		for scaner.Scan() {
			if counter == (int(r) - 31) {
				if c < 8 {
					Map[string(r)] = append(Map[string(r)], scaner.Text())
					c++
				} else {
					break
				}
				// fmt.Println("ysf")
			} else if scaner.Text() == "" {
				counter++
			}
		}
		//fmt.Println(counter)
	}
	return Map
}
