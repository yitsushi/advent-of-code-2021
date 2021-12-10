package day10

import "fmt"

// UnknownCharacterError occurs when an unknown character was found ina
// sequence.
type UnknownCharacterError struct {
	Character rune
}

func (e UnknownCharacterError) Error() string {
	return fmt.Sprintf("unknown character: %c", e.Character)
}
