package inmem

import 
	 "database/sql/driver"
type inMemCoonection struct{}                                      //driver.Conn
func (*inMemCoonection) Prepare(query string) (driver.Stmt, error) {}
func (*inMemCoonection) Close() error {
	return nil
}
func (*inMemCoonection) Begin() (driver.Tx, error) {
	return new(inMemTransaction), nil
}

type inMemStatement struct{} //driver.Stmt
// Close closes the statement.
func (*inMemStatement) Close() error {}

// NumInput()  the sql package will sanity check
// argument counts from callers and return errors to the callers
// before the statement's Exec or Query methods are called.
func (*inMemStatement) NumInput() int {
	return 0
}

// Exec executes a query that doesn't return rows, such
// as an INSERT or UPDATE.
func (*inMemStatement) Exec(args []driver.Value) (driver.Result, error)

// Query executes a query that may return rows, such as a
// SELECT.
func (*inMemStatement) Query(args []driver.Value) (driver.Rows, error)

type inMemResult struct{} //Result interface
func (*inMemResult) LastInsertId() (int64, error)
func (*inMemResult) RowsAffected() (int64, error)

type inMemRows struct{}

// Columns returns the names of the columns. The number of
// columns of the result is inferred from the length of the
// slice. If a particular column name isn't known, an empty
// string should be returned for that entry.
func (*inMemRows) Columns() []string

// Close closes the rows iterator.
func (*inMemRows) Close() error

// Next is called to populate the next row of data into
// the provided slice. The provided slice will be the same
// size as the Columns() are wide.
// Next should return io.EOF when there are no more rows.
func (*inMemRows) Next(dest []driver.Value) error

//inMem dosn't support transaction
type inMemTransaction struct{}

func (inMemTransaction) Commit() error   { return nil }
func (inMemTransaction) Rollback() error { return nil }

func (memDB *inMemDB) Open(string) (driver.Conn, error) {
}
