package hub

import (
	"fmt"
	"projet-red/src/charactermenu"
	combat "projet-red/src/fight"
	"projet-red/src/model"
)

func Hub(perso *model.Personnage) {
	fmt.Println("Que voulez-vous faire ?")
	fmt.Println("1 - Partir à l'aventure")
	fmt.Println("2 - Ouvrir le menu")

	var choix int
	fmt.Scan(&choix)

	switch choix {
	case 1:
		fmt.Println("L'aventure commence !")
		combat.Combat(perso)
	case 2:
		charactermenu.AfficherMenu(perso)
	default:
		fmt.Println("Choix invalide, réessayez.")
	}
}
