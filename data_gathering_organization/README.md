## Introduction to Go

blah blah, general description...

## Notes

- A matrix is a rectangular array representation of numbers, expressions, etc.
- Elements in a matrix are referenced by a row and column index.
- `gonum.org/v1/gonum/mat` provides functionality to create, modify, and manipulate matrices made up of float64 values.

## Links

[The Matrix Cookbook](http://www.math.uwaterloo.ca/~hwolkowi/matrixcookbook.pdf)  
[Khan Academy - Matrices](https://www.khanacademy.org/math/algebra-home/precalculus/precalc-matrices)  
[Khan Academy - Linear Algebra](https://www.khanacademy.org/math/linear-algebra)

## Code Review

[gonum.org/v1/mat docs](https://godoc.org/gonum.org/v1/gonum/mat)    
[Form a float64 matrix](example1/example1.go)  
[Modify a matrix](example2/example2.go)  
[Access values in a matrix](example3/example3.go)  
[Format matrix output](example4/example4.go)  

## Exercises

### Exercise 1

Create a matrix from [diabetes.csv](../data_versioning/data/diabetes.csv) using `gonum.org/v1/gonum/mat`. Format and output the first 10 rows to standard out.

[Template](exercises/template1/template1.go) |
[Answer](exercises/exercise1/exercise1.go)

