package fight

import (
	"fmt"
	"projet-red/src/inventory"
	"projet-red/src/items"
	"projet-red/src/model"
)

// Combat d'un joueur contre un ennemi avec pattern Gobelin
func Combat(perso *model.Personnage, ennemi *model.Ennemi) {
	turn := 1

	for ennemi.PVActuels > 0 && perso.PVActuels > 0 {
		fmt.Println("\n--- Tour de combat ---")
		fmt.Printf("Vous : %d/%d PV | Ennemi %s : %d/%d PV\n",
			perso.PVActuels, perso.PVMax, ennemi.Nom, ennemi.PVActuels, ennemi.PVMax)

		// Si l'ennemi est un gobelin, appliquer son pattern spécial
		if ennemi.Race == "Gobelin" {
			var damage int
			if turn%3 == 0 {
				damage = ennemi.Degats * 2
			} else {
				damage = ennemi.Degats
			}
			perso.PVActuels -= damage
			if perso.PVActuels < 0 {
				perso.PVActuels = 0
			}
			fmt.Printf("%s inflige %d PV à %s ! PV joueur : %d/%d\n",
				ennemi.Nom, damage, perso.Nom, perso.PVActuels, perso.PVMax)
		}

		// Tour joueur
		fmt.Println("\n1 - Attaquer (Coup de poing)")
		fmt.Println("2 - Utiliser un skill")
		fmt.Println("3 - Utiliser un objet")
		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			damage := 20
			ennemi.PVActuels -= damage
			if ennemi.PVActuels < 0 {
				ennemi.PVActuels = 0
			}
			fmt.Printf("Vous utilisez Coup de poing et infligez %d PV. PV ennemi : %d/%d\n",
				damage, ennemi.PVActuels, ennemi.PVMax)

		case 2:
			if len(perso.Skills) == 0 {
				fmt.Println("⚠ Vous n'avez aucune compétence à utiliser !")
				continue
			}
			fmt.Println("Compétences disponibles :")
			for i, skill := range perso.Skills {
				fmt.Printf("%d - %s\n", i+1, skill)
			}
			var skillChoice int
			fmt.Print("Choisissez une compétence : ")
			fmt.Scan(&skillChoice)

			if skillChoice < 1 || skillChoice > len(perso.Skills) {
				fmt.Println("⚠ Choix invalide")
				continue
			}

			skill := perso.Skills[skillChoice-1]

			var damage int
			switch skill {
			case "Coup de poing":
				damage = 20
			case "Boule de feu":
				damage = 50
			default:
				damage = 10
			}

			ennemi.PVActuels -= damage
			if ennemi.PVActuels < 0 {
				ennemi.PVActuels = 0
			}

			fmt.Printf("Vous utilisez %s et infligez %d PV à %s. PV ennemi : %d/%d\n",
				skill, damage, ennemi.Nom, ennemi.PVActuels, ennemi.PVMax)

		case 3:
			inventory.AccessInventory(perso, false, nil, items.UtiliserObjet)

			fmt.Println("Quel objet voulez-vous utiliser ?")
			var nom string
			fmt.Scan(&nom)

		default:
			fmt.Println("⚠ Choix invalide.")
			continue
		}

		// Tour classique de l'ennemi si ce n'est pas un Gobelin (déjà géré ci-dessus)
		if ennemi.Race != "Gobelin" && ennemi.PVActuels > 0 {
			damage := ennemi.Degats
			perso.PVActuels -= damage
			if perso.PVActuels < 0 {
				perso.PVActuels = 0
			}
			fmt.Printf("%s vous attaque et inflige %d PV ! PV joueur : %d/%d\n",
				ennemi.Nom, damage, perso.PVActuels, perso.PVMax)
		}

		// Vérifier la mort du joueur et appliquer revive unique
		if perso.PVActuels <= 0 {
			if !perso.Revived {
				perso.PVActuels = perso.PVMax / 2
				perso.Revived = true
				fmt.Printf("\n💀 Vous êtes mort mais ressuscité ! PV : %d/%d\n", perso.PVActuels, perso.PVMax)
				return // retour au hub après revival
			} else {
				fmt.Println("\n💀 Vous êtes mort pour de bon ! Game Over.")
				return
			}
		}

		turn++
	}
}
