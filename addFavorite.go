package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
)



type DataAddFavorite struct {
	Greet    string

}

type MsgSuccessAddFavorite struct {
	Success bool
	Data    DataAddFavorite
}

type MsgFailAddFavorite struct {
	Success   bool
	Error     string
	ErrorHint string

}



func addFavorite(text echo.Context)  error{



	getId := text.Param("id")
	getGoodName := text.FormValue("goodName")
	getGoodUrl := text.FormValue("goodUrl")
	getImageUrl := text.FormValue("imageUrl")
	getGoodSrc := text.FormValue("goodSrc")
	getPrice := text.FormValue("price")
	getToken := text.Request().Header.Get("token")


	if tokenCheck(getToken,getId){
		dbAddFavorite := connect()

		rows := getRowsAddFavorite(*dbAddFavorite)

		goodId := getGoodIdAddFavorite(rows)

		insertAddFavorite(*dbAddFavorite, goodId, getId, getGoodName, getGoodUrl, getImageUrl, getGoodSrc, getPrice)

		successJson := createSuccessJsonAddFavorite()

		fmt.Println("\n*******add favorite*********\n\n", "\nid:", getId, "\ngoodName:", getGoodName,
			"\ngetPrice:", getPrice, "\ngetGoodSrc:", getGoodSrc, "\ngetGoodUrl:", getGoodUrl,
			"\ngetImageUrl:", getImageUrl, "\nsuccess", "\n\n\n****************")

		return text.JSON(200, successJson)
	} else{
		failJson := createFailJsonAddFavorite()

		fmt.Println("\n*******add favorite*********\n\n", "\nid:", getId, "\ngoodName:", getGoodName,
			"\ngetPrice:", getPrice, "\ngetGoodSrc:", getGoodSrc, "\ngetGoodUrl:", getGoodUrl,
			"\ngetImageUrl:", getImageUrl, "\nfail", "\n\n\n****************")

		return text.JSON(200, failJson)

	}

}

func getGoodIdAddFavorite(rows sql.Rows) int {
	var goodId int

	for rows.Next(){
		err := rows.Scan(&goodId)
		checkErr(err)
	}
	goodId++

	return goodId
}

func insertAddFavorite(dbAddFavorite sql.DB,goodId int,getId string,getGoodName string,getGoodUrl string,getImageUrl string, getGoodSrc string, getPrice string){
	prepare := "INSERT INTO goods(\"goodId\",\"id\",\"goodName\",\"goodUrl\",\"imageUrl\",\"goodSrc\",\"price\") values($1,$2,$3,$4,$5,$6,$7)"

	pre,err := dbAddFavorite.Prepare(prepare)
	checkErr(err)

	ins, err := pre.Exec(goodId,getId,getGoodName,getGoodUrl,getImageUrl,getGoodSrc,getPrice)
	checkErr(err)

	ins.RowsAffected()
}

func getRowsAddFavorite(dbAddFavorite sql.DB) sql.Rows {

	query := "SELECT MAX(\"goodId\") FROM goods"
	rows,err :=dbAddFavorite.Query(query)
	checkErr(err)

	return *rows

}

func createSuccessJsonAddFavorite() MsgSuccessAddFavorite {
	successData := new(DataAddFavorite)
	successData.Greet= "your favorite, guy"
	successJson:=new(MsgSuccessAddFavorite)
	successJson.Success = true
	successJson.Data = *successData
	return *successJson

}

func createFailJsonAddFavorite() MsgFailAddFavorite {
	failJson := new(MsgFailAddFavorite)
	failJson.Success = false
	failJson.Error = "you don't have such right to do so"
	failJson.ErrorHint = "please sign in first"
	return *failJson

}