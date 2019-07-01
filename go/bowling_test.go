package bowling_test

import (
	"bowling"
	"fmt"
	"testing"
)
func TestBownlingUnder10(t *testing.T)  {
	var shouldBeValue int
	var shouldBeString string
	var game *bowling.Game

	for i := 0; i < 10; i++ {
		for j := 0; j < (10 - i); j++ {
			shouldBeValue = 0
			shouldBeString = ""
			game = &bowling.Game{}

			for k := 0; k < 10; k++ {
				shouldBeValue += i + j
				shouldBeString += fmt.Sprintf("%d%d,", i, j)
				
				game.AddFrame(i, j)
			}
			// Remove last ","
			shouldBeString = shouldBeString[:len(shouldBeString)-1]

			asIsValue := game.CalculateScore()
			if shouldBeValue != asIsValue {
				t.Error("Expected:", shouldBeValue, "Result:", asIsValue)
			}

			asIsString := game.Show()
			if shouldBeString != asIsString {
				t.Error("Expected:", shouldBeString, "Result:", asIsString)
			}
		}
	}
}
func TestBownling11(t *testing.T)  {
	shouldBeValue := 11
	shouldBeString := "00,00,00,00,00,00,00,00,00,1/1"

	game := &bowling.Game{}
	game.AddMiss().AddMiss().AddMiss().AddMiss().AddMiss().AddMiss().AddMiss().AddMiss().AddMiss().AddFrameExtra(1, 9, 1)
	
	asIsValue := game.CalculateScore()
	if shouldBeValue != asIsValue {
		t.Error("Expected:", shouldBeValue, "Result:", asIsValue)
	}

	asIsString := game.Show()
	if shouldBeString != asIsString {
		t.Error("Expected:", shouldBeString, "Result:", asIsString)
	}
}

func TestBownling20(t *testing.T)  {
	shouldBeValue := 20
	shouldBeString := "00,00,00,00,00,00,00,00,00,X1/"

	game := &bowling.Game{}
	game.AddMiss().AddMiss().AddMiss().AddMiss().AddMiss().AddMiss().AddMiss().AddMiss().AddMiss().AddFrameExtra(10, 1, 9)
	
	asIsValue := game.CalculateScore()
	if shouldBeValue != asIsValue {
		t.Error("Expected:", shouldBeValue, "Result:", asIsValue)
	}

	asIsString := game.Show()
	if shouldBeString != asIsString {
		t.Error("Expected:", shouldBeString, "Result:", asIsString)
	}
}

func TestBownling270(t *testing.T)  {
	shouldBeValue := 270
	shouldBeString := "X,X,X,X,X,X,X,X,X,9/1"

	game := &bowling.Game{}
	game.AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddFrameExtra(9, 1, 1)
	
	asIsValue := game.CalculateScore()
	if shouldBeValue != asIsValue {
		t.Error("Expected:", shouldBeValue, "Result:", asIsValue)
	}

	asIsString := game.Show()
	if shouldBeString != asIsString {
		t.Error("Expected:", shouldBeString, "Result:", asIsString)
	}
}

func TestBownling271(t *testing.T)  {
	shouldBeValue := 271
	shouldBeString := "X,X,X,X,X,X,X,X,X,1/X"

	game := &bowling.Game{}
	game.AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddFrameExtra(1, 9, 10)
	
	asIsValue := game.CalculateScore()
	if shouldBeValue != asIsValue {
		t.Error("Expected:", shouldBeValue, "Result:", asIsValue)
	}

	asIsString := game.Show()
	if shouldBeString != asIsString {
		t.Error("Expected:", shouldBeString, "Result:", asIsString)
	}
}

func TestBownling300(t *testing.T)  {
	shouldBeValue := 300
	shouldBeString := "X,X,X,X,X,X,X,X,X,XXX"

	game := &bowling.Game{}
	game.AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddFrameExtra(10, 10, 10)
	
	asIsValue := game.CalculateScore()
	if shouldBeValue != asIsValue {
		t.Error("Expected:", shouldBeValue, "Result:", asIsValue)
	}

	asIsString := game.Show()
	if shouldBeString != asIsString {
		t.Error("Expected:", shouldBeString, "Result:", asIsString)
	}
}

func TestBownling10AsStrike(t *testing.T)  {
	shouldBeValue := 10
	shouldBeString := "X"

	game := &bowling.Game{}
	game.AddStrike()
	
	asIsValue := game.CalculateScore()
	if shouldBeValue != asIsValue {
		t.Error("Expected:", shouldBeValue, "Result:", asIsValue)
	}

	asIsString := game.Show()
	if shouldBeString != asIsString {
		t.Error("Expected:", shouldBeString, "Result:", asIsString)
	}
}

func TestBownling10AsSpare(t *testing.T)  {
	shouldBeValue := 10
	shouldBeString := "7/"

	game := &bowling.Game{}
	game.AddSpare(7)
	
	asIsValue := game.CalculateScore()
	if shouldBeValue != asIsValue {
		t.Error("Expected:", shouldBeValue, "Result:", asIsValue)
	}

	asIsString := game.Show()
	if shouldBeString != asIsString {
		t.Error("Expected:", shouldBeString, "Result:", asIsString)
	}
}

func TestBownling245(t *testing.T)  {
	shouldBeValue := 245
	shouldBeString := "X,X,X,X,X,X,X,X,X,11"

	game := &bowling.Game{}
	game.AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddFrame(1, 1)
	
	asIsValue := game.CalculateScore()
	if shouldBeValue != asIsValue {
		t.Error("Expected:", shouldBeValue, "Result:", asIsValue)
	}

	asIsString := game.Show()
	if shouldBeString != asIsString {
		t.Error("Expected:", shouldBeString, "Result:", asIsString)
	}
}

func TestBownling273(t *testing.T)  {
	shouldBeValue := 273
	shouldBeString := "X,X,X,X,X,X,X,X,X,X11"

	game := &bowling.Game{}
	game.AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddFrameExtra(10, 1, 1)
	
	asIsValue := game.CalculateScore()
	if shouldBeValue != asIsValue {
		t.Error("Expected:", shouldBeValue, "Result:", asIsValue)
	}

	asIsString := game.Show()
	if shouldBeString != asIsString {
		t.Error("Expected:", shouldBeString, "Result:", asIsString)
	}
}

func TestBownling15(t *testing.T)  {
	shouldBeValue := 15
	shouldBeString := "11,1/,11"

	game := &bowling.Game{}
	game.AddFrame(1, 1).AddSpare(1).AddFrame(1, 1)
	
	asIsValue := game.CalculateScore()
	if shouldBeValue != asIsValue {
		t.Error("Expected:", shouldBeValue, "Result:", asIsValue)
	}

	asIsString := game.Show()
	if shouldBeString != asIsString {
		t.Error("Expected:", shouldBeString, "Result:", asIsString)
	}
}