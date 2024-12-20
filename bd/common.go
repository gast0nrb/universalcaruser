package bd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gast0nrb/models"
	"github.com/gast0nrb/secretm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	SecretModel models.SecretRDSJson
	err         error
	Db          *sql.DB
)

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err := sql.Open("Mysql", ConnStr(SecretModel))

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexión exitosa a la BD")
	return nil
}

func ConnStr(clave models.SecretRDSJson) string {
	var dbUser, authToken, dbEndPoint, dbName string
	dbUser = clave.Username
	authToken = clave.Password
	dbEndPoint = clave.Host
	dbName = "webuniversalcar"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartetPasswords=true", dbUser, authToken, dbEndPoint, dbName)

	return dsn
}
