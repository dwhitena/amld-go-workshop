# Deploying and Scaling Go-based ML

It's one thing to do some ML experiments on your laptop, but, to create real value in an organization, ML has to be deployed off of your laptop and scaled to production data. Thankfully, there is a variety of Go tooling to accomplish this, and I will illustrate this here.

## Notes

- Most companies are considering or are already using Go-based infrastructure tools (Docker and Kubernetes in particular) to manage their workloads at scale.
- Go-based ML workloads fit nicely into these frameworks and can scale over production data.
- In addition to the integrity benefits discussed earlier, Go automatically gives you the extra benefits of efficiency, built in concurrency, and easy deployment.

## Links

[Introduction to Docker](https://training.docker.com/introduction-to-docker)  
[Building minimal Go Docker images](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/)  
[What is Kubernetes](https://kubernetes.io/docs/concepts/overview/what-is-kubernetes/)  
[Creating Pachyderm analysis pipelines](http://docs.pachyderm.io/en/latest/fundamentals/creating_analysis_pipelines.html)  
[Distributed computing on Pachyderm](http://docs.pachyderm.io/en/latest/fundamentals/distributed_computing.html)  
[Updating Pachyderm pipelines](http://docs.pachyderm.io/en/latest/fundamentals/updating_pipelines.html)  

## Code Review

[Containerize training of a linear regression model](example1)  
[Containerize linear regression prediction](example2)  
[Create a linear regression model training pipeline](example3/train.json)  
[Create a linear regression prediction pipeline](example4/predict.json)  

## Demo

I will give a demo of a completely scalable, reproducible, and easy to manage ML workshop that is build entirely with Go and infrastructure build with Go. To try this out on your own locally, you can follow [this guide](https://github.com/dwhitena/pach-go-regression).

