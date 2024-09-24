package adapter

import "fmt"

type Mobile interface {
	chargeMobile()
}

type ApplePhone struct{}

func (a *ApplePhone) chargeMobile() {
	fmt.Println("Apple phone is charging")
}

type AndroidPhone struct{}

func (a *AndroidPhone) chargeMobile() {
	fmt.Println("Android phone is charging")
}

type WindowsPhone struct {
	IsFastChargeSupported bool
}

func (w *WindowsPhone) fastChargeMobile() {
	fmt.Println("Windows phone is fast charging")
}

func (w *WindowsPhone) chargeMobile() {
	fmt.Println("Windows phone is charging")
}

type WindowsPhoneAdapter struct {
	WindowsPhone *WindowsPhone
}

func (wd *WindowsPhoneAdapter) chargeMobile() {
	if wd.WindowsPhone.IsFastChargeSupported {
		wd.WindowsPhone.fastChargeMobile()
	} else {
		wd.WindowsPhone.chargeMobile()
	}
}

type Client struct{}

func (c *Client) Charge(mobile Mobile) {
	mobile.chargeMobile()
}
