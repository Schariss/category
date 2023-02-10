package handlers

import (
	"encoding/json"
	"github.com/Schariss/category-api/data"
	"io/ioutil"
	"net/http"
	"strconv"
)

func(q *Categories) GetList(rw http.ResponseWriter, r *http.Request){
	q.l.Println("[DEBUG] get all records")
	lc := data.GetCategories()
	for _, c := range lc {
		var pl []*data.Product
		for pi := range c.Products {
			p := &data.Product{}
			productResp, err := http.Get("http://product-api:9090/products/" + strconv.Itoa(c.Products[pi].ID))
			if err != nil {
				http.Error(rw, "Cannot get response from product-api", http.StatusInternalServerError)
			}
			defer productResp.Body.Close()
			prod, err := ioutil.ReadAll(productResp.Body)
			if err != nil {
				q.l.Println("[ERROR] Cannot read response")
				panic(err)
			}
			err = json.Unmarshal(prod, &p)
			if err != nil {
				q.l.Println("[ERROR] serializing product")
				panic(err)
			}
			pl = append(pl, p)
		}
		c.Products = pl
	}
	err := data.ToJSON(lc, rw)
	if err != nil {
		q.l.Println("[ERROR] serializing categories", err)
	}
}

func(q *Categories) GetSingle(rw http.ResponseWriter, r *http.Request){
	id := getCategoryID(r)
	q.l.Println("[DEBUG] get record id", id)
	cat, err := data.GetCategoryByID(id)
	var pl []*data.Product
	for pi := range cat.Products {
		p := &data.Product{}
		productResp, err := http.Get("http://product-api:9090/products/" + strconv.Itoa(cat.Products[pi].ID))
		if err != nil {
			http.Error(rw, "Cannot get response from product-api", http.StatusInternalServerError)
		}
		defer productResp.Body.Close()
		prod, err := ioutil.ReadAll(productResp.Body)
		if err != nil {
			q.l.Println("[ERROR] Cannot read response")
			panic(err)
		}
		err = json.Unmarshal(prod, &p)
		if err != nil {
			q.l.Println("[ERROR] serializing product")
			panic(err)
		}
		pl = append(pl, p)
	}
	cat.Products = pl
	err = data.ToJSON(cat, rw)
	if err != nil {
		q.l.Println("[ERROR] serializing category", err)
	}
}
