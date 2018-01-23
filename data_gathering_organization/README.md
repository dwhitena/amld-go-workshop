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

Create a matrix from [diabetes.csv](../data_versioning/data/diabetes.csv) using `gonum.org/v1/gonum/mat`. Format and output the first 10 rows to standard out.

[Template](exercises/template1/template1.go) |
[Answer](exercises/exercise1/exercise1.go)

### Exercise 2

Divide the following matrix by its norm:


    a = ⎡1  2  3⎤
        ⎢0  4  5⎥
        ⎣0  0  6⎦


[Template](exercises/template2/template2.go) |
[Answer](exercises/exercise2/exercise2.go)

### Exercise 3

Use Gota dataframes to read [iris.csv](data/iris.csv) and output three files corresponding to each Iris species (`setosa.csv`, `versicolor`, and `virginica.csv`), each of the three files containing only the rows corresponding to the respective species.

[Template](exercises/template3/template3.go) |
[Answer](exercises/exercise3/exercise3.go)
