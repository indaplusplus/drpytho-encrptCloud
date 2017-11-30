package main

import (
	"github.com/drpytho/encrptCloud/encrpt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var KEY = [...]byte{4, 46, 53, 180, 151, 173, 194, 191, 100, 17, 107, 253, 230, 35, 19, 155, 161, 229, 146, 117, 11, 38, 221, 194, 234, 157, 204, 210, 26, 247, 37, 190}
var store = make(map[string][]byte)

var singleStore, _ = encrpt.NewRamStoreUnit(14, KEY[:])

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Some text"))
}

func InitializeFS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Hi Low"))
}

func setBlockAt(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("request", err)
	}
	log.Print(len(buf), string(buf))
	block, _ := strconv.Atoi(r.URL.Query().Get("block"))
	err = singleStore.SetBlock(block, buf)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("OK"))
}

func getBlockAt(w http.ResponseWriter, r *http.Request) {
	block, _ := strconv.Atoi(r.URL.Query().Get("block"))
	buf, _, err := singleStore.GetBlock(block)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
	w.Write(buf)
}

func main() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/init", InitializeFS)
	http.HandleFunc("/get", getBlockAt)
	http.HandleFunc("/set", setBlockAt)
	err := http.ListenAndServeTLS(":4433", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
