package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ohhfishal/battleship/engine"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	game := engine.NewGame(
		engine.NewHumanPlayer(os.Stdin, os.Stdout),
		engine.NewHumanPlayer(os.Stdin, os.Stdout),
	)

	if err := game.Start(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "running: %s", err)
	}

}
