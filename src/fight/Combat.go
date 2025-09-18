package combat

import (
	"fmt"
	"projet-red/src/inventory"
	"projet-red/src/model"
)

func Combat(perso *model.Personnage) {
	ennemiPV := 60
	for ennemiPV > 0 && perso.PVActuels > 0 {
		fmt.Println("1 - Attaquer")
		fmt.Println("2 - Utiliser un objet")
		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			ennemiPV -= 20
			fmt.Println("Vous avez attaqué ! PV ennemi :", ennemiPV)
		case 2:
			inventory.AccessInventory()
			fmt.Println("Quel objet voulez-vous utiliser ?")
			var nom string
			fmt.Scan(&nom)
			inventory.UtiliserObjet(nom, perso)
		default:
			fmt.Println("Choix invalide.")
		}

		if ennemiPV > 0 {
			perso.PVActuels -= 15
			fmt.Printf("L'ennemi vous attaque ! PV : %d/%d\n", perso.PVActuels, perso.PVMax)
		}
	}
	fmt.Println("Combat terminé !")
}
