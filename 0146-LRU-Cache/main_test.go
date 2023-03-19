package main

import "testing"

func TestLRUCache(t *testing.T) {
	t.Run("LRU Cache", func(t *testing.T) {
		cache := Constructor(2)
		cache.Put(1, 1)
		cache.Put(2, 2)
		v := cache.Get(1)
		if v != 1 {
			t.Fatalf("actual: %v want: %v", v, 1)
		}
		cache.Put(3, 3)
		v = cache.Get(2)
		if v != -1 {
			t.Fatalf("actual: %v want: %v", v, -1)
		}
		cache.Put(4, 4)
		v = cache.Get(1)
		if v != -1 {
			t.Fatalf("actual: %v want: %v", v, -1)
		}
		v = cache.Get(3)
		if v != 3 {
			t.Fatalf("actual: %v want: %v", v, 3)
		}
		v = cache.Get(4)
		if v != 4 {
			t.Errorf("actual: %v want: %v", v, 4)
		}
	})
}
