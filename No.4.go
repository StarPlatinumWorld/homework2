package main

import(

"database/sql"
"fmt"
"github. com/gin-gonic/gin"
_"github. com/go- sql-driver/ mysql"
)
func main() {
	db, err := sql.Open("mysql", "root: 687211gtcp(127.0 0.1: 3306)/final_ exam?charset=utf8")
	if err!=nil{fmt.Println("open mysql is wrong:",err)}
	engin:=gin.Default()
	engin.POST

}