package target

import "github.com/jmoiron/sqlx"

func deferSqlxOpen() *sqlx.DB {
	pg, err := sqlx.Open("driver", "dsn") // want `\Qcan rewrite as defer 'pg.Close()'`
	if err != nil {
		return nil
	}

	return pg
}
