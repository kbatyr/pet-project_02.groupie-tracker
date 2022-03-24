package main

import (
	"fmt"
	"log"
	"net/http"

	grpt "groupie/src"
)

func main() {
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./style/"))))
	http.Handle("/map/", http.StripPrefix("/map/", http.FileServer(http.Dir("./map/"))))
	http.HandleFunc("/", grpt.IndexHandler)
	http.HandleFunc("/artists/", grpt.ArtistInfoHandler)
	http.HandleFunc("/search/", grpt.SearchHandler)
	http.HandleFunc("/filter/", grpt.FilterHandler)
	fmt.Println("Server is listening at :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
