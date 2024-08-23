package fraction_addition_and_subtraction

import (
	"fmt"
	"strconv"
)

type Fraction struct {
	Up   int
	Down int
}

func NewFraction(expressions []rune) *Fraction {
	fraction := &Fraction{}
	number := make([]rune, 0)
	for _, item := range expressions {
		if item == '/' {
			intNumber, err := strconv.Atoi(string(number))
			if err != nil {
				intNumber = 0
			}
			fraction.Up = intNumber
			number = make([]rune, 0)
		} else {
			number = append(number, item)
		}
	}
	intNumber, err := strconv.Atoi(string(number))
	if err != nil {
		intNumber = 0
	}
	fraction.Down = intNumber
	return fraction
}

func (f *Fraction) ToString() string {
	if f.Up == 0 {
		return "0/1"
	}
	return fmt.Sprintf("%d/%d", f.Up, f.Down)
}

func minimum(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func calc(a *Fraction, b *Fraction, op rune) *Fraction {
	aCount := 1
	bCount := 1
	for a.Down*aCount != b.Down*bCount {
		if a.Down*aCount < b.Down*bCount {
			aCount++
		} else {
			bCount++
		}
	}
	a.Up *= aCount
	b.Up *= bCount
	if op == '+' {
		b.Up += a.Up
	} else {
		b.Up = a.Up - b.Up
	}
	b.Down *= bCount
	canDiv := true
	for canDiv {
		canDiv = false
		for i := 2; i <= abs(minimum(b.Up, b.Down)); i++ {
			if b.Up%i == 0 && b.Down%i == 0 {
				b.Up /= i
				b.Down /= i
				canDiv = true
				break
			}
		}
	}
	return b
}

func fractionAddition(expression string) string {
	expressions := []rune(expression)
	fractionNumbers := make([]*Fraction, 0)
	operations := make([]rune, 0)
	number := make([]rune, 0)
	hasMinus := false
	for _, item := range expressions {
		if item == '+' || item == '-' {
			if len(number) > 0 {
				fraction := NewFraction(number)
				if hasMinus {
					hasMinus = false
					fraction.Up *= -1
				}
				fractionNumbers = append(fractionNumbers, fraction)
				number = make([]rune, 0)
				operations = append(operations, item)
			} else {
				hasMinus = true
			}
		} else {
			number = append(number, item)
		}
	}
	fraction := NewFraction(number)
	if hasMinus {
		hasMinus = false
		fraction.Up *= -1
	}
	fractionNumbers = append(fractionNumbers, fraction)

	answer := fractionNumbers[0]
	for i := 1; i < len(fractionNumbers); i++ {
		answer = calc(answer, fractionNumbers[i], operations[i-1])
	}
	return answer.ToString()
}
