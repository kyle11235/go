
# mod

- enable

        https://github.com/golang/go/wiki/Modules
        export GO111MODULE=on
        export GO111MODULE=auto (reset for original go get to work)

- init

        go mod init / go mod init github.com/kyle11235/go/playground/mod/mod1
        go mod tidy

- use

        - basic (mod1 use quote)
        // automatically add into $GOPATH/pkg/mod
        // in go.mod, highest version of common will be used
        go run hello.go

        - use local module (mod2 use mod1)
        // default is v0.0.0
        require github.com/kyle11235/go/playground/mod/mod1 v0.0.0

        // relative path to mod2
        replace github.com/kyle11235/go/playground/mod/mod1 => ../mod1

        go run world.go

        - use old package online
        import pk "github.com/kyle11235/go/pkpath"
        fmt.Println(pk.Foo("biu biu"))

        // automatically add into $GOPATH/pkg/mod
        // in go.mod, require github.com/kyle11235/go v0.0.0-xxxxxx
        go run world.go

- check

        go list -m all / go list -m -u all / go list -m -versions all / go list -m -u -json all

- update

        update go.mod / go get golang.org/x/text@v0.3.2
        go mod tidy

- v2 or higher

        // declare, v0, v1 are omitted
        module github.com/my/mod/v2

        // use
        require github.com/my/mod/v2 v2.0.0
        import "github.com/my/mod/v2/mypkg"

- vendor

        go mod vendor
        go build -mod=vendor (use vendor rather go.mod)