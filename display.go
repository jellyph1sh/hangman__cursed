package hangman

import (
	"fmt"
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
		if IsInArray(hangman.LetterIndexes, i) {
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
	PrintWord(str, 5, 20, termbox.ColorRed)
}

func DisplayWinLose(hangman *Hangman) {
	ClearAll(0, 4, 29, 5)
	if hangman.Found {
		CreateRect(0, 4, 29, 5, termbox.ColorGreen)
		PrintWord(hangman.Lang.Win, 3, 6, termbox.ColorGreen)
	} else {
		CreateRect(0, 4, 29, 5, termbox.ColorRed)
		PrintWord(hangman.Lang.Lose, 3, 6, termbox.ColorRed)
		PrintWord(hangman.SecretWord, (29-len(hangman.SecretWord))/2, 12, termbox.ColorWhite)
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
	PrintWord("Game", 34, 2, termbox.ColorRed)
	CreateRect(0, 4, 29, 5, termbox.ColorCyan)
	PrintWord(hangman.Lang.ChooseLetter, 7, 4, termbox.ColorCyan)
	CreateRect(30, 4, 39, 20, termbox.ColorRed)
	PrintWord(hangman.Lang.Hangman, 47-(len(hangman.Lang.Hangman)-7)/2, 4, termbox.ColorRed) // Calc the modification value to be center
	PrintWord(hangman.Lang.Attempts, 30+(11-len(hangman.Lang.Attempts))/2, 4, termbox.ColorRed)
	CreateRect(0, 10, 29, 5, termbox.ColorGreen)
	PrintWord(hangman.Lang.SecretWord, 8, 10, termbox.ColorGreen)
	CreateRect(0, 16, 29, 8, termbox.ColorYellow)
	PrintWord(hangman.Lang.LettersUsed, (29-len(hangman.Lang.LettersUsed))/2, 16, termbox.ColorYellow)
	UpdateHangmanMenu(hangman)
}

func DisplayHelpMenu(hangman *Hangman) {
	PrintWord("Help", 41, 2, termbox.ColorRed)
	CreateRect(0, 4, 50, 6, termbox.ColorWhite)
	PrintWord(hangman.Lang.Help, 4, 6, termbox.ColorWhite)
}

func DisplayLanguages(hangman *Hangman) {
	CreateRect(0, 0, 50, 21, termbox.ColorCyan)
	PrintWord(hangman.Lang.LangChoose, 14, 5, termbox.ColorWhite)
	PrintWord("Francais", 10, 10, termbox.ColorWhite)
	PrintWord("English", 32, 10, termbox.ColorWhite)
	CreateRect(9, 9, 9, 2, termbox.ColorGreen)
}

func DisplayScreens(hangman *Hangman) {
	ClearAll(0, 0, 100, 100)
	PrintWord("ESC to quit, LEFT or RIGHT to switch tabs", 2, 0, termbox.ColorRed)
	CreateRect(22, 1, 24, 2, termbox.ColorWhite)
	PrintWord("Welcome", 24, 2, termbox.ColorWhite)
	PrintWord(" | ", 31, 2, termbox.ColorWhite)
	PrintWord("Game", 34, 2, termbox.ColorWhite)
	PrintWord(" | ", 38, 2, termbox.ColorWhite)
	PrintWord("Help", 41, 2, termbox.ColorWhite)
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
	PrintWord("Welcome", 24, 2, termbox.ColorRed)
	CreateRect(14, 4, 40, 5, termbox.ColorBlue)
	PrintWord(hangman.Lang.Welcome, 18, 6, termbox.ColorWhite)
}

func GetAsciiTable(hangman *Hangman, file string) error {
	var asciiLetters [][]string
	fileContent, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	strFileContent := string(fileContent)
	splitFileContent := strings.Split(strFileContent, "\n")
	asciiLetters = ConcatLetters(splitFileContent)
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
	hangmanParts = ConcatHangmanParts(splitHangmanFileContent)
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
	PrintWord(hangman.UserWord, (29-len(hangman.SecretWord))/2, 12, termbox.ColorWhite)
	if hangman.Attempts == 10 {
		PrintWord("10", 34, 6, termbox.ColorWhite)
	} else {
		termbox.SetChar(34, 6, ' ')
		PrintWord(string(rune(hangman.Attempts)+'0'), 35, 6, termbox.ColorWhite)
	}
	PrintWord(hangman.Positions[10-hangman.Attempts], 41, 9, termbox.ColorWhite)
	DisplayUsedLetters(hangman)
}
