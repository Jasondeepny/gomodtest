package main

//
//// 处理函数
////func process(conn net.Conn) {
////	defer conn.Close() // 关闭连接
////	for {
////		reader := bufio.NewReader(conn)
////		var buf [128]byte
////		n, err := reader.Read(buf[:]) // 读取数据
////		if err != nil {
////			fmt.Println("read from client failed, err:", err)
////			break
////		}
////		recvStr := string(buf[:n])
////		fmt.Println("收到client端发来的数据：", recvStr)
////		conn.Write([]byte(recvStr)) // 发送数据
////	}
////}
////
////func main() {
////	listen, err := net.Listen("tcp", "127.0.0.1:20000")
////	if err != nil {
////		fmt.Println("listen failed, err:", err)
////		return
////	}
////	for {
////		conn, err := listen.Accept() // 建立连接
////		if err != nil {
////			fmt.Println("accept failed, err:", err)
////			continue
////		}
////		go process(conn) // 启动一个goroutine处理连接
////	}
////}
//
//func main() {
//	listener, err := net.Listen("tcp","127.0.0.1:20000")
//	if err != nil {
//		return
//	}
//	for  {
//		conn,err := listener.Accept()
//		if err != nil{
//			fmt.Println("listener is error",err)
//			continue
//		}
//		if conn != nil {
//			fmt.Println("listener is work")
//		}
//		go process(conn)
//	}
//}
//
//// TCP server端
//func process(con net.Conn) {
//	defer func(con net.Conn) {
//		err := con.Close()
//		if err != nil {
//			return
//		}
//	}(con)
//
//	for {
//		read := bufio.NewReader(con)
//		var buff [1024]byte
//		n, err := read.Read(buff[:])
//		if err != nil {
//			fmt.Printf("error is %s", err)
//			break
//		}
//		reciver := string(buff[:n])
//		fmt.Println("收到client端发来的数据：", reciver)
//		con.Write([]byte(reciver))
//	}
//}
