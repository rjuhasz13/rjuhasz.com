# rjuhasz.com

- last update: jan 12 2023. Adding the IP project to the main page.

## Contributing

### Overview

A small Golang program generates the static html pages from templates into the `docs` folder.
This folder is set up to be the root directory for GitHub Pages to host the site.
Assets and documents (stylesheet, images, pdfs) are already in this folder to be
accessable by visitors.

Required software:
- [Golang](https://go.dev/doc/install) for html generation, run `go version` to
  check availability
- [Python3](https://www.python.org/downloads/) for running a simple local server, run `python -V` to check
  availability

### Modifying html files

After modifying a template found in the `pages` folder (e.g. `research.tmpl`),
run `make html` to trigger the Go program that generates the html files into `docs`.
If you're adding new research material, place the pdf into `docs/research`
so you can link to them like `/research/new-pdf-file.pdf`.

### Adding a page

1. add a template to /pages/
2. Change the links in the .tmpl file
3. Add page to index, and to Research/index
4. Add page to main.go
5. Add page to base.tmpl
6. Add a folder to /docs, and add an index.html

### Modifying the CV

Delete the old version from `docs/cv` and add the new file.
Run `make html`, this will find the new file and change the link in the menu.
Keep only one file in `docs/cv`, otherwise the program would fail to find the right
one. Don't reuse the filename, always add e.g. a date suffix, otherwise an old
version might get served for the visitors from cache.

### Previewing changes locally

Run a local server via `make local-server` and visit [http://localhost:8080](http://localhost:8080) in a browser to see the site.
The process runs until terminated (e.g. via `Ctrl+C`). While the server runs,
new versions of the html files can be generated (e.g. running `make html` in a different terminal),
the browser will pick it on a refresh.

### Releasing changes

Simply push to the `main` branch. What's on the `main` branch in the `docs`
folder, that's always the live version of the site.


