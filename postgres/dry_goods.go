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


type dryGoodDB struct {
	db *pgxpool.Pool
}


func NewDryGoodStore(db *pgxpool.Pool) dryGoodDB {
	return dryGoodDB{db: db}
}


func (d dryGoodDB) CreateNewDryGood(item food.DryGoods) error {
    query := `INSERT INTO dry_goods (item, unit, quantity, purchase_date) VALUES ($1, $2, $3, $4) RETURNING item, unit, quantity, purchase_date`

    // create argument list to pass into db function
    args := []interface{}{item.Item, item.Unit, item.Quantity, item.Purchase_date}

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    return d.db.QueryRow(ctx, query, args...).Scan(&item.Item, &item.Unit, &item.Quantity, &item.Purchase_date)
}

func (d dryGoodDB) GetDryGoodFromDB(item string) (*food.DryGoods, error) {
    query := `SELECT item, unit, quantity, purchase_date FROM dry_goods WHERE item = $1`

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    var DryGoodItem food.DryGoods

    row := d.db.QueryRow(ctx, query, item)

    err := row.Scan(
        &DryGoodItem.Item,
        &DryGoodItem.Unit,
        &DryGoodItem.Quantity,
        &DryGoodItem.Purchase_date,
    )

    if err != nil {
        switch {
        case errors.Is(err, sql.ErrNoRows):
            return nil, food.ErrDryGoodItemNotFound
        default:
            return nil, err
        }
    }

    return &DryGoodItem, nil
}


func (d dryGoodDB) UpdateDryGoodItem(item *food.DryGoods) error {
    query := `
        UPDATE dry_goods
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

func (d dryGoodDB) DeleteDryGoodItem(item string) error {
    query := `DELETE FROM dry_goods WHERE item = $1`

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    result, err := d.db.Exec(ctx, query, item)
    if err != nil {
        return fmt.Errorf("failed to delete from dry_goods items: %v", result)
    }

    if rows := result.RowsAffected(); rows != 1 {
        return food.ErrDryGoodItemNotFound
    }

    return nil
}
*/
