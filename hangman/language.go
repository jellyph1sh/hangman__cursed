package hangman

import (
	"hangmantools"

	"github.com/nsf/termbox-go"
)

type Language struct {
	Attempts,
	LangChoose,
	LangErr,
	Welcome,
	SecretWord,
	Hangman,
	LettersUsed,
	ChooseLetter,
	Help,
	SaveGame,
	SaveErr,
	LetterErr1,
	LetterErr2,
	LetterErr3,
	Win,
	Lose string
}

func GetLanguage(hangmanStruct *Hangman) {
	var ask bool = true
	DisplayLanguages(hangmanStruct)
	for ask {
		termbox.Flush()
		input := termbox.PollEvent()
		if input.Key == termbox.KeyArrowLeft {
			hangmantools.CreateRect(9, 9, 9, 2, termbox.ColorGreen)
			hangmantools.ClearRect(31, 9, 8, 2)
			hangmanStruct.LangTag = "FR"
		} else if input.Key == termbox.KeyArrowRight {
			hangmantools.CreateRect(31, 9, 8, 2, termbox.ColorGreen)
			hangmantools.ClearRect(9, 9, 9, 2)
			hangmanStruct.LangTag = "EN"
		} else if input.Key == termbox.KeyEnter {
			ask = false
			break
		} else if input.Key == termbox.KeyEsc {
			termbox.Close()

		}
	}
	termbox.Flush()
}

func SetLanguage(hangmanStruct *Hangman) {
	if hangmanStruct.LangTag == "FR" {
		hangmanStruct.Lang = Language{" Essais ", "Choisissez un langage", "Langage invalide [EN/FR]", "Bienvenue dans le jeux du Pendu,\nBon Courage a vous !", " Mot Secret ", " Pendu ", " Lettre utilisee ", " Mot ou lettre ", "1. Esc pour quitter\n2. Entrer pour confirmer votre choix\n3. Retour pour supprimer la derniere lettre", "Vous aves Bien sauvegarder,\nPour relancer avec votre sauvegarde utiliser -startWith", "Vous ne pouvez pas,\nvous avez deja une sauvegarde", "Saisie invalide", "Insérez un caractère valide!", "Lettre déjà utilisé!", "Bien Joué,\nVous avez gagner!", "Dommage,\nVous avez Perdu"}
	} else {
		hangmanStruct.Lang = Language{" Attemps ", "Choose a language", "Invalid language [EN/FR]", "Welcome in the Hangman\nGood Luck !", " Secret Word ", " Hangman ", " Used Letter ", " Word & Letter ", "1. Esc to quit\n2. Enter to confirm your choice\n3. Backspace to delete the last letter", "You have succesfully save you game,\nIf you want to play with your save use -startWith", "You can't,\nThere is an other save", "Empty input", "Please insert valid character!", "Letter already used!", "Good job,\nYou have win!", "Too bad,\nYou lost!"}
	}
}

func InitLanguage(hangmanStruct *Hangman) {
	// Setup temporary language :
	SetLanguage(hangmanStruct)
	// Ask player language :
	GetLanguage(hangmanStruct)
	// Set language of the game :
	SetLanguage(hangmanStruct)
}
