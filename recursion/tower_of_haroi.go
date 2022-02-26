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
