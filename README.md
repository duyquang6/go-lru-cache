# Golang LRU Cache

Simple implementation using container/list and map 

## Example 
```
import github.com/duyquang6/go-lru-cache

func example() {
	cache := lru.Constructor(2)     // initialize new cache with cap = 2
	cache.Put("one", 1)   
	cache.Put("two", 2)
	log.Println(cache.Get("one"))   // return 1
	cache.Put("three", 3)
	log.Println(cache.Get("two"))   // return 2
	cache.Put("four", 4)
	log.Println(cache.Get(1))       // return -1
	log.Println(cache.Get(3))
	log.Println(cache.Get(4))
}
```