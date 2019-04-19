# go

## install

        export GOPATH=~/go
        export PATH=$GOPATH/bin:$PATH
        source /etc/profile

## tour

- go tour

        go get golang.org/x/tour
        tour

- my tour

        cd tour
        go run tour.go (go build tour.go && ./tour)

## install binary

- go get (only checks missing project, -u checks update if project/file exists)

        go get github.com/golang/example/hello (download into $GOPATH/src, install into $GOPATH/bin)
        ./hello

- local

        go install github.com/kyle11235/go/hello (install into $GOPATH/bin)
        ./hello

## install package

- go get

        go get -u --tags nopkcs11 github.com/hyperledger/fabric/core/chaincode/shim

- local

        go install github.com/kyle11235/go/pkpath (install into $GOPATH/pkg)

## others

- build for other OS architecture

        cd hello
        GOOS=linux GOARCH=amd64 go build
        ./hello
