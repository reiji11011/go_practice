package di

import (
	"bytes"
	"fmt"
)

// 参考：スタブ・モック - テスト駆動開発でGO言語を学びましょう
// https://andmorefine.gitbook.io/learn-go-with-tests/go-fundamentals/mocking

func TddDi() {
	buffer := bytes.Buffer{}
	Greet(&buffer, "太郎")
}

func CountDown() {

}

// 以下のテストを書く。
// fmt.Printf()を実行するとstdoutに出力される。
// 依存関係を注入できるようにする。

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
