package handler

import (
	"database/sql"
	"github.com/IshaqNiloy/go-rest-api/applibs"
	"github.com/IshaqNiloy/go-rest-api/database"
	"github.com/IshaqNiloy/go-rest-api/model"
	"github.com/gofiber/fiber/v2"
	"log"
)

func GetAllProducts(c *fiber.Ctx) error {
	// @URL:{{BASE_URL}}/api/

	rows, err := database.DB.Query("SELECT * FROM products ORDER BY name")

	// error while getting rows
	if err != nil {
		err = applibs.Response(c, applibs.RequestFailed, fiber.StatusBadRequest, nil)

		if err != nil {
			log.Printf("getting rows failed: %v", err)
			return err
		}
		return err
	}
	// closes the rows
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Fatalf("error closing rows: %v", err)
		}
	}(rows)

	result := model.Products{}

	// gets rows one by one
	for rows.Next() {
		product := model.Product{}
		err = rows.Scan(&product.Id, &product.Amount, &product.Name, &product.Description, &product.Category)

		// error while scanning rows
		if err != nil {
			log.Printf("scanning rows failed: %v", err)
			// failed response
			err = applibs.Response(c, applibs.RequestFailed, fiber.StatusBadRequest, nil)

			// error while sending failed response
			if err != nil {
				log.Printf("sending failed response failed: %v", err)
				return err
			}
			return err
		}

		result.Products = append(result.Products, product)
	}

	// success response
	err = applibs.Response(c, applibs.RequestSuccess, fiber.StatusOK, &result)

	// error while sending success response
	if err != nil {
		log.Printf("sending success response failed: %v", err)
		// failed response
		err = applibs.Response(c, applibs.RequestFailed, fiber.StatusBadRequest, &result)
		if err != nil {
			log.Printf("sending failed response failed: %v", err)
			return err
		}
	}
	return nil
}

func GetSingleProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	// gets specific row
	row, err := database.DB.Query("SELECT * FROM products where id=$1", id)

	// error while getting specific row
	if err != nil {
		log.Printf("getting row failed: %v", err)
		err = applibs.Response(c, applibs.RequestFailed, fiber.StatusBadRequest, nil)
		if err != nil {
			log.Printf("sending failed response failed: %v", err)
			return err
		}
		return nil
	}

	product := model.Product{}
	result := model.Products{}

	for row.Next() {
		// scans values of the specific row
		err = row.Scan(&product.Id, &product.Amount, &product.Name, &product.Description, &product.Category)

		// error while scanning row
		if err != nil {
			log.Printf("scanning row failed: %v", err)
			err = applibs.Response(c, applibs.RequestFailed, fiber.StatusBadRequest, nil)
			if err != nil {
				log.Printf("sending failed response failed: %v", err)
				return err
			}
			return nil
		}
		result.Products = append(result.Products, product)
	}

	// no data found
	if len(result.Products) < 1 {
		log.Printf("no data found")
		err = applibs.Response(c, applibs.RequestFailed, fiber.StatusBadRequest, nil)
		if err != nil {
			log.Printf("sending failed response failed: %v", err)
			return err
		}
		return nil
	}
	// success response
	err = applibs.Response(c, applibs.RequestSuccess, fiber.StatusOK, &result)

	// error while sending success response
	if err != nil {
		log.Printf("sending success response failed: %v", err)
		err = applibs.Response(c, applibs.RequestFailed, fiber.StatusBadRequest, nil)
		// error while sending failed response
		if err != nil {
			log.Printf("sending failed response failed: %v", err)
			return err
		}
		return nil
	}
	return nil
}

func CreateProduct(c *fiber.Ctx) error {
	// initialize a product
	product := new(model.Product)
	products := model.Products{}

	// parse request body into product struct
	err := c.BodyParser(product)

	// error while parsing request body
	if err != nil {
		log.Printf("parsing request body failed: %v", err)
		// failed response
		err = applibs.Response(c, applibs.RequestFailed, fiber.StatusBadRequest, nil)
		if err != nil {
			log.Printf("sending failed response failed: %v", err)
			return err
		}
		return nil
	}

	// inserting data into db
	_, err = database.DB.Query("INSERT INTO products (amount, name, description, category) "+
		"VALUES ($1, $2, $3, $4)", product.Amount, product.Name, product.Description, product.Category)

	// error while inserting data into db
	if err != nil {
		log.Printf("inserting data into db failed: %v", err)
		// failed response
		err = applibs.Response(c, applibs.RequestFailed, fiber.StatusBadRequest, nil)
		if err != nil {
			log.Printf("sending error response failed: %v", err)
			return err
		}
		return nil
	}

	// log created product
	log.Print("product created successfully")

	products.Products = append(products.Products, *product)

	// success response
	err = applibs.Response(c, applibs.RequestSuccess, fiber.StatusOK, &products)

	// error while sending success response
	if err != nil {
		log.Printf("sending success response failed: %v", err)
		// failed response
		err = applibs.Response(c, applibs.RequestFailed, fiber.StatusBadRequest, nil)
		if err != nil {
			log.Printf("sending failed response failed: %v", err)
			return err
		}
		return nil
	}

	return nil
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	// gets specific row
	row, err := database.DB.Query("SELECT * FROM products where id=$1", id)

	// error while getting specific row
	if err != nil {
		log.Printf("getting row failed: %v", err)
		err = applibs.Response(c, applibs.RequestFailed, fiber.StatusBadRequest, nil)
		if err != nil {
			log.Printf("sending failed response failed: %v", err)
			return err
		}
		return nil
	}

	product := model.Product{}
	result := model.Products{}

	for row.Next() {
		// scans values of the specific row
		err = row.Scan(&product.Id, &product.Amount, &product.Name, &product.Description, &product.Category)

		// error while scanning row
		if err != nil {
			log.Printf("scanning row failed: %v", err)
			err = applibs.Response(c, applibs.RequestFailed, fiber.StatusBadRequest, nil)
			if err != nil {
				log.Printf("sending failed response failed: %v", err)
				return err
			}
			return nil
		}
		result.Products = append(result.Products, product)
	}

	// no data found
	if len(result.Products) < 1 {
		log.Printf("no data found")
		err = applibs.Response(c, applibs.RequestFailed, fiber.StatusBadRequest, nil)
		if err != nil {
			log.Printf("sending failed response failed: %v", err)
			return err
		}
		return nil
	}

	// deletes specific row from db
	row, err = database.DB.Query("DELETE FROM products WHERE id = $1", id)

	// error while sending success response
	if err != nil {
		log.Printf("deleting row failed: %v", err)
		// failed response
		err = applibs.Response(c, applibs.RequestFailed, fiber.StatusBadRequest, nil)
		if err != nil {
			log.Printf("sending failed response failed: %v", err)
			return err
		}
		return nil
	}

	log.Print("product deleted successfully")

	// success response
	err = applibs.Response(c, applibs.RequestSuccess, fiber.StatusOK, &result)

	// error while sending success response
	if err != nil {
		log.Printf("sending error response failed: %v", err)
		// failed response
		err = applibs.Response(c, applibs.RequestFailed, fiber.StatusBadRequest, nil)
		// error while sending failed response
		if err != nil {
			log.Printf("sending error response failed: %v", err)
			return err
		}
		return nil
	}
	return nil
}
