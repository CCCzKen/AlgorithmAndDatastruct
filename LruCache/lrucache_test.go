package lrucache

import "testing"

func TestLruCache(t *testing.T) {
	lru := InitLruCache(5)

	lru.Set("name", "chenzhaokang")
	lru.Set("age", 25)
	lru.Set("sex", "male")
	lru.Set("address", "广州省")
	lru.Set("country", "中国")
	lru.Set("nba", "Golden State")

	length := lru.Size()
	if length != 5 {
		t.Errorf("lru.Set() failed. Got %d, expected 5.", length)
	}

	lru.Delete("age")
	length = lru.Size()
	if length != 4 {
		t.Errorf("lru.Delete() failed. Got %d, expected 2.", length)
	}

	value := lru.Get("nba")
	if value != "Golden State" {
		t.Errorf("lru.Get() failed. Got %s, expected Golden State.", value)
	}

}

