package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Part 1 = %d\n", part1(input))
	fmt.Printf("Part 2 = %d\n", part2(input))
}

func part1(input string) int {
	tiles := parseTiles(input)
	loop := buildLoop(tiles)

	return (loop.len + 1) / 2
}

func part2(input string) int {
	tiles := parseTiles(input)
	loop := buildLoop(tiles)

	enclosedTileCount := 0
	for _, row := range tiles {
		insideLoop := false
		sumVerticalConnections := 0
		for _, tile := range row {
			node := loop.getNodeAt(tile.position)
			if node != nil {
				// track vertical connections, -1 for up and + 1 for down, this helps us track when we pass over the pipe
				// | = -1+1 = 0 (it has an upwards and downwards tick)
				// L--7 = -1+0+0+1 = 0 => passed over pipe
				// L--J = -1+0+0+-1 != 0 => didn't pass over pipe (only traversed)
				sumVerticalConnections += node.pipe().sumVerticalConnections()
				if sumVerticalConnections == 0 {
					// passed over pipe
					insideLoop = !insideLoop
				}
				// always reset at 2, we want to track when we pass through 0,-1,0 or 0,+1,0
				sumVerticalConnections %= 2
				fmt.Print("Loop    ")
			} else if insideLoop {
				fmt.Print("Inside  ")
				enclosedTileCount++
			} else {
				fmt.Print("Outside ")
			}
			fmt.Printf("%v:%v : %v\n", tile.char, tile.position, insideLoop)
		}
	}
	return enclosedTileCount
}

func parseTiles(input string) [][]tile {
	lines := strings.Split(input, "\n")
	var tiles [][]tile
	for y, line := range lines {
		var row []tile
		for x, pipe := range strings.Split(line, "") {
			row = append(row, tile{
				char:     pipe,
				position: position{x, y},
			})
		}
		tiles = append(tiles, row)
	}
	fmt.Printf("Parsed %d rows\n", len(tiles))
	return tiles
}

func buildLoop(tiles [][]tile) loop {
	loop := loop{}
	loop.addNode(getStartTile(tiles))
	loop.addNode(getConnectedTile(tiles, loop.start.tile))

	for {
		nextTilePosition := loop.tail.getNextTilePosition()
		nextTile := tiles[nextTilePosition.y][nextTilePosition.x]
		isEnd := loop.addNode(nextTile)
		if isEnd {
			break
		}
	}

	return loop
}

func getStartTile(tiles [][]tile) tile {
	for _, row := range tiles {
		for _, tile := range row {
			if tile.char == "S" {
				return tile
			}
		}
	}
	return tile{}
}

func getConnectedTile(tiles [][]tile, t tile) tile {
	for _, pipe := range pipes {
		for _, connection := range pipe.connections {
			position := t.position.subtract(connection)
			if position.y < 0 || position.y >= len(tiles) || position.x < 0 || position.x >= len(tiles[position.y]) {
				continue
			}
			if tiles[position.y][position.x].char == pipe.pipe {
				return tiles[position.y][position.x]
			}
		}
	}
	return tile{}
}

type node struct {
	tile tile
	prev *node
	next *node
}

func (n *node) getNextTilePosition() position {
	previousConnection := n.prev.tile.position.subtract(n.tile.position)
	for _, connection := range n.pipe().connections {
		if connection != previousConnection {
			return n.tile.position.add(connection)
		}
	}
	return position{}
}

func (n *node) pipe() pipe {
	p, exists := pipes[n.tile.char]
	if exists {
		return p
	}

	// calculate connections from adjacent nodes
	connections := []position{
		n.next.tile.position.subtract(n.tile.position),
		n.tile.position.subtract(n.prev.tile.position),
	}
	for _, pipe := range pipes {
		if equivalent(pipe.connections, connections) {
			return pipe
		}
	}

	return pipe{}
}

func equivalent(first, second []position) bool {
	if len(first) != len(second) {
		return false
	}
	exists := make(map[position]bool)
	for _, value := range first {
		exists[value] = true
	}
	for _, value := range second {
		if !exists[value] {
			return false
		}
	}
	return true
}

type loop struct {
	start *node
	tail  *node
	len   int
}

func (l *loop) getNodeAt(position position) *node {
	for cur := l.start; ; cur = cur.next {
		if cur.tile.position == position {
			return cur
		}
		if cur.next == l.start || cur.next == nil {
			return nil
		}
	}
}

func (l *loop) addNode(tile tile) (isEnd bool) {
	if l.start != nil && tile.position == l.start.tile.position {
		// back to start - link up loop
		l.tail.next = l.start
		l.start.prev = l.tail
		return true
	}

	newNode := &node{tile: tile}
	if l.start == nil {
		l.start = newNode
		l.tail = newNode
	} else {
		newNode.prev = l.tail
		l.tail.next = newNode
		l.tail = newNode
	}
	l.len++
	return false
}

type position struct {
	x, y int
}

func (p position) add(v position) position {
	return position{p.x + v.x, p.y + v.y}
}

func (p position) subtract(v position) position {
	return position{p.x - v.x, p.y - v.y}
}

type tile struct {
	char string
	position
}

type pipe struct {
	pipe        string
	connections []position
}

func (p pipe) sumVerticalConnections() int {
	sum := 0
	for _, connection := range p.connections {
		sum += connection.y
	}
	return sum
}

var pipes = map[string]pipe{
	"|": {"|", []position{{0, -1}, {0, 1}}},
	"-": {"-", []position{{-1, 0}, {1, 0}}},
	"L": {"L", []position{{0, -1}, {1, 0}}},
	"J": {"J", []position{{0, -1}, {-1, 0}}},
	"7": {"7", []position{{-1, 0}, {0, 1}}},
	"F": {"F", []position{{1, 0}, {0, 1}}},
}
