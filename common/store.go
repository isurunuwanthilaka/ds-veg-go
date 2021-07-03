package common

import (
	"fmt"
)

var (
	store *Store 
)


// Store struct represents a store of vegetables.
type Store struct {
	database map[int]Vegetable // private
}

// Add methods adds a vegetable to the store (procedure).
func (c *Store) Add(vegetable Vegetable) error {

	// check if vegetable already exists in the database
	if _, ok := c.database[vegetable.ID]; ok {
		return fmt.Errorf("vegetable with id '%d' already exists", vegetable.ID)
	}

	// add vegetable `p` in the database
	c.database[vegetable.ID] = vegetable

	// return `nil` error
	return nil
}

// Get methods returns a vegetable with specific id (procedure).
func (c *Store) Get(id int,reply *Vegetable )error{

	// get vegetable with id `p` from the database
	result, ok := c.database[id]

	// check if vegetable exists in the database
	if !ok {
		return fmt.Errorf("vegetable with id '%d' does not exist", id)
	}

	*reply = result

	// return `nil` error
	return nil
}

// GetAll methods returns all vegetable (procedure).
func (c *Store) GetAll(args string ,reply *[]Vegetable ) error{

	var arr []Vegetable

	for _,element := range c.database {
		arr = append(arr, element)
	}

	*reply = arr
	
	return nil
}

// AddVeg methods add a vegetable (procedure).
func (c *Store) AddVeg(vegetable Vegetable,reply *Vegetable )error{

	vegetable.Price = 0.0
	vegetable.Amount = 0.0

	store.Add(vegetable)

	*reply = vegetable

	// return `nil` error
	return nil
}

// Update methods update a vegetable with specific id (procedure).
func (c *Store) Update(vegetable Vegetable,reply *Vegetable )error{

	// get vegetable with id `p` from the database
	_, ok := c.database[vegetable.ID]

	// check if vegetable exists in the database
	if !ok {
		return fmt.Errorf("vegetable with id '%d' does not exist", vegetable.ID)
	}

	c.database[vegetable.ID] =  vegetable

	*reply = vegetable

	// return `nil` error
	return nil
}


func init() {
	store = &Store{
		database: make(map[int]Vegetable),
	}
	store.Add(Vegetable{
		ID:        1,
		Name : "carrot",
		Price: 10.00,
		Amount:  5.25,
	})
	store.Add(Vegetable{
		ID:        2,
		Name : "beetroot",
		Price: 20.00,
		Amount:  10.25,
	})
	store.Add(Vegetable{
		ID:        3,
		Name : "cabbage",
		Price: 15.00,
		Amount:  10.55,
	})
}

// GetStore function returns a new instance of Store (pointer).
func GetStore() *Store {
	return store
}