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


type fruitDB struct {
	db *pgxpool.Pool
}

func NewFruitStore(db *pgxpool.Pool) fruitDB {
	return fruitDB{db: db}
}


func (d fruitDB) CreateNewFruit(item food.Fruit) error {
    query := `INSERT INTO fruit (item, unit, quantity, purchase_date) VALUES ($1, $2, $3, $4) RETURNING item, unit, quantity, purchase_date`

    // create argument list to pass into db function
    args := []interface{}{item.Item, item.Unit, item.Quantity, item.Purchase_date}

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    return d.db.QueryRow(ctx, query, args...).Scan(&item.Item, &item.Unit, &item.Quantity, &item.Purchase_date)
}

// Get an item
func (d fruitDB) GetFruitFromDB(item string) (*food.Fruit, error) {
    query := `SELECT item, unit, quantity, purchase_date FROM fruit WHERE item = $1`

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    var FruitItem food.Fruit

    row := d.db.QueryRow(ctx, query, item)

    err := row.Scan(
        &FruitItem.Item,
        &FruitItem.Unit,
        &FruitItem.Quantity,
        &FruitItem.Purchase_date,
    )

    if err != nil {
        switch {
        case errors.Is(err, sql.ErrNoRows):
            return nil, food.ErrFruitItemNotFound
        default:
            return nil, err
        }
    }

    return &FruitItem, nil
}


// Update an item
func (d fruitDB) UpdateFruitItem(item *food.Fruit) error {
    query := `
        UPDATE fruit
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
func (d fruitDB) DeleteFruitItem(item string) error {
    query := `DELETE FROM fruit WHERE item = $1`

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    result, err := d.db.Exec(ctx, query, item)
    if err != nil {
        return fmt.Errorf("failed to delete from fruit items: %v", result)
    }

    if rows := result.RowsAffected(); rows != 1 {
        return food.ErrFruitItemNotFound
    }

    return nil
}
*/
