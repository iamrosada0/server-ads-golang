package ads

import "github.com/valyala/fasthttp"

type Server struct {
}

func NewService() *Server {
	return &Server{}
}

func (s *Server) Listen() error {
	return fasthttp.ListenAndServe(":8080")
}
