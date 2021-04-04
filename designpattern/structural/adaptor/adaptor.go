package adaptor


type toTypeC interface {
	connect() string
}

type usb1 struct {
}

func (u *usb1) connect() string {
	return "usb1 연결"
}

type thunderbird struct {
}

func (u *thunderbird) connect() string {
	return "썬더보드 연결"
}

type computer struct {

}

func (c *computer) connectCable(typeC toTypeC) string {
	return typeC.connect()
}



