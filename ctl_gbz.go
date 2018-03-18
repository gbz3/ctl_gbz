package ctl_gbz

import (
  "context"
  "fmt"
  "io"
  "time"

  ds "github.com/gbz3/ctl_gbz/data_store"
)

const (
  ctl_name = "ctl_gbz"
)

// コマンド名称とコンストラクタ
type cmdRegister struct {
  name string
  loader func( []string ) SyncCmd
}

var cmdAll []cmdRegister

// 各コマンドのコンストラクタを登録する
// 各コマンドの具象クラスの init() で登録する
func RegistCmd( name string, loader func( []string ) SyncCmd ) {
  cmdAll = append( cmdAll, cmdRegister{ name, loader } )
}

func Main( args []string, stdout io.Writer, stderr io.Writer ) int {
  cmd, err := dispatchCmd( args )
  if err != nil {
    fmt.Fprintf( stdout, "%s %s %s\n", time.Now().Format( "2006/01/02 15:04:05" ), ctl_name, err )
    return 1;
  }

  if err = cmd.CheckArgs( args ); err != nil {
    fmt.Fprintf( stdout, "%s %s %s\n", time.Now().Format( "2006/01/02 15:04:05" ), ctl_name, err )
    return 1;
  }

  ctx, _ := context.WithTimeout( context.Background(), 2*time.Second )
  var id int64
  if id, err = ds.StoreCmd( ctx, cmd.Name(), cmd.Name() ); err != nil {
    fmt.Fprintf( stdout, "%s %s %s\n", time.Now().Format( "2006/01/02 15:04:05" ), ctl_name, err )
    return 1;
  }
  defer ds.ClearCmd( ctx, id )

  var output string
  if output, err = cmd.Execute(); err != nil {
    fmt.Fprintf( stdout, "%s %s %s\n", time.Now().Format( "2006/01/02 15:04:05" ), ctl_name, err )
    return 1;
  }

  fmt.Fprintf( stdout, "%s\n%s %s\n", output, time.Now().Format( "2006/01/02 15:04:05" ), ctl_name )
  return 0;
}

// コマンドのインタフェース宣言
type SyncCmd interface {
  Name() string
  CheckArgs( []string ) error
  Execute() ( string, error )
}

// 第一引数のコマンド名から適切なコマンドオブジェクトを作成
func dispatchCmd( args []string ) ( cmd SyncCmd, err error ) {
  if ( args == nil || len( args ) < 1 ) {
    return nil, fmt.Errorf( "too few arguments." )
  }

  for _, c := range cmdAll {
    if c.name == args[0] {
      return c.loader( args ), nil
    }
  }

  return nil, fmt.Errorf( "unknown command [%s]", args[0] )
}

