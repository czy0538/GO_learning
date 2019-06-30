package main

//首先，实现并执行一个匿名的超时等待函数
timeout:=make(chan bool,1)
go func(){
	time.Sleep(1*time.Second)//等待1s钟
	timeout=true
}()
//然后，把这个timeout利用起来
select{
	case <-ch 
	//从ch中读到数据
	case <-timeout
	//一直没有从ch中读到数据，但从timeout中读到了数据
}
