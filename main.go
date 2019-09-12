package main

import (
	"fmt"
	"log"

	"github.com/alexflint/go-arg"
	"github.com/armon/go-socks5"
)

type config struct {
	Username string
	Password string
	Port     uint16
}

func main() {
	cfg := config{
		Username: "fuckrkn",
		Password: "fuckrkn",
		Port:     22869,
	}
	arg.MustParse(&cfg)

	auth := []socks5.Authenticator{
		socks5.UserPassAuthenticator{
			Credentials: socks5.StaticCredentials{
				cfg.Username: cfg.Password,
			},
		},
	}

	conf := &socks5.Config{
		AuthMethods: auth,
	}

	server, err := socks5.New(conf)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting proxy on port %d...", cfg.Port)
	if err := server.ListenAndServe("tcp", fmt.Sprintf(":%d", cfg.Port)); err != nil {
		log.Fatal(err)
	}
}
