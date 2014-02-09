package main

import (
	"flag"
	"os"
	"log"
	"net/http"
	"html/template"
)

var addr = flag.String("addr", ":8080", "http service address (default: localhost:8080)")
var updateInterval = flag.Int("update", 1, "update interval in seconds (default: 2)")
var historySize = flag.Int("history", 300, "history size in samples (default: 300)")

var dLogger = log.New(os.Stdout, "DEBUG ", log.LstdFlags)
var eLogger = log.New(os.Stderr, "ERROR ", log.LstdFlags)

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
	if (err != nil) {
		hostname = "Unknown"
	}
	data := map[string]string {
		"title": hostname,
		"hostname": r.Host,
	}
	template.Must(template.ParseGlob("web/*")).ExecuteTemplate(w, "index.html", data)
}

func main() {
	flag.Parse()
	monitor.setUpdateInterval(*updateInterval)
	monitor.setHistorySize(*historySize)
	go h.run()
	go monitor.readData()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		eLogger.Fatal("ListenAndServe: ", err)
	}
}
