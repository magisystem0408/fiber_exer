## Go Fiberを学んでみた。(webフレームワーク)

### 使用したソース
- ホットリロードかけられるようにする
  - https://github.com/oxequa/realize
    - 簡単な使い方
      - realize startと出るとホットリロードが開始される
      - yamlファイルができるのでcommandsにrun:{status:true}をしてあげる
    - エラー報告
      - パッケージインストール時にエラー報告以下と出てくる
  >    go get: gopkg.in/urfave/cli.v2@none updating to
      gopkg.in/urfave/cli.v2@v2.3.0: parsing go.mod:
      module declares its path as: github.com/urfave/cli/v2
      but was required as: gopkg.in/urfave/cli.v2

  - 依存関係に問題があるらしく、GO111MODULE=off go get github.com/oxequa/realizeとすれば良い
    - (参考)https://github.com/keitakn/golang-grpc-server/issues/23
  - realize startでnot startedになったが以下で治った
    - http://psychedelicnekopunch.com/archives/1723

- modelのマイグレーション
  - https://gorm.io/ja_JP/docs/models.html

- パスワードをハッシュ化させる
  - https://pkg.go.dev/golang.org/x/crypto/bcrypt