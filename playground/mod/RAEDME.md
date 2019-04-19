
# mod

        export GO111MODULE=on

        go mod init github.com/kyle11235/go/playground/mod
        go build hello.go (automatically add new dependencies with version into $GOPATH/pkg/mod)
        ./hello
        go list -u -m all (check available)
        check/update go.mod
        go mod tidy (remove unused)

        export GO111MODULE=auto (reset for original go get to work)

