## Data Gathering and Organization

Without data, our machine learning models can't learn. This might seem obvious. However, we need to realize that part of the strength of the models that we build is in the data that we feed them.  We need to make sure that we gather relevant, clean data to power our machine learning models, such that they can operate on the data as expected and produce valuable results.

## Notes

- Go's have support for dealing with most any data format/source, and the packages keep getting better
- `gonum` provides matrix/stats capabilities similar to Python's `numpy`
- `gota` provides dataframe capabilities similar to Python's `pandas`

## Links

[Go Slices: Usage and Internals](https://blog.golang.org/go-slices-usage-and-internals)  
[gonum docs](https://godoc.org/gonum.org/v1/gonum)  
[gota docs](https://godoc.org/github.com/kniren/gota/dataframe)  

## Code Review

[Slices](example1/example1.go)  
[Matrices](example2/example2.go)  
[Manipulating a matrix](example3/example3.go)  
[Matrix operations](example4/example4.go)  
[Parsing a CSV into a matrix](example5/example5.go)  
[Parsing a CSV into a dataframe](example6/example6.go)  
[Manipulating a dataframe](example7/example7.go)

## Exercises

### Exercise 1

Implement another way of handling the CSV parsing error we encountered above.  That is, handle the missing value in a way other than throwing an error.

[Template](exercises/template1/template1.go) |
[Answer](exercises/exercise1/exercise1.go)

### Exercise 2

Instead of getting the maximum integer in the integer columns, calculate the sum of integers in the integer column based on [example_clean.csv](data/example_clean.csv) another way of handling the CSV parsing error we encountered above.  

[Template](exercises/template2/template2.go) |
[Answer](exercises/exercise2/exercise2.go)

