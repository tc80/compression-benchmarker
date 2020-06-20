package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

const (
	inputFile = "input.txt" // for now, just reading from input file containing libs
	// environmental variable containing the path to the local cdnjs repo
	// which is the same as the path to the directory https://github.com/cdnjs/cdnjs after cloning
	// 		ex. "/tmp/cdnjs/cdnjs" (inside this dir is the ajax/ dir)
	cdnjsPathEnv = "CDNJS_PATH"
)

var (
	cdnjsPath = os.Getenv(cdnjsPathEnv)
)

// enforce non-nil
func check(e interface{}) {
	if e != nil {
		panic(e)
	}
}

// NOTES:
//	- compression algorithm and flags must be input as command-line arguments
// 		- ex. ./compress brotli -c -q 11
//
//	- to see uncompressed sizes, simply use 'cat' as the algorithm type
// 		- ex. go run compress.go cat
//
//	- the size of the resulting algorithm's STDOUT in bytes will be printed for each library
// 		from the input file, in their respective order
func main() {
	if len(os.Args) == 1 {
		log.Fatalln("Missing compression algorithm command-line arguments. Ex. ./compress zopfli -c -i1000 --gzip")
	}
	alg, flags := os.Args[1], os.Args[2:]

	// open input file containing list of cdnjs library paths (from within $CDNJS_PATH)
	f, err := os.Open(inputFile)
	check(err)
	defer f.Close()

	// read file line by line
	s := bufio.NewScanner(f)
	for s.Scan() {
		lib := s.Text()
		p := path.Join(cdnjsPath, lib)                // full local path to library
		cmd := exec.Command(alg, append(flags, p)...) // execute algorithm with arguments
		var out bytes.Buffer                          // buffer to hold algorithm's STDOUT
		cmd.Stdout = &out                             // NOTE: algorithm must output its result to STDOUT
		err := cmd.Run()                              // 	ex. the -c flag for zopfli or brotli
		check(err)
		fmt.Printf("%s,%d\n", lib, out.Len()) // print lib name, size of each result
		// NOTE: change this to fmt.Println(out.Len()) if you only want the size
	}
	check(s.Err())
}
