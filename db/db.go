package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type DB struct {
	DB *sql.DB
}

func NewDb(cnf mysql.Config) (*DB, error) {
	db, err := sql.Open("mysql", cnf.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &DB{DB: db}, nil
}

func (s *DB) Init() (*sql.DB, error) {
	if err := s.createUsersTable(); err != nil {
		return nil, err
	}

	if err := s.createBooksTable(); err != nil {
		return nil, err
	}

	if err := s.createHighlightsTable(); err != nil {
		return nil, err
	}

	return s.DB, nil
}

func (s *DB) createUsersTable() error {
	_, err := s.DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			email VARCHAR(255) NOT NULL,
			firstName VARCHAR(255) NOT NULL,
			lastName VARCHAR(255) NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (id),
			UNIQUE KEY (email)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)

	return err

}

func (s *DB) createBooksTable() error {
	_, err := s.DB.Exec(`
		CREATE TABLE IF NOT EXISTS books (
			isbn VARCHAR(255) NOT NULL,
			title VARCHAR(255) NOT NULL,
			authors VARCHAR(255) NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (isbn),
			UNIQUE KEY (isbn)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)

	return err
}

func (s *DB) createHighlightsTable() error {
	_, err := s.DB.Exec(`
		CREATE TABLE IF NOT EXISTS highlights (
			id INT NOT NULL AUTO_INCREMENT,
			text TEXT,
			location VARCHAR(255) NOT NULL,
			note TEXT,
			userId INT UNSIGNED NOT NULL,
			bookId VARCHAR(255) NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (id),
			FOREIGN KEY (userId) REFERENCES users(id),
			FOREIGN KEY (bookId) REFERENCES books(isbn)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)

	return err
}
