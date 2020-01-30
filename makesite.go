package main

import ( //format
	"html/template" //allows us to do templating
	"io/ioutil"
	"os"
	// "text/template"
	// "oset/http"
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
		// fmt.Print(line)
	}

	// paths := []string{
	// 	"html/layout.html", //1h24m gotta have a template
	// }

	news := FileLines{
		"EYOOOO", line,
	}

	// t := template.Must(template.ParseFiles("html/layout.html")) //template loader //1h25m is how it is actually read
	// err = t.Execute(os.Stdout, news)                            //1h26m Stdout prints it in the terminal
	// if err != nil {
	// 	panic(err)
	// }

	t, err := template.New("todos").Parse("You have a task titled \"{{ .Title}}\" with message: \"{{ .Message}}\"")
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, news)
	if err != nil {
		panic(err)
	}
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
