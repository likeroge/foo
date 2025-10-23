package ofprepos

import (
	"database/sql"
	"fmt"
	"log"

	"ego.dev21/greetings/internal/entities"
)

type OFPSQLiteRepo struct {
	db *sql.DB
}

func NewOFPSQLiteRepo(db *sql.DB) *OFPSQLiteRepo {
	return &OFPSQLiteRepo{
		db: db,
	}
}

func (o *OFPSQLiteRepo) GetOFPInfoById(id int) *entities.OFP {
	rows := o.db.QueryRow("SELECT id, icao_from, icao_to, etd, eta, flight_number, dof, reg_number FROM ofps WHERE id = ?", id)
	ofp := entities.OFP{}
	if err := rows.Scan(&ofp.Id, &ofp.IcaoFrom, &ofp.IcaoTo, &ofp.ETD, &ofp.ETA, &ofp.FlightNumber, &ofp.DOF, &ofp.RegNumber); err != nil {
		log.Println("Error due to fetch ofp", err)
		return nil
	}
	return &ofp
}

func (o *OFPSQLiteRepo) CreateOFPInfo(ofp *entities.OFP) (int64, error) {
	ofpInDB := o.db.QueryRow("SELECT id FROM ofps WHERE icao_from = ? AND icao_to = ? AND flight_number = ? AND dof = ? AND reg_number = ?", ofp.IcaoFrom, ofp.IcaoTo, ofp.FlightNumber, ofp.DOF, ofp.RegNumber)
	err := ofpInDB.Err()
	if err == nil {
		return -1, fmt.Errorf("OFP already exist")
	}

	result, err := o.db.Exec(`INSERT INTO ofps (icao_from, icao_to, etd, eta, flight_number, dof, reg_number) VALUES (?, ?, ?, ?, ?, ?, ?)`, ofp.IcaoFrom, ofp.IcaoTo, ofp.ETD, ofp.ETA, ofp.FlightNumber, ofp.DOF, ofp.RegNumber)
	if err != nil {
		log.Println("Error due to insert ofp", err)
		return -1, err
	}
	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		log.Println("Error due to insert ofp", err)
		return -1, err
	}
	return lastInsertedId, nil
}

func (o *OFPSQLiteRepo) DeleteOFPInfoById(id int) bool {
	_, err := o.db.Exec("DELETE FROM ofps WHERE id = ?", id)
	if err != nil {
		log.Println("Error due to delete ofp", err)
		return false
	}
	return true
}

func (o *OFPSQLiteRepo) GetAllOFPInfo() ([]*entities.OFP, error) {
	rows, err := o.db.Query("SELECT id, icao_from, icao_to, etd, eta, flight_number, dof, reg_number FROM ofps")
	if err != nil {
		log.Println("Error due to fetch ofp", err)
		return nil, err
	}
	defer rows.Close()
	ofps := []*entities.OFP{}
	for rows.Next() {
		ofp := entities.OFP{}
		if err := rows.Scan(&ofp.Id, &ofp.IcaoFrom, &ofp.IcaoTo, &ofp.ETD, &ofp.ETA, &ofp.FlightNumber, &ofp.DOF, &ofp.RegNumber); err != nil {
			log.Println("Error due to fetch ofp", err)
			return nil, err
		}
		ofps = append(ofps, &ofp)
	}
	return ofps, nil
}
