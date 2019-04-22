package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type API struct {
	Message string "json:message"
}

type Pokemon struct {
	ID    int    "json:id"
	Name  string "json:name"
	Skill string "json:skill"
}

func Hello(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Merhaba %s", r.URL.Path[1:])
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	messageConcat := "Kullanıcı ID " + id
	message := API{messageConcat}
	output, err := json.Marshal(message)
	checkErro(err)
	fmt.Fprintf(w, string(output))
}

func main() {

	//http.HandleFunc("/", handler)
	//http.ListenAndServe(":9000", nil)
	//fmt.Println("WebServer")

	// ... / api

	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/pokemon/{id}", Hello)
	http.Handle("/", gorillaRoute)

	/*
		http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
			message := API{"API Home"}
			output, err := json.Marshal(message)
			checkErro(err)

			fmt.Fprintf(w, string(output))
		})

		// ... / users

		http.HandleFunc(apiRoot+"/pokemons", func(w http.ResponseWriter, r *http.Request) {
			pokemons := []Pokemon{
				Pokemon{ID: 1, Name: "Pikachu", Skill: "Electric"},
				Pokemon{ID: 2, Name: "Onix", Skill: "Stone"},
				Pokemon{ID: 3, Name: "Saydek", Skill: "Duck"},
			}

			message := pokemons
			output, err := json.Marshal(message)
			checkErro(err)
			fmt.Fprintf(w, string(output))
		})
	*/

	http.ListenAndServe(":8080", nil)
}

func checkErro(err error) {
	if err != nil {
		fmt.Println("Fatal Error", err.Error())
		os.Exit(1)
	}
}
