SRC = $(wildcard *.html)

.PHONY = build

build: compiled.js

compiled.js: $(SRC)
	html2js $(SRC) 
