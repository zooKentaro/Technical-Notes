package testgo

import "time"

// 処理利用ファイルにて

// パッケージグローバル変数
var timeNow = time.Now

func TestMethod() {
	// time.Now() は利用せず
	timeNow()         // を利用
	timeNow().Month() // といったtime.Timeの関数利用も可能
}

// ~~~~~
// 利用先テストファイルにて

// テスト用に時刻を固定
var timeTest = time.Date(2023, 1, 2, 3, 4, 5, 6, time.UTC)

func init() {
	timeNow = func() time.Time {
		return timeTest
	}
}

// ============================================================================
// dockerコンテナ内でのログを確認する場合のコマンド
// docker compose logs -f

// 特定のコンテナのみ閲覧したい場合は以下
// docker compose logs コンテナ名 -f

// ============================================================================
// ポインタ型の変数を作らずに直接充てる方法
type Test struct {
	AA *string
}

// 通常
var aaa string = "test"
var test Test = Test{
	AA: &aaa,
}

// テクニック
var testB Test = Test{
	AA: &[]string{"test"}[0],
}
