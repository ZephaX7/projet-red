package hub

import (
	"fmt"
	"projet-red/src/charactermenu"
	"projet-red/src/fight"
	"projet-red/src/model"
)

func Hub(perso *model.Personnage) {
	for {
		fmt.Println("\n🏰 Bienvenue au Hub ! 🏰")
		fmt.Println("1 - Partir à l'aventure")
		fmt.Println("2 - Ouvrir le menu")
		fmt.Println("3 - Quitter")

		var choix int
		fmt.Print("Entrez votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			fmt.Println("L'aventure commence !")
			ennemi := model.RandomEnnemi()
			fmt.Printf("🔹 Un ennemi apparaît : %s ! PV : %d/%d\n", ennemi.Nom, ennemi.PVActuels, ennemi.PVMax)
			fight.Combat(perso, ennemi) // Combat gère la mort / résurrection

		case 2:
			charactermenu.AfficherMenu(perso)

		case 3:
			fmt.Println("Au revoir !")
			return

		default:
			fmt.Println("Choix invalide, réessayez.")
		}
	}
}
