package Youpo

type Client struct {
	db *DB

	dbIndex int

	args []string

	argsNum int

	precess *Process
}
