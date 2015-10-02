package game

import "github.com/thethirdone/connect4-solver/data"

type Board struct {
	Player int8
	board  [7][]int8
}

func (self *Board) Drop(index int8) *Board {

	if len(self.board[index]) > 5 {
		return nil
	}
	asdf := *self
	for i, _ := range asdf.board {
		asdf.board[i] = make([]int8, len(self.board[i]))
		copy(asdf.board[i], self.board[i])
	}
	asdf.board[index] = append(asdf.board[index], asdf.Player)
	asdf.Player = -asdf.Player
	return &asdf
}

//player bit + length of column (3 bits)*7 + moves (1 or 0) (1 bit)
//max length 1 + 21 + 42 = 64 therefore, no hash collisions
func (self *Board) Hash() int64 {
	var out, tmp int64
	for _, col := range self.board {
		out = out << 3
		out += int64(len(col))
	}
	for _, col := range self.board {
		for _, cell := range col {
			tmp = tmp << 1
			if cell == 1 {
				tmp++
			}
		}
	}
	out = out<<42 + tmp
	out *= int64(self.Player)
	return out
}

func (self *Board) win() int8 {
	var tmp int8

	//vertical
	for _, col := range self.board {
		tmp = 0
		for _, temp := range col {
			if tmp*temp > 0 {
				tmp += temp
			} else {
				tmp = temp
			}
			if tmp > 3 {
				return 1
			} else if tmp < -3 {

				return -1
			}
		}
	}

	//horizontal
	for i := 0; i < 6; i++ {
		tmp = 0
		for k := 0; k < 6; k++ {
			if len(self.board[k]) <= i {
				tmp = 0
			} else if tmp*self.board[k][i] >= 0 {
				tmp += self.board[k][i]
			} else {
				tmp = self.board[k][i]
			}
			if tmp > 3 {
				return 1
			} else if tmp < -3 {
				return -1
			}
		}
	}
	//top left to bottom right
	starting := [6][2]int8{{0, 3}, {0, 4}, {0, 5}, {1, 5}, {2, 5}, {3, 5}}
	var index [2]int8
	for i := 0; i < 6; i++ {
		index = starting[i]
		tmp = 0
		for k := 0; k < 6; k++ {

			if len(self.board[index[0]]) <= int(index[1]) {
				tmp = 0
			} else if tmp*self.board[index[0]][index[1]] >= 0 {
				tmp += self.board[index[0]][index[1]]
			} else {
				tmp = self.board[index[0]][index[1]]
			}
			if tmp > 3 {
				return 1
			} else if tmp < -3 {
				return -1
			}
			index[0]++
			index[1]--
			if index[0] > 6 || index[1] < 0 {
				break
			}
		}
	}
	//top right to bottom left
	starting = [6][2]int8{{6, 3}, {6, 4}, {6, 5}, {5, 5}, {4, 5}, {3, 5}}
	for i := 0; i < 6; i++ {
		index = starting[i]
		tmp = 0
		for k := 0; k < 6; k++ {

			if len(self.board[index[0]]) <= int(index[1]) {
				tmp = 0
			} else if tmp*self.board[index[0]][index[1]] >= 0 {
				tmp += self.board[index[0]][index[1]]
			} else {
				tmp = self.board[index[0]][index[1]]
			}
			if tmp > 3 {
				return 1
			} else if tmp < -3 {
				return -1
			}
			index[0]--
			index[1]--
			if index[0] < 0 || index[1] < 0 {
				break
			}
		}
	}
	tmp = 0
	for i := 0; i < 7; i++ {
		tmp += int8(len(self.board[i]))
	}
	if tmp > 41 {
		return 0
	}
	return -2
}

func (self *Board) Evaluate() int8 {
	if value, exists := data.Get(self.Hash()); exists {
		return value
	}

	//checks if the next move is a win
	temp := make([]*Board, 0, 7)

	for i := 0; i < 7; i++ {
		tmp := self.Drop(int8(i))

		//if move is invalid
		if tmp == nil {
			continue
		}

		//if there is a winner
		if winner := tmp.win(); winner != -2 {
			if winner == self.Player {
				data.Set(tmp.Hash(), self.Player)
				return self.Player
			}
			continue
		}

		//otherwise add it to the recursive check list
		temp = append(temp, tmp)
	}

	//recursive check for next moves
	canTie := false

	for _, tmp := range temp {
		score := tmp.Evaluate()
		//if there is a winnning move, take it
		if score == self.Player {
			data.Set(self.Hash(), self.Player)
			return self.Player
		}
		//mark if can tie
		if score == 0 {
			canTie = true
		}
	}
	//if can tie and cannot win, tie
	if canTie {
		data.Set(self.Hash(), 0)
		return 0
	}

	//the opponent wins
	data.Set(self.Hash(), -self.Player)

	return -self.Player
}
func Init(routines int) *Board {
	data.Init(2)
	return new(Board)
}
