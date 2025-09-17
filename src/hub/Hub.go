package hub

import (
	"fmt"
	"os"
	menucharacter "projet-red/src/charactermenu"
)

func Hub() {
	fmt.Println("Que voulez-vous faire ?")

	fmt.Println()

	fmt.Println("1 - Partir a l'aventure ")
	fmt.Println("2 - Ouvrir le menu")

	fmt.Println()
	var choix int
	fmt.Print("Entrez votre choix : ")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		fmt.Println("L'aventure commence !")
	case 2:
		Menu_tab, err := os.ReadFile("src/hub/asciimenu.txt")
		//police small shadow
		if err != nil {
			panic(err)
		}
		fmt.Println(string(Menu_tab))
		fmt.Println()
		menucharacter.AfficherMenu()
	default:
		fmt.Println("Choix invalide, réessayez.")
	}
}
