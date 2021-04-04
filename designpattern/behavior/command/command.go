package command

type parts interface {
	on()
	off()
}

type printer struct {
	status string
}

func (p *printer) on()  {
	p.status = "on"
}

func (p *printer) off()  {
	p.status = "off"
}


type commander interface {
	execute(p parts)
}
type turnonCommand struct {

}

func (t *turnonCommand) execute(p parts)  {
	p.on()
}

type remoteControler struct {
	commander
}

func (r *remoteControler) powerOn(p parts)  {
	r.commander.execute(p)
}
