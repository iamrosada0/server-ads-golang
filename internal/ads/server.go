package ads

import "github.com/valyala/fasthttp"

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Listen() error {
	return fasthttp.ListenAndServe(":8080", s.handleHttp)
}

func (s *Server) handleHttp(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Hello, world")
}
