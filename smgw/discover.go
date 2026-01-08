package smgw

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/pion/mdns/v2"
	"golang.org/x/net/ipv6"
)

const (
	smgwHostname = "smgw.local"
	mDNSTimeout  = 300 * time.Millisecond
)

// Discover discovers the SMGW device using mDNS and returns its address as a string.
// It handles IPv6 link-local addresses by preparing to be used in a URL.
func Discover() (string, error) {
	var packetConnV6 *ipv6.PacketConn
	addr6, err := net.ResolveUDPAddr("udp6", mdns.DefaultAddressIPv6)
	if err != nil {
		return "", err
	}

	l6, err := net.ListenUDP("udp6", addr6)
	if err != nil {
		return "", err
	}

	packetConnV6 = ipv6.NewPacketConn(l6)

	server, err := mdns.Server(nil, packetConnV6, &mdns.Config{})
	if err != nil {
		return "", err
	}
	defer server.Close()

	ctx, cancel := context.WithTimeout(context.Background(), mDNSTimeout)
	defer cancel()

	_, addr, err := server.QueryAddr(ctx, smgwHostname)
	if err != nil {
		return "", err
	}

	if addr.Is6() {
		urlAddr := url.PathEscape(addr.String())
		return fmt.Sprintf("[%s]", urlAddr), nil
	}
	return addr.String(), nil
}
