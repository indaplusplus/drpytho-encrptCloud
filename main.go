package main

import (
	"fmt"
	"github.com/drpytho/encrptCloud/encrpt"
	"net/http"
)

var KEY = [...]byte{4, 46, 53, 180, 151, 173, 194, 191, 100, 17, 107, 253, 230, 35, 19, 155, 161, 229, 146, 117, 11, 38, 221, 194, 234, 157, 204, 210, 26, 247, 37, 190}

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Some text"))
}

func main() {
	buf := make([]byte, 1024)
	dat := []byte("This is some data I want to encrypt")
	for i, val := range dat {
		buf[i] = val
	}
	block, err := EncryptBlock(buf, KEY[:], []byte("block1"))
	if err != nil {
		fmt.Println(err)
	}
	data, err := DecryptBlock(block, KEY[:], []byte("block1"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(block)
	fmt.Println(string(data))

}
