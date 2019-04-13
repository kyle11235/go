# go

## tour

- go tour

        go tool tour

- my tour

        cd tour
        go run tour.go (go build tour.go && ./tour)

## install binary

- go get (only checks missing, -u checks update, go help get)

        go get github.com/golang/example/hello (download/install into $workspace/bin)
        ./hello

- local

        go install github.com/kyle11235/go/hello (install into $workspace/bin)
        ./hello

## install package

- go get

        go get -u --tags nopkcs11 github.com/hyperledger/fabric/core/chaincode/shim (download/build/install into $workspace/pkg)

- local

        go install github.com/kyle11235/go/pkpath (build/install into $workspace/pkg)

## chaincode

        1. mysacc (set/get any key value)
        2. chaincode_example02 (move from A to B)
        3. marbles_chaincode (json struct)
        4. test with shim.NewMockStub
                cd sacc
                go test
        5. developer mode - https://hyperledger-fabric.readthedocs.io/en/release-1.1/chaincode4ade.html
        6. more shim API
                go - https://godoc.org/github.com/hyperledger/fabric/core/chaincode/shim#ChaincodeStub
                node - https://fabric-shim.github.io/ChaincodeStub.html

## others

- build for other OS architecture

        cd hello
        GOOS=linux GOARCH=amd64 go build
        ./hello
