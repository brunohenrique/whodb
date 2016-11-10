package storage

import (
	"errors"
	"strconv"
)

type Storage struct {
	state map[string]string
}

func New() Storage {
	return Storage{state: map[string]string{}}
}

func (s *Storage) Set(k, v string) {
	s.state[k] = v
}

func (s *Storage) SetNX(k, v string) bool {
	if _, ok := s.state[k]; ok {
		return false
	}

	s.state[k] = v
	return true
}

func (s *Storage) Get(k string) string { return s.state[k] }

func (s *Storage) Append(k, endS string) int {
	s.state[k] = s.state[k] + endS
	return len(s.state[k])
}

func (s *Storage) Derc(k string) (int, error) {
	v := s.state[k]
	if v == "" {
		v = "0"
	}

	var converted int
	var err error
	if converted, err = strconv.Atoi(v); err != nil {
		return 0, err
	}

	converted--
	v = strconv.Itoa(converted)
	s.state[k] = v

	return converted, nil
}

func (s *Storage) DercBy(k string, by int) (int, error) {
	v := s.state[k]
	if v == "" {
		v = "0"
	}

	var converted int
	var err error
	if converted, err = strconv.Atoi(v); err != nil {
		return 0, err
	}

	converted -= by
	v = strconv.Itoa(converted)
	s.state[k] = v

	return converted, nil
}

func (s *Storage) Incr(k string) (int, error) {
	v := s.state[k]
	if v == "" {
		v = "0"
	}

	var converted int
	var err error
	if converted, err = strconv.Atoi(v); err != nil {
		return 0, err
	}

	converted++
	v = strconv.Itoa(converted)
	s.state[k] = v

	return converted, nil
}

func (s *Storage) IncrBy(k string, by int) (int, error) {
	v := s.state[k]
	if v == "" {
		v = "0"
	}

	var converted int
	var err error
	if converted, err = strconv.Atoi(v); err != nil {
		return 0, err
	}

	converted += by
	v = strconv.Itoa(converted)
	s.state[k] = v

	return converted, nil
}

func (s *Storage) MSet(kvs ...string) {
	for i, _ := range kvs {
		if (i+1)%2 != 0 {
			s.state[kvs[i]] = kvs[i+1]
		}
	}
}

func (s *Storage) MGet(ks ...string) []string {
	var values []string

	for _, k := range ks {
		values = append(values, s.state[k])
	}
	return values
}

func (s *Storage) MSetNX(kvs ...string) bool {
	for i, _ := range kvs {
		if _, ok := s.state[kvs[i]]; ok {
			return false
		}
	}

	for i, _ := range kvs {
		if (i+1)%2 != 0 {
			s.state[kvs[i]] = kvs[i+1]
		}
	}

	return true
}

func (s *Storage) StrLen(k string) int {
	return len(s.state[k])
}

func (s *Storage) Del(ks ...string) int {
	t := 0

	for i := range ks {
		if _, ok := s.state[ks[i]]; ok {
			t++
			delete(s.state, ks[i])
		}
	}

	return t
}

func (s *Storage) FlushAll() {
	s.state = make(map[string]string)
}

func (s *Storage) Exists(k string) bool {
	_, ok := s.state[k]

	return ok
}

func (s *Storage) GetRange(k string, start, end int) string {
	length := len(s.state[k])
	v, ok := s.state[k]
	if !ok {
		return ""
	}

	if start < 0 {
		start = length + start
		if start < 0 {
			start = 0
		}
	}

	if end < 0 {
		end = length + end
	}

	if end > length {
		end = length - 1
	}

	return v[start : end+1]
}

func (s *Storage) Rename(k, nk string) (bool, error) {
	if _, ok := s.state[k]; !ok {
		return false, errors.New("ERR no such key")
	}

	v := s.state[k]
	delete(s.state, k)
	s.state[nk] = v

	return true, nil
}

func (s *Storage) RenameNX(k, nk string) (bool, error) {
	if _, ok := s.state[nk]; ok {
		return false, nil
	}

	if ok, err := s.Rename(k, nk); !ok {
		return ok, err
	}

	return true, nil
}

func (s *Storage) GetSet(k, v string) string {
	oldV := s.state[k]
	s.state[k] = v

	return oldV
}
