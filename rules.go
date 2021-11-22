package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func deferHttpPackage(m dsl.Matcher) {
	m.Match("$req, $err := http.NewRequest($*_, $*_, $*_); $chk; $close").
		Where(m["req"].Type.Is(`*http.Request`) &&
			m["err"].Type.Implements(`error`) &&
			!m["close"].Text.Matches(`defer .*\.Body.Close()`)).
		Report("can rewrite as 'defer $req.Body.Close()'")

	m.Match("$resp, $err := http.Get($*_); $chk; $close").
		Where(m["resp"].Type.Is(`*http.Response`) &&
			m["err"].Type.Implements(`error`) &&
			!m["close"].Text.Matches(`defer .*\.Body.Close()`)).
		Report("can rewrite as 'defer $resp.Body.Close()'")

	m.Match("$resp, $err := http.Post($*_, $*_, $*_); $chk; $close").
		Where(m["resp"].Type.Is(`*http.Response`) &&
			m["err"].Type.Implements(`error`) &&
			!m["close"].Text.Matches(`defer .*.Body.Close()`)).
		Report("can rewrite as 'defer $resp.Body.Close()'")
}

func deferOsPackage(m dsl.Matcher) {
	m.Match("$f, $err := os.Open($*_); $chk; $close").
		Where(m["f"].Type.Is(`*os.File`) &&
			m["err"].Type.Implements(`error`) &&
			!m["close"].Text.Matches(`defer .*\.Close()`)).
		Report("can rewrite as 'defer $f.Close()'")

	m.Match("$f, $err := os.Create($*_); $chk; $close").
		Where(m["f"].Type.Is(`*os.File`) &&
			m["err"].Type.Implements(`error`) &&
			!m["close"].Text.Matches(`defer .*\.Close()`)).
		Report("can rewrite as 'defer $f.Close()'")
}

func deferSqlPackage(m dsl.Matcher) {
	m.Match("$db, $err := sql.Open($*_, $*_); $chk; $close").
		Where(m["db"].Type.Is(`*sql.DB`) &&
			m["err"].Type.Implements(`error`) &&
			!m["close"].Text.Matches(`defer .*\.Close()`)).
		Report("can rewrite as 'defer $db.Close()'")

	m.Match("$rows, $err := $db.Query($*_, $*_); $chk; $close").
		Where((m["db"].Type.Is(`*sql.DB`) || m["db"].Type.Is(`*sql.Stmt`)) &&
			m["rows"].Type.Is(`*sql.Rows`) &&
			m["err"].Type.Implements(`error`) &&
			!m["close"].Text.Matches(`defer .*\.Close()`)).
		Report("can rewrite as 'defer $rows.Close()'")

	m.Match("$stmt, $err := $db.Prepare($*_, $*_); $chk; $close").
		Where((m["db"].Type.Is(`*sql.DB`)) &&
			m["stmt"].Type.Is(`*sql.Stmt`) &&
			m["err"].Type.Implements(`error`) &&
			!m["close"].Text.Matches(`defer .*\.Close()`)).
		Report("can rewrite as 'defer $stmt.Close()'")
}
