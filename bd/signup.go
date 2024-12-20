package bd

import (
	"fmt"

	"github.com/gast0nrb/models"
	"github.com/gast0nrb/tools"
	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza registro")

	err := DbConnect()

	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + ",'" + tools.FechaMysql() + "'')"
	fmt.Println(sentencia)

	_, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
