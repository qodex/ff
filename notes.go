package ff

import (
	"log"
	"os"
	"path/filepath"
)

type Notes interface {
	Get(folder string, noteId string) string
	All(folder string) map[string]string
	Delete(folder string, noteId string) error
	Set(folder string, noteId string, value string) error
	Append(folder string, noteId string, value string) error
}

type NotesFsImpl struct {
	Dir string
}

func (t NotesFsImpl) All(folder string) map[string]string {
	result := map[string]string{}
	if dir, err := os.Open(filepath.Join(t.Dir, folder)); err != nil {
		// no such folder
	} else if files, err := dir.Readdir(0); err == nil {
		for _, file := range files {
			result[file.Name()] = t.Get(folder, file.Name())
		}
	}
	return result
}

func (t NotesFsImpl) Get(folder string, noteId string) string {
	if data, err := os.ReadFile(filepath.Join(t.Dir, folder, noteId)); err == nil {
		return string(data)
	} else {
		return ""
	}
}

func (t NotesFsImpl) Delete(folder string, noteId string) error {
	if len(noteId) > 0 {
		return os.Remove(filepath.Join(t.Dir, folder, noteId))
	} else {
		return os.RemoveAll(filepath.Join(t.Dir, folder))
	}
}

func (t NotesFsImpl) Set(folder string, noteId string, value string) error {
	if err := CreatePath(filepath.Join(t.Dir, folder, noteId)); err != nil {
		return err
	} else {
		return os.WriteFile(filepath.Join(t.Dir, folder, noteId), []byte(value), 0644)
	}
}

func (t NotesFsImpl) Append(folder string, noteId string, value string) error {
	if err := CreatePath(filepath.Join(t.Dir, folder, noteId)); err != nil {
		return err
	}

	// open for append
	if file, err := os.OpenFile(filepath.Join(t.Dir, folder, noteId), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		log.Println(err)
	} else {
		defer file.Close()
		if _, err := file.WriteString(value); err != nil {
			log.Println(err)
		}
	}
	return nil
}

func NewFsNotes(dir string) NotesFsImpl {
	return NotesFsImpl{
		Dir: dir,
	}
}
