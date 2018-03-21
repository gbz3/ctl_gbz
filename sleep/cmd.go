/*
  sleep
  のコメント
  です。
*/
package sleep

import (
  "context"
  "fmt"
  "strconv"
  "time"

  base "github.com/gbz3/ctl_gbz"
)

func init() {
  base.RegistCmd( "sleep", newCmd )
}

// コマンドの情報を表示
type sleepCmd struct {
  name string
  args []string
}

// コンストラクタ
func newCmd( args []string ) ( cmd base.SyncCmd ) {
  return &sleepCmd{ args[0], args[1:] }
}

// コマンド名
func ( self sleepCmd ) Name() string {
  return self.name
}

// 引数をチェック
func ( self sleepCmd ) CheckArgs( args []string ) ( err error ) {
  if args == nil || len( args ) != 4 {
    return fmt.Errorf( "illegal signature. %v", args )
  }

  if args[0] != self.name {
    return fmt.Errorf( "illegal command name [%s]", args[0] )
  }

  if _, err := strconv.Atoi( args[1] ); err != nil {
    return fmt.Errorf( "arg[1] must be numeric. [%s]", args[1] )
  }

  if args[2] != "remote" {
    return fmt.Errorf( "arg[3] must be [remote]. [%s]", args[2] )
  }

  return nil;
}

// コマンドを実行
func ( self sleepCmd ) Execute( ctx context.Context ) ( r string, err error ) {
  timeSecond, _ := strconv.Atoi( self.args[0] )
  time.Sleep( time.Duration( timeSecond ) * time.Second )
  output := fmt.Sprintf( "# Header #\n%s: %d\n", self.name, timeSecond )
  return output, nil;
}

