package models

import (
	"errors"
	"strings"
)

type Line int

const (
	Red Line = iota
	Blue
	Brown
	Green
	Orange
	Purple
	Pink
	Yellow
)

func (receiver *Line) String() string {
	return [...]string{"Red", "Blue", "Brown", "Green", "Orange", "Purple", "Pink", "Yellow"}[*receiver]
}

func (receiver *Line) MarshalJSON() ([]byte, error) {
	return []byte(`"` + receiver.String() + `"`), nil
}

func (receiver *Line) UnmarshalJSON(data []byte) error {
	lineString := string(data)

	lineString = strings.Trim(lineString, "\"")

	lineString = strings.ToUpper(lineString)

	switch lineString {
	case "RED", "RED LINE":
		*receiver = Red
	case "BLUE", "BLUE LINE":
		*receiver = Blue
	case "BRN", "BROWN LINE":
		*receiver = Brown
	case "G", "GREEN LINE":
		*receiver = Green
	case "ORG", "ORANGE LINE":
		*receiver = Orange
	case "P", "PURPLE LINE":
		*receiver = Purple
	case "PNK", "PINK LINE":
		*receiver = Pink
	case "Y", "YELLOW LINE":
		*receiver = Yellow
	default:
		return errors.New("invalid line")
	}

	return nil
}
