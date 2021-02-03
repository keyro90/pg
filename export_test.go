package pg

import "github.com/keyro90/pg/v10/internal/pool"

func (db *DB) Pool() pool.Pooler {
	return db.pool
}

func (ln *Listener) CurrentConn() *pool.Conn {
	return ln.cn
}
