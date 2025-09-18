package charactermenu

import (
	"fmt"
	combat "projet-red/src/fight"
	"projet-red/src/forgeron"
	"projet-red/src/inventory"
	"projet-red/src/items"
	"projet-red/src/model"
	"projet-red/src/shop"
	"projet-red/src/stats"
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
			stats.AfficherStats(perso)
		case 2:
			inventory.AccessInventory(perso, false, nil, items.UtiliserObjet)
		case 3:
			shop.Shop(perso)
		case 4:
			forgeron.Shop(perso)
		case 5:
			fmt.Println("En avant vers l'aventure !")
			ennemi := model.RandomEnnemi() // ✅ maintenant r est défini
			combat.Combat(perso, ennemi)
			return
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}

func AfficherMenu(p *model.Personnage) {
	Menu_character(p)
}
