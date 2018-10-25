### Golang tutorial

## install
https://golang.org/doc/install

## run program
`
$ go run file.go
`

## build program
`
$ go build file.go
`

### Go Dockerization
I am using to start go application in docker busybox:glibc docker image

## build
```
$ docker build -t hello-indigo:1.0 .
```

## run container
```
$ docker run --rm -p 9000:9000 -it hello-indigo:1.0
```

### to test service
## Just go to page
```
http://localhost:9000/smiles-to-mol?smiles=ONc1cccc1
```
