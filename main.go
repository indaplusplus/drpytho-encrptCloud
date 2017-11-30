package main

import (
	"fmt"
	"github.com/drpytho/encrptCloud/encrpt"
	"net/http"
)

var KEY = [...]byte{4, 46, 53, 180, 151, 173, 194, 191, 100, 17, 107, 253, 230, 35, 19, 155, 161, 229, 146, 117, 11, 38, 221, 194, 234, 157, 204, 210, 26, 247, 37, 190}
var store = make(map[string][]byte)

var singleStore = encrpt.NewRamStoreUnit(14, KEY[:])

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Some text"))
}

func InitializeFS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	size := int(r.URL.Get("size"))
	// Generate a string ID
	x := 14
	ID := "superRandomString"
	size := 4 + 1<<(x-4) - 32 + 1<<x
	store[ID] = make([]byte, size)
	w.Write([]byte(ID))
}

func setBlockAt(w http.ResponceWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/init", InitializeFS)
	err := http.ListenAndServeTLS(":4433", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
