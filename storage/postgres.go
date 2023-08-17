package storage

import (
	"database/sql"
	"fmt"

	"github.com/jbic9832/go-calorie-tracker/types"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=9832-Jb-9832 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createAccountsTable()
}

func (s *PostgresStore) createAccountsTable() error {
	query := `
    create table if not exists accounts (
        id serial not null primary key,
        username varchar(50),
        email varchar(100),
        password varchar(256),
        constrain unique_username unique (username)
    )
    `

	_, err := s.db.Exec(query)
	return err
}


func (s *PostgresStore) CreateAccount(acc *types.Account) error {
    query := `
    insert into accounts (
        username, email, password
    ) values (
        $1, $2, $3
    )
    `

    resp, err := s.db.Query(query, acc.Username, acc.Email, acc.Password)
    if err != nil {
        return err
    }
    fmt.Printf("%+v\n", resp)
    return nil

}

func (s *PostgresStore) DeleteAccount(id int) error {
    query := `delete * from accounts where id = $1`
    _, err := s.db.Query(query, id)
    return err
}

func (s *PostgresStore) UpdateAccount(acc *types.Account) error {
    return nil
}

func (s *PostgresStore) GetAccount(id int) (*types.Account, error) {
    query := `select * from accounts where id = $1`
    rows, err := s.db.Query(query, id)
    if err != nil {
        return nil, err
    }
    for rows.Next() {
        return scanIntoAccount(rows)
    }

    return nil, fmt.Errorf("account %d does not exist.", id)
}

func scanIntoAccount(rows *sql.Rows) (*types.Account, error) {
    account := new(types.Account) 
    err := rows.Scan(&account.ID, &account.Username, &account.Email, &account.Password)
    return account, err 
}































