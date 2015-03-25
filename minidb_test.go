package minidb

import (
	"testing"
)

func TestSimple(t *testing.T) {
	value := 1
	expected := 2
	if value != expected {
		t.Fatalf("Expected %v, but %v:", expected, value)
	}
}

func TestCreateTable(t *testing.T) {
	c, _ := NewClient()
	err := c.Send("create table hoge(name string);")
	if err != nil {
		t.Fatal("Expected table hoge is not create safely.")
	}
	err1 := c.Send("show tables;")
	if err1 != nil {
		t.Fatal("Expected table hoge is not create safely.")
	}
}

func TestSelectData(t *testing.T) {
}
