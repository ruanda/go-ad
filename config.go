package ad

import (
	"errors"
	"fmt"
    "io/ioutil"
	"net"
)

// A Config structure is used to describe Active Directory
type Config struct {
	Domain       string
	BindDN       string
	BindPassword string
	SSL          bool
	RootCA       []byte
}

type ConfigOption func(*Config)

func WithBindDN(binddn, password string) ConfigOption {
	return func(cfg *Config) {
		cfg.BindDN = binddn
		cfg.BindPassword = password
	}
}

func WithInsecure() ConfigOption {
	return func(cfg *Config) {
		cfg.SSL = false
	}
}

func WithCA(rootca []byte) ConfigOption {
	return func(cfg *Config) {
		cfg.RootCA = rootca
	}
}

func WithCAFile(path string) ConfigOption {
    ca, err :=  ioutil.ReadFile(path)
    if err != nil {
        panic(err)
    }
    return WithCA(ca)
}

func NewConfig(domain string, options ...ConfigOption) (*Config, error) {
	var cfg = &Config{
		Domain: domain,
		SSL:    true,
	}
	for _, opt := range options {
		opt(cfg)
	}
	err := cfg.validate()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *Config) validate() error {
	if c.SSL {
		if len(c.RootCA) == 0 {
			return errors.New("define CA or set insecure flag")
		}
	}
	return nil
}

// GetDCs query DNS for domain controllers
func (c *Config) GetDCs() ([]string, error) {
	cname, addrs, err := net.LookupSRV("ldap", "tcp", c.Domain)
	if err != nil {
		return nil, err
	}
	if len(addrs) == 0 {
		return nil, fmt.Errorf("no srv records for domain %s", cname)
	}
	dcs := make([]string, len(addrs))
	for i, v := range addrs {
		dcs[i] = v.Target
	}
	return dcs, nil
}
