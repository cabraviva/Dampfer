//go:build !dev

package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed svelte/dist/*
var embeddedFiles embed.FS

func init() {
	svelteFiles, _ := fs.Sub(embeddedFiles, "svelte/dist")
	svelteFS = http.FileServer(http.FS(svelteFiles))
}
