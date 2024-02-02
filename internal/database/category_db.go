package database

import (
	"database/sql"

	"github.com/leonardohklaus/goapi/internal/entity"
)

type CategoryDB struct {
	db *sql.DB
}

func NewCategoryDB(db *sql.DB) *CategoryDB {
	return &CategoryDB{db: db}
}

func (cd *CategoryDB) GetCategories() ([]*entity.Category, error) {
	rows, err := cd.db.Query("select id, name from categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (cd *CategoryDB) GetCategory(id string) (*entity.Category, error) {
	var category entity.Category
	err := cd.db.QueryRow("select id, name from categories where id = $1", id).Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (cd *CategoryDB) CreateCategory(category *entity.Category) (string, error) {
	_, err := cd.db.Exec("insert into categories (id, name) values ($1, $2)", category.ID, category.Name)
	if err != nil {
		return "teste", err
	}
	return category.ID, nil
}
