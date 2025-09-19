package model

import (
	"fmt"
	"math/rand"
)

type Ennemi struct {
	Nom       string
	Race      string
	PVMax     int
	PVActuels int
	Degats    int
}

// Fonction pour générer un ennemi aléatoire
func RandomEnnemi() *Ennemi {
	ennemis := []Ennemi{
		{"Orc", "Orc", 120, 120, 20},
		{"Gobelin", "Gobelin", 70, 70, 15},
		{"Troll", "Troll", 150, 150, 25},
		{"Loup", "Loup", 80, 80, 18},
	}

	i := rand.Intn(len(ennemis))
	e := ennemis[i]
	return &e
}

func GoblinPattern(player *Personnage, gobelin *Ennemi) {
	turn := 1

	for player.PVActuels > 0 && gobelin.PVActuels > 0 {
		var damage int

		// Tous les 3 tours, dégâts doublés
		if turn%3 == 0 {
			damage = gobelin.Degats * 2
		} else {
			damage = gobelin.Degats
		}

		player.PVActuels -= damage
		if player.PVActuels < 0 {
			player.PVActuels = 0
		}

		fmt.Printf("%s inflige à %s %d de dégâts\n", gobelin.Nom, player.Nom, damage)
		fmt.Printf("PV de %s : %d/%d\n\n", player.Nom, player.PVActuels, player.PVMax)

		turn++

		// Arrêt si le joueur est mort
		if player.PVActuels <= 0 {
			fmt.Printf("%s a été vaincu !\n", player.Nom)
			break
		}
	}
}
