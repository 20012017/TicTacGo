package cli

import "ttt/cli/players"

var GameOptions map[int][]int = map[int][]int{
	1: []int{players.HUMAN, players.HUMAN},
	2: []int{players.HUMAN, players.COMPUTER},
	3: []int{players.COMPUTER, players.HUMAN},
	4: []int{players.COMPUTER, players.COMPUTER},
}
