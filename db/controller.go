package db

import (
	"database/sql"
)

// do not edit and use
func (s Serial) Save() error {

	_, err := DB.Exec(
		"INSERT INTO serials (name, season, episode, resolution) VALUES ($1, $2, $3, $4)",
		s.Name, s.Season, s.Episode, s.Resolution,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s Serial) Get() (Serial, error) {
	err := DB.QueryRow(
		"SELECT (season, episode, resolution) FROM serials WHERE name = $1",
		s.Name,
	).Scan(
		&s.Season,
		&s.Episode,
		&s.Resolution,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return Serial{}, nil
		}
		return Serial{}, err
	}

	return s, nil
}

func (s Serial) UpdateSeasonEpisode() error {
	_, err := DB.Exec(
		"UPDATE serials SET (season =  $1, episode = $2) WHERE name = $3",
		s.Season, s.Episode, s.Name,
	)
	if err != nil {
		return err
	}

	return nil
}
