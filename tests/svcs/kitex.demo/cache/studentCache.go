package cache

import (
	"container/list"
	"sync"

	serverz "day3/kxS/kitex_gen/kitex/serverZ"
)

type Cache struct {
	capacity int
	cache    map[int]*list.Element
	queue    *list.List
	mutex    sync.Mutex
}

func NewCache(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		queue:    list.New(),
	}
}

func (c *Cache) Get(key int) (*serverz.Student, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, ok := c.cache[key]; ok {
		c.queue.MoveToFront(elem)
		return elem.Value.(*serverz.Student), true
	}

	return nil, false
}

func (c *Cache) Put(key int, student *serverz.Student) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, ok := c.cache[key]; ok {
		elem.Value = student
		c.queue.MoveToFront(elem)
	} else {
		if len(c.cache) >= c.capacity {
			// FIFO eviction, remove the least recently used element
			backElem := c.queue.Back()
			if backElem != nil {
				deletedStudent := c.queue.Remove(backElem).(*serverz.Student)
				delete(c.cache, int(deletedStudent.Id))
			}
		}
		elem := c.queue.PushFront(student)
		c.cache[key] = elem
	}
}

// func main() {
// 	cache := NewCache(1000)

// 	// Insert some sample data into the cache
// 	for i := 1; i <= 1500; i++ {
// 		student := &Student{
// 			ID:   i,
// 			Name: fmt.Sprintf("Student-%d", i),
// 		}
// 		cache.Put(i, student)
// 	}

// 	// Try to get some data from the cache
// 	for i := 1; i <= 10; i++ {
// 		if student, ok := cache.Get(i); ok {
// 			fmt.Printf("ID: %d, Name: %s\n", student.ID, student.Name)
// 		} else {
// 			fmt.Println("ID not found in the cache:", i)
// 		}
// 	}
// }
