package main

import (
	"net/http"
	"os"
	"os/exec"
)

func main() {
	// Serve static assets from /web
	http.Handle("/", http.FileServer(http.Dir("web")))

	// Start the gopls server
	// https://github.com/golang/tools/blob/master/gopls/doc/daemon.md
	go startGopls()

	// Start the http server
	http.ListenAndServe(":8080", nil)
}

func startGopls() {
	// Start the gopls server and fwd logs to stdout
	cmd := exec.Command("gopls", "-listen=:37374", "-logfile=auto", "-debug=:0")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}
