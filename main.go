package main

import (
	"fmt"
	"hash/fnv"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type ResearchPageData struct {
	PageTitle     string  `yaml:"page_title"`
	Publications  []Paper `yaml:"publications"`
	WorkingPapers []Paper `yaml:"working_papers"`
}

type Paper struct {
	Title         string   `yaml:"title"`
	URL           string   `yaml:"url"`
	Coauthors     []Author `yaml:"coauthors"`
	MediaCoverage []Medium `yaml:"media_coverage"`
	Abstract      string   `yaml:"abstract"`
}

func (p Paper) ID() uint32 {
	h := fnv.New32a()
	h.Write([]byte(p.Title))
	return h.Sum32()
}

func (p Paper) Name() string {
	return p.Title
}

func (p Paper) HasCoauthors() bool {
	return len(p.Coauthors) > 0
}

type Author struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

type Medium struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

func main() {
	yamlFile, err := ioutil.ReadFile("data/research.yml")
	if err != nil {
		panic(err)
	}

	var data ResearchPageData
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
	tmpl, err := template.New("research.tmpl").Funcs(template.FuncMap{
		"plus1": func(x int) int {
			return x + 1
		},
	}).ParseFiles("html/research.tmpl")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":3000", nil)
}
