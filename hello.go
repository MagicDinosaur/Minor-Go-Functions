package main
import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("Hello guys")
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello guys")
}
