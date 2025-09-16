package Items

import (
	"fmt"
	"time"
)

func takePot() {
	for i, item := range Inventaire {
		if item.Nom == "Potion de soin" && item.Quantite > 0 {
			// Retirer la potion
			Inventaire[i].Quantite -= 1

			// Soigner le joueur
			Player.Joueur.PV += 50
			if Player.Joueur.PV > Player.Joueur.PVMax {
				Player.Joueur.PV = Player.Joueur.PVMax
			}

			fmt.Printf("Vous avez utilisé une Potion de soin. PV actuels : %d/%d\n", Player.Joueur.PV, Player.Joueur.PVMax)
			return
		}
	}
	fmt.Println("Vous n'avez pas de Potion de soin dans votre inventaire.")
}

func PoisonPot(cible string) {
	for i, item := range Inventaire {
		if item.Nom == "Potion de poison" && item.Quantite > 0 {
			// Retirer la potion
			Inventaire[i].Quantite -= 1
			fmt.Println("Vous avez utilisé une Potion de poison !")

			// Inflige 10 dégâts par seconde pendant 3 secondes
			for t := 1; t <= 3; t++ {
				if cible == "joueur" {
					Player.Joueur.PV -= 10
					if Player.Joueur.PV < 0 {
						Player.Joueur.PV = 0
					}
					fmt.Printf("Seconde %d : PV du joueur = %d/%d\n", t, Player.Joueur.PV, Player.Joueur.PVMax)
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
