# What is this repository for?

This repository is a simple key-value store api example using **Clean Architecture** by Robert C. Martin.

# Clean Architecture

Clean Architecture is a software design philosophy that aims to encapsulate business logic other parts of the application. It seperate the software design into different ring levels. Functions, classes or frameworks in the outer layer cannot be accessed inner layer, in other words, inner layer(business logic) cannot be depended to outer layer. This helps developers easily change frameworks in the outer layer by not changing the domain (business logic).

![Clean Architecture](https://netsharpdev.com/images/posts/shape-half.png)

# How to run?

Application can be run in two different ways.

**Run with go run:**

    cd cmd
    go run .

**Run with docker:**

    docker build -t dh-api .
    docker run -dp 8181:8181 dh-api

**Run unit test:**

In order to run unit test, following command can be run.

    go test ./...


## CI/CD and Deployment

This project is using **Github Actions** which is a continues integration and deployment tool offered by Github. After each commit, the pipeline is trigerred and it will automaticly deploy application to **Heroku** after all test pass. Currently, application is running on [this link.](https://dh-api-golang.herokuapp.com/)  