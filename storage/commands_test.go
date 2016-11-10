package storage_test

import (
	"testing"

	"github.com/brunohenrique/whodb/storage"
)

func TestSetAndGet(t *testing.T) {
	s := storage.New()

	if v := s.Get("key"); v != "" {
		t.Errorf("must return an empty string")
	}

	s.Set("key", "value")
	if v := s.Get("key"); v != "value" {
		t.Errorf("must return 'value'")
	}
}

func TestSetNX(t *testing.T) {
	s := storage.New()

	if ok := s.SetNX("key", "value"); !ok {
		t.Errorf("must return true")
	}

	if ok := s.SetNX("key", "new value"); ok {
		t.Errorf("must return false")
	}
}

func TestAppend(t *testing.T) {
	s := storage.New()

	bytesNumber := s.Append("key", " in the end")

	if bytesNumber != 11 {
		t.Errorf("wrong bytes number")
	}

	if v := s.Get("key"); v != " in the end" {
		t.Errorf("wrong string retuned")
	}
}

func TestDecr(t *testing.T) {
	s := storage.New()

	if v, _ := s.Derc("key"); v != -1 {
		t.Errorf("must return -1")
	}

	s.Set("key", "10")
	if v, _ := s.Derc("key"); v != 9 {
		t.Errorf("must return 9")
	}

	s.Set("key", "value")
	if _, err := s.Derc("key"); err == nil {
		t.Errorf("must return an error")
	}
}

func TestDecrBy(t *testing.T) {
	s := storage.New()

	if v, _ := s.DercBy("key", 3); v != -3 {
		t.Errorf("must return -3")
	}

	s.Set("key", "10")
	if v, _ := s.DercBy("key", 3); v != 7 {
		t.Errorf("must return 3")
	}

	s.Set("key", "value")
	if _, err := s.DercBy("key", 3); err == nil {
		t.Errorf("must return an error")
	}
}

func TestIncr(t *testing.T) {
	s := storage.New()

	if v, _ := s.Incr("key"); v != 1 {
		t.Errorf("must return 1")
	}

	s.Set("key", "10")
	if v, _ := s.Incr("key"); v != 11 {
		t.Errorf("must return 11")
	}

	s.Set("key", "value")
	if _, err := s.Incr("key"); err == nil {
		t.Errorf("must return an error")
	}
}

func TestIncrBy(t *testing.T) {
	s := storage.New()

	if v, _ := s.IncrBy("key", 3); v != 3 {
		t.Errorf("must return -3")
	}

	s.Set("key", "10")
	if v, _ := s.IncrBy("key", 3); v != 13 {
		t.Errorf("must return 3")
	}

	s.Set("key", "value")
	if _, err := s.IncrBy("key", 3); err == nil {
		t.Errorf("must return an error")
	}
}

func TestMSetAndMGet(t *testing.T) {
	s := storage.New()

	s.MSet("key1", "value1", "key2", "value2")
	vs := s.MGet("key1", "key2")
	expected := []string{"value1", "value2"}
	for i, _ := range vs {
		if vs[i] != expected[i] {
			t.Errorf("wrong values, got %s expected %s", vs[i], expected[i])
		}
	}
}

func TestMSetNX(t *testing.T) {
	s := storage.New()

	if ok := s.MSetNX("key1", "value1", "key2", "value2"); !ok {
		t.Errorf("must return 1")
	}

	if ok := s.MSetNX("key1", "value1", "key3", "value3"); ok {
		t.Errorf("must return 0")
	}
}

func TestStrLen(t *testing.T) {
	s := storage.New()

	s.Set("key", "value")
	if length := s.StrLen("key"); length != 5 {
		t.Error("must return 5")
	}
}

func TestDel(t *testing.T) {
	s := storage.New()

	s.MSetNX("key1", "value1", "key2", "value2")
	if v := s.Del("key1", "key3"); v != 1 {
		t.Errorf("must return the total of deleted keys")
	}

	if v := s.Get("key1"); v != "" {
		t.Errorf("key1 must have been deleted")
	}

	if v := s.Get("key2"); v != "value2" {
		t.Errorf("key2 must not have been deleted")
	}
}

func TestFlushAll(t *testing.T) {
	s := storage.New()

	s.MSetNX("key1", "value1", "key2", "value2")
	s.FlushAll()

	keys := [2]string{"key1", "key2"}
	for i, _ := range keys {
		if v := s.Get(keys[i]); v != "" {
			t.Errorf("%s must have been deleted", keys[i])
		}
	}
}
