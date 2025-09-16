package Forgeron

import (
	"fmt"

	"github.com/ZephaX7/projet-red/src/Inventory"
)

func Shop() {
	items := []string{
		"Chapeau de l'aventurier",
		"Tunique de l'aventurier",
		"Bottes de l'aventurier",
	}
	prices := []int{5, 5, 5}
	Gold := 100
	bought := []string{}

	for {
		fmt.Printf("\nVous avez %d pièces d'or.\n", Gold)
		for i, item := range items {
			fmt.Printf("%d. %s - %d pièces d'or\n", i+1, item, prices[i])
		}
		fmt.Println("Entrez le numéro de l'article à acheter (ou 0 pour quitter) :")

		var choice int
		fmt.Scan(&choice)

		if choice == 0 {
			break
		}
		if choice < 1 || choice > len(items) {
			fmt.Println("Choix invalide.")
			continue
		}
		// Achat normal
		if Gold >= prices[choice-1] {
			// Vérifier si l'inventaire est plein
			if len(Inventory.Inventaire) >= Inventory.CapaciteMax {
				fmt.Println("Votre inventaire est plein, vous ne pouvez rien acheter de plus.")
				continue // retourne au début de la boucle pour choisir autre chose
			}

			Gold -= prices[choice-1]
			bought = append(bought, items[choice-1])
			fmt.Println("Vous avez acheté", items[choice-1])
		} else {
			fmt.Println("Vous n'avez pas assez d'or.")
		}
	}
}
