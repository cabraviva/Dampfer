//go:build dev

package main

import (
	"net/http"
)

func init() {
	svelteFS = http.FileServer(http.Dir("svelte/dist"))
}
