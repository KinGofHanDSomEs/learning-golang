package main



// Структура кеша
// Обязательно содержит 2 публичных поля 
// UpperBound - верхняя граница размера кеша
// LowerBound - нижняя граница размера кеша
// Если len > UpperBound, кеш автоматически вытеснит значения до нижней границы  
// Если любое из этих значений 0 - то этого не произойдет
type Cache struct {
    UpperBound int
    LowerBound int
}


// Создает инстанс кеша
func New() *Cache

// Проверяет, содержит, ли кэкш ключ
func (c *Cache) Has(key string) bool

// Возвращает значение по ключу, если оно существует
// Возвращает nil, если не существует
func (c *Cache) Get(key string) interface{} {
    c.lock.Lock()
    defer c.lock.Unlock()
    if e, ok := c.values[key]; ok {
        c.increment(e)
        return e.value
    }
    return nil
}

// Сохраняет значение по ключу
func (c *Cache) Set(key string, value interface{})

// Возвращает размер кеша
func (c *Cache) Len() int

// Возвращает частоту обращений к ключу
func (c *Cache) GetFrequency(key string) int

// Возвращает все ключи в кеше
func (c *Cache) Keys() []string

// Удаляет заданное количество наименее часто используемых элементов элементов
// Возвращает количество удаленных элементов
func (c *Cache) Evict(count int) int