package ofprepos

import (
	"database/sql"
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
		log.Println("Error due to fetch ofp: ", err)
		return nil
	}
	return &ofp
}

func (o *OFPSQLiteRepo) CreateOFPInfo(ofp *entities.OFP) (int64, error) {
	sqlStr := `
					INSERT INTO ofps (icao_from, icao_to, etd, eta, flight_number, dof, reg_number, distance, wind, fuel_flow, trip_fuel, flight_time, rte, created_at, updated_at)
					VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?, ?, datetime('now'), datetime('now'))
					ON CONFLICT(icao_from, icao_to, flight_number, dof, reg_number) 
					DO UPDATE SET
						distance = excluded.distance,
						wind = excluded.wind,
						fuel_flow = excluded.fuel_flow,
						trip_fuel = excluded.trip_fuel,
						flight_time = excluded.flight_time,
						etd = excluded.etd,
						eta  = excluded.eta,
						rte = excluded.rte,
						updated_at = datetime('now');
				`

	result, err := o.db.Exec(sqlStr, ofp.IcaoFrom, ofp.IcaoTo, ofp.ETD, ofp.ETA, ofp.FlightNumber, ofp.DOF, ofp.RegNumber, ofp.Distance, ofp.Wind, ofp.FuelFlow, ofp.TripFuel, ofp.FlightTime, ofp.Rte)
	if err != nil {
		log.Println("Error due to insert ofp", err)
		return -1, err
	}
	lastInsertedId, _ := result.LastInsertId()
	return lastInsertedId, nil
}

func (o *OFPSQLiteRepo) DeleteOFPInfoById(id int) bool {
	_, err := o.db.Exec("DELETE FROM ofps WHERE id = ?", id)
	if err != nil {
		log.Println("Error due to delete ofp: ", err)
		return false
	}
	return true
}

func (o *OFPSQLiteRepo) GetAllOFPInfo() ([]*entities.OFP, error) {
	rows, err := o.db.Query("SELECT id, icao_from, icao_to, etd, eta, flight_number, dof, reg_number FROM ofps")
	if err != nil {
		log.Println("Error due to fetch ALL ofp: ", err)
		return nil, err
	}
	defer rows.Close()
	ofps := []*entities.OFP{}
	for rows.Next() {
		ofp := entities.OFP{}
		if err := rows.Scan(&ofp.Id, &ofp.IcaoFrom, &ofp.IcaoTo, &ofp.ETD, &ofp.ETA, &ofp.FlightNumber, &ofp.DOF, &ofp.RegNumber); err != nil {
			log.Println("Error due to Scan row ofp  -> GetAllOFPInfo: ", err)
			return nil, err
		}
		ofps = append(ofps, &ofp)
	}
	return ofps, nil
}
