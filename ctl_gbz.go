package ctl_gbz

import (
  "fmt"
  "time"
)

const (
  ctl_name = "ctl_gbz"
)

var cmdAll []string

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
type syncCmd interface {
  CheckArgs()
  Execute()
}

// 第一引数のコマンド名から適切なコマンドオブジェクトを作成
func dispatchCmd( args []string ) ( cmd *cmdinfoCmd, err error ) {
  if ( args == nil || len( args ) < 1 ) {
    return nil, fmt.Errorf( "too few arguments." )
  }

  for _, c := range cmdAll {
    if c == args[0] {
      return newCmdinfo( args ), nil
    }
  }

  return nil, fmt.Errorf( "unknown command [%s]", args[0] )
}

