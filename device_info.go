package gousb

type DeviceInfo struct {
	AttachedDevice *Device
	DeviceIsOpen   bool
	ProductID      ID
	VendorID       ID
	ProductName    string
	VendorName     string
}

type POSDevices struct {
	printers       []*DeviceInfo
	barcodeReaders []*DeviceInfo
}

func (p *POSDevices) GetPrinters() []*DeviceInfo {
	return p.printers
}

func (p *POSDevices) GetBarcodeReaders() []*DeviceInfo {
	return p.barcodeReaders
}
