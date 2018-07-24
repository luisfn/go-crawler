# Go Crawler

Basic price crawler written in Go as a learning experiment

Each link on links.txt is called and have its content parsed in a separated channel, to improve performance using concurrency capabilities from Go.


## Setup

This Project uses dep as a dependency manager. To install dep, please execute:

```
go get -u github.com/golang/dep/cmd/dep
```

To install needed dependencies execute:


```
dep ensure
```

## Execution

#### Locally
``` 
go run crawler.go
```

#### Docker
``` 
docker-compose up
```