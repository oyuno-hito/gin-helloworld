# gin-helloworld
go+ginを利用したサーバサイドアプリケーションの開発

# 開発環境
## 各バージョン
https://github.com/oyuno-hito/gin-helloworld/issues/1 で設定済み
- go:  go1.21.1 darwin/arm64
- gin: [v1.9.1](https://github.com/gin-gonic/gin/releases/tag/v1.9.1)

００env系を利用したバージョン指定はhttps://github.com/oyuno-hito/gin-helloworld/issues/2 で対応する(優先度低)

# 実行手順
goのコンパイル全然わかっていない
## 実行
cd src
go run .

## ビルド
cd src
go build -o ../build/artifact

# やること検討リスト
- Dockerでの起動
- Sessionログイン
- ユニットテスト
- Redisでのセッション管理
- oapi-codegenの利用
  - cf. https://zenn.dev/rescuenow/articles/3c9a19eb2c0655
