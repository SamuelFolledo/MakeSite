package main

import (
	"fmt" //format
	"html/template" //allows us to do templating
	"io/ioutil"
	"net/http"
	"oset/http"
)

type FileLines struct {
	Title   string //capital means public, small means private
	Message string
}

func main() {
	line, err := readFile("first-post.txt")
	if err != nil {
		panic(err)
	} else {
		fmt.Print(line)
	}

	// paths := []string{
	// 	"html/layout.html", //1h24m gotta have a template
	// }

	// news := FileLines{
	// 	Title: "EYOOOO", Message: line,
	// }
	
	// t := template.ParseFiles("layout.html") //template loader //1h25m is how it is actually read
	// err = t.Execute(os.Stdout, news)        //1h26m Stdout prints it in the terminal
	// if err != nil {
	// 	panic(err)
	// }

	// tmpl := template.Must(template.ParseFiles("html/layout.html"))
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl.Execute(w, news)
	// })
	// http.ListenAndServe(":80", nil)

	// t, err := template.ParseFiles("./html/first.html")
	// if err == nil {
	// 	log.Println("Template parsed successfully....")
	// }
	// err := templates.ExecuteTemplate(w, "html/first.html", nil)
	// if err != nil {
	// 	log.Println("Not Found template")
	// }
	// t.Execute(w, news)
}

func readFile(fileName string) (line string, errorMessage error) { //method that will read a file and return lines or error
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		errorMessage = err
		line = ""
	} else {
		errorMessage = nil
		line = string(fileContents)
	}
	// fmt.Print(string(fileContents))
	return
}
