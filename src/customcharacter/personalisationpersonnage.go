package personalisationpersonnage

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"projet-red/src/hub"
)

//
// --------------------- Modèles (POO) ---------------------
//

type Race int
type Classe int
type Sexe int

const (
	Humain Race = iota + 1
	Elfe
	Nain
)

const (
	Guerrier Classe = iota + 1
	Mage
	Assassin
)

const (
	Masculin Sexe = iota + 1
	Feminin
	Autre
)

func (r Race) String() string {
	switch r {
	case Humain:
		return "Humain"
	case Elfe:
		return "Elfe"
	case Nain:
		return "Nain"
	default:
		return "Inconnue"
	}
}

func (c Classe) String() string {
	switch c {
	case Guerrier:
		return "Guerrier"
	case Mage:
		return "Mage"
	case Assassin:
		return "Assassin"
	default:
		return "Inconnue"
	}
}

func (s Sexe) String() string {
	switch s {
	case Masculin:
		return "Masculin"
	case Feminin:
		return "Féminin"
	case Autre:
		return "Autre"
	default:
		return "Inconnu"
	}
}

// Personnage = entité principale
type Personnage struct {
	Nom       string
	Race      Race
	Classe    Classe
	Sexe      Sexe
	PVActuels int
	PVMax     int
}

// InitPV : initialise PV max / PV actuels selon la race
func (p *Personnage) InitPV() {
	switch p.Race {
	case Humain:
		p.PVMax = 100
		p.PVActuels = 50
	case Elfe:
		p.PVMax = 80
		p.PVActuels = 40
	case Nain:
		p.PVMax = 120
		p.PVActuels = 60
	default:
		p.PVMax = 100
		p.PVActuels = 50
	}
}

// Afficher : résumé lisible
func (p Personnage) Afficher() string {
	return fmt.Sprintf(
		"Nom : %s\nRace : %s\nClasse : %s\nSexe : %s\nPV : %d/%d\n",
		p.Nom, p.Race, p.Classe, p.Sexe, p.PVActuels, p.PVMax,
	)
}

//
// ----------- Constructeur + I/O (Personnalisateur) -----------
//

type Personnalisateur struct {
	reader *bufio.Reader
}

// NewPersonnalisateur : constructeur
func NewPersonnalisateur() *Personnalisateur {
	return &Personnalisateur{
		reader: bufio.NewReader(os.Stdin),
	}
}

// I/O utilitaires (méthodes)
func (pz *Personnalisateur) lireLigne(prompt string) (string, error) {
	fmt.Print(prompt)
	s, err := pz.reader.ReadString('\n')
	return strings.TrimSpace(s), err
}

func (pz *Personnalisateur) lireChoixInt(prompt string, min, max int) int {
	for {
		s, _ := pz.lireLigne(prompt)
		if s == "" {
			continue
		}
		n, err := strconv.Atoi(s)
		if err != nil || n < min || n > max {
			if min == max {
				fmt.Printf("Choix invalide. Veuillez choisir %d.\n", min)
			} else {
				fmt.Printf("Choix invalide. Veuillez choisir un nombre entre %d et %d.\n", min, max)
			}
			continue
		}
		return n
	}
}

// Menus (méthodes)
func (pz *Personnalisateur) ChoixRace() Race {
	fmt.Println("1. Humains (adapté à la classe guerrier)")
	fmt.Println("2. Elfes (adapté à la classe mage)")
	fmt.Println("3. Nains (adapté à la classe assassin)")
	fmt.Println()

	n := pz.lireChoixInt("Choisissez votre race (1, 2 ou 3) : ", 1, 3)
	r := Race(n)
	fmt.Printf("Vous avez choisi la race des %s.\n", r)
	return r
}

func (pz *Personnalisateur) ChoixClasse() Classe {
	fmt.Println("1. Guerrier")
	fmt.Println("2. Mage")
	fmt.Println("3. Assassin")
	fmt.Println()

	n := pz.lireChoixInt("Choisissez votre classe (1, 2 ou 3) : ", 1, 3)
	c := Classe(n)
	fmt.Printf("Vous avez choisi %s.\n", c)
	return c
}

func (pz *Personnalisateur) ChoixSexe() Sexe {
	fmt.Println("1. Masculin")
	fmt.Println("2. Féminin")
	fmt.Println("3. Autre")
	fmt.Println()

	n := pz.lireChoixInt("Choisissez le sexe de votre personnage (1, 2 ou 3) : ", 1, 3)
	s := Sexe(n)
	fmt.Printf("Vous avez choisi %s.\n", s)
	return s
}

func (pz *Personnalisateur) ChoisirNom() string {
	for {
		nom, _ := pz.lireLigne("Entrez le nom de votre personnage : ")
		if nom == "" {
			fmt.Println("Le nom ne peut pas être vide.")
			continue
		}
		fmt.Printf("Votre personnage s'appelle : %s\n", nom)
		reponse, _ := pz.lireLigne("Êtes-vous sûr ? (oui/non) : ")
		switch strings.ToLower(strings.TrimSpace(reponse)) {
		case "oui", "o", "y", "yes":
			fmt.Printf("Bienvenue dans le monde de Aerthar (%s).\n", nom)
			return nom
		case "non", "n", "no":
			fmt.Println("D'accord, recommençons.")
			continue
		default:
			fmt.Println("Réponse non reconnue. Tape 'oui' ou 'non'.")
			continue
		}
	}
}

// Orchestrateur (méthode) : crée un personnage complet
func (pz *Personnalisateur) CreerPersonnageInteractif() *Personnage {
	r := pz.ChoixRace()
	c := pz.ChoixClasse()
	s := pz.ChoixSexe()
	nom := pz.ChoisirNom()

	p := &Personnage{
		Nom:    nom,
		Race:   r,
		Classe: c,
		Sexe:   s,
	}
	p.InitPV()

	// Affichage PV selon la race (équivalent de ton PointdeVie)
	switch p.Race {
	case Humain:
		fmt.Println("Vous avez 50 points de vie actuellement, et vous pouvez en avoir jusqu'à 100 étant humain.")
	case Elfe:
		fmt.Println("Vous avez 40 points de vie actuellement, et vous pouvez en avoir jusqu'à 80 étant elfe.")
	case Nain:
		fmt.Println("Vous avez 60 points de vie actuellement, et vous pouvez en avoir jusqu'à 120 étant nain.")
	}

	return p
}

// -------------- Fonctions de compat / démarrage --------------
//
// StartFlow : conserve ton comportement = crée le perso puis lance le hub,
// et retourne le personnage créé (POO-friendly).
func StartFlow() *Personnage {
	pz := NewPersonnalisateur()
	p := pz.CreerPersonnageInteractif()
	hub.Hub() // ⚠️ Suppose que hub n’importe PAS personalisationpersonnage (sinon cycle)
	return p
}

// (Optionnel) Wrapper si tu veux garder l’ancienne signature de retour (strings)
func CreerPersonnageInteractif() (nom, race, classe, sexe string) {
	pz := NewPersonnalisateur()
	p := pz.CreerPersonnageInteractif()
	hub.Hub()
	return p.Nom, p.Race.String(), p.Classe.String(), p.Sexe.String()
}
