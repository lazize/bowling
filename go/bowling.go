package bowling

import (
	"fmt"
	"strconv"
)

const (
	maxPin    int = 10
	maxFrames int = 10
)

type frameValue interface {
	Value() int
	FirstBall() int
	SecondBall() int
	Next() frameValue
	setNext(frameValue)
	IsStrike() bool
}

type frame struct {
	firstBall  int
	secondBall int

	next frameValue
}

type extraBallFrame struct {
	*frame
	extraBall int
}

// Game : Bownling game
type Game struct {
	head frameValue
	tail frameValue
}

// ### frame
func (f *frame) IsStrike() bool {
	return f.firstBall == maxPin
}

func (f *frame) IsSpare() bool {
	if f.IsStrike() {
		return false
	}
	return (f.firstBall + f.secondBall) == maxPin
}

func (f *frame) String() string {
	if f.IsStrike() {
		return "X"
	}
	if f.IsSpare() {
		return fmt.Sprintf("%v/", f.firstBall)
	}
	return fmt.Sprintf("%v%v", f.firstBall, f.secondBall)
}

func (f *frame) setNext(v frameValue) {
	f.next = v
}

func (f *frame) Next() frameValue {
	return f.next
}

func (f *frame) FirstBall() int {
	return f.firstBall
}

func (f *frame) SecondBall() int {
	return f.secondBall
}

func (f *frame) Value() int {
	// It uses next field, so all values need to comes from interface
	sum := f.FirstBall() + f.SecondBall()
	if f.Next() != nil {
		if f.IsSpare() {
			sum += f.Next().FirstBall()
		}
		if f.IsStrike() {
			sum += f.Next().FirstBall() + f.Next().SecondBall()
			if f.Next().Next() != nil && f.Next().IsStrike() {
				sum += f.Next().Next().FirstBall()
			}
		}
	}
	return sum
}

// ### extraBallFrame
func (e *extraBallFrame) Value() int {
	return e.firstBall + e.secondBall + e.extraBall
}

func (e *extraBallFrame) String() string {
	s := ""

	if e.firstBall == maxPin {
		s += "X"
	} else {
		s += strconv.Itoa(e.firstBall)
	}
	
	if e.secondBall == maxPin {
		s += "X"
	} else if (e.firstBall + e.secondBall) == maxPin {
		s += "/"
	} else {
		s += strconv.Itoa(e.secondBall)
	}

	if e.extraBall == maxPin {
		s += "X"
	} else if (e.firstBall + e.secondBall) != maxPin && (e.secondBall + e.extraBall) == maxPin {
		s += "/"
	} else {
		s += strconv.Itoa(e.extraBall)
	}
	return s
}


// #### Game
func (g *Game) add(frame frameValue) {
	if g.head == nil {
		g.head = frame
	} else {
		g.tail.setNext(frame)
	}
	g.tail = frame
}

func (g *Game) addFrameInternal(ball1, ball2 int) {
	g.add(&frame{ball1, ball2, nil})
}

func (g *Game) addFrameExtraInternal(ball1, ball2, extraBall int) {
	g.add(&extraBallFrame{&frame{ball1, ball2, nil}, extraBall})
}

// AddFrame : Add frame with value from two ball's
func (g *Game) AddFrame(ball1, ball2 int) *Game {
	g.addFrameInternal(ball1, ball2)
	return g
}

// AddFrameExtra : Add frame with value from three ball's. That should be the latest frame
func (g *Game) AddFrameExtra(ball1, ball2, extraBall int) *Game {
 	g.addFrameExtraInternal(ball1, ball2, extraBall)
	return g
}

// AddStrike : Add frame as strike
func (g *Game) AddStrike() *Game {
	g.addFrameInternal(10, 0)
	return g
}

// AddSpare : Add rame as spare, second ball value is calculated
func (g *Game) AddSpare(ball1 int) *Game {
	g.addFrameInternal(ball1, (maxPin - ball1))
	return g
}

// AddMiss : Add frame as missing both ball's
func (g *Game) AddMiss() *Game {
	g.addFrameInternal(0, 0)
	return g
}

// CalculateScore : 
func (g *Game) CalculateScore() int {
	var score int
	for e := g.head; e != nil; e = e.Next() {
		score += e.Value()
	}
	return score
}
// Show : 
func (g *Game) Show() string {
	var result string
	for e := g.head; e != nil; e = e.Next() {
		result += fmt.Sprintf("%v,", e)
	}
	// Remove last ","
	return result[:len(result)-1]
}
