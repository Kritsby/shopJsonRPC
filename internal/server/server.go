package server

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HttpConn struct {
	in  io.Reader
	out io.Writer
}

func (c *HttpConn) Read(p []byte) (n int, err error)  { return c.in.Read(p) }
func (c *HttpConn) Write(d []byte) (n int, err error) { return c.out.Write(d) }
func (c *HttpConn) Close() error                      { return nil }

type Listener struct {
	Listener net.Listener
	rpcS     *rpc.Server
}

func (l *Listener) New(port string, server *rpc.Server) {
	fmt.Println(port)
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal().Err(err)
	}

	l.Listener = listener

	err = http.Serve(listener, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/test" {
			serverCodec := jsonrpc.NewServerCodec(&HttpConn{in: r.Body, out: w})
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(200)
			err := server.ServeRequest(serverCodec)
			if err != nil {
				log.Printf("Error while serving JSON request: %v", err)
				http.Error(w, "Error while serving JSON request, details have been logged.", 500)
				return
			}
		}
	}))

	log.Error().Err(err)
}
