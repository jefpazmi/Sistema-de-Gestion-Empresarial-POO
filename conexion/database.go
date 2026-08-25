package conexion

import (
	"database/sql"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
)

// Conexion representa la conexión con SQL Server.
var DB *sql.DB

// Conectar abre la conexión con SQL Server.
func Conectar() error {

	servidor := `KAYLA03`

	baseDatos := "SISTEMA_GESTION_EMPRESARIAL"

	cadenaConexion := fmt.Sprintf(
		"server=%s;port=1433;database=%s;trusted_connection=yes;encrypt=disable",
		servidor,
		baseDatos,
	)

	var err error

	DB, err = sql.Open("sqlserver", cadenaConexion)
	if err != nil {
		return fmt.Errorf("error al abrir la conexión: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("error al conectarse a SQL Server: %v", err)
	}

	return nil
}