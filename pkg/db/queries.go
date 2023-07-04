package db

const (
	addFeatures = `call features.add_features($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27);`
	//checkUserAccess  = `select test_abac.check_user_access($1, $2, $3, $4, $5);`
	deleteUserAccess = `call features.delete_features($1);`
	updateUserAccess = `call features.update_features($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27);`
)
