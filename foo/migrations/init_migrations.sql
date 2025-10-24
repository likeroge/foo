CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    email TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (email)
);
CREATE TABLE IF NOT EXISTS ofps (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    icao_from TEXT,
    icao_to TEXT,
    etd DATETIME,
    eta DATETIME,
    flight_number TEXT,
    dof DATE,
    reg_number TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    distance INTEGER,
    wind TEXT,
    fuel_flow INTEGER,
    trip_fuel INTEGER,
    flight_time TEXT,
    UNIQUE (icao_from, icao_to, flight_number, dof, reg_number)
);