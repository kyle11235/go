package api

type Args struct {
	A, B int
}

type Result struct {
	A, B int
}

type API struct {
}

type Calculator interface {
	Multiply(args *Args, reply *int) error
	Divide(args *Args, res *Result) error
}
