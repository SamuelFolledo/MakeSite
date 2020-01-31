package main

import ( //format
	// "html/template" //allows us to do templating
	"fmt"
	"io/ioutil"
	"log"
	// "text/template"
	// "oset/http"
)

type FileLines struct {
	Title   string //capital means public, small means private
	Message string
}

func main() {
	directory := "/Users/macbookpro15/Desktop/MakeSite"
	files, err := ioutil.ReadDir(directory) //ReadDir returns a slice of FileInfo structs
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files { //loop through each files
		// fmt.Println(file.Name())
		fmt.Print(file.IsDir(), " = ", file.Name(), "\n")
		if file.Name() == "first-post.txt" { //check if file is what we want
			line, err := readFile("first-post.txt") //read file
			if err != nil {
				panic(err)
			}
			fmt.Print("FILE = ", line) //if no error then print line
			break                      //break from loop
		}
	}

	// line, err := readFile("first-post.txt")
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	// fmt.Print(line)
	// }

	// news := FileLines{
	// 	"EYOOOO", line,
	// }
	// t := template.Must(template.ParseFiles("html/layout.html")) //template loader //1h25m is how it is actually read
	// err = t.Execute(os.Stdout, news)                            //1h26m Stdout prints it in the terminal
	// if err != nil {
	// 	panic(err)
	// }
	// // t, err := template.New("news").Parse("You have a task titled \"{{ .Title}}\" with message: \"{{ .Message}}\"")
	// // if err != nil {
	// // 	panic(err)
	// // }
	// err = t.Execute(os.Stdout, news)
	// if err != nil {
	// 	panic(err)
	// }
	// // code omitted beacuse of brevity
	// newNews := FileLines{"Go", "Contribute to any Go project"}
	// err = t.Execute(os.Stdout, newNews)
}

func readFile(fileName string) (line string, errorMessage error) { //method that will read a file and return lines or error
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		// panic(err)
		errorMessage = err
		line = ""
	} else {
		errorMessage = nil
		line = string(fileContents)
	}
	// fmt.Print(string(fileContents))
	return
}
