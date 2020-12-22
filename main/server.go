package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"woodkv/e"
	"woodkv/store"
	"woodkv/store/skiplist"
	"woodkv/types"
)

type Server struct {
	storage store.Store
}

func main() {
	s := Server{
		storage: skiplist.New(),
	}
	s.serve()
	fmt.Println("server is running...")
}

//Serve start a rpc server
func (s *Server) serve() {
	rpc.HandleHTTP()
	rpc.Register(s)
	l, err := net.Listen("tcp", ":8848")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	go http.Serve(l, nil)
	select {}
}

//Command is a rpc function called by client
//Handle Get, Put, Delete
func (s *Server) Command(request *types.KVRequest, reply *types.KVReply) error {
	var err error
	var val store.Value

	switch request.Method {
	case types.Mget:
		fmt.Printf("%s %s.\n", request.Method, request.Key)
		val, err = s.storage.Get(request.Key)
		if err != nil {
			break
		}
		reply.Value = val.(types.Byte256)
	case types.Mput:
		fmt.Printf("%s %s: (%s).\n", request.Method, request.Key, request.Value)
		err = s.storage.Put(request.Key, request.Value)
	case types.Mdelete:
		fmt.Printf("%s %s.\n", request.Method, request.Key)
		err = s.storage.Delete(request.Key)
	case types.MprefixScan:
		fmt.Printf("%s %s.\n", request.Method, request.Key)
	default:
		err = e.Unknown
	}

	if err != nil {
		log.Printf("[error] %s\n", err)
	} else {
		log.Printf("[%s] succeed.\n", request.Method)
	}
	return err
}

// PrefixScanCommand is a rpc function called by client
// Handle prefixScan method
func (s *Server) PrefixScanCommand(request *types.PrefixScanRequest, reply *types.PrefixScanReply) error {
	keys := s.storage.PrefixScan(request.N)
	for _, key := range keys {
		reply.Keys = append(reply.Keys, key.(types.Byte8))
	}
	return nil
}
