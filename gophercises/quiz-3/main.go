package main

import (
	"fmt"
	"io"
	"log"
	"os"
	_ "reflect"
	_ "strings"
)

func add(x int, y int) func() int {
	return func() int {
		return x - y
	}
}

type multi func(x int) int

var m multi = func(x int) int {
	return x * x
}

func inc(val int) {
	val++
}

func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		file.Close()
	}, nil
}

func bufferToString(f *os.File) (string, bool) {
	if fileContent, err := io.ReadAll(f); err == nil {
		return string(fileContent), true
	}

	return "", false
}

func main() {
	// c := add(2, 4)

	v := 2
	inc(v)
	file := "dump.txt"
	f, closer, err := getFile(file)
	if err != nil {
		log.Fatal(err)
	}

	fileContent, _ := bufferToString(f)
	fmt.Println(fileContent)
	// fmt.Println(reflect.TypeOf(f))
	closer()
}
