package day13

type axis rune

const (
	axisx axis = 'x'
	axisy axis = 'y'
)

type instruction struct {
	Axis  axis
	Value int
}
