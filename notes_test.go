package ff

import (
	"log"
	"testing"
)

func TestNotesFsImpl(t *testing.T) {
	s := NewFsNotes("temp/notes")
	for k, _ := range s.All("mynotes") {
		log.Println(k)
	}
	s.Set("mynotes", "note1", "Test...")
	s.Set("mynotes", "note2", "Test Two...")
	s.Set("mynotes", "note3", "Test Three...")
	log.Println(s.All("mynotes"))
	s.Append("mynotes", "note1", "with an addition")
	log.Println(s.Get("mynotes", "note1"))
	s.Delete("mynotes", "note1")

}
