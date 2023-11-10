package postgres

import "github.com/jackc/pgx/v5/pgxpool"

// ProtienDB represents protein database
type ProteinDB struct {
	db *pgxpool.Pool
}


// returns an instance of ProteinDB
func NewProteinStore(db *pgxpool.Pool) ProteinDB {
	return ProteinDB{db: db}
}


// Create new item


// Get an item


// Update an item


// Delete an item
