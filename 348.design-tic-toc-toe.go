type TicTacToe struct {
    row1,row2 []int
    col1,col2 []int
    slant1, slant2 int
    slant11, slant21 int
    win int
}


/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
    return TicTacToe{
        row1: make([]int, n),
        row2: make([]int, n),
        col1: make([]int, n),
        col2: make([]int, n),
        win: n,
    }
}


/** Player {player} makes a move at ({row}, {col}).
        @param row The row of the board.
        @param col The column of the board.
        @param player The player, can be either 1 or 2.
        @return The current winning condition, can be either:
                0: No one wins.
                1: Player 1 wins.
                2: Player 2 wins. */
func (this *TicTacToe) Move(row int, col int, player int) int {
    if player == 1 {
        this.row1[row]++
        this.col1[col]++
        if row == col {
            this.slant1++
        }
        if row+col == this.win-1 {
            this.slant11++
        }
        if this.row1[row] == this.win || this.col1[col] == this.win || this.slant1 >= this.win || this.slant11 >= this.win {
            return 1
        } 
    } else {
        this.row2[row]++
        this.col2[col]++
        if row == col {
            this.slant2++
        }
        if row+col == this.win-1 {
            this.slant21++
        }
        if this.row2[row] == this.win || this.col2[col] == this.win || this.slant2 >= this.win || this.slant21 >= this.win {
            return 2
        }
    }

    return 0
}


/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */