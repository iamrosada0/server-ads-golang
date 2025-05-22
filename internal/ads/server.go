package ads

import (
	"log"
	"net"
	"net/http"

	realip "github.com/ferluci/fast-realip"
	"github.com/mssola/user_agent"
	"github.com/oschwald/geoip2-golang"
	"github.com/valyala/fasthttp"
)

type Server struct {
	geoip *geoip2.Reader
}

func NewServer(geoip *geoip2.Reader) *Server {
	return &Server{geoip: geoip}
}

func (s *Server) Listen() error {
	return fasthttp.ListenAndServe(":8080", s.handler)
}

func (s *Server) handler(ctx *fasthttp.RequestCtx) {
	remoteIp := realip.FromRequest(ctx)
	ua := string(ctx.Request.Header.UserAgent())

	parsed := user_agent.New(ua)
	browserName, _ := parsed.Browser()

	country, err := s.geoip.Country(net.ParseIP(remoteIp))
	if err != nil {
		log.Printf("Failed to parse country: %v", err)
		return
	}

	user := &User{
		Country: country.Country.IsoCode,
		Browser: browserName,
	}
	campaigns := GetStaticCampaigns()
	log.Printf("User IP: %s | Country: %s | Browser: %s", remoteIp, country.Country.IsoCode, browserName)

	winner := MakeAuction(campaigns, user)
	if winner == nil {
		ctx.Redirect("https://example.com", http.StatusSeeOther)
		return
	}

	//  Verificar se o parâmetro ?preview=true foi passado
	args := ctx.QueryArgs()
	if args.Has("preview") && string(args.Peek("preview")) == "true" {
		ctx.SetStatusCode(http.StatusOK)
		ctx.SetContentType("text/plain")
		ctx.SetBodyString("Winning URL: " + winner.ClickUrl)
		return
	}

	ctx.Redirect(winner.ClickUrl, http.StatusSeeOther)
}
