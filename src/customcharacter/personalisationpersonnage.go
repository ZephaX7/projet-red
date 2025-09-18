package personalisationpersonnage

import (
	"bufio"
	"fmt"
	"math"
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
	Nom    string
	Race   Race
	Classe Classe
	Sexe   Sexe

	// PV
	PVActuels  int
	PVMax      int // PV total = PVBase + BonusPVMax
	PVBase     int // PV propres à la race
	BonusPVMax int // Bonus cumulés via montée de niveau

	// XP & Niveaux
	Niveau      int // Niveau actuel
	NiveauMax   int // Cap de niveau
	ExpActuelle int // XP actuelle dans le niveau courant
	ExpMax      int // XP requise pour passer au prochain niveau
}

// ---- Init PV selon la race ----
func (p *Personnage) InitPV() {
	switch p.Race {
	case Humain:
		p.PVBase = 100
		p.PVActuels = 50
	case Elfe:
		p.PVBase = 80
		p.PVActuels = 40
	case Nain:
		p.PVBase = 120
		p.PVActuels = 60
	default:
		p.PVBase = 100
		p.PVActuels = 50
	}
	p.BonusPVMax = 0
	p.RecalculerPVMax()
}

// ---- Init XP/Niveaux ----
func (p *Personnage) InitXP(niveauInitial, niveauMax int) {
	if niveauInitial < 1 {
		niveauInitial = 1
	}
	if niveauMax < niveauInitial {
		niveauMax = niveauInitial
	}
	p.Niveau = niveauInitial
	p.NiveauMax = niveauMax
	p.ExpActuelle = 0
	p.ExpMax = p.xpRequisePour(p.Niveau)
}

// Courbe d'XP : exponentielle douce (modifiable)
func (p *Personnage) xpRequisePour(niveau int) int {
	if niveau < 1 {
		niveau = 1
	}
	base := 100.0  // XP requise pour passer du niveau 1 au 2
	growth := 1.25 // facteur de croissance par niveau
	return int(math.Round(base * math.Pow(growth, float64(niveau-1))))
}

// Gain d'XP (gère l'overflow et les multi-level-up)
// Retourne (niveaux_gagnés, cap_atteint)
func (p *Personnage) GainXP(amount int) (int, bool) {
	if amount <= 0 || p.Niveau >= p.NiveauMax {
		// Cap atteint -> on ignore l'XP entrante, pour éviter la confusion
		return 0, p.Niveau >= p.NiveauMax
	}

	nivGagnes := 0
	p.ExpActuelle += amount

	for p.Niveau < p.NiveauMax {
		if p.ExpActuelle < p.ExpMax {
			break
		}
		// Consomme l'XP requise pour le niveau en cours
		p.ExpActuelle -= p.ExpMax
		p.Niveau++
		nivGagnes++
		p.onLevelUp()
		// Recalcule l'XP requise du nouveau niveau
		p.ExpMax = p.xpRequisePour(p.Niveau)
	}

	// Si on atteint le cap, on remet l'XP à 0 (plus clair visuellement)
	if p.Niveau >= p.NiveauMax {
		p.ExpActuelle = 0
		p.ExpMax = 0
		return nivGagnes, true
	}
	return nivGagnes, false
}

// Applique les effets d'une montée de niveau (bonus de stats)
func (p *Personnage) onLevelUp() {
	// Exemple de règle : +10 PV Max par niveau (en bonus)
	oldMax := p.PVMax
	p.BonusPVMax += 10
	p.RecalculerPVMax()

	// Soigne de la différence gagnée (optionnel, très classique)
	diff := p.PVMax - oldMax
	if diff > 0 {
		p.PVActuels += diff
		if p.PVActuels > p.PVMax {
			p.PVActuels = p.PVMax
		}
	}
}

// Recalcule PVMax = PVBase + BonusPVMax et clamp PVActuels
func (p *Personnage) RecalculerPVMax() {
	p.PVMax = p.PVBase + p.BonusPVMax
	if p.PVMax < 1 {
		p.PVMax = 1
	}
	if p.PVActuels > p.PVMax {
		p.PVActuels = p.PVMax
	}
	if p.PVActuels < 0 {
		p.PVActuels = 0
	}
}

// Barre d'XP ASCII (ex: [██████░░░░░░] 120/300 (40%)  Niv 3/20)
func (p *Personnage) BarreXP(width int) string {
	if width < 3 {
		width = 3
	}
	fill := '█'
	empty := '░'

	if p.Niveau >= p.NiveauMax {
		// Barre pleine quand cap
		return fmt.Sprintf("[%s] MAX  Niv %d/%d", strings.Repeat(string(fill), width), p.Niveau, p.NiveauMax)
	}

	// Fraction de progression
	f := 0.0
	if p.ExpMax > 0 {
		f = float64(p.ExpActuelle) / float64(p.ExpMax)
		if f < 0 {
			f = 0
		}
		if f > 1 {
			f = 1
		}
	}
	filled := int(math.Round(f * float64(width)))
	if filled < 0 {
		filled = 0
	}
	if filled > width {
		filled = width
	}

	percent := int(math.Round(f * 100))
	return fmt.Sprintf("[%s%s] %d/%d (%d%%)  Niv %d/%d",
		strings.Repeat(string(fill), filled),
		strings.Repeat(string(empty), width-filled),
		p.ExpActuelle, p.ExpMax, percent, p.Niveau, p.NiveauMax,
	)
}

// Afficher : résumé lisible
func (p Personnage) Afficher() string {
	var b strings.Builder
	fmt.Fprintf(&b, "Nom : %s\nRace : %s\nClasse : %s\nSexe : %s\n", p.Nom, p.Race, p.Classe, p.Sexe)
	fmt.Fprintf(&b, "PV : %d/%d (Base:%d + Bonus:%d)\n",
		p.PVActuels, p.PVMax, p.PVBase, p.BonusPVMax)

	// XP / Niveau
	if p.Niveau >= p.NiveauMax {
		fmt.Fprintf(&b, "Niveau : %d (CAP %d) – Expérience : MAX\n", p.Niveau, p.NiveauMax)
	} else {
		fmt.Fprintf(&b, "Niveau : %d/%d – Expérience : %d/%d\n", p.Niveau, p.NiveauMax, p.ExpActuelle, p.ExpMax)
	}
	fmt.Fprintf(&b, "%s\n", p.BarreXP(24))

	return b.String()
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
	fmt.Println("1. Humain (adapté à la classe guerrier)")
	fmt.Println("2. Elfe (adapté à la classe mage)")
	fmt.Println("3. Nain (adapté à la classe assassin)")
	fmt.Println()

	n := pz.lireChoixInt("Choisissez votre race (1, 2 ou 3) : ", 1, 3)
	r := Race(n)
	fmt.Printf("Vous avez choisi la race %s.\n", r)
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
	p.InitPV()      // initialise PVBase/PVMax/PVActuels
	p.InitXP(1, 20) // niveau 1 -> 20 par défaut

	// Affichage PV selon la race (équivalent de ton PointdeVie)
	switch p.Race {
	case Humain:
		fmt.Println("Vous avez 50 points de vie actuellement, et vous pouvez en avoir jusqu'à 100 étant humain (base, hors bonus).")
	case Elfe:
		fmt.Println("Vous avez 40 points de vie actuellement, et vous pouvez en avoir jusqu'à 80 étant elfe (base, hors bonus).")
	case Nain:
		fmt.Println("Vous avez 60 points de vie actuellement, et vous pouvez en avoir jusqu'à 120 étant nain (base, hors bonus).")
	}

	// Petit aperçu XP
	fmt.Printf("Niveau initial : %d/%d – XP : %d/%d\n", p.Niveau, p.NiveauMax, p.ExpActuelle, p.ExpMax)
	fmt.Println(p.BarreXP(24))

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
