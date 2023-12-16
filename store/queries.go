package store

const (
	createQuery  = `INSERT INTO car (id,name,color)values(?,?,?)`
	updateQuery  = `UPDATE car SET name=?, color=? WHERE id=?`
	getByIDQuery = `SELECT id,name,color FROM car WHERE id=?`
	deleteQuery  = `Delete from car where id =?`
	getAllQuery  = `SELECT id,name,color FROM car`
)
