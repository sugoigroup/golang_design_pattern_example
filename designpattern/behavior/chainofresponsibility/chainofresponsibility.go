package chainofresponsibility

type process interface {
	canBill(amount int) bool
	runBill(amount int)
	next() bankinfo
}

type bankinfo interface {
	process
	setNextBank(nb bankinfo)
}

type bank1 struct {
	remain int
	nextBank bankinfo
}
func (b *bank1) next() bankinfo {
	return b.nextBank
}

func (b *bank1) setNextBank(nb bankinfo)  {
	b.nextBank = nb
}

func (b *bank1) canBill(amount int) bool {
	return  b.remain - amount > 0
}

func (b *bank1) runBill(amount int)  {
	b.remain = b.remain - amount
}
type bank2 struct {
	remain int
	nextBank bankinfo
}
func (b *bank2) next() bankinfo {
	return b.nextBank
}


func (b *bank2) setNextBank(nb bankinfo)  {
	b.nextBank = nb
}

func (b *bank2) canBill(amount int) bool {
	return  b.remain - amount > 0
}

func (b *bank2) runBill(amount int)  {
	b.remain = b.remain - amount
}
type bank3 struct {
	remain int
	nextBank bankinfo
}
func (b *bank3) next() bankinfo {
	return b.nextBank
}


func (b *bank3) setNextBank(nb bankinfo)  {
	b.nextBank = nb
}

func (b *bank3) canBill(amount int) bool {
	return  b.remain - amount > 0
}

func (b *bank3) runBill(amount int)  {
	b.remain = b.remain - amount
}

type chainOfBank struct {
	bankinfo
}

func (c *chainOfBank) next()  {
	c.bankinfo = c.bankinfo.next()
}



