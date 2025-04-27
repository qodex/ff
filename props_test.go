package ff

import (
	"testing"
)

func TestFsProps(t *testing.T) {
	p := NewFsProps("test/.env")
	p.Set("one", "One...")
	p.Set("two", "Two...")
	p.Set("three", "Three...")

	if len(p.All()) != 3 {
		t.Errorf("all, unexpected count")
	}

	if p.Get("one") != "One..." {
		t.Errorf("get unexpected val")
	}

	p.Delete("two")

	p2 := NewFsProps("test/.env")
	if len(p2.All()) != 2 {
		t.Errorf("all, unexpected count %s", p2.All())
	}
}
