package main

import (
	"bufio"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Create("compiled.js")
	check(err)

	defer f.Close()

	w := bufio.NewWriter(f)

	for _, file := range os.Args[1:] {
		compileComponent(w, file)
	}
	err = w.Flush()
	check(err)
}

func compileComponent(w *bufio.Writer, filePath string) {
	f, err := os.Open(filePath)
	check(err)
	defer f.Close()

	filePathSplitted := strings.Split(filePath, "/")
	fileName := filePathSplitted[len(filePathSplitted)-1]
	componentName := strings.Split(fileName, ".")[0]
	w.Write([]byte("const "))
	w.WriteString(componentName)
	w.Write([]byte("=()=>{"))

	_ = bufio.NewReader(f)

	w.Write([]byte("};"))
}
