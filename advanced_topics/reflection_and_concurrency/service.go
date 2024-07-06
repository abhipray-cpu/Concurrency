package main

type AddArgs struct {
	A int
	B int
}

type ArithMeticService struct{}

func (s *ArithMeticService) Add(args *AddArgs, result *int) error {
	*result = args.A + args.B
	return nil
}

type RPCService struct{}

type ConcatArgs struct {
	Str1, Str2 string
}

func (s *RPCService) Concat(args *ConcatArgs, reply *string) error {
	*reply = args.Str1 + args.Str2
	return nil
}
