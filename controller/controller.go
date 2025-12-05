package controller

import (
	"Les-Crackhead_groupie_tracker/api"
	"Les-Crackhead_groupie_tracker/structure"
	"Les-Crackhead_groupie_tracker/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var Token string = "CG-5oBeGf9b4qSv7c4ENCCz4rw8"
var Urlapi string = "https://api.coingecko.com/api/v3/"

// CETTE FONCTION REND UN TEMPLATE AVEC DES DONNEES ET L'ECRIT DANS LA REPONSE HTTP
func RenderTemplate(w http.ResponseWriter, filename string, data interface{}) {
	template := template.Must(template.ParseFiles("template/" + filename))

	buf := new(bytes.Buffer)
	if err := template.Execute(buf, data); err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
		return
	}
	w.Write(buf.Bytes())
}

// LA FONCTION GERE L'AFFICHAGE DE HOME
func Home(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	RenderTemplate(w, "home.html", nil)
}

func FetchData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	var data structure.DataReceived

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("Erreur de décodage:", err)
		http.Error(w, "JSON invalide", http.StatusBadRequest)
		return
	}

	toSave := structure.UserData{
		LiveUser: data.Address,
		Address:  data.Address,
	}

	utils.SaveJson(toSave)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "success"}`))
}

func GetData() {}

func Collection(w http.ResponseWriter, r *http.Request) {

	data := api.GetTokenList()

	for i := range data {
		data[i].Id = i + 1
		data[i].FormattedMarketCap = utils.FormatLargeNumber(data[i].MarketCap)
		data[i].FormattedPrice_percentage_24h = fmt.Sprintf("%.2f", data[i].Price_change_percentage_24h)
		if data[i].Price_change_percentage_24h > 0 {
			data[i].IsPricePercentagePositive = true
		} else {
			data[i].IsPricePercentagePositive = false
		}
	}

	RenderTemplate(w, "collection.html", data)
}

func Ressource(w http.ResponseWriter, r *http.Request) {
	symbol := strings.TrimPrefix(r.URL.Path, "/ressource/")

	fmt.Println("Symbole cliqué :", symbol)

	RenderTemplate(w, "ressource.html", nil)
}
