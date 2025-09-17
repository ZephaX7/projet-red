package items

import (
	"fmt"
	"projet-red/src/inventory"
	statspersonnage "projet-red/src/stats"
	"time"
)

func takePot() {
	for i, item := range inventory.Inventaire {
		if item.Nom == "Potion de soin" && item.Quantite > 0 {
			// Retirer la potion
			inventory.Inventaire[i].Quantite -= 1

			// Soigner le joueur
			statspersonnage.Joueur.PvActuels += 50
			if statspersonnage.Joueur.PvActuels > statspersonnage.Joueur.PvMax {
				statspersonnage.Joueur.PvActuels = statspersonnage.Joueur.PvMax
			}

			fmt.Printf("Vous avez utilisé une Potion de soin. PV actuels : %d/%d\n", statspersonnage.Joueur.PvActuels, statspersonnage.Joueur.PvMax)
			return
		}
	}
	fmt.Println("Vous n'avez pas de Potion de soin dans votre inventaire.")
}

func PoisonPot(cible string) {
	for i, item := range inventory.Inventaire {
		if item.Nom == "Potion de poison" && item.Quantite > 0 {
			// Retirer la potion
			inventory.Inventaire[i].Quantite -= 1
			fmt.Println("Vous avez utilisé une Potion de poison !")

			// Inflige 10 dégâts par seconde pendant 3 secondes
			for t := 1; t <= 3; t++ {
				if cible == "Joueur" {
					statspersonnage.Joueur.PvActuels -= 10
					if statspersonnage.Joueur.PvActuels < 0 {
						statspersonnage.Joueur.PvActuels = 0
					}
					fmt.Printf("Seconde %d : PV du Joueur = %d/%d\n", t, statspersonnage.Joueur.PvActuels, statspersonnage.Joueur.PvMax)
				} else if cible == "ennemi" {
					// Ici tu peux gérer les PV de l'ennemi
					fmt.Printf("Seconde %d : L'ennemi perd 10 PV\n", t)
				}

				time.Sleep(1 * time.Second)
			}
			return
		}
	}
	fmt.Println("Vous n'avez pas de Potion de poison dans votre inventaire.")
}
