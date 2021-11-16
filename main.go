package main
import (
	"fmt"
	"log"
	"net/http"
)

fun handler(w http.ReponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá FUll Cycle Developer")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}