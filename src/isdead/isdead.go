package isdead

import (
	"projet-red/src/model"
)

// isdead/check.go
func CheckDeath(perso *model.Personnage) bool {
	if perso.PVActuels <= 0 && !perso.Revived {
		perso.PVActuels = perso.PVMax / 2
		perso.Revived = true
		return true // ressuscité
	}
	return false
}
