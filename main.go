package main

import (
	"fmt"
	"projet-red/src/customcharacter"
	"projet-red/src/hub"
)

func main() {
	// Création perso
	p := customcharacter.StartFlow()
	fmt.Println("\n=== Personnage créé ===")
	fmt.Println(p.Afficher())

	// Entrée dans le hub
	hub.Hub(p)
}
