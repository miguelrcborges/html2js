package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

var elementCount int

func getHTMLElements() [11]string {
	return [11]string {
		"div", "a", "button", "form",
		"h1", "h2", "h3", "h4", "h5", "h6",
		"p", 
	}
}

func isAnHTMLElement(e string) bool {
	elements := getHTMLElements();
	for _, v := range elements {
		if e == v {
			return true
		}
	}
	return false
}

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
	w.Write([]byte("=()=>{let e=document.createElement('div');e.setAtrribute('id','"))
	w.WriteString(componentName)
	w.Write([]byte("');"))

	r := bufio.NewReader(f)
	
	for {
		_, err := r.ReadBytes('<')
		if err != nil {
			break;
		} 	

		num := proccessElement(w, r)
		w.WriteString(fmt.Sprintf("e.appendChild(e%d);", num))
	}

	w.Write([]byte("return e;};"))
}


func proccessElement(w *bufio.Writer, r *bufio.Reader) int {
	elemNumber := elementCount
	elementCount++
	wholeTag, _ := r.ReadString('>')
	stuff := strings.Split(wholeTag, " ")

	// remove the '>' lol
	if len(stuff) > 1 {
		stuff[len(stuff) - 1] = stuff[len(stuff) - 1][:len(stuff[len(stuff) - 1]) - 1]
	} else {
		stuff[0] = stuff[0][:len(stuff[0]) - 1]
	}

	w.WriteString(fmt.Sprintf("let e%d=document.createElement('%s');", elemNumber, stuff[0]))

	textContent, _ := r.ReadString('<')
	textContent = strings.TrimSpace(textContent)

	if len(textContent) > 1 {
		w.WriteString(fmt.Sprintf("e%d.textContent='%s';", elemNumber, textContent[:len(textContent) - 1]))
	}

	for {
		nextTag, _ := r.Peek(len(stuff[0]) + 1)
		fmt.Println(elemNumber, stuff[0], string(nextTag))
		if ; bytes.Compare([]byte("/" + stuff[0]), nextTag) == 0 {
			r.ReadBytes('>')
			break
		}

		num := proccessElement(w, r)
		w.WriteString(fmt.Sprintf("e%d.appendChild(e%d);", elemNumber, num))
		_, err := r.ReadBytes('<')

		if err != nil {
			break;
		} 	
	}

	return elemNumber
}
