package db

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type ItemsDB struct {
	*sqlx.DB
}

type IItemsDB interface {
	AddItem(int, string, string, string, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, float64, string, float64) error
	EditTable(string, interface{}) error
	DeleteItem(int) error
	UpdateItem(int, string, string, string, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, float64, string, float64) error
	CloseDB() error
}

func NewItemDB(db *sqlx.DB) *ItemsDB {
	return &ItemsDB{
		DB: db,
	}
}

func (db *ItemsDB) AddItem(itemID int, name, design, link string, quantity, placeDIN, placeU, portAIN, portAOT, portDIN, portDOT, portEHT, pci, connRJ, con20PIN, con50PIN, conDB37, conDB62, conSNC, wireMGTF, wireCAT6, len, washer, screw, bolt int, power float64, voltage string, price float64) error {
	start := time.Now()
	_, err := db.Exec(addFeatures, itemID, name, design, link, quantity, placeDIN, placeU, portAIN, portAOT, portDIN,
		portDOT, portEHT, pci, connRJ, con20PIN, con50PIN, conDB37, conDB62, conSNC, wireMGTF, wireCAT6, len, washer, screw, bolt, power, voltage, price)
	if err != nil {
		return err
	}
	log.Printf("Adding item access time: %v\n", time.Since(start))
	return nil
}

func (db *ItemsDB) EditTable(newObj string, newType interface{}) error {
	start := time.Now()
	switch dataType := newType.(type) {
	case string:
		if dataType == "string" {
			newType = "varchar"
		} else if dataType == "integer" {
			newType = "integer"
		} else if dataType == "float" {
			newType = "numeric"
		} else {
			return errors.New("Ошибка в вводе типа колонки ")
		}
	default:
		return errors.New("Неверный тип данных ")
	}

	//t := fmt.Sprintf("%s %s", newObj, newType)
	//fmt.Println(t)
	_, err := db.Exec(editTable, newObj, newType)
	if err != nil {
		return err
	}
	log.Printf("Editing table time: %v\n", time.Since(start))
	return nil
}

func (db *ItemsDB) DeleteItem(itemID int) error {
	_, err := db.Exec(deleteUserAccess, itemID)
	if err != nil {
		return err
	}

	return nil
}

func (db *ItemsDB) UpdateItem(itemID int, name, design, link string, quantity, placeDIN, placeU, portAIN, portAOT, portDIN, portDOT, portEHT, pci, connRJ, con20PIN, con50PIN, conDB37, conDB62, conSNC, wireMGTF, wireCAT6, len, washer, screw, bolt int, power float64, voltage string, price float64) error {
	start := time.Now()
	_, err := db.Exec(updateUserAccess, itemID, name, design, link, quantity, placeDIN, placeU, portAIN, portAOT, portDIN,
		portDOT, portEHT, pci, connRJ, con20PIN, con50PIN, conDB37, conDB62, conSNC, wireMGTF, wireCAT6, len, washer, screw, bolt, power, voltage, price)
	if err != nil {
		return err
	}
	log.Printf("Updating item time: %v\n", time.Since(start))
	return nil
}

func (db *ItemsDB) CloseDB() error {
	return db.Close()
}
