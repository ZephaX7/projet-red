package charactermenu

import (
	"fmt"
	combat "projet-red/src/fight"
	"projet-red/src/inventory"
	"projet-red/src/shop"
	statspersonnage "projet-red/src/stats"
)

func Menu_character() {
	for {
		fmt.Println("1- Statistiques")
		fmt.Println()
		fmt.Println("2- Inventaire")
		fmt.Println()
		fmt.Println("3- Boutique")
		fmt.Println()
		fmt.Println("4- Retour")
		fmt.Println()

		var Menu int
		fmt.Println("Entrez votre choix :")
		fmt.Scanln(&Menu)

		switch Menu {
		case 1:
			fmt.Println("ouverture des Statistiques...")
			fmt.Println()
			statspersonnage.AfficherStats()
		case 2:
			fmt.Println("ouverture de l'inventaire...")
			fmt.Println()
			inventory.AccessInventory()
		case 3:
			fmt.Println("ouverture de la Boutique...")
			fmt.Println()
			shop.Shop()
		case 4:
			fmt.Println("De retour pour l'aventure...")
			fmt.Println()
			combat.Combat()
			return
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
			fmt.Println()
		}
	}
}

func AfficherMenu() {
	Menu_character()
}
