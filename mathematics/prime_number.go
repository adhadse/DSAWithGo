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

// IsPrime return a boolean value to check if
// there exists any factors greater than 1.
// a number divisible only by itself and 1.
func IsPrime(x int) bool {
	if x < 2 {
		return false
	} else {
		for i := 2; i < int(math.Sqrt(float64(x))+1); i++ {
			if x%i == 0 {
				return false
			}
		}
	}
	return true
}
