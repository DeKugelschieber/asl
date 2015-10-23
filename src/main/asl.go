package main

import (
	"asl"
	"fmt"
	"io/ioutil"
)

const version = "0.1"

func usage() {
    fmt.Println("Usage: asl [-v|-r|-pretty] <input file/folder> [<output file/folder>]\n")
    fmt.Println("-v (optional) shows asl version")
    fmt.Println("-r (optional) recursivly compile all asl files in folder")
    fmt.Println("-pretty (optional) activates pretty printing\n")
    fmt.Println("<input file/folder> file or directory to compile")
    fmt.Println("<output file/folder> (optional) output file/folder, if not set, files will be created alongside their asl files")
}

func main() {
	// read test file
	code, _ := ioutil.ReadFile("in/complex.asl")
	token := asl.Tokenize(code)
	out := asl.Parse(token, true)

	fmt.Print("OUTPUT:\n-------\n"+out)
}
