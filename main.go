package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Game struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	MainCharacter *Character `json:"mainCharacter"`
}

type Character struct {
	Name string `json:"name"`
}

var games []Game

func getAllGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(games)
}

func getGame(w http.ResponseWriter, r *http.Request) {

}

func createGame(w http.ResponseWriter, r *http.Request) {

}

func deleteGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range games {
		if item.ID == params["id"] {
			games = append(games[:index], games[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode("Game deleted successfully")
}

func updateGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range games {
		if item.ID == params["id"] {
			games = append(games[:index], games[index+1:]...)

			var game Game
			_ = json.NewDecoder(r.Body).Decode(&game)
			game.ID = item.ID

			games = append(games, game)
			break
		}
	}

	json.NewEncoder(w).Encode("Game updated successfully")
}

func main() {
	r := mux.NewRouter()

	games = append(games, Game{ID: "1", Name: "God of War", MainCharacter: &Character{Name: "Kratos"}})
	games = append(games, Game{ID: "2", Name: "The Last of Us", MainCharacter: &Character{Name: "Joel"}})

	r.HandleFunc("/games", getAllGames).Methods("GET")
	r.HandleFunc("/games/{id}", getGame).Methods("GET")
	r.HandleFunc("/games", createGame).Methods("POST")
	r.HandleFunc("/games/{id}", deleteGame).Methods("DELETE")
	r.HandleFunc("/games/{id}", updateGame).Methods("PUT")

	log.Println("Server started on: http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
