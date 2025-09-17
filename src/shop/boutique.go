package shop

import (
	"fmt"

	"github.com/ZephaX7/projet-red/tree/main/src/inventory"
)

func Shop() {
	items := []string{
		"Potion de soin gratuite (cadeau de la maison, mais une seule fois !)",
		"Potion de soin",
		"Potion de poison",
		"Livre de Sort : Boule de feu",
		"Fourrure de Loup",
		"Peau de Troll",
		"Cuir de Sanglier",
		"Plume de Corbeaux",
	}
	prices := []int{0, 3, 6, 25, 4, 7, 3, 1} // la première est gratuite
	Gold := 100
	freeTaken := false // pour l'article gratuit

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

		itemName := items[choice-1]
		itemPrice := prices[choice-1]

		// Article gratuit
		if itemPrice == 0 {
			if freeTaken {
				fmt.Println("⚠ Vous ne pouvez prendre cet article gratuit qu'une seule fois !")
				continue
			}
			if len(inventory.Inventaire) >= inventory.CapaciteMax {
				fmt.Println("Votre inventaire est plein, impossible de prendre l'objet gratuit.")
				continue
			}
			inventory.AddInventory(inventory.Objet{Nom: itemName, Quantite: 1, Type: "Objet"})
			freeTaken = true
			fmt.Println("Vous avez pris :", itemName)
			continue
		}

		// Achat normal
		if Gold >= itemPrice {
			if len(inventory.Inventaire) >= inventory.CapaciteMax {
				fmt.Println("Votre inventaire est plein, vous ne pouvez rien acheter de plus.")
				continue
			}
			Gold -= itemPrice
			inventory.AddInventory(inventory.Objet{Nom: itemName, Quantite: 1, Type: "Objet"})
			fmt.Println("Vous avez acheté :", itemName)
		} else {
			fmt.Println("Vous n'avez pas assez d'or.")
		}
	}
}
