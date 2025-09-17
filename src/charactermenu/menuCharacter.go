package menu

import (
	"fmt"
)
func Menu_character() {

func Menu_character(c *inventaire) {
	var Menu int
	fmt.Println("1- Statistiques")
	fmt.Println("2- Inventaire")
	fmt.Println("3- Boutique")
	fmt.Println("4- Retour")
	switch Menu {
	case 1:
		fmt.Println("ouverture des Statistiques...")

		//appelera la fonction ShowStats
	case 2:
		fmt.Println("ouverture de l'inventaire...")

		//appelera la fonction Inventaire
	case 3:
		fmt.Println("ouverture de la Boutique...")

		//appelera la fonction Boutique
	case 4:
		fmt.Println("De retour pour l'aventure...")

		//appelera la fonction
	}

}

func main() {
	Menu_character()
}
}
