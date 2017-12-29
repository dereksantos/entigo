package entigo

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"testing"
	"time"
)

var db *sql.DB

func init() {
	conn, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	db = conn

	createSchema()
}

func createSchema() {
	q := `CREATE TABLE customers(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(255),
			email VARCHAR(255),
			created DATETIME,
			updated DATETIME
		)`

	err := Exec(db, q)
	if err != nil {
		panic(err)
	}

	q = `CREATE TABLE cars(
			vin VARCHAR(17) NOT NULL PRIMARY KEY,
			color VARCHAR(20),
			make VARCHAR(50),
			model VARCHAR(50)
		)`

	err = Exec(db, q)
	if err != nil {
		panic(err)
	}

}

// A typical model struct.
type Customer struct {
	ID      int64
	Name    string
	Email   string
	Created time.Time
	Updated time.Time
}

func (c *Customer) Entity() *Entity {
	return &Entity{
		Name: "customers",
		Key:  &Field{Name: "id", Value: &c.ID},
		Fields: []*Field{
			&Field{Name: "name", Value: &c.Name},
			&Field{Name: "email", Value: &c.Email},
			&Field{Name: "created", Value: &c.Created},
			&Field{Name: "updated", Value: &c.Updated},
		},
	}
}

// A struct with an alphanumeric primary key
type Car struct {
	VIN   string
	Color string
	Make  string
	Model string
}

func (c *Car) Entity() *Entity {
	return &Entity{
		Name: "cars",
		Key:  &Field{Name: "vin", Value: &c.VIN, NonIncrementing: true},
		Fields: []*Field{
			&Field{Name: "color", Value: &c.Color},
			&Field{Name: "make", Value: &c.Make},
			&Field{Name: "model", Value: &c.Model},
		},
	}
}

// Test inserting, updating, getting and deleting a customer
// with a typical integer auto incrementing key.
func TestCustomerEntity(t *testing.T) {
	customer := &Customer{
		Name:    "John Doe",
		Email:   "john+test@tester.com",
		Created: time.Now(),
		Updated: time.Now(),
	}

	e := customer.Entity()
	id, err := e.Insert(db)
	if err != nil {
		t.Error(err)
	}
	customer.ID = id

	if customer.ID < 1 {
		t.Error("Customer.ID was not set after insert")
	}

	customer.Name = "Jane Fox"
	err = e.Update(db)
	if err != nil {
		t.Error(err)
	}

	getCustomer := &Customer{ID: id}
	err = getCustomer.Entity().Get(db)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", customer)
	t.Logf("%+v", getCustomer)

	if getCustomer.ID != customer.ID ||
		getCustomer.Name != customer.Name ||
		getCustomer.Email != customer.Email ||
		!getCustomer.Created.Equal(customer.Created) ||
		!getCustomer.Updated.Equal(customer.Updated) {
		t.Error("Customer did not scan values correctly.")
	}
}

func TestCarEntity(t *testing.T) {
	vin := "JF2SHBDC5BH745690"
	car := &Car{
		VIN:   vin,
		Color: "red",
		Make:  "honda",
		Model: "crv",
	}

	_, err := car.Entity().Insert(db)
	if err != nil {
		t.Error(err)
	}

	car.Color = "blue"
	err = car.Entity().Update(db)
	if err != nil {
		t.Error(err)
	}

	getCar := &Car{VIN: vin}
	err = getCar.Entity().Get(db)
	if err != nil {
		t.Error(err)
	}

	if car.VIN != getCar.VIN ||
		car.Color != getCar.Color ||
		car.Make != getCar.Make ||
		car.Model != getCar.Model {
		t.Error("Car did not scan values correctly.")
	}

}
