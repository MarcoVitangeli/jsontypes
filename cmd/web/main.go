package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MarcoVitangeli/jsontypes/gen"
	"github.com/a-h/templ"
)

func generateGoCode(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	jsonInput := r.FormValue("json")
	if jsonInput == "" {
		http.Error(w, "Missing 'json' input", http.StatusBadRequest)
		return
	}

	inputBytes := []byte(jsonInput)

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

	w.Header().Set("Content-Type", "text/plain")
	w.Write(bs)
}

func main() {
	component := App()

	http.Handle("/api/generate", http.HandlerFunc(generateGoCode))
	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :80")
	http.ListenAndServe(":3000", nil)
}
