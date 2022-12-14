package maps

import "testing"

func TestSearch(t *testing.T)  {
	 dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got,_ :=dictionary.Search("test")
		want :="this is just a test"
		asserString(t,got,want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_,err :=dictionary.Search("test")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error.")
		}
		asserString(t,err.Error(),want)
	})


}
func asserString(t *testing.T,got,want string)  {

	t.Helper()
	if got !=want{
		t.Errorf("got '%s' want '%s'", got, want)
	}
}