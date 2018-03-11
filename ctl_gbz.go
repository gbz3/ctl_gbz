package ctl_gbz

import (
  "fmt"
  "time"
)

const (
  ctl_name = "ctl_gbz"
)

// コマンド名称とコンストラクタ
type CmdRegister struct {
  name string
  loader func( []string ) SyncCmd
}

var CmdAll []CmdRegister

// 各コマンドのコンストラクタを登録する
// 各コマンドの具象クラスの init() で登録する
func RegistCmd( name string, loader func( []string ) SyncCmd ) {
  CmdAll = append( CmdAll, CmdRegister{ name, loader } )
}

func Main( args []string ) string {
  cmd, err := dispatchCmd( args )
  if err != nil {
    panic( err )
  }

  if err = cmd.CheckArgs( args ); err != nil {
    panic( err )
  }

  var result string
  if result, err = cmd.Execute(); err != nil {
    panic( err )
  }

  return fmt.Sprintf( "%s\n%s %s\n", result, time.Now().Format( "2006/01/02 15:04:05" ), ctl_name )
}

// コマンドのインタフェース宣言
type SyncCmd interface {
  CheckArgs( []string ) error
  Execute() ( string, error )
}

// 第一引数のコマンド名から適切なコマンドオブジェクトを作成
func dispatchCmd( args []string ) ( cmd SyncCmd, err error ) {
  if ( args == nil || len( args ) < 1 ) {
    return nil, fmt.Errorf( "too few arguments." )
  }

  for _, c := range CmdAll {
    if c.name == args[0] {
      return c.loader( args ), nil
    }
  }

  return nil, fmt.Errorf( "unknown command [%s]", args[0] )
}

