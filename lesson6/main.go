package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"lesson6/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	db, err := sql.Open("mysql", "admin_user:admin@tcp(localhost:3306)/migrate_test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// sqlboiler用の設定
	boil.SetDB(db)
	ctx := context.Background()

	// user登録.
	/* user := models.User{
		ID:       "1",
		Username: "test_user",
		Email:    "test@email",
	}
	err = user.Insert(ctx, db, boil.Infer())
	if err != nil {
		log.Fatal(err)
	} */

	// ユーザ複数登録.
	/* userSlice := models.UserSlice{
		{
			ID:       "1",
			Username: "test_user1",
			Email:    "test1@email",
		},
		{
			ID:       "2",
			Username: "test_user2",
			Email:    "test2@email",
		},
		{
			ID:       "3",
			Username: "test_user3",
			Email:    "test3@email",
		},
	}
	err = userSlice.InsertAll(ctx, db, boil.Infer())
	if err != nil {
		log.Fatal(err)
	} */

	// user更新.
	/* user, err := models.FindUser(ctx, db, "1")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)

	user.Username = "update_user"
	_, err = user.Update(ctx, db, boil.Infer())
	if err != nil {
		log.Fatal(err)
	} */

	// user削除.
	/* user, err := models.FindUser(ctx, db, "1")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)

	_, err = user.Delete(ctx, db, false)
	if err != nil {
		log.Fatal(err)
	} */

	// usersテーブルから全ユーザーを取得.
	users, err := models.Users().All(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", users)
}
