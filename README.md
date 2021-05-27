[Test-Driven Development By Example](https://www.amazon.co.jp/-/en/Kent-Beck/dp/0321146530/ref=sr_1_1?dchild=1&keywords=test+driven+development+kent&qid=1622103566&sr=8-1) is written in Java.
But I wanted to TDD in Golang.
In this repository, I read [Test-Driven Development By Example](https://www.amazon.co.jp/-/en/Kent-Beck/dp/0321146530/ref=sr_1_1?dchild=1&keywords=test+driven+development+kent&qid=1622103566&sr=8-1), and translate Java sample codes to Golang.

### Some Points

- I use Go's built-in package, "testing".
- You can trace the process of TDD by commit message.
- You can run test codes in Go, like below.

```bash
go test --cover -v ./money
```

Kent Beckの「Test-Driven Development」はサンプルコードがJavaで書かれています。
GoでもTDDしたいということで、本リポジトリでは、Test-Driven DevelopmentのサンプルコードをGoで実装していきます。
Go標準装備のtestingパッケージを使います。
テストを書く->実装する、の流れをコミット単位で追えます。

テストの実行はリポジトリのルートディレクトリで、以下のコマンドを打ちます。

```bash
go test --cover -v ./money
```
