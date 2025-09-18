package fight

import (
	"fmt"
	"projet-red/src/inventory"
	"projet-red/src/model"
)

// Combat tour par tour joueur vs ennemi
func Combat(joueur *model.Personnage) {
	ennemi := model.RandomEnnemi()
	fmt.Println("🔥 Un ennemi apparaît !")
	fmt.Println(ennemi.Afficher())

	for ennemi.PVActuels > 0 && joueur.PVActuels > 0 {
		fmt.Println("\nVotre tour :")
		fmt.Println("1 - Attaquer")
		fmt.Println("2 - Utiliser un objet")
		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			// Attaque simple : dégâts = joueur Gold / 10 (ou fixe)
			degat := 20
			ennemi.PVActuels -= degat
			if ennemi.PVActuels < 0 {
				ennemi.PVActuels = 0
			}
			fmt.Printf("Vous attaquez l'ennemi et infligez %d dégâts ! PV ennemi : %d/%d\n",
				degat, ennemi.PVActuels, ennemi.PVMax)
		case 2:
			inventory.AccessInventory()
			fmt.Println("Quel objet voulez-vous utiliser ?")
			var nom string
			fmt.Scan(&nom)
			inventory.UtiliserObjet(nom, joueur)
		default:
			fmt.Println("Choix invalide.")
		}

		// Tour de l'ennemi si il est encore vivant
		if ennemi.PVActuels > 0 {
			fmt.Println("\nTour de l'ennemi !")
			degat := ennemi.Force
			joueur.PVActuels -= degat
			if joueur.PVActuels < 0 {
				joueur.PVActuels = 0
			}
			fmt.Printf("L'ennemi attaque et vous inflige %d dégâts ! PV joueur : %d/%d\n",
				degat, joueur.PVActuels, joueur.PVMax)
		}
	}

	// Fin du combat
	if joueur.PVActuels <= 0 {
		fmt.Println("💀 Vous avez été vaincu !")
		joueur.PVActuels = joueur.PVMax / 2
		fmt.Printf("Vous êtes ressuscité avec %d PV.\n", joueur.PVActuels)
	} else {
		fmt.Printf("🎉 Vous avez vaincu %s ! Vous gagnez %d XP et %d pièces d'or.\n",
			ennemi.Race, ennemi.Xp, ennemi.Gold)
		joueur.Xp += ennemi.Xp
		joueur.Gold += ennemi.Gold
	}
}
