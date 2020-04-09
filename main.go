package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"github.com/go-chi/chi"
	"restapi/database"
)

type Product struct {
	ID	int	`json:"id"`
	Product_Code	string `json:"product_code"`
	Description	string `json:"description"`
}

var databaseConnection *sql.DB

func catch(e error)  {
	if e != nil {
		panic(e)
	}
}

func main() {
	/*r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)*/
	databaseConnection = database.InitDB()
	defer databaseConnection.Close()

	// Logic
	r := chi.NewRouter()
	r.Get("/products", AllProducts)
	r.Post("/products", CreateProduct)
	r.Put("/products/{id}", UpdateProduct)
	r.Delete("/products/{id}", DeleteProduct)
	log.Fatalln(http.ListenAndServe(":3000", r))
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	query, err := databaseConnection.Prepare("DELETE FROM products WHERE id=?")
	catch(err)
	_, er := query.Exec(id)
	catch(er)
	query.Close()
	respondwithJSON(w, http.StatusOK, map[string]string{"message":"successfully deleted"})
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	id := chi.URLParam(r, "id")
	er := json.NewDecoder(r.Body).Decode(&product)
	catch(er)

	query, err := databaseConnection.Prepare("UPDATE products SET product_code=?, description=? WHERE id=?")
	catch(err)
	_, errr := query.Exec(product.Product_Code, product.Description, id)
	catch(errr)
	defer query.Close()
	respondwithJSON(w, http.StatusOK, map[string]string{"message":"update successfully"})
}

func CreateProduct(w http.ResponseWriter, r *http.Request)  {
	var product Product
	er := json.NewDecoder(r.Body).Decode(&product)
	catch(er)
	log.Printf("%+v",  product)
	query, err := databaseConnection.Prepare("INSERT products SET product_code=?, description=?")
	catch(err)
	_, err = query.Exec(product.Product_Code, product.Description)
	catch(err)
	defer query.Close()
	respondwithJSON(w, http.StatusCreated, map[string]string{"message":"successfully created"})
}

func AllProducts(w http.ResponseWriter, r *http.Request) {
	const sql = `SELECT id, product_code, COALESCE(description,'') FROM products`
	results, err := databaseConnection.Query(sql)
	catch(err)
	var products []*Product
	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.ID, &product.Product_Code, &product.Description)
		catch(err)
		products = append(products, product)
	}
	respondwithJSON(w, http.StatusOK, map[string]interface{}{"data":products})
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		panic("Error")
	}
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(code)
	w.Write(response)
}
