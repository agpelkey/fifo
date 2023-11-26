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
    query := `
        BEGIN;
        INSERT INTO items (name, type, unit) VALUES ($1, $2, $3);
        INSERT INTO fridge (item_id, quantity) VALUES ((SELECT item_id FROM items WHERE name = $1), $4);
        COMMIT;`

    args := []interface{}{item.Name, item.Type, item.Unit, item.Quantity}

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    return f.db.QueryRow(ctx, query, args...).Scan(&item.Name, &item.Type, &item.Unit, &item.Quantity)
}

// GET
func (f fridgeDB) GetItemFromFridge(name string) (food.Items, error) {
	query := `SELECT fridge.quantity, items.name, items.type, items.unit FROM fridge JOIN items ON items.name = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	var item food.Items

	row := f.db.QueryRow(ctx, query, name)

    err := row.Scan(
        &item.Quantity,
		&item.Name,
		&item.Type,
		&item.Unit,
	)
    if err != nil {
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
