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

func TestInsertData(t *testing.T) {
	c, _ := NewClient()
	_, err := c.Exec("drop table if exists hoge")
	if err != nil {
		t.Fatalf("drop table is failed %v", err)
	}
	_, err = c.Exec("create table hoge(id int)")
	if err != nil {
		t.Fatalf("create table is failed %v", err)
	}
	_, err = c.Exec("insert into hoge(id int) values(1)")
	if err != nil {
		t.Fatalf("insert record is failed %v", err)
	}
	ret, err := c.Query("select * from hoge")
	rows := Rows{
		Row: {
			"id": 1,
		},
	}
	if ret != rows {
		t.Fatalf("Record is not inserted properly: %v", r)
	}
}
