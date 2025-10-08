CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    email TEXT
);
CREATE TABLE IF NOT EXISTS ofps (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    icao_from TEXT,
    icao_to TEXT,
    etd DATETIME,
    eta DATETIME,
    flight_number TEXT,
    dof DATE,
    reg_number TEXT
);