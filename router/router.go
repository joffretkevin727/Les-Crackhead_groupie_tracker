package router

import (
	"Les-Crackhead_groupie_tracker/controller"
	"net/http"
)

// CETTE FONCTION INITIALISE UN SERVEUR MUX, CONFIGURE LES ROUTES ET LES FICHIERS STATIQUES ET LE RETOURNE
func New() *http.ServeMux {
	mux := http.NewServeMux()

	//------------------- ROUTES -----------------------
	mux.HandleFunc("/", controller.Home)
	mux.HandleFunc("/filter", controller.HomeHandler)
	mux.HandleFunc("/home", controller.Home)
	mux.HandleFunc("/list", controller.Collection)
	mux.HandleFunc("/ressource/", controller.Ressource)
	mux.HandleFunc("/api/save-wallet", controller.FetchData)
	mux.HandleFunc("/aboutus", controller.AboutUs)
	mux.HandleFunc("/profil", controller.Profil)
	mux.HandleFunc("/add-favorite", controller.AddFavorite)

	//--------------------------------------------------

	// ---------------- STATIC FILES -------------------
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	//--------------------------------------------------
	return mux
}
