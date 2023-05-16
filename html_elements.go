package main

/**
 * Even though a qsort would be theorically faster,
 * having the string array of html elements in order
 * of popularity and doing a sequential search
 * will be faster in real world uses.
 */
func isAnHTMLElement(e string) bool {
	elements := getHTMLElements()
	for _, v := range elements {
		if e == v {
			return true
		}
	}
	return false
}

func getHTMLElements() [15]string {
	return [15]string{
		"div", "a", "button", "form",
		"h1", "h2", "h3", "h4", "h5", "h6",
		"p", "ul", "ol", "li", "br",
	}
}

