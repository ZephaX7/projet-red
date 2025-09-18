package shop

import (
	"fmt"
	"projet-red/src/inventory"
	"projet-red/src/model"
)

func Shop(perso *model.Personnage) {
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
	prices := []int{0, 3, 6, 25, 4, 7, 3, 1}
	freeTaken := false // pour l'article gratuit

	for {
		fmt.Printf("\nVous avez %d pièces d'or.\n", perso.Gold)
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
			fmt.Println("Choix invalide.\n")
			continue
		}

		itemName := items[choice-1]
		itemPrice := prices[choice-1]

		// Article gratuit
		if itemPrice == 0 {
			if freeTaken {
				fmt.Println("⚠ Vous ne pouvez prendre cet article gratuit qu'une seule fois !\n")
				continue
			}
			if len(inventory.Inventaire) >= inventory.CapaciteMax {
				fmt.Println("Votre inventaire est plein, impossible de prendre l'objet gratuit.\n")
				continue
			}
			inventory.AddInventory(inventory.Objet{Nom: itemName, Quantite: 1, Type: "Objet"})
			freeTaken = true
			fmt.Println("Vous avez pris :", itemName, "\n")
			continue
		}

		// Achat normal
		if perso.Gold >= itemPrice {
			if len(inventory.Inventaire) >= inventory.CapaciteMax {
				fmt.Println("Votre inventaire est plein, vous ne pouvez rien acheter de plus.\n")
				continue
			}
			perso.Gold -= itemPrice
			inventory.AddInventory(inventory.Objet{Nom: itemName, Quantite: 1, Type: "Objet"})
			fmt.Println("Vous avez acheté :", itemName, "\n")
		} else {
			fmt.Println("Vous n'avez pas assez d'or.\n")
		}
	}
}
