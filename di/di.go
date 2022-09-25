package di

import (
	"fmt"
	"time"
)

func Main() {
	if isMorning() {
		fmt.Println("午前中です！")
	}
	fmt.Println("午後です！")
}

// 現在時刻を取得し、0~11なら午前中、12~23なら午後、その他ならエラーを返す関数を作成する。
func isMorning() bool {
	now := time.Now()
	hour := now.Hour()
	return 0 <= hour && hour <= 11
}

// 現状の問題点
// 外部のtimeパッケージを利用している（= 依存している）
// time.Now()の戻り値にisMorning()が依存しているため、テストが書きづらい。
// 午前中であることのテストと午後であることのテストを同時に書くことができない。
