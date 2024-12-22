# Technical-Notes

経験した技術やテクニック等をメモするためのリポジトリ

■git
・空のコミットを行いたい場合のコマンド
git commit --allow-empty -m "空コミット"

■docker
・端末の容量が逼迫してきた場合にdockerの不要ファイルを消すコマンド
docker system prune
docker system volume prune

・dockerコンテナ内でのログを確認する場合のコマンド
docker compose logs -f

・特定のコンテナのみ閲覧したい場合は以下
docker compose logs コンテナ名 -f

・docker ps で確認したコンテナIDでも確認可能
docker ps -a → CONTAINER IDを確認
docker exec -it コンテナID bash

■ツール
・swagger(API設計書作成ツール)
swaggerのhtml出力用コマンド(windows時のコマンドのためmacでは異なる可能性)
npm install -g redoc-cli

swaggerファイル(.yml)のあるディレクトリまで移動し以下のコマンド
redoc-cli bundle swagger.yml(読み込みファイル名) -o index.html(出力ファイル名)

■VSCode
・自動フォーマットを行わずに保存する方法
Ctrl + K → Ctrl + Shift + S

■言語
・go test コマンド
テストファイルのあるディレクトリに移動して以下コマンドを叩く
go test -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html

ルートディレクトリで以下テストを走らせると全て通る テストは一つずつ実行される
go test -p 1 ./... -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html