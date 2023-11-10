package postgres

import "github.com/jackc/pgx/v5/pgxpool"


type proteinstore struct {
    db *pgxpool.Pool
}

// factory to build new protein db
func NewProteinStore(db *pgxpool.Pool) proteinstore {
    return proteinstore{db: db}
}


// function to Insert an item


// function to Get an item


// function to Update an item


// function to delete an item
