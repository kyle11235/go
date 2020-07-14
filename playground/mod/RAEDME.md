
# mod

- enable

        https://github.com/golang/go/wiki/Modules

        Note that outside of GOPATH, you do not need to set GO111MODULE to activate module mode.
        export GO111MODULE=on
        export GO111MODULE=auto (reset for original go get to work)

- init

        go mod init / go mod init github.com/kyle11235/go/playground/mod/mod1
        go mod tidy

- use

        - use module (mod1 use quote)
        
                - update hello.go
                use "rsc.io/quote"
        
                - automatically (save file/go build)
                source code -> $GOPATH/pkg/mod
                require rsc.io/quote v1.5.2 (latest version) -> go.mod

                - manually update version
                e.g. v1.5.2 to v1.5.1
                save hello.go will download new declared version
                go mod tidy (clear pkg usage history in go.sum)

                - or use go get to manually download
                e.g. 
                go get example.com/package (download latest pkg)
                go get -u example.com/package (download latest pkg + all its latest dependencies)
                go get -u github.com/hyperledger/fabric/core/chaincode/shim@v1.4.4

        - use old package online (mod2 use my old pkg)
        
                - update world.go
                import pk "github.com/kyle11235/go/pkpath"
                fmt.Println(pk.Foo("biu biu"))

                - automatically
                souce code -> $GOPATH/pkg/mod
                require github.com/kyle11235/go v0.0.0-20190729030720-xxxxxx -> go.mod

        - use local module (mod2 use mod1)

                once local pkg changed to module, have to declare require/replace
       
                - update word.go
                use "github.com/kyle11235/go/playground/mod/mod1/foo"
                
                - update go.mod manually
                e.g.
                require github.com/kyle11235/go/playground/mod/mod1 v0.0.0 (default is v0.0.0 if not declared)
                replace github.com/kyle11235/go/playground/mod/mod1 => ../mod1 (relative path to mod2)

        - v2 or higher (mod2 use mod1/v2)

                - module/import name vs release
                
                        github.com/kyle11235/go/playground/mod/mod1 (v0, v1 are omitted, no release required)
                        github.com/kyle11235/go/playground/mod/mod1/v2 (release should be go/playground/mod/mod1/v2.0.0 required)
                        github.com/kyle11235/go/playground/mod/mod1/v2.1.6 (release should be go/playground/mod/mod1/v2.1.6/v2.1.6 required)

                - update to v2
                
                        - option 1
                        update code in current directory
                        update module to /v2
                        update import to /v2 (comsumer side needs to have Go versions 1.9.7+, 1.10.3+, and 1.11)
                        tag to release v2.0.0

                        - option 2
                        copy code to sub folder ./v2
                        create module /v2 in sub folder
                        tag to release v2.0.0

- check version

        - View final versions that will be used in a build for all direct and indirect dependencies
        go list -m all
        go list -m -json all

        - View available minor and patch upgrades for all direct and indirect dependencies
        go list -m -u all
        
        
        - view all version
        go list -m -versions all
        
- vendor

        go mod vendor
        go build -mod=vendor (use vendor rather go.mod)