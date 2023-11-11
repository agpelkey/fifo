package postgres

import (
	"context"
	"time"

	"github.com/agpelkey/food"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ProtienDB represents protein database
type proteinDB struct {
	db *pgxpool.Pool
}

// returns an instance of ProteinDB
func NewProteinStore(db *pgxpool.Pool) proteinDB {
	return proteinDB{db: db}
}


func (p proteinDB) CreateNewProtein(item food.Protein) error {
    query := `INSERT INTO protein (item, unit, quantity) VALUES ($1, $2, $3) RETURNING item, unit, quantity`

    // create argument list to pass into db function
    args := []interface{}{item.Item, item.Unit, item.Quantity}

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    return p.db.QueryRow(ctx, query, args...).Scan(&item.Item, &item.Unit, &item.Quantity)
}

// Get an item


// Update an item


// Delete an item
