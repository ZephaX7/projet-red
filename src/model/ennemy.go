package model

import (
	"math/rand"
	"time"
)

// Ennemi struct
type Ennemi struct {
	Nom       string
	Race      string
	PVMax     int
	PVActuels int
	// Ajoute d'autres stats si nécessaire
}

// Liste des ennemis possibles
var EnnemiList = []Ennemi{
	{Nom: "Gobelin", Race: "Gobelin", PVMax: 50, PVActuels: 50},
	{Nom: "Orc", Race: "Orc", PVMax: 80, PVActuels: 80},
	{Nom: "Troll", Race: "Troll", PVMax: 120, PVActuels: 120},
}

// RandomEnnemi retourne un ennemi aléatoire
func RandomEnnemi() *Ennemi {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(EnnemiList))
	e := EnnemiList[index]

	// On crée une vraie copie pour que chaque combat soit indépendant
	ennemi := Ennemi{
		Nom:       e.Nom,
		Race:      e.Race,
		PVMax:     e.PVMax,
		PVActuels: e.PVMax, // On initialise les PV à max
	}
	return &ennemi
}
