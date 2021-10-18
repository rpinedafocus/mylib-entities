package security

import (
	conn "github.com/rpinedafocus/mylib-dbconn"
)

type InfoLogin struct {
	full_name string
	user_name string
	role_id   int
}

func GetAccess(myLevel int, endpoint string) bool {

	db, errc := conn.GetConnection()

	if errc {
		return false
	} else {
		var level int
		stmt, err := db.Prepare(`SELECT level FROM university.access WHERE endpoint = $1`)
		err = stmt.QueryRow(endpoint).Scan(&level)

		if err != nil {
			return false
		} else {
			if myLevel == level {
				return true
			} else if level == 3 {
				return true
			} else {
				return false
			}
		}
	}
}

func Login(user string, password string) (InfoLogin, bool) {

	db, errc := conn.GetConnection()

	var infoLogin InfoLogin

	if errc {
		panic(errc)
		return infoLogin, true
	}

	stmt, err := db.Prepare(`SELECT CONCAT(trim(first_name),' ',trim(last_name)) as full_name, user_name as user, role_id level FROM university.users WHERE trim(user_name) = $1 AND trim(password) = $2`)
	if err != nil {
		panic(err)
		return infoLogin, true
	}
	err = stmt.QueryRow(user, password).Scan(&infoLogin.full_name, &infoLogin.user_name, &infoLogin.role_id)

	return infoLogin, false
}

/*
func main() {

	test, err := Login("root", "root")

	if err {
		fmt.Println("you don't have access")
	} else {
		fmt.Print(test)
	}
}
*/
