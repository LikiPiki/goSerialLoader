package db

func (s Serial) Put() error {

	_, err := DB.Exec(
		"INSERT INTO serials (name, season, episode) VALUES ($1, $2, $3)",
		s.Name, s.Season, s.Episode,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s Serial) Get() (int, int, error) {
	err := DB.QueryRow(
		"SELECT season, episode FROM serials WHERE name = $1",
		s.Name,
	).Scan(
		&s.Season,
		&s.Episode,
	)
	if err != nil {
		return 0, 0, err
	}

	return s.Season, s.Episode, nil
}
