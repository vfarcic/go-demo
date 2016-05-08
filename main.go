package main
import (
	"net/http"
	"log"
	"io"
)

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

	func HelloServer(w http.ResponseWriter, req *http.Request) {
		log.Printf("Request to %s\n", req.RequestURI)
		io.WriteString(w, "hello, world!\n")
	}