Go Training Course
==================

### 概要

本リポジトリは、第1期 (2016年)「Go 基本技術習得コース」の @mikan の課題を保管・共有するためのリポジトリです。

#### 課題本

[The Go Programming Language](http://www.gopl.io/)
 (著者： Alan A. A. Donovan, Brian W. Kernighan ISBN：978-0134190440)

### 実施説明

#### プロジェクト構成

ソースコードディレクトリ:

:file_folder: [src/github.com/mikan/gopl](src/github.com/mikan/gopl)

#### 開発環境

golang:

* 1.5.1

IDE:

* JetBrains IntelliJ IDEA Ultimate 15 with [Go plugin](https://github.com/go-lang-plugin-org)
* Atom 1.4 with [go-plus plugin](https://atom.io/packages/go-plus)

OS:

* OS X 10.11
* Windows 10 Pro 64bit

#### メモ

##### 現在位置を GOPATH にする

BASH の場合 ([完全版](gopath.sh)):

```bash
export GOPATH="$(cd "$(dirname "${BASH_SOURCE:-$0}")"; pwd)"
```

バッチファイルの場合 ([完全版](gopath.bat)):

```
for /f "delims=" %%a in ('@cd') do setx GOPATH %%a
```

setx の都合上、コンソール起動前に実行する必要がある。

##### gofmt/goimports 全部適用

バッチファイルの場合 ([完全版](format.bat)):

```
for /r %%i in (*.go) do bin\goimports -l -w "%%i"
```

##### ベンチマークテストの実行

ch02/ex03 の例:

```bash
go test github.com/mikan/gopl/ch02/ex03 -bench=.
```

### 謝辞

#### 研修指導

@YoshikiShibata
