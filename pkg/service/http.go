package service

type HttpConf struct {
	Addr      string `toml:"addr"`
	AddrHttps string `toml:"addrHttps"`
}
