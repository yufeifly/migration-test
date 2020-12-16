package validate

import "fmt"

// VerifyOptions ...
type VerifyOptions struct {
	Addr  string
	Range string
}

func VerifyResult(opts VerifyOptions) error {
	fmt.Printf("addr: %v, range: %v\n", opts.Addr, opts.Range)

	return nil
}
