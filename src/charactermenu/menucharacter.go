package menucharacter

import (
	"fmt"
	"projet-red/src/inventory"
	statspersonnage "projet-red/src/stats"
)

func Menu_character() {
	fmt.Println("1- Statistiques")
	fmt.Println("2- Inventaire")
	fmt.Println("3- Boutique")
	fmt.Println("4- Retour")

	var Menu int
	fmt.Println("Entrez votre choix :")
	fmt.Scanln(&Menu)

	switch Menu {
	case 1:
		fmt.Println("ouverture des Statistiques...")
		statspersonnage.AfficherStats()
		//appelera la fonction ShowStats
	case 2:
		fmt.Println("ouverture de l'inventaire...")
		inventory.AccessInventory()
		//appele la fonction Inventaire
	case 3:
		fmt.Println("ouverture de la Boutique...")

		//appelera la fonction Boutique
	case 4:
		fmt.Println("De retour pour l'aventure...")

		//appelera la fonction
	}

}

func AfficherMenu() {
	Menu_character()
}
