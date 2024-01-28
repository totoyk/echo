# trial-api-golang

## Overview

### What is this?
GolangでAPI実装することが増えたので、改めてGolangでAPI実装するときのBestPracticeを考察したり、まだ知らないことたくさんあるので見識深めるためにいじってるモノ。

## Usage

### docker-compose
- [app](#app)
- [db](#db)
- [codegen](#codegen) (depends_on: [oapigen](#oapigen))
- [oapigen](#oapigen)

### app
Echo(Golang)でつくったWebAPIアプリケーションのコンテナ。

### db
ローカル開発用に用意したmysql8.0のコンテナ。

### codegen
oapi-codegenを利用してopenapi.yamlからappのRouterとStructを生成するコンテナ。

### oapigen
外部参照しているspec/openapi.yamlを.distに統合するためのコンテナ。

## Roadmap
- ORMはほんとにGorm一択なのか
- テストコードってコレいるのか（一般論ではなく）
- ログってどうするのがよいのか
- その他いろいろ
