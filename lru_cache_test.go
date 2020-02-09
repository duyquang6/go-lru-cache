package lru

import "testing"

func TestLRUCache_Get(t *testing.T) {
	cache := Constructor(2)
	cache.Put("1", 1)
	cache.Put("2", 2)
	cache.Put("3", 2)
	type args struct {
		key string
	}
	tests := []struct {
		name string
		this *LRUCache
		args args
		want int
	}{
		{
			name: "TC01_GetEvictedKey",
			this: &cache,
			args: args{key: "1"},
			want: -1,
		},
		{
			name: "TC02_GetGoodKey",
			this: &cache,
			args: args{key: "2"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Get(tt.args.key); got != tt.want {
				t.Errorf("LRUCache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
