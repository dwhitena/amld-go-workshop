## More sophisticated models

We should all strive for simplicity (which is a mantra of the Go community), but sometimes we do need a model that is more complicated than linear regression or kNN. Go has us covered here. We can interface with major frameworks like TensorFlow, utilize more Go-centric frameworks like Gorgonia, or you can utilize services like MachineBox to manage your ML models.

## Notes

- The Go bindings for Tensorflow allow for easy inference based on models trained in Python. You can technically train a model directly in Go, but it is not recommended quite yet. These binding are actively being developed.
- Because Go provides built in HTTP support and `cgo`, it is relatively easy to interact with a variety of frameworks that have a REST or C/C++ interface (e.g., H2O or Intel's Deep Learning SDK).
- Gorgonia remains the largest Go-native effort to enable many of the same workflows that are associated with TensorFlow.

## Links

[TensorFlow in Go](https://www.tensorflow.org/install/install_go)  
[Gorgonia](https://github.com/gorgonia/gorgonia)  
[MachineBox (which includes a very nice Go SDK)](https://machinebox.io/)  
[GoCV (which we will use here)](https://gocv.io/)  

## Code Review

[Object detection with TensorFlow and GoCV](example1/example1.go)

## Exercise

### Exercise 1

Recreate the object detection example on your local machine by:

1. [Installing GoCV](https://gocv.io/getting-started/)
2. Downloading the TensorFlow Inception model from [here](https://storage.googleapis.com/download.tensorflow.org/models/inception5h.zip)
3. Unzip the Inception model files.
4. Build the example with `go build`.
5. Run the example with `./example1 0 <modelfile> <descriptionsfile>`
