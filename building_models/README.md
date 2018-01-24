# Building ML Models

Go actually has a [sizable ecoystems for machine learning and data science](https://github.com/gopherdata/resources/tree/master/tooling). There are packages like `golearn` that allow you to quickly implement common algorithms, such as decision trees and kNN, and there are more advanced packages like `gorgonia` that provided flexible foundations for ML (similar to TensorFlow). Finally, there are easy ways to integrate Go with widely used ML frameworks like TensorFlow, H2O, etc. We will explore a couple of common ML examples here, and a more advanced one in the [bonus section](../bonus), assuming we have time. 

## Notes

- Many ML techniques are already implemented in Go, [this list](https://github.com/gopherdata/resources/tree/master/tooling) of tooling is a good starting point.
- If you can't find what you need, make sure and ping all the data science Gophers in the #data-science channel on [Gophers Slack](https://invite.slack.golangbridge.org/).
- ML in Go is ready for production, many companies (including the New York Times, Dell/EMC, and others) are already using it at scale.

## Links

[Machine Learning with Go (book)](https://www.packtpub.com/big-data-and-business-intelligence/machine-learning-go)  
[Building a neural network from scratch in Go](http://www.datadan.io/building-a-neural-net-from-scratch-in-go/)  
[golearn docs](https://godoc.org/github.com/sjwhitworth/golearn)  
[sajari/regression docs](https://godoc.org/github.com/sajari/regression)  
[Other Go tooling for ML and data science in general](https://github.com/gopherdata/resources/tree/master/tooling)  

## Code Review

### Regression Example:

In the following examples, we will implement a linear regression model that predicts diabetes disease progression based on the [diabetes data set](data/diabetes.csv).

[Profile the data](example1/example1.go)   
[Visualize correlations](example2/example2.go)  
[Split the data into training and test sets](example3/example3.go)  
[Train and evaluate a linear regression model](example4/example4.go)  

### Classification Example:

In these next examples, we will implement a kNN model that can classify Iris flower species based on measurements of Iris flowers:

[Profile the data](example5/example5.go)  
[Train and use cross-validation to validate a kNN model](example6/example6.go)

## Exercises

### Exercise 1

**Part A** Look again at the scatter plots created in [example2](example2/example2.go).  Select another feature (other than `bmi`) to add to our regression model.

**Part B** Add your selected feature to the model trained in [example4](example4/example4.go) making it a "multiple linear regression" model.  Evaluate your choice.  You may be to try a few different features choices, but try to choose one that improves your model evaluation.

[Template](exercises/template1b/template1b.go) |
[Answer](exercises/exercise1b/exercise1b.go)

### Exercise 2

Find an optimal `k` value for the above model of iris species. That is, search over various `k` values and evaluate the predictions.

[Template](exercises/template2/template2.go) |
[Answer](exercises/exercise2/exercise2.go)
