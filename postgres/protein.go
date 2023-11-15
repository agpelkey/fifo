package postgres

import (
	"context"
	"database/sql"
	"errors"
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
    query := `INSERT INTO protein (item, unit, quantity, purchase_date) VALUES ($1, $2, $3, $4) RETURNING item, unit, quantity, purchase_date`

    // create argument list to pass into db function
    args := []interface{}{item.Item, item.Unit, item.Quantity, item.Purchase_date}

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    return p.db.QueryRow(ctx, query, args...).Scan(&item.Item, &item.Unit, &item.Quantity, &item.Purchase_date)
}

// Get an item
func (p proteinDB) GetProteinFromDB(item string) (*food.Protein, error) {
    query := `SELECT item, unit, quantity, purchase_date FROM protein WHERE item = $1`

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    var ProtienItem food.Protein

    row := p.db.QueryRow(ctx, query, item)

    err := row.Scan(
        &ProtienItem.Item,
        &ProtienItem.Unit,
        &ProtienItem.Quantity,
        &ProtienItem.Purchase_date,
    )

    if err != nil {
        switch {
        case errors.Is(err, sql.ErrNoRows):
            return nil, food.ErrProteinItemNotFound
        default:
            return nil, err
        }
    }

    return &ProtienItem, nil
}


// Update an item
func (p proteinDB) UpdateProteinItem(item *food.Protein) error {
    query := `
        UPDATE protein 
        SET item = $1, unit = $2, quantity = $3
        WHERE item = $1
    `

    queryArguments := []interface{}{
        item.Item,
        item.Unit,
        item.Quantity,
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    row, err := p.db.Query(ctx, query, queryArguments...)
    if err != nil {
        switch {
        case errors.Is(err, sql.ErrNoRows):
            return food.ErrEditConflict
        default:
            return err
        }
    }

    return row.Err()
}

// Delete an item










