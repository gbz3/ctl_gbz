package ctl_gbz_cmdctl

import (
  "context"
  "fmt"

  base "github.com/gbz3/ctl_gbz"
)

func init() {
  base.RegistCmd( "cmdctl", newCmd )
}

// コマンドの情報を表示
type cmdctlCmd struct {
  name string
  args []string
}

// コンストラクタ
func newCmd( args []string ) ( cmd base.SyncCmd ) {
  return &cmdctlCmd{ args[0], args[1:] }
}

// コマンド名                                        
func ( self cmdctlCmd ) Name() string {
  return self.name
}

// 引数をチェック
func ( self cmdctlCmd ) CheckArgs( args []string ) ( err error ) {
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
func ( self cmdctlCmd ) Execute( ctx context.Context ) ( r string, err error ) {
  return "# Header #\ncmdctl: 9999\n", nil;
}

