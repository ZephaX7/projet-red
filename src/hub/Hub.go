package hub

import (
	"fmt"
	"projet-red/src/charactermenu"
	"projet-red/src/fight"
	"projet-red/src/model"
)

func Hub(perso *model.Personnage) {
	gameOver := false

	for !gameOver {
		fmt.Println("\n🏰 Bienvenue au Hub ! 🏰")
		fmt.Println("1 - Partir à l'aventure")
		fmt.Println("2 - Ouvrir le menu")
		fmt.Println("3 - Quitter")

		var choix int
		fmt.Print("Entrez votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			ennemi := model.RandomEnnemi()
			fmt.Printf("\nUn ennemi apparaît : %s ! PV : %d/%d\n", ennemi.Nom, ennemi.PVActuels, ennemi.PVMax)

			fight.Combat(perso, ennemi)

			if perso.PVActuels <= 0 {
				if !perso.Revived {
					perso.PVActuels = perso.PVMax / 2
					perso.Revived = true
					fmt.Printf("\n💀 Vous êtes mort mais ressuscité ! PV : %d/%d\n", perso.PVActuels, perso.PVMax)
				} else {
					fmt.Println("\n💀 Vous êtes mort pour de bon ! Game Over.")
					gameOver = true // stoppe le hub
				}
			}

		case 2:
			charactermenu.AfficherMenu(perso)
		case 3:
			fmt.Println("Au revoir !")
			gameOver = true
		default:
			fmt.Println("Choix invalide, réessayez.")
		}
	}
}
