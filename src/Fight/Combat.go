package combat

import (
	"fmt"
	"strings"
)

// Combat tour par tour entre le joueur et un ennemi
func Combat() {
	playerHP := personalisationpersonnage.PlayerHP
	enemyHP := 40
	playerAttack := 10
	enemyAttack := 8

	fmt.Println("Un combat commence !")
	for playerHP > 0 && enemyHP > 0 {
		fmt.Printf("\nVotre vie : %d | Vie de l'ennemi : %d\n", playerHP, enemyHP)
		fmt.Println("A vous de jouer ! (1: Attaquer, 2: Défendre)")
		input, _ := personalisationpersonnage.Reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "1" {
			fmt.Println("Vous attaquez l'ennemi !")
			enemyHP -= playerAttack
		} else if input == "2" {
			fmt.Println("Vous vous défendez et subissez moins de dégâts.")
			playerHP -= enemyAttack / 2
		} else {
			fmt.Println("Action invalide, vous subissez l'attaque de l'ennemi.")
			playerHP -= enemyAttack
		}

		if enemyHP > 0 {
			fmt.Println("L'ennemi vous attaque !")
			playerHP -= enemyAttack
		}
		// Mise à jour de la vie globale du joueur
		personalisationpersonnage.PlayerHP = playerHP
	}

	if playerHP > 0 {
		fmt.Println("\nVous avez gagné le combat !")
	} else {
		fmt.Println("\nVous avez perdu le combat...")
	}
}
