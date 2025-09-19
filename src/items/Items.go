package items

import (
	"fmt"
	"projet-red/src/inventory"
	"projet-red/src/model"
	"time"
)

// TakePot utilise une potion de soin
func TakePot(perso *model.Personnage, inventaire *[]inventory.Objet) {
	for i, obj := range *inventaire {
		if obj.Nom == "Potion de soin" && obj.Quantite > 0 {
			(*inventaire)[i].Quantite--
			if (*inventaire)[i].Quantite <= 0 {
				*inventaire = append((*inventaire)[:i], (*inventaire)[i+1:]...)
			}

			soin := 50
			perso.PVActuels += soin
			if perso.PVActuels > perso.PVMax {
				perso.PVActuels = perso.PVMax
			}

			fmt.Printf("Vous avez utilisé une Potion de soin ! PV actuels : %d/%d\n",
				perso.PVActuels, perso.PVMax)
			return
		}
	}
	fmt.Println("⚠ Vous n'avez pas de Potion de soin dans votre inventaire.")
}

// PoisonPot inflige 10 PV de dégâts par seconde pendant 3 secondes
// PoisonPot inflige 10 PV de dégâts par seconde pendant 3 secondes à un ennemi
func PoisonPot(target *model.Ennemi, inventaire *[]inventory.Objet) {
	for i := 1; i <= 3; i++ {
		damage := 10
		target.PVActuels -= damage
		if target.PVActuels < 0 {
			target.PVActuels = 0
		}
		fmt.Printf("💀 Poison ! %s subit %d PV de dégâts. PV actuels : %d/%d\n",
			target.Nom, damage, target.PVActuels, target.PVMax)
		time.Sleep(1 * time.Second)
		if target.PVActuels <= 0 {
			fmt.Printf("%s est mort !\n", target.Nom)
			break
		}
	}

	// Retirer la potion de l'inventaire
	for i, obj := range *inventaire {
		if obj.Nom == "Potion de poison" && obj.Quantite > 0 {
			(*inventaire)[i].Quantite--
			if (*inventaire)[i].Quantite <= 0 {
				*inventaire = append((*inventaire)[:i], (*inventaire)[i+1:]...)
			}
			break
		}
	}

	fmt.Println("La potion de poison a été consommée.")
}

// UtiliserObjet permet d'utiliser un objet depuis l'inventaire
func UtiliserObjet(nom string, perso *model.Personnage, inventaire *[]inventory.Objet, target any, inCombat bool) {
	for _, item := range *inventaire {
		if item.Nom == nom {
			switch item.Nom {
			case "Potion de soin":
				TakePot(perso, inventaire)

			case "Potion de poison":
				if inCombat {
					if ennemi, ok := target.(*model.Ennemi); ok {
						PoisonPot(ennemi, inventaire)
					} else {
						fmt.Println("⚠ Potion de poison utilisable uniquement sur un ennemi.")
					}
				} else {
					fmt.Println("⚠ Potion de poison utilisable uniquement en combat.")
				}

			case "Chapeau de l'aventurier", "Tunique de l'aventurier", "Bottes de l'aventurier":
				perso.EquipItem(item.Nom)

			default:
				fmt.Println("⚠ Objet inconnu :", item.Nom)
			}
			return
		}
	}
	fmt.Println("⚠ Vous n'avez pas cet objet dans l'inventaire :", nom)
}
