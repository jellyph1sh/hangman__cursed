package hangman

import (
	"fmt"
	"hangmantools"
	"os"
	"strings"

	"github.com/nsf/termbox-go"
)

type TextColor struct {
	Reset,
	Red,
	Green,
	Yellow,
	Blue,
	Purple,
	Cyan,
	Gray,
	White string
}

func Display(hangman *Hangman) {
	DisplayWord(hangman)
	DisplayUsedLetters(hangman)
	fmt.Print("\n")
}

func DisplayRandomLetter(hangman *Hangman) {
	for i := 0; i < len(hangman.SecretWord); i++ {
		if hangmantools.IsInArray(hangman.LetterIndexes, i) {
			hangman.UserWord += string(hangman.SecretWord[i])
		} else {
			hangman.UserWord += "_"
		}
	}
}

func DisplayUsedLetters(hangman *Hangman) {
	var str string = ""
	for _, letter := range hangman.UsedLetters {
		str += letter + " "
	}
	hangmantools.PrintWord(str, 5, 20, termbox.ColorRed)
}

func DisplayWinLose(hangman *Hangman) {
	hangmantools.ClearAll(0, 4, 29, 5)
	if hangman.Found {
		hangmantools.CreateRect(0, 4, 29, 5, termbox.ColorGreen)
		hangmantools.PrintWord(hangman.Lang.Win, 3, 6, termbox.ColorGreen)
	} else {
		hangmantools.CreateRect(0, 4, 29, 5, termbox.ColorRed)
		hangmantools.PrintWord(hangman.Lang.Lose, 3, 6, termbox.ColorRed)
		hangmantools.PrintWord(hangman.SecretWord, (29-len(hangman.SecretWord))/2, 12, termbox.ColorWhite)
	}
	termbox.Flush()
	termbox.PollEvent()
}

func DisplayWord(hangman *Hangman) {
	var lettersPrint [][]string
	var word string
	if hangman.Attempts == 0 || hangman.Found {
		word = hangman.SecretWord
	} else {
		word = hangman.UserWord
	}
	for _, char := range word {
		if char == '_' {
			tmp := hangman.AsciiLetters[int(char)-32][2:]
			tmp = append(tmp, "         \n")
			tmp = append(tmp, "         ")
			lettersPrint = append(lettersPrint, tmp)
		} else {
			lettersPrint = append(lettersPrint, hangman.AsciiLetters[int(char)-32])
		}
	}
	for line := 0; line < 8; line++ {
		for letter := 0; letter < len(lettersPrint); letter++ {
			fmt.Print(hangman.Colors.Cyan + strings.Replace(lettersPrint[letter][line], "\r", "", -1) + hangman.Colors.Reset)
		}
		fmt.Println()
	}
}

func DisplayHangmanMenu(hangman *Hangman) {
	hangmantools.PrintWord("Game", 34, 2, termbox.ColorRed)
	hangmantools.CreateRect(0, 4, 29, 5, termbox.ColorCyan)
	hangmantools.PrintWord(hangman.Lang.ChooseLetter, 7, 4, termbox.ColorCyan)
	hangmantools.CreateRect(30, 4, 39, 20, termbox.ColorRed)
	hangmantools.PrintWord(hangman.Lang.Hangman, 47-(len(hangman.Lang.Hangman)-7)/2, 4, termbox.ColorRed) // Calc the modification value to be center
	hangmantools.PrintWord(hangman.Lang.Attempts, 30+(11-len(hangman.Lang.Attempts))/2, 4, termbox.ColorRed)
	hangmantools.CreateRect(0, 10, 29, 5, termbox.ColorGreen)
	hangmantools.PrintWord(hangman.Lang.SecretWord, 8, 10, termbox.ColorGreen)
	hangmantools.CreateRect(0, 16, 29, 8, termbox.ColorYellow)
	hangmantools.PrintWord(hangman.Lang.LettersUsed, (29-len(hangman.Lang.LettersUsed))/2, 16, termbox.ColorYellow)
	UpdateHangmanMenu(hangman)
}

func DisplayHelpMenu(hangman *Hangman) {
	hangmantools.PrintWord("Help", 41, 2, termbox.ColorRed)
	hangmantools.CreateRect(0, 4, 50, 6, termbox.ColorWhite)
	hangmantools.PrintWord(hangman.Lang.Help, 4, 6, termbox.ColorWhite)
}

func DisplayLanguages(hangman *Hangman) {
	hangmantools.CreateRect(0, 0, 50, 21, termbox.ColorCyan)
	hangmantools.PrintWord(hangman.Lang.LangChoose, 14, 5, termbox.ColorWhite)
	hangmantools.PrintWord("Francais", 10, 10, termbox.ColorWhite)
	hangmantools.PrintWord("English", 32, 10, termbox.ColorWhite)
	hangmantools.CreateRect(9, 9, 9, 2, termbox.ColorGreen)
}

func DisplayScreens(hangman *Hangman) {
	hangmantools.ClearAll(0, 0, 100, 100)
	hangmantools.PrintWord("ESC to quit, LEFT or RIGHT to switch tabs", 2, 0, termbox.ColorRed)
	hangmantools.CreateRect(22, 1, 24, 2, termbox.ColorWhite)
	hangmantools.PrintWord("Welcome", 24, 2, termbox.ColorWhite)
	hangmantools.PrintWord(" | ", 31, 2, termbox.ColorWhite)
	hangmantools.PrintWord("Game", 34, 2, termbox.ColorWhite)
	hangmantools.PrintWord(" | ", 38, 2, termbox.ColorWhite)
	hangmantools.PrintWord("Help", 41, 2, termbox.ColorWhite)
	switch hangman.Menu {
	case "Game":
		DisplayHangmanMenu(hangman)
	case "Welcome":
		DisplayWelcomeMenu(hangman)
	case "Help":
		DisplayHelpMenu(hangman)
	}

}

func DisplayWelcomeMenu(hangman *Hangman) {
	hangmantools.PrintWord("Welcome", 24, 2, termbox.ColorRed)
	hangmantools.CreateRect(14, 4, 40, 5, termbox.ColorBlue)
	hangmantools.PrintWord(hangman.Lang.Welcome, 18, 6, termbox.ColorWhite)
}

func GetAsciiTable(hangman *Hangman, file string) error {
	var asciiLetters [][]string
	fileContent, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	strFileContent := string(fileContent)
	splitFileContent := strings.Split(strFileContent, "\n")
	asciiLetters = hangmantools.ConcatLetters(splitFileContent)
	hangman.AsciiLetters = asciiLetters
	return nil
}

func GetHangmanParts(hangman *Hangman, file string) error {
	var hangmanParts []string
	hangmanFileContent, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	strHangmanFileContent := string(hangmanFileContent)
	splitHangmanFileContent := strings.Split(strHangmanFileContent, "\n")
	hangmanParts = hangmantools.ConcatHangmanParts(splitHangmanFileContent)
	hangman.Positions = hangmanParts
	return nil
}

func InitDisplay(hangman *Hangman) bool {
	// Initialize Hangman :
	hangmanErr := GetHangmanParts(hangman, "../dependencies/hangman.txt")
	// Check if no error :
	if hangmanErr != nil {
		fmt.Println(hangmanErr)
		return false
	}
	return true
}

func UpdateHangmanMenu(hangman *Hangman) {
	hangmantools.PrintWord(hangman.UserWord, (29-len(hangman.SecretWord))/2, 12, termbox.ColorWhite)
	if hangman.Attempts == 10 {
		hangmantools.PrintWord("10", 34, 6, termbox.ColorWhite)
	} else {
		termbox.SetChar(34, 6, ' ')
		hangmantools.PrintWord(string(rune(hangman.Attempts)+'0'), 35, 6, termbox.ColorWhite)
	}
	hangmantools.PrintWord(hangman.Positions[10-hangman.Attempts], 41, 9, termbox.ColorWhite)
	DisplayUsedLetters(hangman)
}
