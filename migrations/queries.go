package migrations

const (
	createCarTable = "CREATE TABLE IF NOT EXISTS car (ID INT, NAME VARCHAR(255), COLOR VARCHAR(255));"
	dropCarTable   = "DROP TABLE IF EXISTS car"
)
