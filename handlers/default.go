package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spmccann/dictionary/components"
	"github.com/spmccann/dictionary/services"
)

var (
	result [][]string
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	page := components.Page(result)
	page.Render(r.Context(), w)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	search := r.FormValue("search")

	if r.Form.Has("search") && search != "" {
		result = services.BinarySearch(search)
	}

	components.SearchResults(result).Render(r.Context(), w)
}

func Run() {
	services.SetupDictionary()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			postHandler(w, r)
			return
		}
		getHandler(w, r)
	})
	fmt.Println("Listening on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Printf("error listening: %v", err)
	}
}
