package ad_test

import (
    "bytes"
	"testing"

	"github.com/ruanda/go-ad"
)

func TestNewConfig(t *testing.T) {
	_, err := ad.NewConfig("some.domain")
	if err == nil {
		t.Error("NewConfig should return error")
	}
	_, err = ad.NewConfig("some.domain", ad.WithInsecure())
	if err != nil {
		t.Errorf("NewConfig should work, got: %v", err)
	}
    _, err = ad.NewConfig("some.domain", ad.WithCA([]byte{}))
    if err == nil {
        t.Error("NewConfig should return error")
    }
    cfg, err := ad.NewConfig("some.domain", ad.WithCAFile("_test/bogus_ca.pem"))
    if err != nil {
		t.Errorf("NewConfig should work, got: %v", err)
    }
    test_ca := []byte("test")
    if !bytes.Equal(test_ca, cfg.RootCA) {
        t.Errorf("NewConfig did not set RootCA: got %v, should be: %v", cfg.RootCA, test_ca)
    }
}
