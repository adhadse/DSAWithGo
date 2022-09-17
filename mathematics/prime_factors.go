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

package mathematics

import "math"

// PrimeFactors return a list of prime factors of an integer by
// first checking if the modulo of the value of d and number is equal to 0.
func PrimeFactors(x int) []int {
	var divisors []int
	var primeFactors []int
	for d := 2; d < int(math.Floor(float64(x)/2))+1; d++ {
		if x%d == 0 {
			divisors = append(divisors, d)
		}
	}
	// Then applying the for loop inside the for loop to count a factor only once.
	for _, d := range divisors {
		for _, od := range divisors {
			if od != d {
				if d%od != 0 {
					continue
				} else {
					break
				}
			}
			primeFactors = append(primeFactors, d)
		}
	}
	return primeFactors
}
