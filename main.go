package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"os"
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

var sides = []rune{'F', 'B', 'L', 'R', 'U', 'D'}

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
	gin.SetMode(gin.ReleaseMode)
	rand.Seed(time.Now().UnixMicro())
	r := gin.Default()
	r.LoadHTMLFiles("index.gohtml")
	r.GET("/scramble", func(c *gin.Context) {
		moves, err := RandomMoves(20)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		res := make([]string, len(moves))
		for i := range res {
			res[i] = moves[i].String()
		}
		c.JSON(http.StatusOK, res)
	})
	r.GET("/", func(c *gin.Context) {
		moves, err := RandomMoves(20)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		res := make([]string, len(moves))
		for i := range res {
			res[i] = moves[i].String()
		}
		c.HTML(http.StatusOK, "index.gohtml", gin.H{
			"scramble": strings.Join(res, " "),
		})
	})
	port := "8080"
	if len(os.Args) > 2 {
		fmt.Println("USAGE: ./cubiks <port>")
		fmt.Println("       ./cubiks (defaults to 8080)")
		return
	}
	if len(os.Args) == 2 {
		port = os.Args[1]
	}
	err := r.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
