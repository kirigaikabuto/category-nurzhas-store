package category_nurzhas_store

import (
	"database/sql"
	"log"
)

var ordersQueries = []string{
	`create table if not exists nurzhas_orders(
		id text,
		building_type text,
		width text,
		height text,
		length text,
		panel_type text,
		insulation_type text,
		panel_depth text,
		layout_type text,
		panel_width text,
		color text,
		primary key(id)
	);`,
}

type ordersStore struct {
	db *sql.DB
}

func NewPostgresOrdersStore(cfg PostgresConfig) (OrderStore, error) {
	db, err := getDbConn(getConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range ordersQueries {
		_, err = db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	db.SetMaxOpenConns(10)
	store := &ordersStore{db: db}
	return store, nil
}

func (c *ordersStore) CreateOrder(order *Order) (*Order, error) {
	result, err := c.db.Exec(
		"INSERT INTO nurzhas_orders "+
			"(id, building_type, width, height, "+
			"length, panel_type, insulation_type, panel_depth, "+
			"layout_type, panel_width, color) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		order.Id, order.BuildingType, order.Width, order.Height,
		order.Length, order.PanelType, order.InsulationType,
		order.PanelDepth, order.LayoutType, order.PanelWidth, order.Color,
	)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrCreateCategoryUnknown
	}
	return order, nil
}

func (c *ordersStore) ListOrder() ([]Order, error) {
	orders := []Order{}
	var values []interface{}
	q := "select " +
		"id, building_type, width, height, " +
		"length, panel_type, insulation_type, panel_depth, " +
		"layout_type, panel_width, color " +
		"from nurzhas_orders"
	//cnt := 1
	rows, err := c.db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		order := Order{}
		err = rows.Scan(
			&order.Id, &order.BuildingType, &order.Width, &order.Height,
			&order.Length, &order.PanelType, &order.InsulationType, &order.PanelDepth,
			&order.LayoutType, &order.PanelWidth, &order.Color)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
