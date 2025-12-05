package utils

import (
	"Les-Crackhead_groupie_tracker/structure"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
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
