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

import (
	"fmt"
	"testing"
)

func TestSolveNQueensProblem(t *testing.T) {
	board := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	board, solved := SolveNQueensProblem(board, 0)
	if solved == false {
		println("Solution does not exist")
		return
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf(" %d ", board[i][j])
		}
		println()
	}
}