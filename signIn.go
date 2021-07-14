package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)


type DataSignIn struct {
	Username string
	Greet    string
	Id int

}

type MsgSuccessSignIn struct {
	Success bool
	Data    DataSignIn
	Token string
}

type MsgFailSignIn struct {
	Success   bool
	Error     string
	ErrorHint string

}


func signIn(text echo.Context) error {

	getEmail := text.FormValue("email")
	getPassword := text.FormValue("password")


	dbSignIn := connect()

	rows := getRowsSignIn(*dbSignIn, getEmail)


	//check if this name exists and get the true pw
	getJudge, password, id ,name:= judgeSignIn(rows)


	// match name and pw
	if getJudge==true&&getPassword==password{

		neoId := fmt.Sprintf("%d",id)

		token := createToken(neoId)
		fmt.Println(token)

		//json creation
		successJson := createSuccessJsonSignIn(name,id,token)

		fmt.Println("\n********sign in********\n\n","\nname:", name,"\nemail:", getEmail,"\npassword:",getPassword,"\nsuccess","\n\n\n****************")

		return text.JSON(200,successJson)

	}else {

		//json creation
		failJson := createFailJsonSignIn()

		fmt.Println("\n********sign in********\n\n","\nname:", getEmail,"\npassword:",getPassword,"\nfail","\n\n\n****************")

		return text.JSON(200,failJson)
	}

}


func createSuccessJsonSignIn(getName string,id int, t string) MsgSuccessSignIn{
	successData := new(DataSignIn)
	successData.Username = getName
	successData.Id = id
	successData.Greet = "welcome back "+getName
	successJson := new(MsgSuccessSignIn)
	successJson.Success=true
	successJson.Data=*successData
	successJson.Token=t
	return *successJson
}


func createFailJsonSignIn() MsgFailSignIn {
	failJson := new(MsgFailSignIn)
	failJson.Error="your name doesn't match the password"
	failJson.ErrorHint="please recheck you name and password"
	return *failJson
}

func judgeSignIn(rows sql.Rows) (bool,string,int,string){
	getJudge := false

	var id int
	var password string
	var name string

	for rows.Next() {

		var err error

		err = rows.Scan(&password,&id,&name)
		checkErr(err)

		getJudge = true
	}

	return getJudge,password,id,name
}

func getRowsSignIn(dbSignIn sql.DB, getEmail string) sql.Rows{
	var err error
	query := "SELECT password,id,name FROM users WHERE email =  "+"'"+ getEmail +"'"
	rows, err := dbSignIn.Query(query)
	checkErr(err)

	return *rows
}
