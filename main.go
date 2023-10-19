package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
}

func main() {
	f, err := os.Create("compiled.js")
	check(err)

	defer f.Close()

	w := bufio.NewWriter(f)
	ch := make(chan string)

	for _, file := range os.Args[1:] {
		go compileComponent(ch, file)
	}

	for range os.Args[1:] {
		w.WriteString(<-ch)
		w.Flush()
	}
}

func compileComponent(ch chan string, filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file %s.\n", filePath)
		ch <- ""
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)

	filePathSplitted := strings.Split(filePath, "/")
	fileName := filePathSplitted[len(filePathSplitted)-1]
	componentName := strings.Split(fileName, ".")[0]
	variables, err := r.ReadString('<')
	if err != nil {
		fmt.Printf("Couldn't compile file %s.\n", fileName)
		return
	}

	output := fmt.Sprintf("/**\n * Generates %s component.\n", componentName)

	var variablesNames string
	for _, v := range strings.Split(variables, "\n") {
		if v[0] != '-' {
			continue
		}
		output += fmt.Sprintf(" * @param %s\n", v[2:])
		variablesNames += strings.Split(v, " ")[2] + ","
	}

	if len(variablesNames) > 0 {
		variablesNames = variablesNames[:len(variablesNames)-1]
	}

	output += fmt.Sprintf(" * @return Component\n */\nconst %s=(%s)=>{const e=document.createElement('div');e.setAttribute('id','%s');", componentName, variablesNames, componentName)

	count := 0
	for {
		output, count = proccessElement(r, output, count)
		output += fmt.Sprintf("e.appendChild(e%d);", count)

		_, err := r.ReadBytes('<')
		if err != nil {
			break
		}
	}

	output += "return e;};\n"
	ch <- output
}

func proccessElement(r *bufio.Reader, out string, count int) (string, int) {
	wholeTag, _ := r.ReadString('>')
	stuff := strings.Split(wholeTag, " ")
	this_id := count
	count += 1

	// remove the '>' lol
	if len(stuff) > 1 {
		stuff[len(stuff)-1] = stuff[len(stuff)-1][:len(stuff[len(stuff)-1])-1]
	} else {
		stuff[0] = stuff[0][:len(stuff[0])-1]
	}

	if isAnHTMLElement(stuff[0]) {
		out += fmt.Sprintf("const e%d=document.createElement('%s');", this_id, stuff[0])
	} else {
		out += fmt.Sprintf("const e%d=%s();", this_id, stuff[0])
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

		out += fmt.Sprintf("e%d.setAttribute('%s','%s');", this_id, split[0], split[1])
	}

	textContent, _ := r.ReadString('<')
	textContent = strings.TrimSpace(textContent)
	textContent = strings.ReplaceAll(textContent, "\n", " ")
	textContent = strings.ReplaceAll(textContent, "\\", "\\\\")
	textContent = strings.ReplaceAll(textContent, "`", "\\`")

	if len(textContent) > 1 {
		out += fmt.Sprintf("e%d.textContent=`%s`;", this_id, textContent[:len(textContent)-1])
	}

	for {
		if nextTag, _ := r.Peek(len(stuff[0]) + 1); bytes.Equal([]byte("/"+stuff[0]), nextTag) {
			r.ReadBytes('>')
			break
		}

		out, count = proccessElement(r, out, count)
		out += fmt.Sprintf("e%d.appendChild(e%d);", this_id, count)
		_, err := r.ReadBytes('<')

		if err != nil {
			break
		}
	}

	return out, count
}
