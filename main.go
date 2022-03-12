package main

import (
	"fmt"
	"todo_app/app/controllers"
	"todo_app/app/models"
)

func main() {

	// データベースを作成するためにinit関数を呼び出す
	fmt.Println(models.Db)

	// サーバの立ち上げ
	controllers.StartMainServer()

	/*
		// sessionの取得
		user, _ := models.GetUserByEmail("sample@sample.com")
		fmt.Println("GETUSER :", user)

		session, err := user.CreateSession()

		if err != nil {
			log.Println(err)
		}
		fmt.Println("GET SESSION :", session)

		valid, err := session.CheckSession()
		fmt.Println("Valid", valid)
	*/

	//参考：configファイルの中身

	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/

	// 参考：userテーブルのCRUD操作

	/*
		// userを作成する
		u := &models.User{}
		u.Name = "test2"
		u.Email = "test2@example.com"
		u.Password = "testtest"
		fmt.Println(u)

		u.CreateUser()
	*/

	/*
		// userを取得する
		u, _ := models.GetUser(1)

		fmt.Println(u)

		// userを更新する
		u.Name = "Test2"
		u.Email = "test2@example.com"
		u.UpdateUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)

		// userを削除する
		u.DeleteUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)
	*/

	/*
		// userの取得、userのメソッドとしてCreateTodoを実行
		user, _ := models.GetUser(2)
		user.CreateTodo("First Todo")
	*/

	// 参考：todoテーブルのCRUD操作

	/*
		// Todoを一件取得
		t, _ := models.GetTodo(1)
		fmt.Println(t)
	*/

	/*
		// Todoを全件取得
		todos, _ := models.GetTodos()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	/*
		// 特定のユーザのTodoを全件取得
		user2, _ := models.GetUser(3)
		todos, _ := user2.GetTodosByUser()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	/*
		// Todoを更新する
		t, _ := models.GetTodo(1)
		t.Content = "Update Todo"
		t.UpdateTodo()
	*/

	/*
		//Todoを削除する
		t, _ := models.GetTodo(3)
		t.DeleteTodo()
	*/

}
