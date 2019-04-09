# go

## tour

        cd tour
        go build basic.go && ./basic

## install

- go hello

        go get github.com/golang/example/hello (fetch/build/install into $workspace/bin)
        ./hello

- my hello

        go install github.com/kyle11235/go/hello (build/install into $workspace/bin)
        ./hello

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

        - complile on mac, execute on linux:
        GOOS=linux GOARCH=amd64 go build ./hello.
