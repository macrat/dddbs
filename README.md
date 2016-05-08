dddbs: Dame Dame na DataBase System
===================================

golangで書かれたインメモリの全文検索エンジン(笑)。

メモリを猛烈に食い、規模が大きくなると多分壊滅的に遅くなる。
小規模な利用ならきっとそこそこ速い。

ユニグラムの転置データベースを使って絞り込んで、順次走査で更に絞り込む。

## 使い方
### 簡単な使い方
なんとなく以下のような感じで使います。

``` go
	import "dddbs"

	func main() {
		db := dddbs.NewDataBase()

		db.Add("key", "this is text")

		for _, x := range db.Search("text").Sort() {
			fmt.Println(x.Key, "=>", x.Score)
		}
	}
```

スコアの値は含まれているキーワードの数です。
**Sort**メソッドを使うことでこれによってソートされた値が得られます。

**Search**関数に渡すクエリはスペース区切りにするとアンド検索になります。
スペースもクエリとして検索したい場合は**SearchSingleQuery**が使えます。

### 複数の要素を検索する
dddbsはkey-value形式なので、キーと値しか入れることが出来ません。
また、検索対象になるのは値だけです。

どうしても複数の要素について検索したい場合、以下のようにすることで検索結果を結合することが出来ます。

``` go
	import "dddbs"

	func main() {
		dbA := dddbs.NewDataBase()
		dbB := dddbs.NewDataBase()

		dbA.Add("key", "this is text")
		dbB.Add("key", "hello world")

		resultA := dbA.Search("text")
		resultB := dbB.Search("hello")

		result := resultA.And(resultB)

		for _, x := range result.Sort() {
			fmt.Println(x.Key, "=>", x.Score)
		}
	}
```

**dbA**と**dbB**という二つのデータベースを作り、それらを個別に検索した結果を**And**メソッドで結合しています。
かなり効率が悪い気がしますが、ダメダメで元々なので諦めてください。

## ライセンス
[MIT License](https://opensource.org/licenses/MIT)です。
