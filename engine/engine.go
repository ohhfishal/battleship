package engine

import (
	"context"
	"fmt"
)

// No. 	Class of ship 	Size
// 1 	Carrier 	5
// 2 	Battleship 	4
// 3 	Destroyer 	3
// 4 	Submarine 	3
// 5 	Patrol Boat 	2

type Move int64
type Player interface {
	GetInitialShipPlacement(context.Context) []Ship
	Move(context.Context, PlayerView) (Move, error)
}

type Game struct {
	players       []Player
	playerViews   []PlayerView
	currentPlayer int
}

func NewGame(first, second Player) *Game {
	return &Game{
		players: []Player{first, second},
		playerViews: []PlayerView{
			{},
			{},
		},
	}
}

func (game *Game) Start(ctx context.Context) error {
	// TODO: Handle context closing
	for {
		select {
		case <-ctx.Done():
			// TODO: Handle better
			return nil
		default:
			cur := game.currentPlayer
			next := (cur + 1) % 2

			move, err := game.players[cur].Move(ctx, game.playerViews[cur])
			if err != nil {
				return fmt.Errorf(`getting move: %w`, err)
			} else if !game.playerViews[cur].Board.Valid(move) {
				return fmt.Errorf(`got invalid move: %d`, move)
			}

			hit := -1
			for i, ship := range game.playerViews[next].Board.Ships {
				if ship.HitBy(move) {
					hit = i
					break
				}
			}

			if hit == -1 {
				// TODO: Miss for both players
			} else {
				// TODO: Hit for both players
			}

			game.currentPlayer = next
		}
	}
}

type Ship struct {
	Name     string
	Location uint128
}

func (ship Ship) HitBy(move Move) bool {
	return true
}

type PlayerView struct {
	Board PlayerBoard
	Enemy EnemyBoard
}

type PlayerBoard struct {
	Ships   []Ship
	Guesses GuessBoard
}

func (board PlayerBoard) Valid(move Move) bool {
	return true
}

type EnemyBoard struct {
	Board GuessBoard
}
type GuessBoard struct {
	Hits   uint128
	Misses uint128
}
