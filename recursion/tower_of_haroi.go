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

package recursion

import "fmt"

// TowerOfHanoi objective is to move the entire stack of
// small to larger radius discs to another rod.
// On the condition that:
//  - 1. Only one disk can be moved at a time
//  - 2. A disk can only be moved if it is the uppermost disk on a stack
//  - 3. No disk may be placed on top of a smaller disk
func TowerOfHanoi(numOfDiscs int, source, destination, auxiliary string) {
	if numOfDiscs == 1 {
		fmt.Printf("Move disk 1 from %s -> %s\n", source, destination)
		return
	}
	TowerOfHanoi(numOfDiscs-1, source, auxiliary, destination)
	fmt.Printf("Move disk %d from %s -> %s\n", numOfDiscs, source, destination)
	TowerOfHanoi(numOfDiscs-1, auxiliary, destination, source)
}
