package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/skratchdot/open-golang/open"
)

var token string

func ServeTts(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(filepath.Join("tmpls", "demo.html"))
	if err != nil {
		panic(err)
	}
	data := make(map[string]string)
	data["Token"] = token
	err = tpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func main() {
	var fn = flag.String("i", "app.json", "baidu app descipt file")
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	data, err := ioutil.ReadFile(*fn)
	if err != nil {
		log.Fatal(err)
	}
	var app AppIDs
	err = json.Unmarshal(data, &app)
	if err != nil {
		log.Fatal(err)
	}
	token, _, err = getToken(app.ApiKey, app.SecretKey)
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("tcp", ":10000")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	http.HandleFunc("/tts", ServeTts)
	fsHandler := http.StripPrefix("/static", http.FileServer(http.Dir("static")))
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".wasm") {
			w.Header().Set("content-type", "application/wasm")
		}
		fsHandler.ServeHTTP(w, r)
	})

	fmt.Println("URL: http://localhost:10000/tts")

	go func() {
		time.Sleep(time.Second * 2)
		open.Run("http://localhost:10000/tts")
	}()

	http.Serve(l, nil)
}
