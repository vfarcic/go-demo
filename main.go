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
		log.Println("Request to %s", req.RequestURI)
		io.WriteString(w, "hello, world!\n")
	}