module github.com/kyle11235/go/playground/mod/mod2

go 1.12

require (
	github.com/kyle11235/go v0.0.0-20190729030720-dc104ed0938e

	// github.com/kyle11235/go/playground/mod/mod1 v0.0.0
	github.com/kyle11235/go/playground/mod/mod1/v2 v2.0.0
)

// replace github.com/kyle11235/go/playground/mod/mod1 => ../mod1
replace github.com/kyle11235/go/playground/mod/mod1/v2 => ../mod1/v2
