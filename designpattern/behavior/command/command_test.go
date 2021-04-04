package command


import "testing"


func TestFactory(t *testing.T) {
	remoteController := &remoteControler{
		commander: &turnonCommand{},
	}
	printer := &printer{ status: "off"}
	remoteController.powerOn(printer)
	if printer.status != "on" {
		t.Error("on")
	}
}

