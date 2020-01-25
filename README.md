# GolangGinTemplate
Gin Mix RestAPI template modularized

*Built with: go version go1.13.1 darwin/amd64*

- Requires you have golang installed: https://golang.org/doc/install
- Requires you have dep installed(easy to manage dependencies): https://golang.github.io/dep/docs/installation.html
- Make sure this repo is downloaded within the go ENV Path

## Get Dependencies 
`dep ensure`

## Build exectuable
`go build -o restapi`

## Execute
`./restapi`


## Dockerized Solution

### Build Container
`docker build --tag go_gin .`

### Run Container
`docker run -p 10000:10000 --name go_gin -d go_gin`


