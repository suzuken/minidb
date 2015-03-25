// mini database implementation
package minidb

type MiniDB struct {
}

// REPL of minidb
type Console struct {
	Session *Session
}

// persistent session of minidb
type Session struct {
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
func NewClient() (*Client, error){
    return nil, nil
}

// just send query or definition for minidb
func (c *Client) Send(q string) error {
    return nil
}

// saving file
func (c *Client) CreateFile() error {
    return nil
}

// delete db file itself
func (c *Client) DeleteFile() error {
    return nil
}

// create table and write schema into file
func (c *Client) CreateTable(name string, tableinfo TableInfo) (Table, error) {
    return Table{}, nil
}

func (c *Client) DeleteTable() error {
    return nil
}
