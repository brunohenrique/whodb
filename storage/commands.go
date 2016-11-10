package storage

import (
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

	for i, _ := range ks {
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
