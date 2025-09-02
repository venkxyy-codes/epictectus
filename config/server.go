package config

type Server struct {
	ListenAddress string `yaml:"LISTEN_ADDRESS" env:"http_listen_address"`
	DebugMode     bool   `yaml:"DEBUG_MODE" env:"http_debug_mode"`
}

func (s *Server) SetDefault() {
	s.ListenAddress = ":9090"
}
