package iface

import (
	"fmt"

	"github.com/pion/transport/v2"
)

// NewWGIFace Creates a new WireGuard interface instance
func NewWGIFace(ifaceName string, address string, mtu int, transportNet transport.Net, args *MobileIFaceArguments) (*WGIface, error) {
	wgAddress, err := parseWGAddress(address)
	if err != nil {
		return nil, err
	}

	wgIFace := &WGIface{
		tun:           newTunDevice(wgAddress, mtu, transportNet, args.TunAdapter),
		userspaceBind: false,
	}
	return wgIFace, nil
}

// CreateOnAndroid creates a new Wireguard interface, sets a given IP and brings it up.
// Will reuse an existing one.
func (w *WGIface) CreateOnAndroid(routes []string, dns string, searchDomains []string) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	cfgr, err := w.tun.Create(routes, dns, searchDomains)
	if err != nil {
		return err
	}
	w.configurer = cfgr
	return nil
}

// Create this function make sense on mobile only
func (w *WGIface) Create() error {
	return fmt.Errorf("this function has not implemented on mobile")
}
