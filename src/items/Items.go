package items

import (
	"fmt"

	statspersonnage "projet-red/src/stats"
	"time"
)

func TakePot() {
	// Soigner le joueur
	statspersonnage.Joueur.PvActuels += 50
	if statspersonnage.Joueur.PvActuels > statspersonnage.Joueur.PvMax {
		statspersonnage.Joueur.PvActuels = statspersonnage.Joueur.PvMax
	}

	fmt.Printf("Vous avez utilisé une Potion de soin. PV actuels : %d/%d\n", statspersonnage.Joueur.PvActuels, statspersonnage.Joueur.PvMax)
	return
}

func PoisonPot(cible string) {
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
