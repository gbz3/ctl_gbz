/*
  cmdinfo
  のコメント
  です。
*/
package cmdinfo

import (
  "fmt"

  base "github.com/gbz3/ctl_gbz"
)

func init() {
  base.RegistCmd( "cmdinfo", newCmd )
}

// コマンドの情報を表示
type cmdinfoCmd struct {
  name string
  args []string
}

// コンストラクタ
func newCmd( args []string ) ( cmd base.SyncCmd ) {
  return &cmdinfoCmd{ args[0], args[1:] }
}

// 引数をチェック
func ( self cmdinfoCmd ) CheckArgs( args []string ) ( err error ) {
  if args == nil || len( args ) != 5 {
    return fmt.Errorf( "illegal signature. %v", args )
  }

  if args[0] != self.name {
    return fmt.Errorf( "illegal command name [%s]", args[0] )
  }

  if args[1] != "cmd" && args[1] != "id" {
    return fmt.Errorf( "arg[1] must be [cmd] or [id]. [%s]", args[1] )
  }

  if args[3] != "remote" {
    return fmt.Errorf( "arg[3] must be [remote]. [%s]", args[3] )
  }

  return nil;
}

// コマンドを実行
func ( self cmdinfoCmd ) Execute() ( r string, err error ) {
  return "# Header #\ncmdinfo: 9999\n", nil;
}

