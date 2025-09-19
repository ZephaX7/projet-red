package forgeron

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

type Material struct {
	Name     string
	Quantity int
}

type ItemCost struct {
	Gold      int
	Materials []Material
}

func musiqueShop() (beep.StreamSeekCloser, beep.Format) {
	f, err := os.Open("asset/Forge.mp3")
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
		Volume:   -1,  // diminue le volume de 5 dB
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
		"Chapeau de l'aventurier",
		"Tunique de l'aventurier",
		"Bottes de l'aventurier",
	}

	costs := []ItemCost{
		{Gold: 5, Materials: []Material{{"Plume de Corbeaux", 1}, {"Cuir de Sanglier", 1}}}, // ⚡ pluriel ici
		{Gold: 5, Materials: []Material{{"Fourrure de Loup", 2}, {"Peau de Troll", 1}}},
		{Gold: 5, Materials: []Material{{"Fourrure de Loup", 1}, {"Cuir de Sanglier", 1}}},
	}

	for {
		fmt.Printf("\n💰 Vous avez %d pièces d'or.\n", perso.Gold)
		for i, item := range items {
			fmt.Printf("%d. %s - %d pièces d'or + matériaux : ", i+1, item, costs[i].Gold)
			for _, m := range costs[i].Materials {
				fmt.Printf("%d %s ", m.Quantity, m.Name)
			}
			fmt.Println()
		}
		fmt.Println("Entrez le numéro de l'article à fabriquer (ou 0 pour quitter) :")

		var choice int
		fmt.Scan(&choice)

		if choice == 0 {
			streamer.Close()
			break
		}
		if choice < 1 || choice > len(items) {
			fmt.Println("⚠ Choix invalide.")
			continue
		}

		cost := costs[choice-1]

		// Vérifier l'or
		if perso.Gold < cost.Gold {
			fmt.Println("⚠ Vous n'avez pas assez d'or.")
			continue
		}

		// Vérifier les matériaux
		canCraft := true
		for _, m := range cost.Materials {
			total := 0
			for _, obj := range inventory.Inventaire {
				if obj.Nom == m.Name {
					total += obj.Quantite
				}
			}
			if total < m.Quantity {
				fmt.Printf("⚠ Il vous manque %d %s.\n", m.Quantity-total, m.Name)
				canCraft = false
			}
		}
		if !canCraft {
			continue
		}

		// Vérifier la capacité de l'inventaire
		if len(inventory.Inventaire) >= inventory.CapaciteMax {
			fmt.Println("⚠ Votre inventaire est plein, impossible de fabriquer cet objet.")
			continue
		}

		// Retirer l'or
		perso.Gold -= cost.Gold

		// Retirer les matériaux
		for _, m := range cost.Materials {
			inventory.RemoveInventory(inventory.Objet{
				Nom:      m.Name,
				Quantite: m.Quantity,
			})
		}

		// Ajouter l'objet fabriqué
		inventory.AddInventory(inventory.Objet{
			Nom:      items[choice-1],
			Quantite: 1,
			Type:     "Équipement",
		})
		fmt.Println("Vous pouvez maintenant équiper cet objet. Voulez-vous l'équiper ? (oui/non)")
		var rep string
		fmt.Scan(&rep)
		if rep == "oui" {
			perso.EquipItem(items[choice-1])
		}

		fmt.Printf("✅ Vous avez fabriqué : %s\n", items[choice-1])
	}
}
