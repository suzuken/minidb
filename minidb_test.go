package minidb

import (
	"testing"
)

func TestCreateTable(t *testing.T) {
	c, _ := NewClient()
	_, err := c.Exec("create table hoge(name string);")
	if err != nil {
		t.Fatal("Expected table hoge is not create safely.")
	}
	_, err1 := c.Exec("show tables;")
	if err1 != nil {
		t.Fatal("Expected table hoge is not create safely.")
	}
}

func TestSelectData(t *testing.T) {
	c, _ := NewClient()
	resp, err := c.Exec("create table hoge(name string)")
	if err != nil {
		t.Fatal("Request create table statement is not worked proper.")
	}
}
