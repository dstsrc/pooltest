package pooltest

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type A struct {
	Str string
	Int int
}

type PoolStorage struct {
	sp *sync.Pool
}

func New() *PoolStorage {
	sp := &sync.Pool{
		New: func() interface{} { return &A{} },
	}
	return &PoolStorage{
		sp: sp,
	}
}

func (p *PoolStorage) Put(s *A) {
	p.sp.Put(s)
	return
}

func (p *PoolStorage) Get() *A {
	s := p.sp.Get().(*A)
	return s
}

func (a *A) reset() {
	a.Str = ""
	a.Int = 0
}

func (p *PoolStorage) sc1() {
	a := &A{}
	a.Str = "test"
	p.Put(a)

	time.Sleep(time.Millisecond) // for race detector

	newA := p.Get()
	newA.Str = "change"

	if a.Str == "change" {
		log.Println("str field changed")
	}
}

func (p *PoolStorage) sc2() {
	a := p.Get()
	r := rand.Int()
	a.Int = r
	p.Put(a)

	time.Sleep(time.Millisecond) // for race detector

	if a.Int != r {
		log.Println("int field changed")
	}
}

func (p *PoolStorage) sc3() {
	a := p.Get()
	if a.Str == "old" {
		log.Println("no zero value") // is not problem
	}

	time.Sleep(time.Millisecond)

	a.Str = "old"
	p.Put(a)
}
