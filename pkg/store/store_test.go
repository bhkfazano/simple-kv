package store

import "testing"

func TestNewSimpleKV(t *testing.T) {
	s := NewSimpleKV[string, string]()

	if s == nil {
		t.Error("NewSimpleKV() returned nil")
	}
}

func TestPut(t *testing.T) {
	s := NewSimpleKV[string, string]()

	var err = s.Put("key", "value")
	if err != nil {
		t.Errorf("Put() returned an error: %s", err)
	}
}

func TestGet(t *testing.T) {
	var s = NewSimpleKV[string, string]()

	var key string = "foo"
	var value string = "bar"

	s.Put(key, value)
	value, err := s.Get(key)

	if err != nil {
		t.Errorf("Get() returned an error: %s", err)
	}
}

func TestGetNotFound(t *testing.T) {
	var s = NewSimpleKV[string, string]()

	var key string = "foo"
	_, err := s.Get(key)

	if err == nil {
		t.Error("Get() did not return an error for a missing key")
	}
}

func TestUpdate(t *testing.T) {
	var s = NewSimpleKV[string, string]()

	var key string = "foo"
	var value string = "bar"

	s.Put(key, value)

	var newValue string = "baz"
	err := s.Update(key, newValue)

	if err != nil {
		t.Errorf("Update() returned an error: %s", err)
	}
}

func TestUpdateNotFound(t *testing.T) {
	var s = NewSimpleKV[string, string]()

	var keyNotFound string = "baz"
	var value string = "bar"

	err := s.Update(keyNotFound, value)

	if err == nil {
		t.Error("Update() did not return an error for a missing key")
	}
}

func TestDelete(t *testing.T) {
	var s = NewSimpleKV[string, string]()

	var key string = "foo"
	var value string = "bar"

	s.Put(key, value)

	_, err := s.Delete(key)
	if err != nil {
		t.Errorf("Delete() returned an error: %s", err)
	}
}

func TestDeleteNotFound(t *testing.T) {
	var s = NewSimpleKV[string, string]()

	var keyNotFound string = "baz"

	_, err := s.Delete(keyNotFound)
	if err == nil {
		t.Error("Delete() did not return an error for a missing key")
	}
}
