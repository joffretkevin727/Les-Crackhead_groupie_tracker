package utils

import (
	"Les-Crackhead_groupie_tracker/structure"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Sort(list []structure.Token) []structure.Token { // Filtre les tokens dont le prix est strictement supérieur à 100 USD.
	newList := []structure.Token{}

	for i := range list {
		if list[i].CurrentPrice > 100 {
			newList = append(newList, list[i])
		}
	}

	return newList
}

func FormatLargeNumber(n float64) string { // Convertit les grands nombres en chaînes abrégées (T, B, M, K) pour la lisibilité financière.
	abs := math.Abs(n)

	switch {
	case abs >= 1_000_000_000_000:
		return fmt.Sprintf("%.2fT", n/1_000_000_000_000)
	case abs >= 1_000_000_000:
		return fmt.Sprintf("%.2fB", n/1_000_000_000)
	case abs >= 1_000_000:
		return fmt.Sprintf("%.2fM", n/1_000_000)
	case abs >= 1_000:
		return fmt.Sprintf("%.2fK", n/1_000)
	default:
		return fmt.Sprintf("%.0f", n)
	}
}

func saveJSON(fileName string, data interface{}) error { // Encode n'importe quelle structure en JSON indenté et l'écrit physiquement sur le disque.

	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, bytes, 0644) // Écrire dans le fichier
}

func AddToJSON(toAdd structure.UserData) { // Gère l'historique des connexions en ajoutant les nouvelles données utilisateur au fichier JSON existant.
	file, err := os.ReadFile("userConnexion.json")
	if err != nil {
		if os.IsNotExist(err) {
			file = []byte("[]")
		} else {
			fmt.Println("Erreur de lecture fichier :", err)
			return
		}
	}

	if len(file) == 0 {
		file = []byte("[]")
	}

	var historic []structure.UserData
	err = json.Unmarshal(file, &historic)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON :", err)
		return
	}

	historic = append(historic, toAdd)

	err = saveJSON("userConnexion.json", historic)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture JSON :", err)
		return
	}

	fmt.Println("JSON mis à jour")
}

func FormatLargeNumberInt(s string) int { // Parse une chaîne numérique et retourne sa valeur abrégée sous forme d'entier.
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}

	abs := math.Abs(n)

	switch {
	case abs >= 1_000_000_000_000: // T
		return int(n / 1_000_000_000_000)
	case abs >= 1_000_000_000: // B
		return int(n / 1_000_000_000)
	case abs >= 1_000_000: // M
		return int(n / 1_000_000)
	case abs >= 1_000: // K
		return int(n / 1_000)
	default:
		return int(n)
	}
}

func LoadFavorites() map[string]bool { // Charge la map des favoris depuis le disque ou initialise une map vide en cas d'absence de fichier.
	file, err := os.ReadFile("favorites.json")
	if err != nil {
		return make(map[string]bool) // Retourne une map vide si le fichier n'existe pas
	}
	var favs map[string]bool
	json.Unmarshal(file, &favs)
	return favs
}

func SaveFavorites(favorites map[string]bool) { // Sérialise et sauvegarde la map des favoris en format JSON compact.
	data, _ := json.Marshal(favorites)
	os.WriteFile("favorites.json", data, 0644) // Écrit la map sur le disque
}

func SyncFavorites(tokens []structure.Token, userFavs map[string]bool) { // Parcourt une liste de tokens pour injecter l'état booléen IsFavorite basé sur la map utilisateur.
	for i := range tokens {
		tokens[i].IsFavorite = userFavs[tokens[i].FullName]
	}
}

func Research(allTokens []structure.Token, query string) []structure.Token { // Filtre les tokens par nom ou symbole en priorisant les résultats commençant par la requête.
	query = strings.ToLower(strings.TrimSpace(query))
	if query == "" {
		return allTokens
	}

	var startsWith []structure.Token
	var containsOnly []structure.Token

	for _, token := range allTokens {
		name := strings.ToLower(token.FullName)
		symbol := strings.ToLower(token.Symbol)

		if strings.HasPrefix(name, query) || strings.HasPrefix(symbol, query) {
			// Priorité 1 : Commence par la recherche
			startsWith = append(startsWith, token)
		} else if strings.Contains(name, query) || strings.Contains(symbol, query) {
			// Priorité 2 : Contient la recherche (mais ne commence pas par elle)
			containsOnly = append(containsOnly, token)
		}
	}

	// Fusionne les deux listes : les "commence par" en premier
	return append(startsWith, containsOnly...)
}
