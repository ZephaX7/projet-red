package isdead

import (
	"fmt"
	"projet-red/src/hub"
	"projet-red/src/model"
)

func IsDead(perso *model.Personnage) {
	if perso.PVActuels <= 0 {
		fmt.Println("💀 Vous êtes mort... mais vous êtes ressuscité avec la moitié de vos PV.")
		perso.PVActuels = perso.PVMax / 2
		hub.Hub(perso) // ✅ On relance le hub avec le personnage
	}
}
