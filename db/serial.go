package db

type SerialDB struct {
	Serial
	ID uint
}

type Serial struct {
	Name    string
	Season  int
	Episode int
}
