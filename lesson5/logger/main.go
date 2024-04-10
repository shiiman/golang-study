package main

import (
	"log"
	"os"
)

func main() {
	// フォーマットの変更
	log.SetFlags(log.Llongfile | log.LUTC | log.LstdFlags)
	log.SetPrefix("[my-app] ")
	log.Println("Hello", "世界")
	log.Print("Hello", "地球")

	// ロガーのインスタンスを生成
	myLogger := log.New(
		os.Stdout, "[my-app] ",
		log.Ldate|log.Lmicroseconds|log.Lshortfile|log.LUTC|log.Lmsgprefix,
	)
	myLogger.Println("Hey", "世界")
	myLogger.Print("Hey", "地球")
}
