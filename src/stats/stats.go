package stats

import (
	"fmt"
	"projet-red/src/model"
)

func AfficherStats(p *model.Personnage) {
	fmt.Println("=== Stats du perso ===")
	fmt.Println(p.Afficher())
}
