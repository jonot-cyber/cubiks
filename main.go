package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

const (
	X = iota
	Y = iota
	Z = iota
)

type Move struct {
	Side  rune
	Axis  int
	Count int
}

func (m Move) String() string {
	var b strings.Builder
	b.WriteRune(m.Side)
	switch m.Count {
	case 1: // do nothing
	case 2:
		b.WriteRune('2')
	case 3:
		b.WriteRune('\'')
	}
	return b.String()
}

var sides []rune = []rune{'F', 'B', 'L', 'R', 'U', 'D'}

func SideAxis(s rune) (int, error) {
	switch s {
	case 'L', 'R':
		return X, nil
	case 'U', 'D':
		return Y, nil
	case 'F', 'B':
		return Z, nil
	}
	return 0, errors.New("invalid rune")
}

func RandomBool() bool {
	return rand.Intn(2) == 1
}

func RandomMove(lastMove Move) (Move, error) {
	var availableSides []rune
	defaultMove := Move{}
	if lastMove == defaultMove {
		availableSides = sides
	} else {
		for _, v := range sides {
			x, err := SideAxis(v)
			if err != nil {
				return Move{}, err
			}
			if x != lastMove.Axis {
				availableSides = append(availableSides, v)
			}
		}
	}
	side := availableSides[rand.Intn(len(availableSides))]
	axis, err := SideAxis(side)
	if err != nil {
		return Move{}, err
	}
	return Move{
		Side:  side,
		Axis:  axis,
		Count: rand.Intn(3) + 1,
	}, nil
}

func RandomMoves(count int) ([]Move, error) {
	res := make([]Move, count)
	lastMove := Move{}
	for i := 0; i < count; i++ {
		move, err := RandomMove(lastMove)
		lastMove = move
		if err != nil {
			return nil, err
		}
		res[i] = move
	}
	return res, nil
}

func main() {
	defer func(start time.Time) {
		fmt.Println(time.Since(start))
	}(time.Now())
	rand.Seed(time.Now().UnixMicro())
	for i := 0; i < 1_000_000; i++ {
		moves, err := RandomMoves(20)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(moves)
	}
}
