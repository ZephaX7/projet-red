package model

import (
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
