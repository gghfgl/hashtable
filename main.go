package main

import(
	"fmt"
	"sync"
	"strconv"
	"os"
	"text/tabwriter"
)

type HashTable struct {
	items map[int]map[string]string
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
		ht.items = make(map[int]map[string]string)
	}

	ht.items[i] = map[string]string{k: v}
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

func (ht *HashTable) Get(k string) map[string]string {
	ht.lock.Lock()
	defer ht.lock.Unlock()

	i := hash(k)
	v, exists := ht.items[i]
	if exists {
		return v
	}
	return nil
}

func (ht *HashTable) Dump() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "hash*\tbuckets*\t")
	for k, v := range ht.items {
		fmt.Fprintln(w, k, "\t", v, "\t")
	}
	fmt.Fprintln(w)
	w.Flush()
}

// **********************************
var r int = 10

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
