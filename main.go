package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() error {
	if err := createPages(); err != nil {
		return err
	}

	return nil
}

func createPages() error {
	homePage := Page{
		tmpl: template.Must(template.ParseFiles("pages/base.tmpl", "pages/home.tmpl")),
		slug: "",
	}

	researchPage := Page{
		tmpl: template.Must(template.ParseFiles("pages/base.tmpl", "pages/research.tmpl")),
		slug: "research",
	}

	pages := []Page{homePage, researchPage}

	for _, page := range pages {
		if err := page.create(); err != nil {
			return err
		}
	}

	return nil
}

type Page struct {
	tmpl *template.Template
	slug string
}

func (p Page) create() error {
	cv, err := findCV()
	if err != nil {
		return err
	}

	var processed bytes.Buffer
	p.tmpl.ExecuteTemplate(&processed, "base", Data{Slug: p.slug, CV: "/cv/" + cv})

	outputPath := "./docs/index.html"
	if p.slug != "" {
		outputPath = "./docs/" + p.slug + "/index.html"
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	w.WriteString(string(processed.Bytes()))
	w.Flush()
	if err != nil {
		return err
	}

	return nil
}

func findCV() (string, error) {
	files, err := os.ReadDir("./docs/cv")
	if err != nil {
		return "", err
	}
	// get only the ones with .pdf suffix
	var cvFiles []fs.DirEntry
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if file.Name()[len(file.Name())-4:] == ".pdf" {
			cvFiles = append(cvFiles, file)
		}
	}

	if len(cvFiles) == 0 {
		return "", errors.New("no pdf files in 'docs/cv', please place one here")
	}
	if len(cvFiles) > 1 {
		return "", errors.New("there are multiple pdf files in 'docs/cv', please place only one here")
	}

	return cvFiles[0].Name(), nil
}

type Data struct {
	CV   string
	Slug string
}
