# Contributing Guidelines
## Filing an issue
- Check existing issues, comment if you find something worth discussing.
- File a Bug report or Feature request fill out the template as much as possible.

## Submitting a pull request
### Do
- Make sure you format your code using `gofmt` tool
- Do comment your code wherever necessary especially for functions
  explaining what it does and what it returns and expects
  
  Example:
  ```go
  // BubbleSort returns a sorted slice using bubble sort algorithm
  //
  // It has an outer loop looping over whole list index i and
  // in the inner loop of j from 0->n-i-1 repetitively creates "bubbles"; item at j and j+1
  // exchanging position for incorrectly ordered elements in the "bubble"
  func BubbleSort(list []int) []int {
      ...  
  }
  ```

  Use names for variables that signifies what it does:
  ```go
  visited := make([]bool, len(list))  
  ```

  and comment hard to understand parts like this (from Dijkstra algorithm):
  ```go
  // Second condition of OR applies when we are not able to choose next 'UNEXPLORED'
  // node; i.e., we reached a node from where all nodes are already visited
  if nextToVisit == destination || nextToVisit == currentlyVisited {
      break
  }
  ```
- Do submit only a single example/feature per pull-request
- Do include a description of what your example is expected to do
- Do add tests to your Code in `<package>_test.go` 

### Don't
- Avoid commenting every single line/variable and don't give variable meaningless names like this:
  ```go
  // variable that store if item visited or not
  a := make([]bool, len(list))  
  ```
- Don't change the license information
- Don't try to do too much at once, change/create one or two files only
- Don't submit multiple variations of the same example, demonstrate one thing concisely

## Suggesting a feature
- Check existing issues, open and closed to make sure it hasn't already been suggested
- Considered if it's necessary in the library, or is an advanced technique that could be separately explained in an example