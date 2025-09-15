package Shop

import (
	"fmt"
)

func Shop() {
	items := []string{
		"Potion de soin gratuite(C'est pour toi bg,cadeau de la maison. Par contre si tu l'uses pour rien, je vais te retrouver donc fais gaffe à toi)",
		"Potion de soin",
		"Potion de poison",
		"Livre de Sort : Boule de feu",
		"Fourrure de Loup",
		"Peau de Troll",
	}
	prices := []int{0, 3, 6, 25, 4, 7, 3, 1} // la première est gratuite
	gold := 100
	bought := []string{}
	freeTaken := false // pour l'article gratuit

	for {
		fmt.Printf("\nVous avez %d pièces d'or.\n", gold)
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

		// Vérification de l'article gratuit
		if prices[choice-1] == 0 {
			if freeTaken {
				fmt.Println("Vous ne pouvez prendre cet article gratuit qu'une seule fois !")
				continue
			} else {
				bought = append(bought, items[choice-1])
				freeTaken = true
				fmt.Println("Vous avez pris", items[choice-1])
				continue
			}
		}

		// Achat normal
		if gold >= prices[choice-1] {
			gold -= prices[choice-1]
			bought = append(bought, items[choice-1])
			fmt.Println("Vous avez acheté", items[choice-1])
		} else {
			fmt.Println("Vous n'avez pas assez d'or.")
		}
	}

	fmt.Println("Vos articles achetés :", bought)
}
