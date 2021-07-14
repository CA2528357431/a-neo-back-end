package main

import (
"database/sql"
"fmt"
"github.com/labstack/echo"
)




type DataDeleteFavorite struct {
	Greet    string
}

type MsgSuccessDeleteFavorite struct {
	Success bool
	Data    DataDeleteFavorite
}

type MsgFailDeleteFavorite struct {
	Success   bool
	Error     string
	ErrorHint string

}


func deleteFavorite(text echo.Context)  error{


	getId := text.Param("id")
	getGoodId := text.FormValue("goodId")
	getToken := text.Request().Header.Get("token")

	if tokenCheck(getToken,getId){
		dbDeleteFavorite := connect()

		deleteDeleteFavorite(*dbDeleteFavorite, getGoodId, getId)

		successJson := createSuccessJsonDeleteFavorite()

		fmt.Println("\n*******delete favorite*********\n\n", "\nid:", getId, "\ngoodId:", getGoodId, "\nsuccess", "\n\n\n****************")

		return text.JSON(200, successJson)
	}else{
		failJson := createFailJsonDeleteFavorite()

		fmt.Println("\n*******delete favorite*********\n\n", "\nid:", getId, "\ngoodId:", getGoodId, "\nfail", "\n\n\n****************")

		return text.JSON(200, failJson)
	}


}

func deleteDeleteFavorite(dbDeleteFavorite sql.DB,getGoodId string, getId string){
	sql := "DELETE FROM goods WHERE \"goodId\" = "+getGoodId+" AND \"id\" = "+getId
	fmt.Println(sql)

	_, err := dbDeleteFavorite.Exec(sql)
	checkErr(err)

}


func createSuccessJsonDeleteFavorite() MsgSuccessDeleteFavorite {
	successData := new(DataDeleteFavorite)
	successData.Greet= "delete finish, bro"
	successJson:=new(MsgSuccessDeleteFavorite)
	successJson.Success = true
	successJson.Data = *successData
	return *successJson

}

func createFailJsonDeleteFavorite() MsgFailDeleteFavorite {
	failJson := new(MsgFailDeleteFavorite)
	failJson.Success = false
	failJson.Error = "you don't have such right to do so"
	failJson.ErrorHint = "please sign in first"
	return *failJson

}