package game

import (
	"errors"
	"fmt"
)

type Player struct {
	board *Board
}

// BOARD SECTION

var ErrInvalidCoordinates = errors.New("placing ships on invalid coordinates")

// Holds the user board information during the game
type Board struct {
	tiles [10][10]Tile
	ships map[string]*Ship
}

func (b *Board) PlaceShip(s *Ship) error {
	// Checks:
	//	If the ship is out of bound
	if err := s.ValidateShip(); err != nil {
		return fmt.Errorf("error placing ship: %w", err)
	}

	//	If the ship doesn't overlap other ships
	//	If the adjacent tiles doesn't have other ship
	x := s.startTile.xCoordinate
	y := s.startTile.yCoordinate
	sLen := s.length
	o := s.orientaion

	var h, v int
	switch o {
	case Horizontal:
		h = sLen + 2
		v = 3
	case Vertical:
		h = 3
		v = sLen + 2
	}

	checkX := x - 1
	checkY := y - 1
	for range h {
		for range v {
			if b.tiles[checkY][checkX].occupied != nil {
				return ErrInvalidCoordinates
			}

			checkX = checkX + 1
		}

		checkY = checkY + 1
	}

	// Populate the tile
	for range sLen {
		b.tiles[y][x].occupied = s

		if o == Horizontal {
			x = x + 1
		} else {
			y = y + 1
		}
	}

	return nil
}

type Tile struct {
	isHit    bool
	occupied *Ship
}

// SHIP SECTION

var (
	ErrAlreadyDestroyed = errors.New("ship already destroyed")
	ErrShipOutOfBound   = errors.New("ship is out of bound")
)

// TODO: Generate ships with coordinates.
func CreateShips() []*Ship {
	for {
	}
}

type Ship struct {
	class      ShipClass
	health     int
	length     int
	isAlive    bool
	startTile  Coordinate
	orientaion Orientation
}

// Checks if the ship coordinate is within the 10x10 board
// Returns an error if invalid, nil otherwise
func (s *Ship) ValidateShip() error {
	var tNum int

	switch s.orientaion {
	case Horizontal:
		tNum = s.startTile.xCoordinate
	case Vertical:
		tNum = s.startTile.yCoordinate
	}

	for range s.length {
		if tNum > 9 {
			return ErrShipOutOfBound
		}

		tNum = tNum + 1
	}

	return nil
}

func (s *Ship) IsAlive() bool {
	return s.isAlive
}

func (s *Ship) GetHealth() int {
	return s.health
}

func (s *Ship) Hit() error {
	if !s.isAlive {
		return ErrAlreadyDestroyed
	}

	s.health = s.health - 1

	if s.health <= 0 {
		s.isAlive = false
	}

	return nil
}

type ShipClass int

const (
	Destroyer ShipClass = iota
	Submarine
	Cruiser
	Battleship
	AircraftCarrier
)

func (sc ShipClass) String() string {
	switch sc {
	case Destroyer:
		return "Destroyer"
	case Submarine:
		return "Submarine"
	case Cruiser:
		return "Cruiser"
	case Battleship:
		return "Battleship"
	case AircraftCarrier:
		return "AircraftCarrier"

	default:
		return ""
	}
}

type Coordinate struct {
	xCoordinate int
	yCoordinate int
}

type Orientation int

func (o Orientation) String() string {
	switch o {
	case Horizontal:
		return "horizontal"
	case Vertical:
		return "vertical"
	default:
		return ""
	}
}

const (
	Horizontal Orientation = iota
	Vertical
)
