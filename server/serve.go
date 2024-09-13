package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() *mux.Router {
	// Create router instance
	router := mux.NewRouter()

	// Static response
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Http Server by ZxFae33!"))
	})

	// Call configurationServer to work with my router instance
	// Can check testfile at server_test.go file
	configurations := configurationServer(router)
	log.Println("Starting server at port 8080")
	if err := configurations.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	return router
}
