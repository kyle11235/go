# go

## install

        export GOROOT=/usr/local/go
        export PATH=$PATH:$GOROOT/bin
        export GOPATH=$HOME/go
        export PATH=$PATH:$GOPATH/bin
        source /etc/profile

## install tour

        go get golang.org/x/tour (download whole x project into $GOPATH/src, install tour into $GOPATH/bin)
        tour

## install my hello

        go get github.com/kyle11235/go/hello
        hello

## run my tour

        cd tour
        go run tour.go (go build tour.go && ./tour)

## others

- build for other OS architecture

        cd hello
        GOOS=linux GOARCH=amd64 go build
        ./hello
