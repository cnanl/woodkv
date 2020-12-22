package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"woodkv/types"
)

type client struct {
	method string
	key    []byte
	value  []byte
	n      int
	stdin  *bufio.Reader
}

func main() {
	cl := client{}
	cl.start()
}
func (cl *client) start() {
	cl.stdin = bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(">")
		fmt.Fscan(cl.stdin, &cl.method)
		cl.parse()
	}
}

func (cl *client) parse() {
	switch cl.method {
	case "GET", "Get", "get":
		fmt.Fscan(cl.stdin, &cl.key)
		cl.handle(types.Mget)
	case "PUT", "Put", "put":
		fmt.Fscan(cl.stdin, &cl.key, &cl.value)
		cl.handle(types.Mput)
	case "DELETE", "Delete", "delete":
		fmt.Fscan(cl.stdin, &cl.key)
		cl.handle(types.Mdelete)
	case "PREFIXSCAN", "PrefixScan", "prefixscan":
		fmt.Fscan(cl.stdin, &cl.n)
		cl.handle(types.MprefixScan)
	case "q", "quit", "exit":
		os.Exit(0)
	default:
		fmt.Println("error")
		os.Exit(1)
	}
	cl.stdin.ReadString('\n')
}

func (cl *client) handle(m types.Method) {
	//handle prefixScan
	if m == types.MprefixScan {
		req := &types.PrefixScanRequest{
			N: cl.n,
		}
		reply := &types.PrefixScanReply{}
		err := call("Server.PrefixScanCommand", req, reply)
		if err != nil {
			fmt.Println(err)
		} else {
			for _, key := range reply.Keys {
				fmt.Printf("%s ", key)
			}
			fmt.Println()
		}
		return
	}
	//handle Get,Put,Delete
	req := &types.KVRequest{
		Method: m,
	}
	copy(req.Key[:], cl.key)
	if m == types.Mput {
		//req.Value = cl.value
		copy(req.Value[:], cl.value)
	}
	reply := &types.KVReply{}

	err := call("Server.Command", req, reply)
	if err != nil {
		fmt.Printf("[error] %s\n", err)
	} else if m == types.Mget {
		fmt.Printf("(%s)\n", reply.Value)
	} else {
		fmt.Printf("succeed.\n")
	}
}

//
// send an RPC request to the server, wait for the response.
//
func call(rpcname string, args interface{}, reply interface{}) error {
	c, err := rpc.DialHTTP("tcp", ":8848")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	return err
}
