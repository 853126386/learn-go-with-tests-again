//更多的 TDD 练习
//学习了 for 循环
//学习了如何编写基准测试

package iteration

import "testing"
func TestIteration(t *testing.T)  {
	repeated :=Repeat("a")
	expected :="aaa"

	if repeated != expected{
		t.Errorf("expected '%q' but got '%q'",expected,repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i:=0;i<b.N;i++ {
		Repeat("a")
	}
}