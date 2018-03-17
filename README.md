# ctl_gbz

## 呼び出し側main

```
package main

import (
  "os"

  "github.com/gbz3/ctl_gbz"
)

func main() {
  os.Exit( ctl_gbz.Main( os.Args[1:], os.Stdout, os.Stderr ) )
}
```

## WSLの設定

### bashの設定

~/.bashrc を編集する。

### apt

apt コマンドの基本的な使い方。

| コマンド | 動作 |
|---|---
| apt-get update | パッケージリスト情報の更新 |
| apt-get upgrade | 更新されているパッケージを最新に更新 |
| apt-get install パッケージ名 | 指定パッケージをインストール |
| apt-get remove パッケージ名 | 指定パッケージをアンインストール。設定ファイルは残す |
| apt-get purge パッケージ名 | 指定パッケージをインストール。設定ファイルも削除 |
| apt-cache search パッケージ名 | インストール可能なパッケージを検索。正規表現可 |
| apt-cache show パッケージ名 | パッケージの情報を表示 |
| dpkg -l [パッケージ名] | インストールされているパッケージをリスト表示 |
| dpkg -L パッケージ名 | パッケージでインストールされたファイルをリスト表示 |

### Go設定

```
export GOPATH=/home/kashiba/go
PATH="$PATH:$GOPATH/bin"
umask 022
→ ~/.bashrc にも設定

# Goパッケージ取得
go get github.com/gbz3/ctl_gbz
go get golang.org/x/crypto/ssh
```

### Git設定

```
git config --global user.name "First-name Family-name"
git config --global user.email username@example.com
git config --global core.editor 'vim -c "set fenc=utf-8"'
```

## ssh関連情報取得

```
## サーバに接続して公開鍵を表示  ※複数の鍵を利用可能であれば、鍵の種類分出力
$ ssh-keyscan -p <ポート番号> <IPアドレス>
# IPアドレス:ポート番号 SSH/OpenSSHのバージョン
[IPアドレス]:ポート番号 ssh-rsa 公開鍵
# IPアドレス:ポート番号 SSH/OpenSSHのバージョン
[IPアドレス]:ポート番号 ecdsa-sha2-nistp256 公開鍵
...

## known_hosts から指定サーバの公開鍵情報を表示
$ ssh-keygen -F [IPアドレス]:ポート番号
# Host  [IPアドレス]:ポート番号 found: line 1
|1|ホスト情報のハッシュ値？ ecdsa-sha2-nistp256 公開鍵

## known_hosts から指定サーバの公開鍵情報を削除
$ ssh-keygen -R [IPアドレス]:ポート番号
```
