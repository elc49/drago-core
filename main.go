package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/elc49/drago/assets"
	"github.com/elc49/drago/ui/pages/landing"
	"github.com/templui/templui/utils"
)

func main() {
	mux := http.NewServeMux()
	setupAssetsRoutes(mux)
	mux.Handle("GET /", templ.Handler(landing.Landing()))

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatalln(http.ListenAndServe(":8080", mux))
}

func setupAssetsRoutes(mux *http.ServeMux) {
	isDev := true

	// App assets(CSS, fonts, images, ...)
	assetHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isDev {
			w.Header().Set("Cache-Control", "no-store")
		} else {
			w.Header().Set("Cache-Control", "public, max-age=60")
		}

		var fs http.Handler
		if isDev {
			fs = http.FileServer(http.Dir("./assets"))
		} else {
			fs = http.FileServer(http.FS(assets.Assets))
		}

		fs.ServeHTTP(w, r)
	})

	mux.Handle("GET /assets/", http.StripPrefix("/assets/", assetHandler))

	// templUI embedded component scrips
	utils.SetupScriptRoutes(mux, isDev)
}
