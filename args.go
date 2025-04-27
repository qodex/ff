package ff

import (
	"slices"
	"strings"
)

type OsArgs struct {
	args []string
}

func NewOsArgs(args []string) OsArgs {
	osargs := OsArgs{
		args: args,
	}
	return osargs
}

func (t OsArgs) Val(name, value string) string {
	name = strings.TrimLeft(name, "-")
	for i, arg := range t.args {
		arg = strings.TrimLeft(arg, "-")
		if arg == name && len(t.args) > i+1 {
			return t.args[i+1]
		}
	}
	return value
}

func (t OsArgs) WithPrefix(prefix, value string) string {
	for _, arg := range t.args {
		if strings.HasPrefix(arg, prefix) {
			return arg
		}
	}
	return value
}

func (t OsArgs) WithSuffix(suffix, value string) string {
	for _, arg := range t.args {
		if strings.HasSuffix(arg, suffix) {
			return arg
		}
	}
	return value
}

func (t OsArgs) Has(name string) bool {
	return slices.Contains(t.args, name)
}

func (t OsArgs) Find(fn func(arg string) bool, value string) string {
	for _, arg := range t.args {
		if fn(arg) {
			return arg
		}
	}
	return value
}

func (t OsArgs) RestAfter(name, value string) string {
	for i, arg := range t.args {
		if arg == name && len(t.args) > i {
			return strings.Join(t.args[i+1:], " ")
		}
	}
	return value
}
