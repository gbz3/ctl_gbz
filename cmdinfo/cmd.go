/*
  cmdinfo
  のコメント
  です。
*/
package cmdinfo

import (
  "context"
  "fmt"
  "strconv"

  base "github.com/gbz3/ctl_gbz"
  ds "github.com/gbz3/ctl_gbz/dstore"
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

// コマンド名
func ( self cmdinfoCmd ) Name() string {
  return self.name
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
func ( self cmdinfoCmd ) Execute( ctx context.Context ) ( r string, err error ) {
fmt.Printf( "debug: args[1]=%s args[2]=%s\n", self.args[0], self.args[1] )

  var output string
  switch self.args[0] {
    case "cmd":
      rows, err := ds.GetCmdAll( ctx, self.args[1] )
      if err != nil {
        return "", err
      }
      for _, row := range rows {
        output += fmt.Sprintf( "%s: %d\n", row.Name, row.Id )
      }
    case "id":
      var result *ds.Cmd
      i, _ := strconv.Atoi( self.args[1] )
      if result, err = ds.GetCmdOne( ctx, int64( i ) ); err != nil {
        return "", err
      }
      output = fmt.Sprintf( "%s: %d\n", result.Name, result.Id )
    default:
      return "", fmt.Errorf( "arg[1] must be [cmd] or [id]. [%s]", self.args[1] )
  }
  r = fmt.Sprintf( "# Header #\n%s", output )
  return r, nil
}

