package utils

import (
	"Les-Crackhead_groupie_tracker/structure"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
)

func Sort(list []structure.Token) []structure.Token {
	newList := []structure.Token{}

	for i := range list {
		if list[i].CurrentPrice > 100 {
			newList = append(newList, list[i])
		}
	}

	return newList
}

func FormatLargeNumber(n float64) string {
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

func saveJSON(fileName string, data interface{}) error { // Convertir les données en JSON

	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, bytes, 0644) // Écrire dans le fichier
}

func AddToJSON(toAdd structure.UserData) {
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

	err = saveJSON("userConnexion", historic)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture JSON :", err)
		return
	}

	fmt.Println("JSON mis à jour avec succès !")
}

func FormatLargeNumberInt(s string) int {
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

func LoadFavorites() map[string]bool {
	file, err := os.ReadFile("favorites.json")
	if err != nil {
		return make(map[string]bool) // Retourne une map vide si le fichier n'existe pas
	}
	var favs map[string]bool
	json.Unmarshal(file, &favs)
	return favs
}

func SaveFavorites(favorites map[string]bool) {
	data, _ := json.Marshal(favorites)
	os.WriteFile("favorites.json", data, 0644) // Écrit la map sur le disque
}
