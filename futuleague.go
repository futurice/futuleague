package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	var port = flag.Int("port", 4000, "Port for listening.")
	var dbPath = flag.String("db", "./data/database.db", "Path to SQLite database, will be seeded if not found.")
	flag.Parse()

	fmt.Println("Welcome to the Futu League backend!")
	initDB(*dbPath)

	r := defineRoutes()
	fmt.Printf("Listening at http://localhost:%d\n", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), r)
}

func defineRoutes() *mux.Router {
	r := mux.NewRouter()
	routeAPIv1(r)
	r.HandleFunc("/", ServerRoot)
	return r
}

func ServerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to FutuLeague.")
}
