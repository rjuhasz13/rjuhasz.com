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

type Data struct {
	CV            string  `yaml:"cv"`
	Publications  []Paper `yaml:"publications"`
	WorkingPapers []Paper `yaml:"working_papers"`
	Slug          string
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
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/research", researchHandler)
	http.ListenAndServe("localhost:"+os.Getenv("PORT"), nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl := template.Must(template.ParseFiles("html/base.tmpl", "html/home.tmpl", "html/style.tmpl"))

	data := data()
	data.Slug = ""
	err := tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		panic(err)
	}
}

func researchHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/research" {
		http.NotFound(w, r)
		return
	}

	data := data()
	data.Slug = "research"
	tmpl, err := template.New("base.tmpl").Funcs(template.FuncMap{
		"plus1": func(x int) int {
			return x + 1
		},
	}).ParseFiles("html/base.tmpl", "html/research.tmpl", "html/style.tmpl")
	if err != nil {
		panic(err)
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		panic(err)
	}
}

func data() Data {
	yamlFile, err := ioutil.ReadFile("data/research.yml")
	if err != nil {
		panic(err)
	}

	var data Data
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
	return data
}
