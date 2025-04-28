package ff

import (
	"strings"
	"testing"
)

func TestArgsVal(t *testing.T) {
	ss := strings.Split("./id1 --url https://api.id1.au create test1", " ")
	args := NewOsArgs(ss)
	if val := args.Val("create", ""); val != "test1" {
		t.Errorf("unexpected create: %s", val)
	}
	if val := args.Val("url", ""); val != "https://api.id1.au" {
		t.Errorf("unexpected url: %s", val)
	}
}

func TestArgsKeyVal(t *testing.T) {
	ss := strings.Split("./id1 env set id=test1 and get var1 uno=\"dos\" --var prop1=\"Ein Zwei\" \"\"...\"", " ")
	args := NewOsArgs(ss)
	if key, val := args.KeyVal("set", "a", "b"); key != "id" || val != "test1" {
		t.Errorf("unexpected key/value: %s=%s", key, val)
	}
	if key, val := args.KeyVal("get", "a", "b"); key != "var1" || val != "" {
		t.Errorf("unexpected key/value: %s=%s", key, val)
	}
	if key, val := args.KeyVal("var", "a", "b"); key != "prop1" || val != "Ein Zwei" {
		t.Errorf("unexpected key/value: %s=%s", key, val)
	}

}

func TestArgs(t *testing.T) {
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

func TestArgsLast(t *testing.T) {
	ss := strings.Split("./id1 https://api.id1.au --id test1 --key test.pem set:/test1/pub/name Test One", " ")
	args := NewOsArgs(ss)
	if val := args.Last(); val != "One" {
		t.Errorf("unexpected last: %s", val)
	}

}
