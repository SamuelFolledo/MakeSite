package main

import ( //format
	// "html/template" //allows us to do templating

	"html/template"
	"io/ioutil"
	"log"
	"os"
	//"reflect" //package has TypeOf() which returns the Type of an object
	// "text/template"
	// "oset/http"
)

type FileLines struct {
	Title   string //capital means public, small means private
	Message string
}

func main() {
	line := populateLine()
	news := FileLines{
		"EYOOOO", line,
	}
	// t := template.Must(template.ParseFiles("html/layout.html")) //template loader //1h25m is how it is actually read
	// err = t.Execute(os.Stdout, news)                            //1h26m Stdout prints it in the terminal
	// if err != nil {
	// 	panic(err)
	// }
	t, err := template.New("news").Parse("You have a task titled\n\"{{ .Title}}\"\n\"{{ .Message}}\"")
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, news)
	if err != nil {
		panic(err)
	}
	// newNews := FileLines{"Go", "Contribute to any Go project"}
	// err = t.Execute(os.Stdout, newNews)
}

func populateLine() (line string) {
	directory := "/Users/macbookpro15/Desktop/MakeSite"
	fileName := "sample.txt" //file we will be searching for
	file := findFile(fileName, directory)
	line = ""
	if file != nil {
		line = readFile(fileName)
	}
	return
}

func findFile(fileName, directory string) (fileResult os.FileInfo) { //func that finds a filename from a directory and returns the file found. //[]os.FileInfo is a slice of interfaces
	files, err := ioutil.ReadDir(directory) //ReadDir returns a slice of FileInfo structs
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files { //loop through each files
		// print("File: ", file.Name(), " is ")
		if file.IsDir() { //skip if file is directory
			continue
		}
		// fmt.Print(file.IsDir(), " = ", file.Name(), "\n")
		if file.Name() == fileName {
			// println("\n\nFound", fileName)
			fileResult = file
			return
		}
	}
	return
}

func readDirectory(directory string) []os.FileInfo { //method that takes a directory and returns a list of files and directories
	files, err := ioutil.ReadDir(directory) //ReadDir returns a slice of FileInfo structs
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func writeFile(fileName, data string) {
	bytesToWrite := []byte("hello\ngo\n")                       //data written
	err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644) //filename, byte array (binary representation), and 0644 which represents permission number. (0-777) //will create a new text file if that text file does not exist yet
	if err != nil {
		panic(err)
	}
	print("Successful at writing file")
}

func readFile(fileName string) (content string) { //method that will read a file and return lines or error
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	// fmt.Print("READ FILE = \n", string(fileContents))
	content = string(fileContents)
	return
}
