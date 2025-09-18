package charactermenu

import (
	"fmt"
	combat "projet-red/src/fight"
	"projet-red/src/forgeron"
	"projet-red/src/inventory"
	"projet-red/src/model"
	"projet-red/src/shop"

	statspersonnage "projet-red/src/stats"
)

func Menu_character(perso *model.Personnage) {
	for {
		fmt.Println("1- Statistiques")
		fmt.Println("2- Inventaire")
		fmt.Println("3- Boutique")
		fmt.Println("4- Forge")
		fmt.Println("5- Retour")
		fmt.Println()

		var Menu int
		fmt.Println("Entrez votre choix :")
		fmt.Scanln(&Menu)

		switch Menu {
		case 1:
			fmt.Println("ouverture des Statistiques...")
			fmt.Println()
			statspersonnage.AfficherStats(perso)
		case 2:
			fmt.Println("ouverture de l'inventaire...")
			fmt.Println()
			inventory.AccessInventory()
		case 3:
			fmt.Println("ouverture de la Boutique...")
			fmt.Println()
			shop.Shop()
		case 4:
			fmt.Println("Ouverture de la forge...")
			fmt.Println()
			forgeron.Shop()
		case 5:
			fmt.Println("En avant vers l'aventure !")
			fmt.Println()
			combat.Combat(perso)
			return
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}

func AfficherMenu(p *model.Personnage) {
	Menu_character(p)
}
