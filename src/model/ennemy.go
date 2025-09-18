package model

import (
	"fmt"
	"math/rand"
)

type Ennemi struct {
	Race      string
	PVActuels int
	PVMax     int
	Gold      int
	Xp        int
	Force     int
}

// Affichage lisible
func (e *Ennemi) Afficher() string { // ⚡ pointeur pour pouvoir modifier les PV
	return fmt.Sprintf("Race : %s\nPV : %d/%d\nForce : %d\nGold : %d\nXP : %d",
		e.Race, e.PVActuels, e.PVMax, e.Force, e.Gold, e.Xp)
}

// Liste des ennemis possibles
var ListeEnnemis = []Ennemi{
	{"Gobelin", 50, 50, 10, 20, 10},
	{"Orc", 80, 80, 15, 30, 15},
	{"Troll", 120, 120, 25, 50, 20},
	{"Squelette", 40, 40, 5, 15, 8},
}

// Génère un ennemi aléatoire
func RandomEnnemi() *Ennemi { // retourne un pointeur
	e := ListeEnnemis[rand.Intn(len(ListeEnnemis))]
	return &e
}
