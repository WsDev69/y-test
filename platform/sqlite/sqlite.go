package sqlite

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"y-test/internal/constant/query"
)

const driverName = "sqlite3"

type connectionStr struct {
	dbPath string
}

func (c *connectionStr) Open() *Sqlite {
	db, err := sqlx.Connect(driverName, c.dbPath)
	if err != nil {
		logrus.Fatalln("SQLITE Open connections failed: ", err)
	}

	db.MustExec(query.Schema)
	return &Sqlite{
		DB: db,
	}
}

type Connections interface {
	Open() *Sqlite
}

type Sqlite struct {
	DB *sqlx.DB
}

func InitConnections(dbPath string) Connections {
	return &connectionStr{dbPath: dbPath}
}
