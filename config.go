package ad

import (
    "fmt"
    "net"
)

// A Config structure is used to describe Active Directory
type Config struct {
    Domain  string
}

// GetDCs query DNS for domain controllers
func (c Config) GetDCs() ([]string, error) {
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
