package main

import (
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var root = flag.String("root", "./cache_assets", "file system path")

func main() {
	h := sha1.New()

	http.Handle("/", http.FileServer(http.Dir(*root)))

	count := 0
	value1, value2 := "hello world1", "hello world2"
	http.HandleFunc("/age-test", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("come to server")

		count++
		writer.Header().Add("Cache-Control", "max-age=3")

		var value string
		if count == 300 {
			value = value2
		} else {
			value = value1
		}
		fmt.Println(value)
		h.Write([]byte(value))
		hashValue := hex.EncodeToString(h.Sum(nil))
		h.Reset()
		writer.Header().Add("ETag", hashValue)

		log.Println("req header: ", request.Header.Get("If-None-Match"), " hash: ", hashValue)
		if request.Header.Get("If-None-Match") == hashValue {
			writer.WriteHeader(http.StatusNotModified)
			return
		}

		writer.Write([]byte(value)) // 제일 마지막에 써줘야
	})

	log.Println("Listening on 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
