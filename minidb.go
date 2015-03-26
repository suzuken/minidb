// mini database implementation
package minidb

type DB struct {
}

// REPL of minidb
type Console struct {
	Session *Session
}

// persistent session of minidb
type Session struct {
}

// following database/sql
type Client interface {
	Close() error
	Exec(query string, args ...interface{}) (Result, error)
	Prepare(query string) (*Stmt, error)
	Query(query string, args ...interface{}) (*Rows, error)
}

type Stmt struct {
}

type Rows struct {
}

// managing connection with minidb
type Client struct {
}

// connection object of minidb
type Connection struct {
}

type Query struct {
}

type Page struct {
}

// each record
type Record struct {
}

// table
type Table struct {
	TableInfo TableInfo
}

// meta data of table
type TableInfo struct {
	FieldInfo []FieldInfo
}

type FieldInfo struct {
	name string
	// int, float, etc. it supports all data types in golang
	// TODO cast to datatype string to golang datatype itself
	// only support boolean, string, numeric
	DataType string
}

// for saving
type File struct {
	fd   int
	path string
}

// create new client
func NewClient() (*Client, error) {
	return nil, nil
}

// just send query or definition for minidb
func (c *Client) Exec(q string) (Result, error) {
    // TODO parsing query, and transfer responsibility to each function
	return Result{}, nil
}

// just send query or definition for minidb
func (c *Client) Query(q string) (*Rows, error) {
    // TODO parsing query, and transfer responsibility to each function
	return &Rows{}, nil
}

// create table and write schema into file
func (c *Client) CreateTable(name string, tableinfo TableInfo) (Table, error) {
	return Table{}, nil
}

func (c *Client) DeleteTable() error {
	return nil
}

// saving file
func (d *DB) CreateFile() error {
	return nil
}

// delete db file itself
func (d *DB) DeleteFile() error {
	return nil
}

