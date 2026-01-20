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

var data = &structure.Data{
	Tokens: api.GetTokenList(),
}
var UserFavorites = utils.LoadFavorites()

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
		Address:  "data.Address",
	}

	utils.AddToJSON(toSave)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "success"}`))
}

func Collection(w http.ResponseWriter, r *http.Request) {

	for i := range data.Tokens {
		data.Tokens[i].Id = i + 1
		data.Tokens[i].FormattedMarketCap = utils.FormatLargeNumber(data.Tokens[i].MarketCap)
		data.Tokens[i].FormattedPrice_percentage_24h = fmt.Sprintf("%.2f", data.Tokens[i].Price_change_percentage_24h)
		data.Tokens[i].Type = "layer1"
		if data.Tokens[i].Price_change_percentage_24h > 0 {
			data.Tokens[i].IsPricePercentagePositive = true
		} else {
			data.Tokens[i].IsPricePercentagePositive = false
		}
		data.Tokens[i].IsFavorite = UserFavorites[data.Tokens[i].FullName]
	}
	fmt.Println("collection")
	fmt.Println(data)
	RenderTemplate(w, "collection.html", data)
}

func Ressource(w http.ResponseWriter, r *http.Request) {
	symbol := strings.TrimPrefix(r.URL.Path, "/ressource/")

	fmt.Println("Symbole cliqué :", symbol)

	data := api.GetTokenInfo(symbol)

	data.Supply = utils.FormatLargeNumber(data.MarketData.TotalSupply)
	data.VolumeUSD = utils.FormatLargeNumber(data.Tickers[0].ConvertedVolume.USD)
	data.MarketCap = utils.FormatLargeNumber(data.MarketData.MarketCap.USD)

	RenderTemplate(w, "ressource.html", data)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	filterSup1B := r.URL.Query().Has("sup1b")
	filterInf1B := r.URL.Query().Has("inf1b")
	filterPos24h := r.URL.Query().Has("positive24h")
	filtered := []structure.Token{}
	if !filterSup1B && !filterInf1B && !filterPos24h {
		filtered = data.Tokens
	}

	for _, t := range data.Tokens {
		keep := false
		if filterSup1B && t.MarketCap >= 1000000000 {
			keep = true
		}
		if filterInf1B && t.MarketCap < 1000000000 {
			keep = true
		}

		if filterPos24h {
			if t.Price_change_percentage_24h > 0 {

				if (filterSup1B || filterInf1B) && !keep {
				} else {
					keep = true
				}
			} else {
				keep = false
			}
		}

		if keep {
			filtered = append(filtered, t)
		}

	}
	for i := range filtered {
		filtered[i].IsFavorite = UserFavorites[filtered[i].FullName]
	}

	pageData := structure.Data{
		Tokens: filtered,
	}
	RenderTemplate(w, "collection.html", pageData)
}

var favorites = []string{}

func FavoritesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	token := r.FormValue("token")
	if token == "" {
		http.Error(w, "Token missing", http.StatusBadRequest)
		return
	}

	// Ajouter le token à la liste des favoris
	favorites = append(favorites, token) // ou utiliser saveJSON pour persister

	fmt.Println(favorites)

	// Rediriger vers la page actuelle pour recharger la table
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
}

func AboutUs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	RenderTemplate(w, "aboutus.html", nil)
}

func Profil(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var favoriteTokens []structure.Token

	for _, t := range data.Tokens {
		if UserFavorites[t.FullName] {
			favoriteTokens = append(favoriteTokens, t)
		}
	}

	pageData := structure.Data{
		Tokens: favoriteTokens,
	}

	RenderTemplate(w, "profil.html", pageData)
}

func AddFavorite(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/list", http.StatusSeeOther)
		return
	}

	// Récupère l'ID envoyé par le formulaire
	tokenName := r.FormValue("tokenName")

	if tokenName != "" {
		// Toggle les favoris
		if UserFavorites[tokenName] {
			delete(UserFavorites, tokenName)
		} else {
			UserFavorites[tokenName] = true
		}
	}
	utils.SaveFavorites(UserFavorites)
	// Redirige l'utilisateur vers la même page pour "rafraîchir" l'affichage
	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
}

func FilterResearch(w http.ResponseWriter, r *http.Request) {
	// 1. Récupérer la saisie utilisateur (attribut 'name' du input HTML)
	query := r.FormValue("search")
	var textMessage string
	// 2. Appeler ta fonction de filtrage (assure-toi qu'elle accepte query en paramètre)
	// On suppose que 'data.AllTokens' contient ta liste complète initiale
	filtered := utils.Research(data.Tokens, query)

	if (len(filtered) == 0) && (query != "") {
		textMessage = "Aucun résultat trouvé pour :" + " '' " + query + " '' "
	}

	for i := range filtered {
		filtered[i].IsFavorite = UserFavorites[filtered[i].FullName]
	}
	// 3. Préparer les données pour le template
	output := struct {
		Tokens      []structure.Token
		SearchQuery string
		Message     string
	}{
		Tokens:      filtered,
		SearchQuery: query,
		Message:     textMessage,
	}

	// 4. Exécuter le template de la page collection
	RenderTemplate(w, "research.html", output)
}
