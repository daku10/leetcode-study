package main

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	t.Run("LRU Cache 1", func(t *testing.T) {
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

	t.Run("LRU Cache 2", func(t *testing.T) {
		cache := Constructor(2)
		cache.Get(2)
		cache.Put(2, 6)
		cache.Get(1)
		cache.Put(1, 5)
		cache.Put(1, 2)
		if v := cache.Get(1); v != 2 {
			t.Fatalf("actual: %v want: %v", v, 2)
		}
		if v := cache.Get(2); v != 6 {
			t.Errorf("actual: %v want: %v", v, 6)
		}
		// cache.Put(1, 1)
		// cache.Put(2, 2)
		// v := cache.Get(1)
		// if v != 1 {
		// 	t.Fatalf("actual: %v want: %v", v, 1)
		// }
		// cache.Put(3, 3)
		// v = cache.Get(2)
		// if v != -1 {
		// 	t.Fatalf("actual: %v want: %v", v, -1)
		// }
		// cache.Put(4, 4)
		// v = cache.Get(1)
		// if v != -1 {
		// 	t.Fatalf("actual: %v want: %v", v, -1)
		// }
		// v = cache.Get(3)
		// if v != 3 {
		// 	t.Fatalf("actual: %v want: %v", v, 3)
		// }
		// v = cache.Get(4)
		// if v != 4 {
		// 	t.Errorf("actual: %v want: %v", v, 4)
		// }
	})

	t.Run("LRU Cache 3", func(t *testing.T) {
		// 		["LRUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]
		// [[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]
		cache := Constructor(10)
		cache.Put(10, 13)
		cache.Put(3, 17)
		cache.Put(6, 11)
		cache.Put(10, 5)

		cache.Put(9, 10)
		cache.Get(13)
		cache.Put(2, 19)
		cache.Get(2)
		cache.Get(3)
		cache.Put(5, 25)
		cache.Get(8)
		cache.Put(9, 22)
		cache.Put(5, 5)
		cache.Put(1, 30)
		cache.Get(11)
		cache.Put(9, 12)
		cache.Get(7)
		cache.Get(5)
		cache.Get(8)
		cache.Get(9)
		cache.Put(4, 30)
		cache.Put(9, 3)
		cache.Get(9)
		cache.Get(10)
		cache.Get(10)
		cache.Put(6, 14)
		cache.Put(3, 1)
		cache.Get(3)
		cache.Put(10, 11)
		cache.Get(8)
		cache.Put(2, 14)
		cache.Get(1)
		cache.Get(5)
		cache.Get(4)
		cache.Put(11, 4)
		cache.Put(12, 24)
		cache.Put(5, 18)
		cache.Get(13)
		cache.Put(7, 23)
		cache.Get(8)
		cache.Get(12)
		cache.Put(3, 27)
		cache.Put(2, 12)
		cache.Get(5)
		cache.Put(2, 9)
		cache.Put(13, 4)
		cache.Put(8, 18)
		cache.Put(1, 7)
		cache.Get(6)
		cache.Put(9, 29)
		cache.Put(8, 21)
		cache.Get(5)
		cache.Put(6, 30)
		cache.Put(1, 12)
		cache.Get(10)
		cache.Put(4, 15)
		cache.Put(7, 22)
		cache.Put(11, 26)
		cache.Put(8, 17)
		cache.Put(9, 29)
		cache.Get(5)
		cache.Put(3, 4)
		cache.Put(11, 30)
	})

}
