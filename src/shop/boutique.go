package shop

import (
	"fmt"
	"log"
	"os"
	"projet-red/src/inventory"
	"projet-red/src/model"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// Fonction pour lancer la musique d'accueil
func musiqueShop() (beep.StreamSeekCloser, beep.Format) {
	f, err := os.Open("src/asset/Shop.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// Créer un effet de volume
	volume := &effects.Volume{
		Streamer: streamer,
		Base:     2.8, // base logarithmique
		Volume:   -2,  // diminue le volume de 5 dB
		Silent:   false,
	}

	// Lancer en arrière-plan
	go func() {
		speaker.Play(volume)
	}()

	return streamer, format
}

func Shop(perso *model.Personnage) {

	streamer, _ := musiqueShop()

	items := []string{
		"Potion de soin gratuite (cadeau de la maison, mais une seule fois !)",
		"Potion de soin",
		"Potion de poison",
		"Livre de Sort : Boule de feu",
		"Fourrure de Loup",
		"Peau de Troll",
		"Cuir de Sanglier",
		"Plume de Corbeaux",
		"Upgrade inventaire",
	}
	prices := []int{0, 3, 6, 25, 4, 7, 3, 1, 30}
	freeTaken := false

	// Compteur pour les upgrades, max 3
	var upgradeCount int

	for {
		fmt.Printf("\nVous avez %d pièces d'or.\n", perso.Gold)
		for i, item := range items {
			fmt.Printf("%d. %s - %d pièces d'or\n", i+1, item, prices[i])
		}
		fmt.Println("Entrez le numéro de l'article à acheter (ou 0 pour quitter) :")

		var choice int
		fmt.Scan(&choice)

		if choice == 0 {
			streamer.Close()

			break
		}
		if choice < 1 || choice > len(items) {
			fmt.Println("Choix invalide.\n")
			continue
		}

		itemName := items[choice-1]
		itemPrice := prices[choice-1]

		// Article gratuit
		if itemPrice == 0 {
			if freeTaken {
				fmt.Println("⚠ Vous ne pouvez prendre cet article gratuit qu'une seule fois !\n")
				continue
			}
			if len(inventory.Inventaire) >= inventory.CapaciteMax {
				fmt.Println("Votre inventaire est plein, impossible de prendre l'objet gratuit.\n")
				continue
			}
			inventory.AddInventory(inventory.Objet{Nom: itemName, Quantite: 1, Type: "Objet"})
			freeTaken = true
			fmt.Println("Vous avez pris :", itemName, "\n")
			continue
		}

		// Achat normal
		if perso.Gold >= itemPrice {
			// Gestion de l'upgrade
			if itemName == "Upgrade inventaire" {
				if upgradeCount >= 3 {
					fmt.Println("⚠ Vous ne pouvez plus améliorer votre inventaire ! Limite atteinte.\n")
					continue
				}
				perso.Gold -= itemPrice
				upgradeCount++
				inventory.CapaciteMax += 10
				fmt.Printf("✅ Votre inventaire a été amélioré ! Nouvelle capacité : %d (Améliorations restantes : %d)\n\n", inventory.CapaciteMax, 3-upgradeCount)
				continue
			}
			if itemName == "Livre de Sort : Boule de feu" {
				// Vérifie si le joueur connaît déjà le sort
				alreadyKnown := false
				for _, s := range perso.Skills {
					if s == "Boule de feu" {
						alreadyKnown = true
						break
					}
				}
				if !alreadyKnown {
					perso.Skills = append(perso.Skills, "Boule de feu")
					fmt.Println("🔥 Vous avez appris le sort : Boule de feu !")
				} else {
					fmt.Println("⚠ Vous connaissez déjà ce sort.")
				}
				perso.Gold -= itemPrice
				continue
			}

			if len(inventory.Inventaire) >= inventory.CapaciteMax {
				fmt.Println("Votre inventaire est plein, vous ne pouvez rien acheter de plus.\n")
				continue
			}

			perso.Gold -= itemPrice
			inventory.AddInventory(inventory.Objet{Nom: itemName, Quantite: 1, Type: "Objet"})
			fmt.Println("Vous avez acheté :", itemName, "\n")
		} else {
			fmt.Println("Vous n'avez pas assez d'or.\n")
		}
	}
}
