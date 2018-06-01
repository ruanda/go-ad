package ad_test

import (
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
		t.Errorf("Config should work, got: %v", err)
	}
}
