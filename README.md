# ctl_gbz

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
```

### Git設定

```
git config --global user.name "First-name Family-name"
git config --global user.email username@example.com
git config --global core.editor 'vim -c "set fenc=utf-8"'
```
