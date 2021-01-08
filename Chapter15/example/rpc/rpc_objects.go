package rpc_objects

// Args 类型
type Args struct {
	N, M int
}

// Multiply 方法
func (t *Args) Multiply(args *Args, reply *int) error {
	*reply = args.N * args.M
	return nil
}
