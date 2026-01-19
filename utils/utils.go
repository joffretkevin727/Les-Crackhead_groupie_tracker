package utils

import (
	"Les-Crackhead_groupie_tracker/structure"
	"encoding/json"
	"fmt"
	"log"
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

func SaveJson(data structure.UserData) {
	existing := make(map[string]structure.UserData)

	// Lire le fichier existant
	jsonBytes, err := os.ReadFile("coins.json")
	if err == nil {
		err = json.Unmarshal(jsonBytes, &existing)
		if err != nil {
			log.Fatal(err)
		}
	} else if !os.IsNotExist(err) {
		log.Fatal(err)
	}

	// Ajouter / mettre à jour la donnée
	existing[data.Address] = data

	// Réécrire le JSON
	newJsonBytes, err := json.MarshalIndent(existing, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("coins.json", newJsonBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ajout ou mise à jour dans coins.json")
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
