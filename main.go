package main

import (
	"log"
	"net/http"

	"github.com/markbates/pkger"
	"github.com/webview/webview"
)

func main() {
	// start the webserver
	go runWebserver()

	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Backpacker")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://localhost:3000/")
	w.Run()

}

func runWebserver() {
	if err := openPortAndListen(); err != nil {
		log.Fatal(err)
	}
}

func openPortAndListen() error {
	dir := http.FileServer(pkger.Dir("/public"))
	return http.ListenAndServe(":3000", dir)
}
