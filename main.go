package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

// Embed files in the svelte/public directory
//
//go:embed svelte/dist/*
var embeddedFiles embed.FS

func main() {
	// Serve embedded files as an HTTP file system
	svelteFiles, _ := fs.Sub(embeddedFiles, "svelte/dist")
	fs := http.FileServer(http.FS(svelteFiles))
	http.Handle("/", fs)

	fmt.Println("Server starting on port 13777...")
	log.Fatal(http.ListenAndServe(":13777", nil))
}
