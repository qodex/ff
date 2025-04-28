package ff

import (
	"fmt"
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

func (t OsArgs) KeyVal(name, key, value string) (string, string) {
	val := t.Val(name, fmt.Sprintf("%s=%s", key, value))
	key = strings.Split(val, "=")[0]
	value = strings.TrimLeft(val[len(key):], "=")
	if strings.HasPrefix(value, "\"") {
		value = strings.Split(t.RestAfter(name, value), "\"")[1]
	}
	return key, value
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
	name = fmt.Sprintf("%s ", name)
	str := fmt.Sprintf("%s ", strings.Join(t.args, " "))
	if strings.Contains(str, name) {
		return strings.Trim(str[strings.Index(str, name)+len(name):], " ")
	} else {
		return value
	}
}

func (t OsArgs) Last() string {
	if len(t.args) > 0 {
		return t.args[len(t.args)-1]
	} else {
		return ""
	}
}
