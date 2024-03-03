package server

type Config struct {
	ListenAddrGrpc string `yaml:"listen_addr_grpc"`
	ListenAddrHttp string `yaml:"listen_addr_http"`
	ListenGrpc bool `yaml:"listen_grpc"`
	ListenHttp bool `yaml:"listen_http"`
}