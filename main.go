package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: Must use \"okra [script]\" to run a .okr file")
		os.Exit(-1)
	}
	runFile(os.Args[1])
}

func runFile(path string) {
	bytes, err := ioutil.ReadFile(path)
	check(err)
	run(string(bytes))
}

func run(source string) {
	// scanner := &Scanner{}
	// tokens := scanner.scanTokens()

	// for i, token := range tokens {
	// 	fmt.Println(i, token)
	// }
}

func reportError(line int, loc string, message string) {
	fmt.Println("[line", line, "] Error", loc, ":", message)
}

func check(e error) {
	if e != nil {
		fmt.Println("Error:", e)
		os.Exit(-1)
	}
}
