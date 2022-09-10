package gotest

import "testing"

func TestCalc(t *testing.T) {
	got := Calc(1, 3)
	want := uint32(4)
	if got != want {
		t.Errorf("エラー got %d want %d", got, want)
	}
}

// ベンチマークテストを行う。
// 関数のシグネチャはfunc BenchmarkXxx(b *testing.B)
// テスト対象の関数をb.N回実行する（b.Nは十分な長さになるよう自動で調整される）
func BenchmarkCalc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calc(1, 2)
	}
}
