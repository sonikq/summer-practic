package db

const (
	addFeatures = `call items.add_items($1, $2::varchar, $3::varchar, $4::varchar, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28);`
	//checkUserAccess  = `select test_abac.check_user_access($1, $2, $3, $4, $5);`
	deleteUserAccess = `call items.delete_item($1);`
	updateUserAccess = `call items.update_items($1, $2::varchar, $3::varchar, $4::varchar, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28);`
)
