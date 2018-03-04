# ctl_gbz

## WSLの設定

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
