package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var iface = flag.String("interface", "localhost", "http service interface (default: localhost)")
var port = flag.String("port", "8080", "http service port (default: 8080)")
var htmlPort = flag.String("htmlport", "", "http service port for html if different")
var updateInterval = flag.Int("update", 1, "update interval in seconds (default: 2)")
var historySize = flag.Int("history", 300, "history size in samples (default: 300)")
var debugging = flag.Bool("debug", false, "enable debugging output (default: false)")

var dLogger, eLogger *log.Logger

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method nod allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var hostname, err = os.Hostname()
	if err != nil {
		hostname = "Unknown"
	}
	var wsURL = r.Host
	var wsPort = "80"
	if strings.Contains(r.Host, ":") {
		wsURL = strings.Split(r.Host, ":")[0]
		wsPort = strings.Split(r.Host, ":")[1]
	}
	if *htmlPort != "" {
		wsPort = *htmlPort
	}
	wsURL = fmt.Sprintf("%s:%s", wsURL, wsPort)
	data := map[string]string{
		"title":    hostname,
		"hostname": wsURL,
	}
	template.Must(template.ParseGlob("web/*")).ExecuteTemplate(w, "index.html", data)
}

func main() {
	flag.Parse()
	eLogger = log.New(os.Stderr, "ERROR ", log.LstdFlags)
	if *debugging {
		dLogger = log.New(os.Stderr, "DEBUG ", log.LstdFlags)
	} else {
		dLogger = log.New(ioutil.Discard, "DEBUG ", log.LstdFlags)
	}
	monitor.setUpdateInterval(*updateInterval)
	monitor.setHistorySize(*historySize)
	go h.run()
	go monitor.readData()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", *iface, *port), nil)
	if err != nil {
		eLogger.Fatal("ListenAndServe: ", err)
	}
}
