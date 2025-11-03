package cache

import "sync"

// Interface - реализуйте этот интерфейс
type Interface interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

// Не меняйте названия структуры и название метода создания экземпляра Cache, иначе не будут проходить тесты

type Cache struct {
	m   map[string]string
	rmu sync.RWMutex
}

func NewCache() Interface {
	return &Cache{
		m: make(map[string]string),
	}
}

func (c *Cache) Set(k, v string) {
	c.rmu.Lock()
	defer c.rmu.Unlock()
	c.m[k] = v
}

func (c *Cache) Get(k string) (string, bool) {
	c.rmu.RLock()
	defer c.rmu.RUnlock()
	v, ok := c.m[k]
	return v, ok
}

/*

Почему использовать RWMutex, а не Mutex?
	- RWMutex позволяет нескольким горутинам одновременно читать кеш (RLock()), пока никто не пишет.
	- Mutex блокирует и чтение, и запись, поэтому снижает параллелизм.
	- Для кеша, где чтений обычно больше, чем записей, RWMutex ускоряет работу.

Если память не бесконечна
Проблемы:
	- Кеш может расти бесконтрольно → использование всей памяти → падение программы.

Решения:
	- Удалять старые или редко используемые элементы.

Алгоритмы выселения:
	- LRU (Least Recently Used) — удаляем наименее недавно использованные.
	- FIFO (First In, First Out) — удаляем самые старые по добавлению.
	- LFU (Least Frequently Used) — удаляем самые редко используемые.
*/
