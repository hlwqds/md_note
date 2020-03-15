package main

type Func func(key string) (interface{}, error)

type result struct{
	value struct{}
	err error
}
type entry struct{
	res result
	ready chan struct{}
}

type request struct{
	key string
	resp chan<-result
}

type Memo struct{
	requests chan request
}

//创建某个函数的函数记忆
func New(f Func)*Memo{
	me := &Memo{
		requests: make(chan request),
	}
	go me.server(f)
	return &me
}

func (memo *Memo) Get(s string)(interface{}, error){
	req := request{
		key: s,
		resp: make(chan result),
	}
	memo.requests <- req
	return <-req.resp
}
func (memo *Memo) Close(){
	close(memo.requests)
}

func (me *Memo)server(f Func){
	ha := make(map[string]*entry)
	for req := range me.requests{
		e, ok := ha[req.key]
		if e == nil{
			e = &entry{
				ready: make(chan struct{}),
			}
			ha[req.key] = e
			e.call(f, req.key)
		}
		e.deliver(req.resp)
	}
}
func (e *entry) call(f Func, s string){
	e.res.value, e.res.err := f(s)
	close(e.ready)
}

func (e *entry) deliver(resp chan<-result){
	<- e.ready
	resp <- e.res
	close(resp)
}
