package category_nurzhas_store

import (
	"database/sql"
	"log"
	"strconv"
	"strings"
)

var categoryQueries = []string{
	`create table if not exists category(
		id text,
		name text,
		small_description text,
		big_description text,
		image_url text,
		primary key(id)
	);`,
}

type categoryStore struct {
	db *sql.DB
}

func NewPostgresCategoryStore(cfg PostgresConfig) (CategoryStore, error) {
	db, err := getDbConn(getConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range categoryQueries {
		_, err = db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	db.SetMaxOpenConns(10)
	store := &categoryStore{db: db}
	return store, nil
}

func (c *categoryStore) CreateCategory(category *Category) (*Category, error) {
	result, err := c.db.Exec(
		"INSERT INTO categories (id, name, small_description, big_description, image_url) "+
			"VALUES ($1, $2, $3, $4, $5)",
		category.Id, category.Name, category.SmallDescription, category.BigDescription, category.ImageUrl,
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
	return category, nil
}

func (c *categoryStore) UpdateCategory(category *CategoryUpdate) (*Category, error) {
	q := "update categories set "
	parts := []string{}
	values := []interface{}{}
	cnt := 0
	if category.Name != nil {
		cnt++
		parts = append(parts, "name = $"+strconv.Itoa(cnt))
		values = append(values, category.Name)
	}
	if category.SmallDescription != nil {
		cnt++
		parts = append(parts, "small_description = $"+strconv.Itoa(cnt))
		values = append(values, category.SmallDescription)
	}
	if category.BigDescription != nil {
		cnt++
		parts = append(parts, "big_description = $"+strconv.Itoa(cnt))
		values = append(values, category.BigDescription)
	}
	if category.ImageUrl != nil {
		cnt++
		parts = append(parts, "image_url = $"+strconv.Itoa(cnt))
		values = append(values, category.ImageUrl)
	}
	if len(parts) <= 0 {
		return nil, ErrNothingToUpdate
	}
	cnt++
	q = q + strings.Join(parts, " , ") + " WHERE id = $" + strconv.Itoa(cnt)
	values = append(values, category.Id)
	result, err := c.db.Exec(q, values...)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrCategoryNotFound
	}
	return c.GetCategory(category.Id)
}

func (c *categoryStore) ListCategory() ([]Category, error) {
	categories := []Category{}
	var values []interface{}
	q := "select id, name, small_description, big_description, image_url from categories"
	//cnt := 1
	rows, err := c.db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		category := Category{}
		err = rows.Scan(&category.Id, &category.Name, &category.SmallDescription, &category.BigDescription, &category.ImageUrl)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (c *categoryStore) GetCategory(id string) (*Category, error) {
	category := &Category{}
	err := c.db.QueryRow("select id, name, small_description, big_description, image_url from categories where id = $1 limit 1", id).
		Scan(&category.Id, &category.Name, &category.SmallDescription, &category.BigDescription, &category.ImageUrl)
	if err == sql.ErrNoRows {
		return nil, ErrCategoryNotFound
	} else if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *categoryStore) DeleteCategory(id string) error {
	result, err := c.db.Exec("delete from categories where id= $1", id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n <= 0 {
		return ErrCategoryNotFound
	}
	return nil
}
