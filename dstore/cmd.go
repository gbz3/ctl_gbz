package dstore

import (
  "context"
  "database/sql"
  "fmt"

  _ "github.com/mattn/go-sqlite3"
)

// コマンド情報登録
func StoreCmd( ctx context.Context, name string, lockKey string ) ( id int64, err error ) {
fmt.Printf( "debug: sql.Drivers=%v\n", sql.Drivers() )
fmt.Printf( "debug: StoreCmd( name=%s, lockKey=%s )\n", name, lockKey )

  // DB接続
  db, err := sql.Open( "sqlite3", "./test.db" )
  if err != nil { return -1, err }
  defer db.Close()

  // テーブル作成
  _, err = db.ExecContext( ctx, `CREATE TABLE IF NOT EXISTS "CMD" ( "ID" INTEGER PRIMARY KEY AUTOINCREMENT, "NAME" VARCHAR(255), "LOCK_KEY" VARCHAR(255) )`, )
  if err != nil { return -1, err }

  // トランザクション開始
  tx, err := db.BeginTx( ctx, &sql.TxOptions{} )
  if err != nil { return -1, err }
  defer tx.Rollback()

  // SQL準備
  stmt, err := tx.PrepareContext( ctx, `INSERT INTO CMD ( NAME, LOCK_KEY ) VALUES ( ?, ? )` )
  if err != nil { return -1, err }
  defer stmt.Close()

  // レコード挿入
  res, err := stmt.ExecContext( ctx, name, lockKey )
  if err != nil { return -1, err }

  i, err := res.LastInsertId()
  if err != nil { return -1, err }

  tx.Commit()
  return i, nil
}

// コマンド情報削除
func ClearCmd( ctx context.Context, id int64 ) {
fmt.Printf( "debug: ClearCmd( id=%d )\n", id )

  // DB接続
  db, err := sql.Open( "sqlite3", "./test.db" )
  if err != nil { panic( err ) }
  defer db.Close()

  // トランザクション開始
  tx, err := db.BeginTx( ctx, &sql.TxOptions{} )
  if err != nil { panic( err ) }
  defer tx.Rollback()

  // SQL準備
  stmt, err := tx.PrepareContext( ctx, `DELETE FROM CMD WHERE ID=?` )
  if err != nil { panic( err ) }
  defer stmt.Close()

  // レコード削除
  res, err := stmt.ExecContext( ctx, id )
  if err != nil { panic( err ) }

  i, _ := res.RowsAffected()
fmt.Printf( "debug: Delete %d rows.\n", i )

  tx.Commit()
}

type Cmd struct {
  Id int64
  Name string
  LockKey string
}

// コマンド情報取得
func GetCmdOne( ctx context.Context, id int64 ) ( *Cmd, error ) {
fmt.Printf( "debug: GetCmdOne( id=%d )\n", id )

  // DB接続
  db, err := sql.Open( "sqlite3", "./test.db" )
  if err != nil { panic( err ) }
  defer db.Close()

  // トランザクション開始
  tx, err := db.BeginTx( ctx, &sql.TxOptions{} )
  if err != nil { panic( err ) }
  defer tx.Rollback()

  // 1件取得
  r := &Cmd{}
  if err = db.QueryRowContext( ctx, `SELECT ID, NAME, LOCK_KEY FROM CMD WHERE ID = ?`, id ).Scan( &r.Id, &r.Name, &r.LockKey ); err != nil {
    return nil, err
  }

fmt.Printf( "debug: Id=%d Name=%s LockKey=%s\n", r.Id, r.Name, r.LockKey )
  return r, nil
}

// コマンド情報取得
func GetCmdAll( ctx context.Context, name string ) ( []Cmd, error ) {
fmt.Printf( "debug: GetCmdAll( name=%s )\n", name )

  // DB接続
  db, err := sql.Open( "sqlite3", "./test.db" )
  if err != nil { panic( err ) }
  defer db.Close()

  // トランザクション開始
  tx, err := db.BeginTx( ctx, &sql.TxOptions{} )
  if err != nil { panic( err ) }
  defer tx.Rollback()

  rows, err := db.Query( `SELECT ID, NAME, LOCK_KEY FROM CMD WHERE NAME = ?`, name )
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  var result []Cmd
  for rows.Next() {
    cmd := Cmd{}
    if err = rows.Scan( &cmd.Id, &cmd.Name, &cmd.LockKey ); err != nil {
      return nil, err
    }
    result = append( result, cmd )
  }

  return result, nil
}

