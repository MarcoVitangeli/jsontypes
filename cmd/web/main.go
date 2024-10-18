package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MarcoVitangeli/jsontypes/gen"
	"github.com/a-h/templ"
)

func generateGoCode(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Json string `json:"json"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	// Convert the input JSON string into bytes
	inputBytes := []byte(requestData.Json)
	log.Printf("INPUT_BYTES: %d\n", len(inputBytes))

	// TODO: use output buffer instead of reading file
	if err := gen.Gen(inputBytes); err != nil {
		log.Println(err)
		http.Error(w, "Error generating Go code", http.StatusInternalServerError)
		return
	}

	bs, err := os.ReadFile("type_gen.go")
	if err != nil {
		log.Println(err)
		http.Error(w, "Error generating Go code", http.StatusInternalServerError)
		return
	}

	log.Println(string(bs))

	w.Header().Set("Content-Type", "text/plain")
	w.Write(bs)
}

func main() {
	component := App()

	http.Handle("/api/generate", http.HandlerFunc(generateGoCode))
	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
