package main

import ( //format
	// "html/template" //allows us to do templating
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath" //to use filepath.Ext(*fileFlag) to trim file extension
	"strings"
	//"reflect" //package has TypeOf() which returns the Type of an object
	// "text/template"
	// "oset/http"
)

var paths = []string{
	"template.tmpl", //1h24m gotta have a template
}

type FileLines struct {
	Title   string //capital means public, small means private
	Message string
	Done    bool
}

type Article struct {
	Author   string
	NewsList []FileLines
}

// note, that variables are pointers
var fileFlag = flag.String("file", "", "Name of file")
var boolFlag = flag.Bool("bool", false, "Description of flag") //bool flag

//init() which gets called before main()
func init() {
	flag.StringVar(fileFlag, "", "no file name passed", "Description") // example of flag, which is better to do in init that main
}

func main() {
	saveFileFlag()
}

// function used to input filename to generate a new HTML file. Example: `latest-post.txt` flag will generate a `latest-post.html`
func saveFileFlag() {
	flag.Parse()                                                          //parse flags
	fmt.Println("File flag =", *fileFlag)                                 //after flag.Parse(), *fileFlag is now user's --file= input
	var fileName = strings.TrimSuffix(*fileFlag, filepath.Ext(*fileFlag)) //trims the file's extension
	var htmlFileName = "html/" + fileName + ".html"                       //takes a fileName with no extension and add a .html at the end
	//create what we will be storing to html file
	var line = populateLine()
	var news = []FileLines{
		FileLines{Title: "Title 1", Message: line, Done: true},
		FileLines{Title: "Title 2", Message: "MESSAGEE 2", Done: false},
		FileLines{Title: "Title 3", Message: "MESSAGEEE 3", Done: false},
	}
	var articles = Article{Author: "Kobe", NewsList: news}        //contain news to articles variable
	readTmplAndWriteHtml(articles, "template.tmpl", htmlFileName) //create and save an html file from whatever user named the .txt file
}

func readTmplAndWriteHtml(articles Article, tmplName, htmlName string) {
	var t = template.Must(template.New(tmplName).ParseFiles(paths...)) //1) parse files //template loader //1h25m is how it is actually read
	var htmlFile = createFile(htmlName)                                //2) Create html file we will be saving to
	var err = t.Execute(htmlFile, articles)                            //3) execute //1h26m Stdout prints it in the terminal
	if isError(err) {
		return
	}
}

func writeToFile(fileName, lines string) {
	bytesToWrite := []byte(lines)                         //data written
	err := ioutil.WriteFile(fileName, bytesToWrite, 0644) //filename, byte array (binary representation), and 0644 which represents permission number. (0-777) //will create a new text file if that text file does not exist yet
	if isError(err) {
		return
	}
}

func createFile(fileName string) (returnedFile *os.File) {
	// check if file exists
	var _, err = os.Stat(fileName)
	if os.IsNotExist(err) == false { //if file exist, then delete file first
		print(fileName, " exist\n")
		deleteFile(fileName)
	}
	//create file
	var file, errr = os.Create(fileName)
	if isError(errr) {
		return
	}
	returnedFile = file
	fmt.Println("File Created Successfully", fileName)
	return
}

func deleteFile(fileName string) {
	// delete file
	var err = os.Remove(fileName)
	if isError(err) {
		return
	}

	fmt.Println("File Deleted")
}

func printLines(news FileLines) {
	// t := template.Must(template.ParseFiles("html/layout.html")) //template loader //1h25m is how it is actually read
	// err = t.Execute(os.Stdout, news)                            //1h26m Stdout prints it in the terminal
	// if err != nil {
	// 	panic(err)
	// }
	t, err := template.New("news").Parse("You have a task titled\n\"{{ .Title}}\"\n\"{{ .Message}}\"") //1) Parse files
	if isError(err) {
		return
	}

	err = t.Execute(os.Stdout, news) //3) Execute and save the parsedFiles to file
	if isError(err) {
		return
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
	if isError(err) {
		return
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
	if isError(err) {
		return nil
	}
	return files
}

func writeFile(fileName, data string) {
	bytesToWrite := []byte("hello\ngo\n")                       //data written
	err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644) //filename, byte array (binary representation), and 0644 which represents permission number. (0-777) //will create a new text file if that text file does not exist yet
	if isError(err) {
		return
	}
	print("Successful at writing file")
}

func readFile(fileName string) (content string) { //method that will read a file and return lines or error
	fileContents, err := ioutil.ReadFile(fileName)
	if isError(err) {
		return
	}
	// fmt.Print("READ FILE = \n", string(fileContents))
	content = string(fileContents)
	return
}

func isError(err error) bool { //error helper
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return (err != nil)
}
