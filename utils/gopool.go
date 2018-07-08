package gopool


import (
	"sync/atomic"
	"sync"
)
type GoPool struct{
	Locker *sync.Mutex
	Cond  *sync.Cond
	Max int32			//最大goroutine个数
	Running int32 		//正在运行goroutine个数 
}

func NewGoPool(max int) *GoPool{
 	L := new(sync.Mutex)
	return &GoPool{
		Locker:L,
		Cond:sync.NewCond(L),
		Max:int32(max),
		Running:0,
	}
}
//设置最大goroutine
func (gp *GoPool) Set(max int){
	gp.Max = int32(max)
	gp.Cond.Signal()
}

//新增一个goroutine
func (gp *GoPool) Incr(){
	gp.Locker.Lock()
	defer gp.Locker.Unlock()
	for{
		if gp.Max>gp.Running{
			break
		}
		gp.Cond.Wait()
	}
	atomic.AddInt32(&gp.Running, 1)

}
//结束一个goroutine
func (gp *GoPool)Dec(){
	atomic.AddInt32(&gp.Running, -1) 
	gp.Cond.Signal()
}
