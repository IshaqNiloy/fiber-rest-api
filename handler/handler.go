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
		return err
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
			return err
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
		return err
	}
	return nil
}
