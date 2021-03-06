package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/realtimdunbar/go_game/models"
)

type server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// New will create a new server struct for the API configuration
func New(dialect, conn string) (server, error) {
	db, err := gorm.Open(dialect, conn)

	s := server{
		Router: mux.NewRouter(),
	}

	if err != nil {
		return s, err
	}

	s.DB = db

	return s, nil
}

func (s *server) routes() {
	s.Router.HandleFunc("/players", s.IndexPlayers).Methods("GET")
	s.Router.HandleFunc("/players/{id}", s.ShowPlayer).Methods("GET")
	s.Router.HandleFunc("/players", s.CreatePlayer).Methods("POST")
	s.Router.HandleFunc("/players/{id}", s.DeletePlayer).Methods("DELETE")
}

func (s *server) IndexPlayers(w http.ResponseWriter, r *http.Request) {
	var player []models.Player
	s.DB.Find(&player)
	json.NewEncoder(w).Encode(&player)
}

func (s *server) ShowPlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var player models.Player
	s.DB.First(&player, params["id"])
	json.NewEncoder(w).Encode(&player)
}

func (s *server) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var player models.Player
	json.NewDecoder(r.Body).Decode(&player)
	s.DB.Create(&player)
	json.NewEncoder(w).Encode(&player)
}

func (s *server) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var player models.Player
	s.DB.First(&player, params["id"])
	s.DB.Delete(&player)

	var players []models.Player
	s.DB.Find(&players)
	json.NewEncoder(w).Encode(&players)
}

func (s *server) IndexGames(w http.ResponseWriter, r *http.Request) {
	var games []models.Game
	s.DB.Find(&games)
	json.NewEncoder(w).Encode(&games)
}

func (s *server) ShowGame(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var game models.Game
	s.DB.First(&game, params["id"])
	json.NewEncoder(w).Encode(&game)
}

func (s *server) CreateGame(w http.ResponseWriter, r *http.Request) {
	var game models.Game
	json.NewDecoder(r.Body).Decode(&game)
	s.DB.Create(&game)
	json.NewEncoder(w).Encode(&game)
}

func (s *server) DeleteGame(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var game models.Game
	s.DB.First(&game, params["id"])
	s.DB.Delete(&game)

	var games []models.Game
	s.DB.Find(&games)
	json.NewEncoder(w).Encode(&games)
}

func (s *server) IndexStones(w http.ResponseWriter, r *http.Request) {
	var stones []models.Stone
	s.DB.Find(&stones)
	json.NewEncoder(w).Encode(&stones)
}

func (s *server) ShowStone(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var stone models.Stone
	s.DB.First(&stone, params["id"])
	json.NewEncoder(w).Encode(&stone)
}

func (s *server) CreateStone(w http.ResponseWriter, r *http.Request) {
	var stone models.Stone
	json.NewDecoder(r.Body).Decode(&stone)
	s.DB.Create(&stone)
	json.NewEncoder(w).Encode(&stone)
}

func (s *server) DeleteStone(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var stone models.Stone
	s.DB.First(&stone, params["id"])
	s.DB.Delete(&stone)

	var stones []models.Stone
	s.DB.Find(&stones)
	json.NewEncoder(w).Encode(&stones)
}
