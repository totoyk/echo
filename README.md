# sample echo webapp

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

## Development

### Run
```bash
docker-compose up
```

### Coding Step
1. spec/openapi.yamlを編集する
2. `docker-compose up codegen`でRouterとStructを生成する
3. internal/*を作成、編集する
4. `docker-compose up`でAPIサーバを起動する
