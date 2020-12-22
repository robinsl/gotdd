package maps

import "testing"

func TestDictionary_Search(t *testing.T) {
	dictionary := Dictionary{"test":  "this is test"}

	t.Run("Known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is test"

		assertStrings(t, got, want)
		assertNoError(t, err)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknow")
		if err == nil {
			t.Fatal("Expect an error to be fired")
		}

		assertError(t, err, ErrNotFound)

	})
}

func TestDictionary_Add(t *testing.T) {
	t.Run("Simple Add", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "Test"
		definition := "Test add"
		err := dictionary.Add(word, definition)

		assertNoError(t, err)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Existing word Add", func(t *testing.T) {
		word := "Test"
		definition := "Test add"
		definition2 := "Test add other definitions"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition2)

		assertError(t, err, ErrWordExist)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestDictionary_Update(t *testing.T) {
	t.Run("Update existing word", func(t *testing.T) {
		word := "Test"
		definition := "Test add"
		definition2 := "Test add other definitions"

		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, definition2)

		assertNoError(t, err)
		assertDefinition(t, dictionary, word, definition2)
	})

	t.Run("Update not existing word", func(t *testing.T) {
		word := "Test"
		definition := "Test add"

		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDictionary_Delete(t *testing.T) {
	t.Run("Delete existing word", func(t *testing.T) {
		word := "Test"
		definition := "Test add"

		dictionary := Dictionary{word: definition}

		err := dictionary.Delete(word)
		got, errSearch := dictionary.Search(word)

		assertStrings(t, got, "")
		assertNoError(t, err)
		assertError(t, errSearch, ErrNotFound)
	})

	t.Run("Delete not existing word", func(t *testing.T) {
		word := "Test"
		dictionary := Dictionary{}

		err := dictionary.Delete(word)

		assertError(t, err, ErrWordDoesNotExistDelete)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("Unexpected error")
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	assertStrings(t, got, definition)
	assertNoError(t, err)
}