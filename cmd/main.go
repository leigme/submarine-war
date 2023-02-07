package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	submarine_war "github.com/leigme/submarine-war"
	"log"
)

func main() {
	game := submarine_war.NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
