package strategy

type iDiscount interface {
	discount() int
}

type ruleChild struct {

}

func (c *ruleChild) discount() int {
	return -1000
}

type ruleAdult struct {

}

func (c *ruleAdult) discount() int {
	return 0
}

type busChager struct {
	basePrice int
}

func (b *busChager) getPrice(d iDiscount) int {
	return b.basePrice - d.discount()
}