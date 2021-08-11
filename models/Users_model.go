package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	UsersId  int    `json:"user_id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Uid      string `json:"uid"`
}

func UserData() ([]Users, error) {
	// https://tutorialedge.net/golang/golang-mysql-tutorial/
	sqlType := "mysql"
	sqlUsn := "root"
	sqlPass := ""
	sqlDatabase := "liff_db"

	con, err := sql.Open(sqlType, sqlUsn+":"+sqlPass+"@/"+sqlDatabase)
	// con, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/liff_db")
	if err != nil {
		return []Users{}, err
	}

	defer con.Close()

	result, err := con.Query("SELECT * FROM users")
	if err != nil {
		return []Users{}, err
	}
	fmt.Println(result)
	arrStruc := []Users{}
	// https://tutorialedge.net/golang/golang-mysql-tutorial/
	// Populating Structs from Results
	for result.Next() {
		var u Users
		err := result.Scan(
			&u.UsersId,
			&u.Fullname,
			&u.Email,
			&u.Mobile,
			&u.Uid,
		)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(u.Fullname)
		fmt.Println(u.Email)
		// https://stackoverflow.com/questions/38449863/how-to-add-a-struct-to-an-array-of-structs-in-go
		// user_value := Users{
		// 	UsersId:  u.UsersId,
		// 	Fullname: u.Fullname,
		// 	Email:    u.Email,
		// 	Mobile:   u.Mobile,
		// 	Uid:      u.Uid,
		// }
		// user_value := u
		arrStruc = append(arrStruc, u)
	}

	return arrStruc, nil
}
