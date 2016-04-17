package dddbs

import (
	"testing"
)

func TestSearch(t *testing.T) {
	db := NewDataBase()

	db.Add("0 hoge", "hello world")
	db.Add("1 fuga", "welcome to world world")

	r := db.Search("world").Sort()
	if len(r) != 2 {
		t.Error("Invalid number of entries in result of \"world\"")
	}

	if len(r) >= 1 && r[0].Key != "1 fuga" {
		t.Error("First entry expected fuga but got " + r[0].Key)
	}

	if len(r) >= 2 && r[1].Key != "0 hoge" {
		t.Error("Second entry expected hoge but got " + r[1].Key)
	}

	r = db.Search("hello").Sort()
	if len(r) != 1 {
		t.Error("Invalid number of entries in result of \"hello\"")
	}

	if len(r) >= 1 && r[0].Key != "0 hoge" {
		t.Error("The entry expected hoge but got " + r[0].Key)
	}
}

func TestJapaneseSearch(t *testing.T) {
	db := NewDataBase()

	db.Add("0 日本語", "これは日本語です")
	db.Add("1 テスト", "こいつは日本語だ")

	r := db.Search("日本語").Sort()
	if len(r) != 2 {
		t.Error("Invalid number of entries in result of \"日本語\"")
	}

	if len(r) >= 1 && r[0].Key != "0 日本語" {
		t.Error("First entry expected 日本語 but got " + r[0].Key)
	}

	if len(r) >= 2 && r[1].Key != "1 テスト" {
		t.Error("Second entry expected テスト but got " + r[1].Key)
	}
}
