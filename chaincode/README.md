# chaincode

- mysacc (set/get any key value)
- chaincode_example02 (move from A to B)
- marbles_chaincode (json struct)
- test with shim.NewMockStub
  
        go version -> go version go1.13.4 darwin/amd64
        export GO111MODULE=on
        go get -u github.com/hyperledger/fabric/core/chaincode/shim@v1.4.4

        cd sacc
        go mod init / go mod init example.com/m (if out of GOPATH)
        go test