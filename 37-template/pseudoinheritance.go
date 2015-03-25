package main

import "fmt"
import "html/template"
import "os"

func main() {
	tmpl := make(map[string]*template.Template)
	tmpl["index.html"] = template.Must(template.ParseFiles("view/index.html", "view/base.html"))
	tmpl["other.html"] = template.Must(template.ParseFiles("view/other.html", "view/base.html"))

	data := map[string]string{"aaa": "bbb"}

	out := tmpl["index.html"].ExecuteTemplate(os.Stdout, "base", data)
	fmt.Println(out)
	out = tmpl["other.html"].ExecuteTemplate(os.Stdout, "base", data)
	fmt.Println(out)
}
