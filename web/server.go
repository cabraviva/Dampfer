package web

import (
	"Dampfer/utils"
	"net/http"
)

func InitServer(svelteFS http.Handler) {
	// Serve embedded files as an HTTP file system
	http.Handle("/", svelteFS)
}

func StartServer() {
	utils.Log.Info("Server started on port 13777")
	utils.Log.Fatal(http.ListenAndServe(":13777", nil))
}
