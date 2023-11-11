package postgres

import (
	"context"
	"time"

	"github.com/agpelkey/food"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ProtienDB represents protein database
type ProteinDB struct {
	db *pgxpool.Pool
}

// returns an instance of ProteinDB
func NewProteinStore(db *pgxpool.Pool) ProteinDB {
	return ProteinDB{db: db}
}


// Create new item
func (p ProteinDB) CreateNewProtein(item food.Protein) error {
    query := `INSERT INTO protein VALUES ($1, $2, $3, $4)`

    // create argument list to pass into db function
    args := []interface{}{item.Item, item.Unit, item.Quantity, item.Purchase_date}

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    return p.db.QueryRow(ctx, query, args...).Scan(&item.Item, &item.Unit, &item.Quantity, &item.Purchase_date)
}

// Get an item


// Update an item


// Delete an item
