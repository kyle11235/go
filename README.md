# mygo

## go hello

        go get github.com/golang/example/hello (fetch/build/install into $workspace/bin)
        hello

## my hello

        go install github.com/kyle11235/mygo/stringutil
        go install github.com/kyle11235/mygo/hello
        hello

## tour

        cd tour
        go build test.go
        ./test

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

        - auto complete issue
        manual suggestion trigger has beed set to cmd + 1 in keyboard shortcuts
        fix auto complete issue:
                Run Go: Install/Update Tools in VSCode
                cd ~/go/bin
                ./gocode close


        - complile on mac, execute on linux:
        GOOS=linux GOARCH=amd64 go build ./hello.

        - invoke
        curl -H "Content-type:application/json" -X POST http://129.213.107.166:4100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"myorgorderer","chaincode":"mysacc","method":"set","args":["a", "200"],"chaincodeVer":"v2"}'

        - query
        curl -H "Content-type:application/json" -X POST http://129.213.107.166:4100/bcsgw/rest/v1/transaction/query -d '{"channel":"myorgorderer","chaincode":"mysacc","method":"query","args":["a"],"chaincodeVer":"v2"}'