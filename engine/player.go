package engine

import (
	"bufio"
	"context"
	"fmt"
	"io"
)

var _ Player = &ReaderPlayer{}

func NewHumanPlayer(stdin io.Reader, stdout io.Writer) Player {
	return ReaderPlayer{stdin: stdin, stdout: stdout}
}

type ReaderPlayer struct {
	stdin  io.Reader
	stdout io.Writer
}

func (player ReaderPlayer) GetInitialShipPlacement(ctx context.Context) []Ship {
	var ships []Ship
	return ships
}
func (player ReaderPlayer) Move(ctx context.Context, board PlayerView) (Move, error) {
	// TODO: Print a better board
	io.WriteString(player.stdout, "Make a move: ")
	for {
		line, err := ReadLine(ctx, player.stdin)
		if err != nil {
			return -1, fmt.Errorf(`reading from stdin: %w`, err)
		}
		io.WriteString(player.stdout, "Got: "+line)
		// TODO: Continue processing string to see if its a valid move
	}
	return 0, nil
}

func ReadLine(ctx context.Context, stdin io.Reader) (string, error) {
	type message struct {
		Line string
		Err  error
	}

	ch := make(chan message, 1)
	defer close(ch)

	go func(channel chan<- message) {
		var msg message
		msg.Line, msg.Err = bufio.NewReader(stdin).ReadString('\n')
		channel <- msg
	}(ch)

	select {
	case <-ctx.Done():
		return ``, ctx.Err()
	case msg := <-ch:
		return msg.Line, msg.Err
	}
}
