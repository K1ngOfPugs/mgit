package main

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Entry struct {
	Name   string `json:"name"`
	Folder string `json:"folder"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(color.BlueString("Input repo url: "))
	line, _ := reader.ReadString('\n')

	site := strings.TrimSpace(line)

	repo_url := site + "/mgit.json"

	response := string(request(repo_url))

	fmt.Print(color.GreenString("Available repos:\n\n"))

	var data []Entry
	_ = json.Unmarshal([]byte(response), &data)

	for i, object := range data {
		fmt.Printf("%d: ", i+1)
		fmt.Print(color.CyanString(object.Name) + "\n")
	}

	fmt.Print(color.YellowString("\nEnter the number of the repo you would like to clone: "))

	line, _ = reader.ReadString('\n')

	selected, _ := strconv.Atoi(strings.TrimSpace(line))

	selected = selected - 1

	if selected >= 0 && selected < len(data) {
		printData(data, selected)
		download(data, selected, site)

	} else {
		fmt.Print(color.RedString("Error: Input did not match an option."))
		os.Exit(0)
	}

}

func request(url string) []byte {
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func printData(data []Entry, selected int) {
	fmt.Print("\nSelected:\n", color.CyanString("Name: "+data[selected].Name+"\nFolder: "+data[selected].Folder+"\n\n"))
}

func download(data []Entry, selected int, site string) {
	zipfile := data[selected].Folder + ".zip"
	//Download the zip file
	fmt.Print("Downloading...\n")
	out, err := os.Create(zipfile)
	if err != nil {
		log.Fatal(err)
	}

	url := site + "/" + data[selected].Folder + "/dist/" + zipfile

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	out.Close()
	resp.Body.Close()

	//Extract the zip file
	r, err := zip.OpenReader(zipfile)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer rc.Close()

		newFilePath := f.Name

		if f.FileInfo().IsDir() {
			err = os.MkdirAll(newFilePath, 0777)
			if err != nil {
				log.Fatal(err)
			}
			continue
		}

		uncompressedFile, err := os.Create(newFilePath)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(uncompressedFile, rc)
		if err != nil {
			log.Fatal(err)
		}
	}

	//Clean
	r.Close()
	err = os.Remove(zipfile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(color.GreenString("Done!"))
}
