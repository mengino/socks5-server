package main

import (
	"log"
	"net"
	"os"

	"github.com/mengino/go-socks5"
)

type config struct {
	User     string
	Password string
	Port     string
	IP       string
}

func main() {
	cfg := config{
		User:     os.Getenv("PROXY_USER"),
		Password: os.Getenv("PROXY_PASSWORD"),
		Port:     os.Getenv("PROXY_PORT"),
		IP:       os.Getenv("PROXY_IP"),
	}

	socsk5conf := &socks5.Config{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}

	if cfg.User+cfg.Password != "" {
		creds := socks5.StaticCredentials{
			cfg.User: cfg.Password,
		}
		cator := socks5.UserPassAuthenticator{Credentials: creds}
		socsk5conf.AuthMethods = []socks5.Authenticator{cator}
	}

	if cfg.IP != "" {
		bindIP, _, err := net.ParseCIDR(cfg.IP + "/24")
		if err != nil {
			log.Fatal(err)
		}

		socsk5conf.BindIP = bindIP
	}

	server, err := socks5.New(socsk5conf)
	if err != nil {
		log.Fatal(err)
	}

	if cfg.Port == "" {
		cfg.Port = "1080"
	}

	log.Printf("Start listening proxy service on port %s\n", cfg.Port)
	if err := server.ListenAndServe("tcp", ":"+cfg.Port); err != nil {
		log.Fatal(err)
	}
}
