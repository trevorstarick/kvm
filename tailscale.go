package kvm

import (
	"cmp"

	"tailscale.com/tsnet"
)

func RunTailscaleServer() {
	LoadConfig()

	if config.Tailscale == nil {
		return
	}

	hostname := cmp.Or(config.Tailscale.Hostname, "jetkvm")
	addr := cmp.Or(config.Tailscale.Addr, ":80")

	s := new(tsnet.Server)
	s.Hostname = hostname
	defer s.Close()
	ln, err := s.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	r := setupRouter()

	err = r.RunListener(ln)
	if err != nil {
		panic(err)
	}
}
