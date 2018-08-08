package db

func (s Serial) Put(resolution string) error {

	_, err := DB.Exec(
		"INSERT INTO serials (name, season, episode, resolution) VALUES ($1, $2, $3, $4)",
		s.Name, s.Season, s.Episode, resolution,
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

func (s Serial) GetResolution() (string, error) {
	var resolution string

	err := DB.QueryRow(
		"SELECT resolution FROM serials WHERE name = $1",
		s.Name,
	).Scan(
		&resolution,
	)
	if err != nil {
		return "", err
	}

	return resolution, nil
}

func (s Serial) Set() error {
	_, err := DB.Exec(
		"UPDATE serials SET season =  $1, episode = $2 WHERE name = $3",
		s.Season, s.Episode, s.Name,
	)
	if err != nil {
		return err
	}

	return nil
}
