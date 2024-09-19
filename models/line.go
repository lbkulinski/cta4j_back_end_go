package models

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

func (line Line) String() string {
    return [...]string{"Red", "Blue", "Brown", "Green", "Orange", "Purple", "Pink", "Yellow"}[line]
}
