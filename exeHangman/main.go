package main

import (
	"hangman"
	"os"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	args := os.Args[1:]
	isGameValid, hangmanStruct := hangman.InitGame(args)
	if isGameValid {
		hangman.InitLanguage(hangmanStruct)
		hangman.InitDisplay(hangmanStruct)
		if !hangmanStruct.Stop {
			hangman.SelectRandomLettersInWord(hangmanStruct)
			hangman.DisplayRandomLetter(hangmanStruct)
		} else {
			hangmanStruct.Stop = false
		}
		hangman.DisplayScreens(hangmanStruct)
		hangmanStruct.Menu = "Game"
		for hangmanStruct.Attempts > 0 && !hangmanStruct.Found && !hangmanStruct.Stop {
			termbox.Flush()
			hangman.GetPlayerInput(hangmanStruct)
			if hangmanStruct.UserWord == hangmanStruct.SecretWord {
				hangmanStruct.Found = true
			}
		}
		if hangmanStruct.Stop {
			hangman.Save(hangmanStruct)
		}
		if !hangmanStruct.Stop {
			hangman.DisplayWinLose(hangmanStruct)
		}
		termbox.Close()
	}
}
