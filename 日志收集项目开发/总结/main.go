package 总结

//context的五个注意事项:
//推荐以参数的方式显示传递Context
//以Context作为参数的函数方法，应该把Context作为第一个参数。
//给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO()
//Context的Value相关方法应该传递请求域的必要数据，不应该用于传递可选参数
//Context是线程安全的，可以放心的在多个goroutine中传递
//func main() {
//	context.WithCancel()   //调用cancel就回给ctx.Done()发信号
//	context.WithDeadline() //超时了也要调用cancel()
//	context.WithTimeout()  //超时了也要调用cancel()
//	context.WithValue()    //注意 key 要使用自定义类型
//}
