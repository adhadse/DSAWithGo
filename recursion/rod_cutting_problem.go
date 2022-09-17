// Copyright (C) 2022   Anurag Dhadse. All Rights Reserved.
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

package recursion

import "math"

// RodCutting is a problem where a given rod of length n
// and price list of all piece length < n; we need to find
// the maximum value obtainable by selling up cut rod pieces.
func RodCutting(pricePerLength []int, length int) int {
	if length <= 0 {
		return 0
	}
	maxVal := -math.MaxInt64
	for i := 0; i < length; i++ {
		maybeMaxVal := pricePerLength[i] + RodCutting(pricePerLength, length-i-1)
		if maxVal <= maybeMaxVal {
			maxVal = maybeMaxVal
		}
	}
	return maxVal
}
