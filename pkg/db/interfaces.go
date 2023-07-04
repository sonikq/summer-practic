package db

import (
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type ItemsDB struct {
	*sqlx.DB
}

type IItemsDB interface {
	AddItem(string, string, string, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, float64, string, float64) error
	//CheckUserAccess(int, int, int, int, string) (bool, error)
	DeleteItem(string) error
	UpdateItem(string, string, string, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, float64, string, float64) error
	CloseDB() error
}

func NewItemDB(db *sqlx.DB) *ItemsDB {
	return &ItemsDB{
		DB: db,
	}
}

func (db *ItemsDB) AddItem(name, design, link string, quantity, placeDIN, placeU, portAIN, portAOT, portDIN, portDOT, portEHT, pci, connRJ, con20PIN, con50PIN, conDB37, conDB62, conSNC, wireMGTF, wireCAT6, len, washer, screw, bolt int, power float64, voltage string, price float64) error {
	start := time.Now()
	_, err := db.Exec(addFeatures, name, design, link, quantity, placeDIN, placeU, portAIN, portAOT, portDIN,
		portDOT, portEHT, pci, connRJ, con20PIN, con50PIN, conDB37, conDB62, conSNC, wireMGTF, wireCAT6, len, washer, screw, bolt, power, voltage, price)
	if err != nil {
		return err
	}
	log.Printf("Adding item access time: %v\n", time.Since(start))
	return nil
}

//func (db *ItemsDB) CheckUserAccess(userID int, okud int, reportID int, chapterID int, operation string) (bool, error) {
//	var exists bool
//	row, err := db.Query(checkUserAccess, userID, okud, reportID, chapterID, operation)
//	if err != nil {
//		return false, err
//	}
//	defer row.Close()
//
//	for row.Next() {
//		if err = row.Scan(&exists); err != nil {
//			return false, err
//		}
//	}
//
//	return exists, nil
//}

func (db *ItemsDB) DeleteItem(design string) error {
	_, err := db.Exec(deleteUserAccess, design)
	if err != nil {
		return err
	}

	return nil
}

func (db *ItemsDB) UpdateItem(name, design, link string, quantity, placeDIN, placeU, portAIN, portAOT, portDIN, portDOT, portEHT, pci, connRJ, con20PIN, con50PIN, conDB37, conDB62, conSNC, wireMGTF, wireCAT6, len, washer, screw, bolt int, power float64, voltage string, price float64) error {
	start := time.Now()
	_, err := db.Exec(updateUserAccess, name, design, link, quantity, placeDIN, placeU, portAIN, portAOT, portDIN,
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
