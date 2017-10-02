package main

import (
	"fmt"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"./lib"
)

var mux sync.Mutex
var count = 0

func incCount() {
	mux.Lock()
	count++
	mux.Unlock()
}

func main() {
	var port = 3500

	http.HandleFunc("/gif", handleGif)
	http.HandleFunc("/svg", handleSvg)
	http.HandleFunc("/fractal", handleFractal)
	http.HandleFunc("/count", handleCount)
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/github", handleGithub)

	if len(os.Args) > 1 {
		if port2, err := strconv.ParseInt(os.Args[1], 10, 32); err != nil {
			port = int(port2)
		}
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil))
}

func handleCount(w http.ResponseWriter, r *http.Request) {
	incCount()
	mux.Lock()
	fmt.Fprintf(w, "Total requests: %d", count)
	mux.Unlock()
}

func handleGif(w http.ResponseWriter, r *http.Request) {
	incCount()
	lib.Lissajous(w)
}

func handleFractal(w http.ResponseWriter, r *http.Request) {
	incCount()

	wstr := r.URL.Query().Get("w")
	hstr := r.URL.Query().Get("h")
	xstr := r.URL.Query().Get("x")
	ystr := r.URL.Query().Get("y")

	width, err := strconv.ParseInt(string(wstr), 10, 32)
	if err != nil {
		width = 100
	}
	height, err := strconv.ParseInt(string(hstr), 10, 32)
	if err != nil {
		height = 100
	}

	xpos, err := strconv.ParseFloat(string(xstr), 64)
	if err != nil {
		xpos = 0
	}
	ypos, err := strconv.ParseFloat(string(ystr), 64)
	if err != nil {
		ypos = 0
	}

	start := time.Now()
	png.Encode(w, lib.Mandelbrot(int(width), int(height), xpos, ypos))
	fmt.Printf("%.2fs for mandelbrot [%v x %v]\n", time.Since(start).Seconds(), wstr, hstr)
}

func handleSvg(w http.ResponseWriter, r *http.Request) {
	param := ""
	if fun := r.URL.Query()["fun"]; len(fun) > 0 {
		param = fun[0]
	}

	fmt.Fprintf(w, "<html><body><h1>SVG: %s</h1>", param)
	fmt.Fprintf(w, lib.GenSvg(param))
	fmt.Fprintf(w, "</body></html>")
}

func handleGithub(w http.ResponseWriter, r *http.Request) {
	repo := r.FormValue("repo")
	f := r.FormValue("f")

	if len(repo) == 0 {
		fmt.Fprintf(w, "Empty request. Use ?repo=xx&f=yy")
		return
	}

	lib.SearchGithub([]string{repo, f}, w)
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	incCount()

	fmt.Fprintf(w, "Hola\n\n")
	fmt.Println("serving: ", r.URL.Path)
}
