package fight

import (
	"fmt"
	"os"
	"projet-red/src/inventory"
	"projet-red/src/model"
)

func Combat(perso *model.Personnage, ennemi *model.Ennemi) {
	for ennemi.PVActuels > 0 && perso.PVActuels > 0 {
		fmt.Println("\n1 - Attaquer")
		fmt.Println("2 - Utiliser un objet")
		fmt.Println("3 - Utiliser une compétence magique (si livre dans l'inventaire)")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			damage := ennemi.Degats
			perso.PVActuels -= damage
			if ennemi.PVActuels < 0 {
				ennemi.PVActuels = 0
			}
			fmt.Printf("Vous attaquez %s et infligez %d PV. PV ennemi : %d/%d\n",
				ennemi.Nom, damage, ennemi.PVActuels, ennemi.PVMax)

		case 2:
			inventory.AccessInventory()
			fmt.Println("Quel objet voulez-vous utiliser ?")
			var nom string
			fmt.Scan(&nom)
			inventory.UtiliserObjet(nom, perso)

		case 3:
			if inventory.HasMagicBook(perso) { // fonction à créer pour vérifier l'inventaire
				damage := 50 // dégâts du sort
				ennemi.PVActuels -= damage
				if ennemi.PVActuels < 0 {
					ennemi.PVActuels = 0
				}
				fmt.Printf("💥 Vous utilisez votre sort magique et infligez %d PV à %s ! PV ennemi : %d/%d\n",
					damage, ennemi.Nom, ennemi.PVActuels, ennemi.PVMax)
			} else {
				fmt.Println("⚠ Vous n'avez pas de livre magique dans votre inventaire !")
			}

		default:
			fmt.Println("⚠ Choix invalide.")
			continue
		}

		// Attaque de l'ennemi
		if ennemi.PVActuels > 0 {
			damage := 15
			perso.PVActuels -= damage
			if perso.PVActuels < 0 {
				perso.PVActuels = 0
			}
			fmt.Printf("%s vous attaque ! PV : %d/%d\n", ennemi.Nom, perso.PVActuels, perso.PVMax)

			// Résurrection unique
			if perso.PVActuels <= 0 {
				if !perso.Revived {
					// Première résurrection
					perso.PVActuels = perso.PVMax / 2
					perso.Revived = true
					fmt.Printf("💀 Vous êtes mort mais ressuscité avec %d/%d PV !\n", perso.PVActuels, perso.PVMax)
					return // retour au hub
				} else {
					// Deuxième mort → game over
					fmt.Println("☠️ Vous êtes mort définitivement. Fin du jeu.")
					os.Exit(0) // ou return pour terminer le hub
				}
			}
		}
	}
}
