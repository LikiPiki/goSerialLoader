package db

type Serial struct {
	Name    string
	Season  int
	Episode int
}

type SerialDB struct {
	Serial
	ID uint
}
