package main

import "text/template"
import "strings"
import "log"
import "os"
import "fmt"

// Whoa! my own evil function
func countSpaces(str string) string {
	count := strings.Count(str, " ")
	return fmt.Sprintf("Text \"%s\" has %d spaces inside", str, count)
}

func main() {
	// First we create a FuncMap with which to register the function.
	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"title":       strings.Title,
		"up":          strings.ToUpper,
		"countSpaces": countSpaces,
	}

	// A simple template definition to test our function.
	// We print the input text several ways:
	// - the original
	// - title-cased
	// - title-cased and then printed with %q
	// - printed with %q and then title-cased.
	const templateText = `
Input: {{printf "%q" .}}
Output 0: {{title .}}
Output 1: {{title . | printf "%q"}}
Output 2: {{printf "%q" . | title}}
Output 3: {{printf "%q" . | countSpaces}}
Output 4: {{countSpaces .}}
Output 5 {{ . | up}} is obvious
`

	// Create a template, add the function map, and parse the text.
	tmpl, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	// Run the template to verify the output.
	err = tmpl.Execute(os.Stdout, "the go programming language")
	if err != nil {
		log.Fatalf("execution: %s", err)
	}
}
