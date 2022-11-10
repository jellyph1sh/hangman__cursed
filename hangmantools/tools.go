package hangmantools

import "github.com/nsf/termbox-go"

func ClearAll(x, y, length, height int) {
	for i := 0; i <= length; i++ {
		for j := 0; j <= height; j++ {
			termbox.SetCell(x+i, y+j, ' ', termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

func ClearRect(x, y, length, height int) {
	// ┌╌╌╌╌╌╌╌╌┐
	// │		╎
	// ╎		╎
	// └╌╌╌╌╌╌╌╌┘
	termbox.SetCell(x, y, ' ', termbox.ColorWhite, termbox.ColorDefault)
	termbox.SetCell(x, y+height, ' ', termbox.ColorWhite, termbox.ColorDefault)
	for i := 0; i < length; i++ {
		termbox.SetCell(x+i+1, y, ' ', termbox.ColorWhite, termbox.ColorDefault)
		termbox.SetCell(x+i+1, y+height, ' ', termbox.ColorWhite, termbox.ColorDefault)
	}
	termbox.SetCell(x+length, y, ' ', termbox.ColorWhite, termbox.ColorDefault)
	termbox.SetCell(x+length, y+height, ' ', termbox.ColorWhite, termbox.ColorDefault)
	for i := 1; i < height; i++ {
		termbox.SetCell(x, y+i, ' ', termbox.ColorWhite, termbox.ColorDefault)
		termbox.SetCell(x+length, y+i, ' ', termbox.ColorWhite, termbox.ColorDefault)
	}
}

func ConcatHangmanParts(splitHangmanFile []string) []string {
	var hangmanParts []string
	strHangmanPart := ""
	hangmanParts = append(hangmanParts, "")
	for i := 0; i < len(splitHangmanFile); i++ {
		if splitHangmanFile[i] == "" {
			hangmanParts = append(hangmanParts, strHangmanPart)
			strHangmanPart = ""
		} else {
			strHangmanPart += splitHangmanFile[i] + "\n"
		}
	}
	return hangmanParts
}

func ConcatLetters(splitLetters []string) [][]string {
	var (
		asciiLetters    [][]string
		asciiLettersTab []string
	)
	for i := 0; i < len(splitLetters); i++ {
		if i%9 == 0 {
			asciiLetters = append(asciiLetters, asciiLettersTab)
			asciiLettersTab = nil
		} else {
			asciiLettersTab = append(asciiLettersTab, splitLetters[i])
		}
	}
	return asciiLetters
}

func CreateRect(x, y, length, height int, color termbox.Attribute) {
	// ┌╌╌╌╌╌╌╌╌┐
	// │		╎
	// ╎		╎
	// └╌╌╌╌╌╌╌╌┘
	termbox.SetCell(x, y, '┌', color, termbox.ColorDefault)
	termbox.SetCell(x, y+height, '└', color, termbox.ColorDefault)
	for i := 0; i < length; i++ {
		termbox.SetCell(x+i+1, y, '─', color, termbox.ColorDefault)
		termbox.SetCell(x+i+1, y+height, '─', color, termbox.ColorDefault)
	}
	termbox.SetCell(x+length, y, '┐', color, termbox.ColorDefault)
	termbox.SetCell(x+length, y+height, '┘', color, termbox.ColorDefault)
	for i := 1; i < height; i++ {
		termbox.SetCell(x, y+i, '│', color, termbox.ColorDefault)
		termbox.SetCell(x+length, y+i, '│', color, termbox.ColorDefault)
	}
}

func InsertAtIndex(word, letter string, index int) string { // Replace a letter at an index given
	return word[:index] + letter + word[index+1:]
}

func IsInputIsAlphabet(input string) bool {
	if (input[0] < 'a' || input[0] > 'z') && (input[0] < 'A' || input[0] > 'Z') {
		return false
	}
	return true
}

func IsInArray(list []int, value int) bool {
	for _, i := range list {
		if i == value {
			return true
		}
	}
	return false
}

func IsLetterAlreadyUse(usedLetters []string, inputLetter string) bool {
	for _, letter := range usedLetters {
		if letter == inputLetter {
			return true
		}
	}
	return false
}

func IsUppercase(inputLetter string) bool {
	for _, i := range inputLetter {
		if i < 'A' || i > 'Z' {
			return false
		}
	}
	return true
}

func PrintWord(word string, x, y int, color termbox.Attribute) {
	var backspace int = 0
	for index, letter := range word {
		if letter == '\n' {
			y++
			backspace = (index + 1) * (-1)
		}
		termbox.SetCell(x+index+backspace, y, letter, color, termbox.ColorDefault)
	}
	termbox.Flush()
}

func RemoveCharAtIndex(word string, index int) string {
	return word[:index] + word[index+1:]
}

func ToUppercase(characters string) string {
	for i := 0; i < len(characters); i++ {
		if !IsUppercase(string(characters[i])) {
			characters = InsertAtIndex(characters, string(characters[i]-32), i)
		}
	}
	return characters
}
