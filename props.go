package ff

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

type Props interface {
	Get(name string) string
	All() map[string]string
	Delete(name string) error
	Set(name, value string) error
}

type propsFsImpl struct {
	path  string
	props map[string]string
	mu    *sync.Mutex
}

func NewFsProps(path string) Props {
	props := map[string]string{}
	if str, err := ReadString(path); err == nil {
		lines := strings.Split(str, "\n")
		for _, line := range lines {
			kv := strings.Split(line, "=")
			if len(kv) > 1 {
				k := kv[0]
				v := line[len(k)+1:]
				props[k] = v
			}
		}
	}

	fsProps := propsFsImpl{
		path:  path,
		props: props,
		mu:    &sync.Mutex{},
	}

	return fsProps
}

func (t propsFsImpl) save() error {

	data := []byte{}
	for k, v := range t.props {
		data = fmt.Appendf(data, "%s=%v\n", k, v)
	}

	t.mu.Lock()
	defer t.mu.Unlock()
	CreatePath(t.path)
	os.WriteFile(t.path, data, os.ModePerm)

	return nil
}

func (t propsFsImpl) All() map[string]string {
	return t.props
}

func (t propsFsImpl) Get(name string) string {
	return t.props[name]
}

func (t propsFsImpl) Delete(name string) error {
	delete(t.props, name)
	return t.save()
}

func (t propsFsImpl) Set(name, value string) error {
	t.props[name] = value
	return t.save()
}
