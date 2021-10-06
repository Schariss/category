package data

import (
	"fmt"
)

var ErrCategoryNotFound = fmt.Errorf("Product not found")

type Category struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Products []*Product `json:"products"`
}

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	SKU string `json:"sku"`
}

type Categories []*Category

func GetCategories() Categories{
	return categoryList
}

func GetCategoryByID(id int) (*Category, error) {
	i := findIndexByCategoryID(id)
	if id == -1 {
		return nil, ErrCategoryNotFound
	}
	return categoryList[i], nil
}

func AddCategory(q Category){
	i := categoryList[len(categoryList) -1].ID
	q.ID = i+1
	categoryList = append(categoryList, &q)
}

func DeleteCategory(id int) error {
	i := findIndexByCategoryID(id)
	if i == -1 {
		return ErrCategoryNotFound
	}
	categoryList = append(categoryList[:i], categoryList[i+1])
	return nil
}

func findIndexByCategoryID(id int) int {
	for i, c := range categoryList {
		if c.ID == id {
			return i
		}
	}
	return -1
}

var categoryList = []*Category{
	&Category{
		ID: 1,
		Name: "Coffee",
		Products: []*Product{
			&Product{
				ID: 1,
			},
			&Product{
				ID: 2,
			},
		},
	},
	&Category{
		ID: 2,
		Name: "Food",
		Products: []*Product{
			&Product{
				ID: 3,
			},
		},
	},
}
