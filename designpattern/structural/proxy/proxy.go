package proxy

type iProxyService interface {
	connectSite() string
}

type proxyService struct {

}

func (p *proxyService) connectSite() string {
	return "connected"
}

type proxyPattern struct {
	proxy iProxyService
}

func (p *proxyPattern) run() string {
	return p.proxy.connectSite()
}