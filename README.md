# Go Simple REST API With MongoDB and Docker

Simple Go REST API for CRUD operations using MongoDB as database backend

Endpoints:
```
/movies
/movies/create/
/movies/get/{id}
/movies/update/{id}
/movies/delete/{id}
```

I have applied Test Driven Development (TDD) for this project, you can find unit test cases in goapi_test.go file.

The whole application has been Dockerized using docker-compose for easy deployment. Both the Go App and the Mongo Database
Just build the docker images with docker-compose and run it.
Then you can access it from :8076 port on the above mentioned endpoints.

Optimization: Deploying the whole app into a docker container requires almost 800+ MB space. But, single compiled binary takes only 20 MB
So I have deployed this app as single binary file. You can find compiled binaries in the bin/ directory. I compiled for windows and Linux

DevOps approach: I came from almost 1 year of Python/Django web development and DevOps background.
So I developed this project with DevOps in mind, such as I have implemented Jenkins for CI/CD, testing etc.

PS: Pardon me for the .idea directory, I've added it to the .gitignore file. But it's not working for some strange reason.