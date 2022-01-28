# DSAWithGo <img src="https://img.icons8.com/color/48/000000/golang.png"/>
A DSA repository but everything is in Go.

Reimplementation of this [DSA repo](https://github.com/thisisshub/DSA) by [thisisshub](https://github.com/thisisshub/)

# Contents
1. [mathematics](https://github.com/adhadse/DSAWithGo/tree/master/mathematics)
2. [sorting](https://github.com/adhadse/DSAWithGo/tree/master/sorting)

# FAQshttps://github.com/adhadse/DSAWithGo/tree/master/mathematics)
1. **Why in Go? Why not in Python/C/C++ etc?**

As for C/C++, you can, you probably should! But when the focus is learning DSA, C/C++'s complexity can become a barrier. As for Python, it's already done by [thisisshub](https://github.com/thisisshub/).

Then why Go? One reason is for pure performance comparable to C, combined with the simplicity and ease like python.

2. **How do I use it?**

Go through them one by one, and reimplement in your favourite language.

# Prerequisites
Install Go for your relevant platform. If you are using editors or IDEs it's fairly easy to play through them, otherwise you can run these commands:

- Running all tests 

   ```go
   go test ./...
   ```
- Running all benchmark
  
  ```go
  go test ./... -bench ./...
  ```
- Benchmarking a package, go into that package and,

  ```go
  go test -bench .
  ```
- Testing a package, go into that package and

  ```go
  go test .
  ```
- Running a specific test, go into that specific directory and,

   ```go
   go test run=<TestFunctionName>
   ```
- Benchmarking a function with a specific number of loop count
   
   ```go
   go test -run=<BenchmarkFunctionName> -bench=10000x
   ```
   
Created By Anurag Dhadse
