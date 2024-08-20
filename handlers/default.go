package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/spmccann/clickable-dictionary/components"
	"github.com/spmccann/clickable-dictionary/services"
)

func Run() {
	services.SetupDictionary()
	page := components.Page()
	result := services.BinarySearch("bacon")
	definitions := components.Definitions(result)

	http.Handle("/", templ.Handler(page))
	http.Handle("/search", templ.Handler(definitions))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
