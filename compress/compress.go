package compress

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// AAAZZZZEEED [A-Z]
// A3Z4E3D, nil

// AAAA4D
// , err

func Compress(data string) (string, error) {
	if data == "" {
		return "", errors.New("error")
	}

	var result strings.Builder
	var previousRune rune
	var counter int

	dataInRunes := []rune(data)

	for key, currentRune := range dataInRunes {
		if !unicode.IsLetter(currentRune) || unicode.IsNumber(currentRune) {
			return "", errors.New("error")
		}

		if currentRune == 0 || string(previousRune) != "" && unicode.IsDigit(previousRune) {
			return "", errors.New("error")
		}

		// если первый символ
		if key == 0 {
			counter++
		} else {
			// если кончились дубли
			if currentRune != previousRune {
				// если один такой символ
				if counter == 1 {
					result.WriteRune(previousRune)
				} else {
					result.WriteRune(previousRune)

					counterStr := strconv.Itoa(counter)
					result.WriteString(counterStr)

				}
				counter = 1

			} else {
				counter++
			}
		}

		previousRune = currentRune
	}

	result.WriteRune(previousRune)
	counterStr := strconv.Itoa(counter)
	result.WriteString(counterStr)

	return result.String(), nil

}
