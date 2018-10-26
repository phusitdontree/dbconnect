package main
import (
_ "github.com/go-sql-driver/mysql"
"database/sql"
"fmt"
)
func main() {
   db, err := sql.Open("mysql", "root:phusitdontree@localhost/GOWAI")
   checkErr(err)
        // query
        rows, err := db.Query("SELECT * FROM users")
        checkErr(err)

        for rows.Next() {
            var uid int
            var uname string
            var fullname string
            err = rows.Scan(&uid, &uname, &fullname)
            checkErr(err)
            fmt.Println(uid)
            fmt.Println(uname)
            fmt.Println(fullname)
        }

   db.Close()
}

func checkErr(err error) {
   if err != nil {
   panic(err)
}
    }
