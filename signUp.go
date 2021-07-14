package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

type DataSignUp struct {
	Username string
	Greet    string

}

type MsgSuccessSignUp struct {
	Success bool
	Data    DataSignUp
}

type MsgFailSignUp struct {
	Success   bool
	Error     string
	ErrorHint string

}




func signUp(text echo.Context) error {

	getName := text.FormValue("name")
	getPassword := text.FormValue("password")
	getEmail := text.FormValue("email")
	getCode := text.FormValue("code")

	code := redisGet(getEmail)

	dbSignUp := connect()

	rows := getRowsSignUp(*dbSignUp,getEmail)

	getJudgeEmail := judgeSignUp(rows)
	getJudgeCode := judgeCode(getCode,code)



	successJson := createSuccessJsonSignUp(getName)

	failJson := createFailJsonSignUp()



	if getJudgeEmail && getJudgeCode {

		neoId := getNeoIdSignUp(*dbSignUp)

		insertSignUp(*dbSignUp,neoId,getName,getPassword,getEmail)

		fmt.Println("\n*******sign up*********\n\n","\nname:",getName,"\npassword:",getPassword,"\nsuccess","\n\n\n****************")

		return text.JSON(200,successJson)
	}else {

		fmt.Println("\n*******sign up*********\n\n","\nname:",getName,"\npassword:",getPassword,"\nfail","\n\n\n****************")

		return text.JSON(200,failJson)
	}



}


func getRowsSignUp(dbSignUp sql.DB, getEmail string) sql.Rows{

	query := "SELECT * FROM users WHERE email =  "+"'"+getEmail+"'"
	rows, err := dbSignUp.Query(query)
	checkErr(err)

	return *rows
}

func judgeSignUp(rows sql.Rows) bool{
	getJudge := true
	for rows.Next() {
		getJudge = false
	}
	return getJudge
}

func createFailJsonSignUp() MsgFailSignUp{
	failJson := new(MsgFailSignUp)
	failJson.Success =false
	failJson.Error ="this email is uesd or code is wrong"
	failJson.ErrorHint ="try another"
	return *failJson
}

func createSuccessJsonSignUp(getName string) MsgSuccessSignUp  {
	successData := new(DataSignUp)
	successData.Greet ="welcome to join us, "+getName
	successData.Username =getName
	successJson := new(MsgSuccessSignUp)
	successJson.Success =true
	successJson.Data = *successData
	return *successJson
}

func getNeoIdSignUp(dbSignUp sql.DB) int {
	query := "SELECT MAX(id) FROM users "
	rows, err := dbSignUp.Query(query)
	checkErr(err)
	var neoId int
	for rows.Next(){
		_ = rows.Scan(&neoId)
	}
	neoId++
	return neoId
}

func insertSignUp(dbSignUp sql.DB,neoId int,getName string, getPassword string, getEmail string){
	prepare := "INSERT INTO users(id,name,password,deleted,email) values($1,$2,$3,$4,$5)"

	pre, err := dbSignUp.Prepare(prepare)
	checkErr(err)

	ins, err := pre.Exec(neoId, getName, getPassword,"false",getEmail)
	checkErr(err)

	affect, err := ins.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}