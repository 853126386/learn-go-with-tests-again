package main

import "testing"

func TestHello(t *testing.T) {
	//test1(t)
	//test2(t)
	//test3(t)
	test4(t)
}

//基础测试----版本1
func test1(t *testing.T)  {
	got :=Hello("Soleil")
	want :="Hello,Soleil"
	if got !=want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}

//子测试t.Run----版本1
func test2(t *testing.T)  {
	t.Run("saying hello to people", func(t *testing.T) {
		got :=Hello("Soleil")
		want :="Hello,Soleil"
		if got !=want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got :=Hello("")
		want :="Hello,world"
		if got !=want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})
}

//重构子测试t.Run，使用闭包函数重构重复的代码----版本1
func test3(t *testing.T)  {
	assertCorrectMessage := func(t *testing.T,got,want string) {
		//t.Helper() 需要告诉测试套件这个方法是辅助函数（helper）。通过这样做，当测试失败时所报告的行号将在函数调用中而不是在辅助函数内部。
		//这将帮助其他开发人员更容易地跟踪问题。如果你仍然不理解，请注释掉它，使测试失败并观察测试输出。
		t.Helper()
		if got != want {
			t.Errorf("got '%q',want '%q'",got,want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got :=Hello("Soleil")
		want :="Hello,Soleil"
		assertCorrectMessage(t,got,want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got :=Hello("")
		want :="Hello,world"
		assertCorrectMessage(t,got,want)
	})
}

//继续迭代----版本2
func test4(t *testing.T)  {
	assertCorrectMessage := func(t *testing.T,got,want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q',want '%q'",got,want)
		}
	}
	t.Run("in Spanish", func(t *testing.T) {
		got :=Hello2("Elodie","Spanish")
		want :="Halo,Elodie"
		assertCorrectMessage(t,got,want)
	})

	t.Run("in French", func(t *testing.T) {
		got :=Hello2("Elodie","French")
		want :="Bonjour,Elodie"
		assertCorrectMessage(t,got,want)
	})
}



