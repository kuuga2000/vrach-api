package models

import (
	"carilokak/configs"
)

func CreateObject(object Object) (Object, error) {
	db := configs.GetDB()
	sqlStatement := `INSERT INTO objects (object_name,gender_id) VALUES ($1, $2) RETURNING object_id`
	err := db.QueryRow(sqlStatement, object.ObjectName, object.ObjectGender).Scan(&object.ObjectId)
	if err != nil {
		return object, err
	}
	return object, nil
}

//func GetObjects() (Objects, error) {
/*	db := configs.GetDB()
	sqlStatement := `SELECT * FROM objects`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return err.Error()
	}

	return rows.Scan()*/
/*result := Objects{}
for rows.Next() {
	list := Objects{}
	err2 := rows.Scan(&list.ObjectId, list.ObjectName, list.ObjectId)
	if err2 != nil {
		return "asd"
	}
	result.Objects = append(result.Objects, list)
}
return result, nil*/

//}

func UpdateObject(object Object) (Object, error) {
	db := configs.GetDB()
	sqlStatement := `UPDATE objects SET object_name=$1, gender_id=$2 WHERE object_id = $3 RETURNING object_id`
	db.Query(sqlStatement, object.ObjectName, object.ObjectGender, object.ObjectId)

	return object, nil
}

func DeleteObject(objectId string) string {
	db := configs.GetDB()
	sqlStatement := `DELETE FROM objects WHERE object_id = $1`
	db.Query(sqlStatement, objectId)

	return objectId
}
