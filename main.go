package main
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
func main(){

	fmt.Println("GO - MSQL Server Connection!")
	db,err:=sql.Open("mysql","root:root@tcp(localhost:3306)/goblog")
	if err !=nil{
		panic(err.Error())
	}
	fmt.Println("Connected Succesfull!")
	defer db.Close()
	insert,err:=db.Query("insert into employee(name,city) values('cesar','callao'),('miguel','lima');")
	if err !=nil{
		panic(err.Error())
	}
	fmt.Println("Succesfull insert employees for goblog db!")
	defer insert.Close()
}