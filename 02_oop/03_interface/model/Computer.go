package model

type Computer struct {
}

func (c Computer) Work(usb Usb) {
	usb.Start()
	usb.Stop()
}
