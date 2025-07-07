// Server1 是一个最小的"回声"服务器
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // 每个请求都调用handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler 回应请求的URL r.URL.Path
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
} 