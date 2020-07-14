package main

import (
	"fmt"

	pk "github.com/kyle11235/go/pkpath"

	// "github.com/kyle11235/go/playground/mod/mod1/foo"
	"github.com/kyle11235/go/playground/mod/mod1/v2/foo"
)

func main() {
	// module
	fmt.Println(foo.Foo())

	// old online package
	fmt.Println(pk.Foo("biu biu"))
}
