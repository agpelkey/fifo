package postgres

/*
import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/agpelkey/food"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ProtienDB represents protein database
type dairyDB struct {
	db *pgxpool.Pool
}

// returns an instance of DairyDB
func NewDairyStore(db *pgxpool.Pool) dairyDB {
	return dairyDB{db: db}
}


func (d dairyDB) CreateNewDairy(item food.Dairy) error {
    query := `INSERT INTO dairy (item, unit, quantity, purchase_date) VALUES ($1, $2, $3, $4) RETURNING item, unit, quantity, purchase_date`

    // create argument list to pass into db function
    args := []interface{}{item.Item, item.Unit, item.Quantity, item.Purchase_date}

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    return d.db.QueryRow(ctx, query, args...).Scan(&item.Item, &item.Unit, &item.Quantity, &item.Purchase_date)
}

// Get an item
func (d dairyDB) GetDairyFromDB(item string) (*food.Dairy, error) {
    query := `SELECT item, unit, quantity, purchase_date FROM dairy WHERE item = $1`

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    var DairyItem food.Dairy

    row := d.db.QueryRow(ctx, query, item)

    err := row.Scan(
        &DairyItem.Item,
        &DairyItem.Unit,
        &DairyItem.Quantity,
        &DairyItem.Purchase_date,
    )

    if err != nil {
        switch {
        case errors.Is(err, sql.ErrNoRows):
            return nil, food.ErrDairyItemNotFound
        default:
            return nil, err
        }
    }

    return &DairyItem, nil
}


// Update an item
func (d dairyDB) UpdateDairyItem(item *food.Dairy) error {
    query := `
        UPDATE dairy 
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

    row, err := d.db.Query(ctx, query, queryArguments...)
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
func (d dairyDB) DeleteDairyItem(item string) error {
    query := `DELETE FROM dairy WHERE item = $1`

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    result, err := d.db.Exec(ctx, query, item)
    if err != nil {
        return fmt.Errorf("failed to delete from dairy items: %v", result)
    }

    if rows := result.RowsAffected(); rows != 1 {
        return food.ErrDairyItemNotFound
    }

    return nil
}
*/
