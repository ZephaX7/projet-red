package items

import (
	"fmt"
	"projet-red/src/model"
	"time"
)

// Potion de soin
func TakePot(perso *model.Personnage) {
	perso.PVActuels += 50
	if perso.PVActuels > perso.PVMax {
		perso.PVActuels = perso.PVMax
	}
	fmt.Printf("Vous avez utilisé une Potion de soin. PV actuels : %d/%d\n", perso.PVActuels, perso.PVMax)
}

// Potion de poison
func PoisonPot(perso *model.Personnage, cible string) {
	for t := 1; t <= 3; t++ {
		if cible == "Joueur" {
			perso.PVActuels -= 10
			if perso.PVActuels < 0 {
				perso.PVActuels = 0
			}
			fmt.Printf("Seconde %d : PV du Joueur = %d/%d\n", t, perso.PVActuels, perso.PVMax)
		} else if cible == "ennemi" {
			fmt.Printf("Seconde %d : L'ennemi perd 10 PV\n", t)
		}
		time.Sleep(1 * time.Second)
	}
}
