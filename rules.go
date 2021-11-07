package gorules

import (
	_ "github.com/jmoiron/sqlx"
	"github.com/quasilyte/go-ruleguard/dsl"
)

func deferSqlxOpen(m dsl.Matcher) {
	m.Import("github.com/jmoiron/sqlx")

	m.Match("$db, $err := sqlx.Open($*_, $*_); $chk; $close").
		Where(m["db"].Type.Is(`*sqlx.DB`) &&
			m["err"].Type.Implements(`error`) &&
			!m["close"].Text.Matches(`defer .*\.Close()`)).
		Report("can rewrite as defer '$db.Close()'")
}
