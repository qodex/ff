package ff

import (
	"strings"
	"testing"
)

func TestArgsCreate(t *testing.T) {
	ss := strings.Split("./id1 --url https://api.id1.au create test1", " ")
	args := NewOsArgs(ss)
	if val := args.Val("create", ""); val != "test1" {
		t.Errorf("unexpected create: %s", val)
	}
	if val := args.Val("url", ""); val != "https://api.id1.au" {
		t.Errorf("unexpected url: %s", val)
	}
}

func TestArgsSet(t *testing.T) {
	ss := strings.Split("./id1 https://api.id1.au --id test1 --key test.pem set:/test1/pub/name Test One", " ")
	args := NewOsArgs(ss)
	if val := args.Val("id", ""); val != "test1" {
		t.Errorf("unexpected create: %s", val)
	}
	if val := args.WithPrefix("http", ""); val != "https://api.id1.au" {
		t.Errorf("unexpected url: %s", val)
	}
	if val := args.Find(func(arg string) bool {
		return strings.HasPrefix(arg, "set:/")
	}, ""); val != "set:/test1/pub/name" {
		t.Errorf("unexpected cmd: %s", val)
	}
	if val := args.RestAfter("set:/test1/pub/name", ""); val != "Test One" {
		t.Errorf("unexpected cmd data: %s", val)
	}
}
