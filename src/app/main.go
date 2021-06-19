package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/joho/godotenv"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	targetFolder   string
	targetFile     string
	searchResult   []string
	commandApp     string
	folderDistPath []string
	rootPath       string
	rootSrcPath    string
	rootDistPath   string
)

func findFile(path string, fileInfo os.FileInfo, err error) error {

	if err != nil {
		fmt.Println(err)
		return nil
	}

	// get absolute path of the folder that we are searching
	absolute, err := filepath.Abs(path)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	if fileInfo.IsDir() {
		fmt.Println("Searching directory ... ", absolute)

		// correct permission to scan folder?
		testDir, err := os.Open(absolute)

		if err != nil {
			if os.IsPermission(err) {
				fmt.Println("No permission to scan ... ", absolute)
				fmt.Println(err)
			}
		}
		testDir.Close()
		return nil
	} else {
		// ok, we are dealing with a file
		// is this the target file?

		// yes, need to support wildcard search as well
		// https://www.socketloop.com/tutorials/golang-match-strings-by-wildcard-patterns-with-filepath-match-function

		matched, err := filepath.Match(targetFile, fileInfo.Name())
		if err != nil {
			fmt.Println(err)
		}

		if matched {
			// yes, add into our search result
			add := absolute
			searchResult = append(searchResult, add)
		}
	}

	return nil
}

func buildConfigurations(searchResult []string, rootDistPath string) {
	for _, v := range searchResult {
		fmt.Println("Processing: " + v)
		indexPath := strings.Index(v, "/config") + 7
		fileUri := v[indexPath:]
		fileDistPath := rootDistPath + fileUri

		b, _ := ioutil.ReadFile(v) // just pass the file name
		tpl := template.Must(
			template.New("base").Funcs(sprig.FuncMap()).Parse(string(b)),
		)

		// open output file
		dirPath := filepath.Dir(fileDistPath)
		folderDistPath = append(folderDistPath, dirPath)
		os.MkdirAll(dirPath, os.ModePerm)
		fo, err := os.Create(fileDistPath)
		if err != nil {
			panic(err)
		}
		// close fo on exit and check for its returned error
		defer func() {
			if err := fo.Close(); err != nil {
				panic(err)
			}
		}()
		// make a write buffer
		w := bufio.NewWriter(fo)

		vars := map[string]interface{}{}

		err = tpl.Execute(w, vars)
		if err != nil {
			fmt.Printf("Error during template execution: %s", err)
			return
		}

		if err = w.Flush(); err != nil {
			panic(err)
		}

		fmt.Println("----------> " + fileDistPath)
	}
}

func checkConfigurations() {
	fmt.Println("")
	fmt.Println("Processing KrakenD JSON file!!!")
	fmt.Println("----------------------------------------")
	comamnd, _ := ioutil.ReadFile(rootSrcPath + "/command_check.stub")
	command := "DISTDIR=" + rootDistPath + " SRCDIR=" + rootSrcPath + " && " + string(comamnd)
	cmd := exec.Command("sh", "-c", command)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println("")
	fmt.Println("Cleaning up dist folder!!!")
	fmt.Println("----------------------------------------")
	for _, v := range folderDistPath {
		delCmd := exec.Command("sh", "-c", "rm -rf "+v)
		errDel := delCmd.Run()
		if errDel != nil {
			fmt.Println(errDel.Error())
			panic(errDel)
		}
	}

	fmt.Println("All are ready in [" + rootDistPath + "]")
}

// func run() {
// 	fmt.Println("")
// 	fmt.Println("Run KrakenD!!!")
// 	fmt.Println("----------------------------------------")

// 	cmd := exec.Command("./run.sh")
// 	cmd.Dir = rootDistPath

// 	f, err := pty.Start(cmd)
// 	if err != nil {
// 		panic(err)
// 	}

// 	io.Copy(os.Stdout, f)
// }

func run() {
	fmt.Println("")
	fmt.Println("Starting KrakenD!!!")
	fmt.Println("----------------------------------------")

	cmd := exec.Command("./run.sh")
	cmd.Dir = rootDistPath

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	// Execute the command
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	log.Println(stdBuffer.String())
}

func main() {

	args := os.Args
	if len(args) == 2 {
		commandApp = args[1]
		if commandApp != "run" {
			fmt.Println("Invalid param " + commandApp)
			os.Exit(0)
		}
	}

	executeFilePath := filepath.Dir(args[0])
	targetFolder, _ = filepath.Abs(executeFilePath)
	targetFolder = targetFolder + "/../config"
	targetFile = "*"
	rootSrcPath = filepath.Dir(targetFolder) + "/app"
	rootPath, _ = filepath.Abs(rootSrcPath + "/../../")
	rootDistPath = rootPath + "/dist"

	_ = godotenv.Load(rootSrcPath + "/../.env")

	// sanity check
	testFile, err := os.Open(targetFolder)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer testFile.Close()

	testFileInfo, _ := testFile.Stat()
	if !testFileInfo.IsDir() {
		fmt.Println(targetFolder, " is not a directory!")
		os.Exit(-1)
	}

	err = filepath.Walk(targetFolder, findFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	// display our search result

	fmt.Println("\nFound ", len(searchResult), " hits!")
	fmt.Println("----------------------------------------")

	buildConfigurations(searchResult, rootDistPath)

	checkConfigurations()

	if commandApp != "" {
		run()
	}
}
