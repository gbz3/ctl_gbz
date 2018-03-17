package data_store

import (
  "fmt"
)

// コマンド情報登録
func StoreCmd( lockKey string ) ( id int ) {
  fmt.Printf( "debug: StoreCmd( lockKey=%s ) %d\n", lockKey, len( lockKey ) )
  return len( lockKey )
}

// コマンド情報削除
func ClearCmd( id int ) {
  fmt.Printf( "debug: ClearCmd( id=%d )\n", id )
}
