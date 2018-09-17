package main

import(
	"fmt"
	"sync"
	"strconv"
)

type HashTable struct {
	items map[int]string
	lock sync.RWMutex
}

func hash(k string) int {
	key := fmt.Sprintf("%s", k)
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
	}
	return h
}

func (ht *HashTable) Add(k string, v string) {
	ht.lock.Lock()
	defer ht.lock.Unlock()

	i := hash(k)
	if ht.items == nil {
		ht.items = make(map[int]string)
	}

	ht.items[i] = v
}

func (ht *HashTable) Remove(k string) {
	ht.lock.Lock()
	defer ht.lock.Unlock()

	i := hash(k)
	_,  exists := ht.items[i]
	if exists {
		delete(ht.items, i)
	}
}

func (ht *HashTable) Get(k string) string {
	ht.lock.Lock()
	defer ht.lock.Unlock()

	i := hash(k)
	v, exists := ht.items[i]
	if exists {
		return v
	}
	return ""
}

func (ht *HashTable) Dump() {
	fmt.Println("hash\tbuckets\n")
	for k, v := range ht.items {
		fmt.Printf("%d\t%s\n", k, v)
	}
}

// **********************************
var r int = 1000

// MAIN.
func main() {
	input := make(map[string]string)
	for i := 0; i < r; i++ {
		input["key"+strconv.Itoa(i)] = "value"+strconv.Itoa(i)
	}
	
	hashtable := HashTable{}
	for k, v := range input {
		hashtable.Add(k, v)
	}

	hashtable.Dump()
}
