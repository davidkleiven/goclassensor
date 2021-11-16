package server

import "testing"

func TestDetectTestContext(t *testing.T) {
	got := dbName()
	expect := "file::memory:?cache=shared"

	if got != expect {
		t.Errorf("Expected %s got %s\n", expect, got)
	}
}

func TestInitDB(t *testing.T) {
	db := connect()
	defer db.Close()

	for i := 0; i < 2; i++ {
		initDb(db)
		res := query(db, "SELECT count(*) FROM sqlite_master")
		var num int
		for res.Next() {
			res.Scan(&num)
		}

		if num != 1 {
			t.Errorf("Test #%d: Expected 1 got %d\n", i, num)
		}
	}

}

func TestInsert(t *testing.T) {
	db := connect()
	defer db.Close()
	initDb(db)

	data := TemperatureData{
		Timestamp:   "21:04",
		DataSource:  "unittest",
		OutdoorTemp: 4.0,
		IndoorTemp:  22.0,
	}

	insert(db, data)
	insert(db, data)

	// Count number of rows
	res := query(db, "SELECT COUNT(*) FROM temperatures")
	var num int
	for res.Next() {
		res.Scan(&num)
	}

	if num != 2 {
		t.Errorf("Expected 2 rows got %d\n", num)
	}
}
