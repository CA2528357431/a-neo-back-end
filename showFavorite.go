package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
)

type DataItem struct {
	GoodId int
	GoodName string
	GoodUrl string
	ImageUrl string
	GoodSrc int
	GoodPrice int

}
type MsgSuccessShowFavorite struct {
	Success bool
	Data    []DataItem
}

type MsgFailShowFavorite struct {
	Success   bool
	Error     string
	ErrorHint string

}


func showFavorite(text echo.Context)  error{

	getId := text.Param("id")
	getToken := text.Request().Header.Get("token")


	if tokenCheck(getToken,getId){
		dbShowFavorite := connect()

		rows := getRowsShowFavorite(*dbShowFavorite, getId)

		sli := getDataShowFavorite(rows)

		successJson := createSuccessJsonShowFavorite(sli)

		fmt.Println("\n*******show favorite*********\n\n")
		for i := 0; i < len(sli); i++ {
			fmt.Println(sli[i])
		}
		fmt.Println("\nsuccess")
		fmt.Println("\nsuccess\n\n\n****************")

		return text.JSON(200, successJson)
	}else {
		failJson := createFailJsonShowFavorite()

		fmt.Println("\n*******show favorite*********\n\n")
		fmt.Println("\nfail")
		fmt.Println("\nsuccess\n\n\n****************")

		return text.JSON(200,failJson)

	}


}

func getRowsShowFavorite(dbShowFavorite sql.DB,getId string)  sql.Rows{
	query := "SELECT \"goodId\",\"goodName\",\"goodUrl\",\"imageUrl\",\"goodSrc\",\"price\" FROM goods WHERE id = "+getId
	fmt.Println(query)

	rows, err := dbShowFavorite.Query(query)
	checkErr(err)

	return *rows

}

func createSuccessJsonShowFavorite (sli []DataItem) MsgSuccessShowFavorite {
	successJson := new(MsgSuccessShowFavorite)
	successJson.Data = sli
	successJson.Success = true
	return *successJson
}

func getDataShowFavorite(rows sql.Rows)  []DataItem{
	sli := make([]DataItem,0)

	for rows.Next(){
		item := new(DataItem)

		err := rows.Scan(&item.GoodId,&item.GoodName,&item.GoodUrl,&item.ImageUrl,&item.GoodSrc,&item.GoodPrice)
		checkErr(err)

		sli = append(sli,*item)
	}

	return sli
}

func createFailJsonShowFavorite() MsgFailShowFavorite {
	failJson := new(MsgFailShowFavorite)
	failJson.Success = false
	failJson.Error = "you don't have such right to do so"
	failJson.ErrorHint = "please sign in first"
	return *failJson

}