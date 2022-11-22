package database

func (db *DataStore) DbHealthCheck() (int, error) {
	var count int
	sqlStr := "select id from p_test where id = 1"
	err := db.Get(&count, sqlStr)
	return count, err
}