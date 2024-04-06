package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
)

const (
	content = "content.md"
	index   = "assets/index.html"
)

func main() {
	fs := http.FileServer(http.Dir("./assets"))

	http.HandleFunc("/", root)
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func root(w http.ResponseWriter, _ *http.Request) {
	md, err := read(content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	html, err := convert(md)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	fmt.Fprint(w, template(html))
}

func convert(b []byte) (string, error) {
	var buf bytes.Buffer

	g := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)

	if err := g.Convert(b, &buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func template(html string) string {
	t, err := read(index)
	if err != nil {
		return ""
	}

	return fmt.Sprintf(string(t), html)
}

func read(f string) ([]byte, error) {
	b, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}

	return b, nil
}
