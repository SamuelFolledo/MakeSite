# 🔗 makesite

[![Go Report Card](https://goreportcard.com/badge/github.com/SamuelFolledo/MakeSite)](https://goreportcard.com/report/github.com/SamuelFolledo/MakeSite)

_Create your own custom Static Site Generator (like [Jekyll](https://jekyllrb.com/) or [Hugo](https://gohugo.io/)) by cloning and fulfilling the requirements in this repo!_

<p align="center">
  <img width="230" height="230" src="https://github.com/SamuelFolledo/MakeSite/blob/master/static/myGopher.png">
</p>

### 📚 Table of Contents

1. [Demo](#demo)
    - [MVP and v1.1 Requirements](#mvp-and-v1.1)
    - [Third Party API: Google Cloud Translator](#api-translator)
2. [Project Structure](#project-structure)
3. [Getting Started](#getting-started)
4. [Deliverables](#deliverables)
5. [Resources](#resources)

## Demo
<a name="mvp-and-v1.1"></a>
### MVP and v1.1 Requirements
- Takes a directory as a flag
- Read contents of text files only inside the directory
- Create an HTML file with the same name as the text file found
- Populate HTML file with codes needed using a [Go Template](template.tmpl)
<img src="https://github.com/SamuelFolledo/MakeSite/blob/master/static/gifs/flagDirOfTextsToHtml.gif" width="1000" height="712">

------

<a name="api-translator"></a>
### Third Party API: Google Cloud Translator
##### To Setup
1. Fork or download this project
2. Follow basic setup [instructions](https://cloud.google.com/translate/docs/basic/setup-basic)
    - [ ] Create or select a project
    - [ ] Enable the Cloud Translation API for this project
    - [ ] Download a private key as JSON
3. Install the Cloud Translation - [Basic client library](https://cloud.google.com/translate/docs/reference/libraries/v2/go)
    ```
    go get -u cloud.google.com/go/translate
    ```
#### To Run
1. __Important:__ Run this command __once__ each time the project starts to set the environment variable. Replace the ```[PATH]``` to the path of the ```.json``` file downloaded from setup 2 step 3. [Instructions](https://cloud.google.com/docs/authentication/production) for more info or for Windows setup
    ```
    export GOOGLE_APPLICATION_CREDENTIALS="[PATH]"
    ```
    For example:
    ```
    export GOOGLE_APPLICATION_CREDENTIALS="/home/user/Downloads/[FILE_NAME].json"
    ```
2. Once ```export...``` command is done, run the app with this command
    ```
    go build && ./makesite
    ```
<img src="https://github.com/SamuelFolledo/MakeSite/blob/master/static/gifs/googleCloudTranslationDemo.gif" width="1000" height="712">

------

## Project Structure

```bash
📂 makesite
├── README.md
├── first-post.txt
├── latest-post.txt
├── makesite.go
└── template.tmpl
```

## Getting Started

1. Visit [github.com/new](https://github.com/new) and create a new repository named `makesite`.
2. Run each command line-by-line in your terminal to set up the project:

```bash
$ cd ~/go/src
$ git clone git@github.com:Make-School-Labs/makesite.git
$ cd makesite
$ git remote rm origin
$ git remote add origin git@github.com:YOUR_GITHUB_USERNAME/makesite.git
```

## Deliverables

**For each task**:

- Complete each task in the order they appear.
- Use [GitHub Task List](https://help.github.com/en/github/managing-your-work-on-github/about-task-lists) syntax to update the task list.

### MVP

Complete the MVP as If you finish early, move on to the stretch challenges.

If you get stuck on any step, be sure to print the output to `stdout`!

#### v1.0 Requirements

- [X] Edit line `4` of `README.md`. Change this line to the following, replacing `YOUR_USERNAME` and `YOUR_REPONAME` with your GitHub username and repository name respectively.
- [X] Read in the contents of the provided `first-post.txt` file.
- [X] Edit the provided HTML template (`template.tmpl`) to display the contents of `first-post.txt`.
- [X] Render the contents of `first-post.txt` using Go Templates and print it to stdout.
- [X] Write the HTML template to the filesystem to a file. Name it `first-post.html`.
- [X] Manually test the generated HTML page by running `/.makesite`. Double-click the `first-post.html` file that appears in your directory after running the command to open the generated page in your browser.
- [X] **Add, commit, and push to GitHub**.
- [X] Add a new flag to your command named `file`. This flag represents the name of any `.txt` file in the same directory as your program. Run `./makesite --file=latest-post.txt` to test.
- [X] Update the `save` function to use the input filename to generate a new HTML file. For example, if the input file is named `latest-post.txt`, the generated HTML file should be named `latest-post.html`.
- [X] **Add, commit, and push to GitHub**.

#### v1.0 Stretch Challenges

- [ ] Use Bootstrap, or another CSS framework, to enhance the style and readability of your template. _Get creative! Writing your very own website generator is a great opportunity to broadcast your style, personality, and development preferences to the world!_

### v1.1

#### v1.1 Requirements

- [X] Create 3 new `.txt` files for testing in the same directory as your project.
- [X] Add a new flag to the `makesite` command named `dir`.
- [X] Use the flag to find all `.txt` files in the given directory. Print them to `stdout`.
- [X] With the list of `.txt` files you found, generate an HTML page for each.
- [X] Run `./makesite --dir=.` to test in your local directory.
- [X] **Add, commit, and push to GitHub**.

#### v1.1 Stretch Challenges

- [ ] Recursively find all `.txt` files in the given directory, as well as it's subdirectories. Print them to `stdout` to make sure. Generate an HTML page for each.
- [ ] When your program finishes, print: `Success! Generated 5 pages.` The `Success!` substring must be <span style="color: green; font-weight:bold;">bold green</span>, and the count (`5`) must be **bold**.
- [ ] Modify the success message to read: `Success! Generated 5 pages (18.2kB total).` Calculate the total by summing the size of each HTML file, then converting the total to kilobytes. Always return one significant digit after the decimal point.
- [ ] Determine how long it took to execute your static site generator. Modify the success message to read: `Success! Generated 5 pages (18.2kB total) in 3.25 seconds.` Always return two significant digits after the decimal point.
- [ ] Test your solutions to these stretch challenges on many different directories containing `.txt` files. Are there any ways to make your code faster?

### v1.2

#### v1.2 Requirements

- [X] Initialize Go modules in your project.
- [X] Add any third party library to your project to enhance it's functionality. Some ideas you might consider include **_(CHOOSE ONLY ONE)_**:
    - [X] Translating page content using Google Translate.
    - [ ] Parse Markdown (`.md`) files and transform them into HTML. `#` through `######` should translate to `<h1>` through `<h6>` elements.
    - [ ] **_FILL IN THE BLANK_**: `I will use the __________________________________ library. The documentation is located at __________________________________. My goal is to use it to __________________________________.`
    - [X] I used __Google Cloud Translation__ library. The documentation is located at [Cloud Translation API](https://console.cloud.google.com/apis/library/translate.googleapis.com?q=translation&project=go-makesite&folder&organizationId), [GitHub googleapis](https://github.com/googleapis/google-cloud-go/blob/master/translate/translate.go), and [pkg.go.dev](https://pkg.go.dev/cloud.google.com/go/translate/apiv3?tab=doc). My goal is to use it to translate any strings to whatever language I want to support.
- [X] **Add, commit, and push to GitHub**.

## Resources
- [Google Cloud Translation Reference](https://cloud.google.com/translate/docs/apis)
- [Setting the environment variable](https://cloud.google.com/docs/authentication/production)

### Lesson Plans

- [**BEW 2.5**: Project #1 - SSGs](https://make-school-courses.github.io/BEW-2.5-Strongly-Typed-Ecosystems/#/Lessons/SSGProject): Code samples you can use to complete the MVP requirements.
- [**BEW 2.5**: Files & Directories](https://make-school-courses.github.io/BEW-2.5-Strongly-Typed-Ecosystems/#/Lessons/FilesDirectories): Code samples you can use to complete v1.1 requirements.
- [**BEW 2.5**: Files & Directories](https://make-school-courses.github.io/BEW-2.5-Strongly-Typed-Ecosystems/#/Lessons/3rdPartyLibs): Code samples you can use to complete v1.2 requirements.

### Example Code

- [**Go By Example**: Reading Files](https://gobyexample.com/reading-files)
- [**Go By Example**: Writing Files](https://gobyexample.com/writing-files)
- [**Go By Example**: Panic](https://gobyexample.com/panic)
- [**GopherAcademy**: Using Go Templates](https://blog.gopheracademy.com/advent-2017/using-go-templates/)
- [**rapid7.com**: Building a Simple CLI Tool with Golang](https://blog.rapid7.com/2016/08/04/build-a-simple-cli-tool-with-golang/)

## Tips
- __Run commands from a .sh file__
    1. create a ```.sh``` file
    2. run this in the terminal at first, once after creating the ```.sh``` file
        ```
        $ chmod +x fileName.sh
        ```
    3. now to run the file simply type this
        ```
        ./run.sh
        ```