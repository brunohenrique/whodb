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

func TestExists(t *testing.T) {
	s := storage.New()

	s.Set("key", "value")

	if v := s.Exists("key"); !v {
		t.Errorf("must return true when existing key")
	}

	if v := s.Exists("nkey"); v {
		t.Errorf("must return false when non existent key")
	}
}

func TestGetRange(t *testing.T) {
	s := storage.New()

	s.Set("key", "This is a string")

	if v := s.GetRange("key", 0, 3); v != "This" {
		t.Errorf("must return 'This' but '%s' was returned", v)
	}

	if v := s.GetRange("key", -3, -1); v != "ing" {
		t.Errorf("must return 'ing' but '%s' was returned", v)
	}

	if v := s.GetRange("key", 0, -1); v != "This is a string" {
		t.Errorf("must return 'This is a string' but '%s' was returned", v)
	}

	if v := s.GetRange("key", 10, 100); v != "string" {
		t.Errorf("must return 'string' but '%s' was returned", v)
	}

	if v := s.GetRange("key", -100, 100); v != "This is a string" {
		t.Errorf("must return 'This is a string' but '%s' was returned", v)
	}

	if v := s.GetRange("key", 0, 0); v != "T" {
		t.Errorf("must return 'T' but '%s' was returned", v)
	}

	if v := s.GetRange("nkey", 0, 2); v != "" {
		t.Errorf("must return an empty string but '%s' was returned", v)
	}
}

func TestRename(t *testing.T) {
	s := storage.New()

	s.Set("key", "value")

	if ok, _ := s.Rename("key", "newKey"); !ok {
		t.Errorf("must return true and renamed key")
	}

	if v := s.Get("newKey"); v != "value" {
		t.Errorf("newKey must contain the old key value")
	}

	if v := s.Get("key"); v != "" {
		t.Errorf("key must have been deleted")
	}

	if _, err := s.Rename("nkey", "key1"); err == nil {
		t.Errorf("must return an error when non existent key")
	}
}

func TestRenameNX(t *testing.T) {
	s := storage.New()

	s.MSet("key1", "value1", "key2", "value2", "key3", "value3")
	if ok, _ := s.RenameNX("key1", "newKey"); !ok {
		t.Errorf("must return true and renamed key")
	}

	if ok, _ := s.RenameNX("key2", "key3"); ok {
		t.Errorf("must return false when have conflicting keys")
	}

	if v := s.Get("newKey"); v != "value1" {
		t.Errorf("newKey must contain the old key value")
	}

	if v := s.Get("key1"); v != "" {
		t.Errorf("key1 must have been deleted")
	}

	if _, err := s.RenameNX("nkey", "key1"); err == nil {
		t.Errorf("must return an error when non existent key")
	}
}

func TestGetSet(t *testing.T) {
	s := storage.New()

	if v := s.GetSet("key", "value1"); v != "" {
		t.Errorf("must return an empty string")
	}

	if v := s.Get("key"); v != "value1" {
		t.Errorf("must return 'value1'")
	}

	if v := s.GetSet("key", "value2"); v != "value1" {
		t.Errorf("must return old value")
	}
}
