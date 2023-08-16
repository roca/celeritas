package cache

type BadgerCache struct{}

func (b *BadgerCache) Has(str string) (bool, error)
func (b *BadgerCache) Get(str string) (any, error)
func (b *BadgerCache) Set(str string, value any, expires ...int) error
func (b *BadgerCache) EmptyByMatch(str string) error
func (b *BadgerCache) Empty() error

func (b *BadgerCache) makeKey(str string) string
func (b *BadgerCache) getKeys(pattern string) ([]string, error)
