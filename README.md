Kent Beckの「Test-Driven Development」はサンプルコードがJavaで書かれています。
GoでもTDDしたいということで、本リポジトリでは、Test-Driven DevelopmentのサンプルコードをGoで実装していきます。
Go標準装備のtestingパッケージを使います。
テストを書く->実装する、の流れをコミット単位で追えます。

テストの実行はリポジトリのルートディレクトリで、以下のコマンドを打ちます。

```bash
go test --cover -v money
```
