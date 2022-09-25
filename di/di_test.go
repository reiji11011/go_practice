package di

import "testing"

func TestIsMorning(t *testing.T) {
	// 実行する時間帯によってテストの結果が変わる。
	got := true
	want := isMorning()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
