// Copyright 2022 The DSAWithGo Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package backtracking

// SolveNQueensProblem A recursive function to solve N Queen's problem
// The problem is placing N chess queens on an NxN
// chessboard so that no two queens attack each other.
//
// BackTracking Solution
//  1. Start in leftmost column
//  2. If all queens are placed
//	  return true
//  3. Try all rows in the current column.
//    Do following for every tried row.
//    a. If the queen can be placed safely in this row
// 		 then mark this [row, column] as part of the solution
// 		 and recursively check if placing queen here leads to a
//		 solution
//    b. If placing the queen in [row, column] leads to a solution
//       then return true
//    c. If placing queen doesn't lead to a solution then
// 		 unmark this [row, column] (backtrack) and go to step (a)
//		 to try other rows
//  4. If all rows have been tried and nothing worked,
// 	   return false to trigger backtracking.
func SolveNQueensProblem(board [][]int, col int) ([][]int, bool) {
	N := 8
	// base case: If all queens are placed then return true
	if col >= N {
		return board, true
	}
	// consider this column and try placing this queen in all rows one by one
	for i := 0; i < N; i++ {
		// check if the queen can be placed on board[i][col]
		if isSafe(board, i, col, N) {
			// place this queen in board[i][col]
			board[i][col] = 1

			// recur to place rest of the queens
			board, solved := SolveNQueensProblem(board, col+1)
			if solved {
				return board, true
			}

			// if placing queen in board[i][col] doesn't lead to
			// a solution, then
			// remove queen from board[i][col]
			board[i][col] = 0
		}
	}
	// if queen cannot be placed in any row in this column then return false
	return board, false
}

// isSafe checks if a queen can be placed on board[row][col]
// Note that this function is called when "col" queens are
// already placed in columns from 0 to col-1.
// So we need to check only left side for attacking queens.
func isSafe(board [][]int, row, column, N int) bool {
	// check this row on left side
	for col := 0; col < column; col++ {
		if board[row][col] == 1 {
			return false
		}
	}

	// Check upper diagonal on left side
	for i, j := row, column; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	// Check lower diagonal on left side
	for i, j := row, column; j >= 0 && i < N; i, j = i+1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}
	return true
}
