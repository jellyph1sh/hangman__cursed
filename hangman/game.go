package hangman

import (
	"encoding/json"
	"fmt"
	"hangmantools"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/nsf/termbox-go"
)

type Hangman struct {
	AsciiLetters  [][]string
	Attempts      int
	ChoosedLetter string
	Colors        TextColor
	Found         bool
	Lang          Language
	LangTag       string
	LetterIndexes []int
	Menu          string
	Positions     []string
	SecretWord    string
	Stop          bool
	UsedLetters   []string
	UserWord      string
	UserInput     string
	WordsFile     string
}

func CheckInputLetter(hangman *Hangman) bool {
	if !hangmantools.IsInputIsAlphabet(hangman.UserInput) {
		hangmantools.PrintWord(hangman.Lang.LetterErr2, 3, 5, termbox.ColorRed)
		return false
	}
	if !hangmantools.IsUppercase(hangman.UserInput) {
		hangman.UserInput = hangmantools.ToUppercase(hangman.UserInput)
	}
	if hangmantools.IsLetterAlreadyUse(hangman.UsedLetters, hangman.UserInput) {
		hangmantools.PrintWord(hangman.Lang.LetterErr3, 3, 5, termbox.ColorRed)
		return false
	}
	hangman.ChoosedLetter = hangman.UserInput
	return true
}

func CheckInputStop(hangman *Hangman) bool {
	if !hangmantools.IsInputIsAlphabet(hangman.UserInput) {
		hangmantools.PrintWord(hangman.Lang.LetterErr2, 3, 5, termbox.ColorRed)
		return false
	}
	if !hangmantools.IsUppercase(hangman.UserInput) {
		hangman.UserInput = hangmantools.ToUppercase(hangman.UserInput)
	}
	if hangman.UserInput != "STOP" {
		return false
	}
	return true
}

func CheckInputWord(hangman *Hangman) bool {
	if !hangmantools.IsInputIsAlphabet(hangman.UserInput) {
		hangmantools.PrintWord(hangman.Lang.LetterErr2, 3, 5, termbox.ColorRed)
		return false
	}
	if !hangmantools.IsUppercase(hangman.UserInput) {
		hangman.UserInput = hangmantools.ToUppercase(hangman.UserInput)
	}
	if hangman.UserInput == hangman.SecretWord {
		hangman.Found = true
	}
	hangman.ChoosedLetter = ""
	return true
}

func GetGameParameters(args []string) (string, []string, []string) {
	if len(args) == 0 {
		return "", nil, nil
	}
	var params []string
	var paramsArgs []string
	for i := 0; i < len(args); i += 2 {
		if strings.HasPrefix(args[i], "--") {
			_, param, found := strings.Cut(args[i], "--")
			if found {
				params = append(params, param)
			}
		} else {
			return "Parameters error !", params, paramsArgs
		}
	}
	for i := 1; i < len(args); i += 2 {
		paramsArgs = append(paramsArgs, args[i])
	}
	return "", params, paramsArgs
}

func GetPlayerInput(hangman *Hangman) { // Function ask the user an letter in input
	hangmantools.ClearAll(3, 5, 24, 0)
	input := termbox.PollEvent()
	if input.Key == termbox.KeyArrowLeft && hangman.Menu != "Welcome" {
		switch hangman.Menu {
		case "Game":
			hangman.Menu = "Welcome"
		case "Help":
			hangman.Menu = "Game"
		}
		hangman.UserInput = ""
		DisplayScreens(hangman)
	} else if input.Key == termbox.KeyArrowRight && hangman.Menu != "Help" {
		switch hangman.Menu {
		case "Game":
			hangman.Menu = "Help"
		case "Welcome":
			hangman.Menu = "Game"
		}
		hangman.UserInput = ""
		DisplayScreens(hangman)
	} else if input.Key == termbox.KeyEsc {
		termbox.Close()
	}
	switch hangman.Menu {
	case "Game":
		if (input.Ch >= 'a' && input.Ch <= 'z') || (input.Ch >= 'A' && input.Ch <= 'Z') {
			if len(hangman.UserInput) <= 21 {
				hangman.UserInput += string(input.Ch)
				hangmantools.PrintWord(hangman.UserInput, 7, 6, termbox.ColorWhite)
			}
		} else if input.Key == termbox.KeyEnter && len(hangman.UserInput) > 0 {
			for i := 0; i < len(hangman.UserInput); i++ {
				termbox.SetChar(7+i, 6, ' ')
			}
			if len(hangman.UserInput) > 1 {
				if CheckInputStop(hangman) {
					hangman.Stop = true
				} else if !CheckInputWord(hangman) {
					hangman.UserInput = ""
				} else if !hangman.Found {
					hangman.Attempts -= 2
				}
			} else if len(hangman.UserInput) == 1 {
				if !CheckInputLetter(hangman) {
					hangman.UserInput = ""
				} else if !InsertLetter(hangman) {
					hangman.Attempts--
					hangman.UsedLetters = append(hangman.UsedLetters, hangman.UserInput)
				}
			} else {
				hangmantools.PrintWord(hangman.Lang.LetterErr2, 2, 6, termbox.ColorRed)
			}
			hangman.UserInput = ""
		} else if input.Key == termbox.KeyBackspace2 && len(hangman.UserInput) > 0 {
			hangman.UserInput = hangman.UserInput[:len(hangman.UserInput)-1]
			termbox.SetChar(7+len(hangman.UserInput), 6, ' ')
		}
		UpdateHangmanMenu(hangman)
	}
}

func InsertLetter(hangman *Hangman) bool {
	// Variable declaration
	find := false
	for index, letter := range hangman.SecretWord { // Check if the letter is in the word
		if letter == rune(hangman.ChoosedLetter[0]) { // Place the letter at the index
			hangman.UserWord = hangmantools.InsertAtIndex(hangman.UserWord, hangman.ChoosedLetter, index)
			find = true
		}
	}
	return find
}

func Save(hangman *Hangman) {
	isSave := SaveGame(hangman)
	hangmantools.ClearAll(0, 0, 100, 100)
	hangmantools.CreateRect(0, 0, 59, 7, termbox.ColorBlue)
	if isSave {
		hangmantools.PrintWord(hangman.Lang.SaveGame, 3, 3, termbox.ColorGreen)
	} else {
		hangmantools.PrintWord(hangman.Lang.SaveErr, 3, 3, termbox.ColorRed)
	}
	termbox.Flush()
	termbox.PollEvent()
}

func StartSavedGame(fileName string) *Hangman {
	var hangmanStruct Hangman
	jsonHangmanSave, err := os.ReadFile(fileName)
	if err != nil {
		return nil
	}
	err2 := json.Unmarshal(jsonHangmanSave, &hangmanStruct)
	if err2 != nil {
		return nil
	}
	return &hangmanStruct
}

func SaveGame(hangman *Hangman) bool {
	jsonHangmanSave, err := json.Marshal(hangman)
	if err != nil {
		return false
	}
	err2 := os.WriteFile("save.txt", jsonHangmanSave, 777)
	return err2 == nil
}

func SelectRandomWord(hangman *Hangman) error {
	words, err := os.ReadFile(hangman.WordsFile)
	if err != nil {
		return err
	}
	strWords := string(words)
	splitWords := strings.Split(strWords, "\n")
	rand.Seed(time.Now().UnixNano())
	word := string(splitWords[rand.Intn(len(splitWords)-1)])
	word = hangmantools.ToUppercase(word)
	hangman.SecretWord = word
	return nil
}

func SelectRandomLettersInWord(hangman *Hangman) {
	nLetters := len(hangman.SecretWord)/2 - 1
	hangman.LetterIndexes = make([]int, nLetters)
	word := hangman.SecretWord
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < nLetters; i++ {
		randNb := rand.Intn(len(word) - 1)
		hangman.LetterIndexes[i] = randNb
		word = hangmantools.RemoveCharAtIndex(word, randNb)
	}
}

func SetDefaultGameStructure(wordFileLocation string) (bool, *Hangman) {
	hangmanStruct := &Hangman{
		Attempts:  10,
		Found:     false,
		LangTag:   "FR",
		Stop:      false,
		WordsFile: wordFileLocation,
		Menu:      "Game",
	}
	err := SelectRandomWord(hangmanStruct)
	if err != nil {
		fmt.Println(err)
		return true, nil
	}
	return false, hangmanStruct
}

func InitGame(args []string) (bool, *Hangman) {
	errParams, params, paramsArgs := GetGameParameters(args)
	if errParams != "" {
		fmt.Println("Error during initialize game parameters!")
		return false, nil
	}
	isError, hangmanStruct := InitParams(params, paramsArgs)
	if isError {
		fmt.Println("Error during execute game parameters!")
		return false, nil
	}
	return true, hangmanStruct
}

func InitParams(params []string, paramsArgs []string) (bool, *Hangman) {
	for i := 0; i < len(params); i++ {
		switch params[i] {
		case "startWith":
			hangmanStruct := StartSavedGame(paramsArgs[i])
			if hangmanStruct == nil {
				return true, hangmanStruct
			}
			SetLanguage(hangmanStruct)
			return false, hangmanStruct
		case "wordsFile":
			isError, hangmanStruct := SetDefaultGameStructure(paramsArgs[i])
			return isError, hangmanStruct
		default:
			isError, hangmanStruct := SetDefaultGameStructure("../dependencies/words.txt")
			return isError, hangmanStruct
		}
	}
	isError, hangmanStruct := SetDefaultGameStructure("../dependencies/words.txt")
	return isError, hangmanStruct
}
