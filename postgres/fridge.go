package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/agpelkey/food"
	"github.com/jackc/pgx/v5/pgxpool"
)


type fridgeDB struct {
	db *pgxpool.Pool
}

func NewFridgeDB(db *pgxpool.Pool) fridgeDB {
	return fridgeDB{db: db}
}

// Insert
func (f *fridgeDB) InsertIntoFridge(item food.Items) error {
	return nil
}

// GET
func (f fridgeDB) GetItemFromFridge(name string) (food.Items, error) {
	query := `SELECT name, type, quantity, unit FROM items WHERE name = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	var item food.Items

	if err := f.db.QueryRow(ctx, query, name).Scan(
		item.Name,
		item.Type,
		item.Quantity,
		item.Unit,
	); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return food.Items{}, food.ErrItemNotFound
		default:
			return food.Items{}, err
		}		
	}

	return item, nil
}


// UPDATE


// DELETE
