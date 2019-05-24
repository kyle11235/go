
# mod

- enable

        https://github.com/golang/go/wiki/Modules
        export GO111MODULE=on
        export GO111MODULE=auto (reset for original go get to work)

- init

        go mod init / go mod init github.com/kyle11235/go/playground/mod1
        go mod tidy

- usage

        go run hello.go (automatically add new dependencies with version into $GOPATH/pkg/mod, highest version of common will be used)

- check

        go list -m all / go list -m -u all / go list -m -versions all / go list -m -u -json all

- update

        update go.mod / go get golang.org/x/text@v0.3.2
        go mod tidy

- replace local module

        default is v0.0.0

        require github.com/kyle11235/go/playground/mod/mod1 v0.0.0
        replace github.com/kyle11235/go/playground/mod/mod1 => ../mod1

        go run world.go

- v2 or higher

        v0, v1 are omitted
        module github.com/my/mod/v2
        require github.com/my/mod/v2 v2.0.0
        import "github.com/my/mod/v2/mypkg

- vendor

        go mod vendor
        go build -mod=vendor (use vendor rather go.mod)