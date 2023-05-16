package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

var elementCount int


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
		elementCount = 0
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
	w.Write([]byte("=()=>{const e=document.createElement('div');e.setAttribute('id','"))
	w.WriteString(componentName)
	w.Write([]byte("');"))

	r := bufio.NewReader(f)

	for {
		_, err := r.ReadBytes('<')
		if err != nil {
			break
		}

		num := proccessElement(w, r)
		w.WriteString(fmt.Sprintf("e.appendChild(e%d);", num))
	}

	w.Write([]byte("return e;};\n"))
}

func proccessElement(w *bufio.Writer, r *bufio.Reader) int {
	elemNumber := elementCount
	elementCount++
	wholeTag, _ := r.ReadString('>')
	stuff := strings.Split(wholeTag, " ")

	// remove the '>' lol
	if len(stuff) > 1 {
		stuff[len(stuff)-1] = stuff[len(stuff)-1][:len(stuff[len(stuff)-1])-1]
	} else {
		stuff[0] = stuff[0][:len(stuff[0])-1]
	}

	if isAnHTMLElement(stuff[0]) {
		w.WriteString(fmt.Sprintf("const e%d=document.createElement('%s');", elemNumber, stuff[0]))
	} else {
		w.WriteString(fmt.Sprintf("const e%d=%s();", elemNumber, stuff[0]))
	}

	for _, prop := range stuff[1:] {
		split := strings.Split(prop, "=")
		if len(split) < 2 {
			continue
		}
		// Remove "" of the value
		if len(split[1]) > 2 {
			split[1] = split[1][1 : len(split[1])-1]
		}

		w.WriteString(fmt.Sprintf("e%d.setAttribute('%s','%s');", elemNumber, split[0], split[1]))
	}

	textContent, _ := r.ReadString('<')
	textContent = strings.TrimSpace(textContent)
	textContent = strings.ReplaceAll(textContent, "\n", " ")
	textContent = strings.ReplaceAll(textContent, "'", "\\'")

	if len(textContent) > 1 {
		w.WriteString(fmt.Sprintf("e%d.textContent='%s';", elemNumber, textContent[:len(textContent)-1]))
	}

	for {
		if nextTag, _ := r.Peek(len(stuff[0]) + 1); bytes.Equal([]byte("/"+stuff[0]), nextTag) {
			r.ReadBytes('>')
			break
		}

		num := proccessElement(w, r)
		w.WriteString(fmt.Sprintf("e%d.appendChild(e%d);", elemNumber, num))
		_, err := r.ReadBytes('<')

		if err != nil {
			break
		}
	}

	return elemNumber
}
