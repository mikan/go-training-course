package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mikan/go-training-course/ch05/ex13/links"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // トークンを獲得
	list, err := links.Extract(url)
	<-tokens // トークンを開放
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	workList := make(chan []string)  // URL のリスト、重複を含む
	unseenLinks := make(chan string) // 重複していない URL

	// コマンドライン引数を workList へ追加する
	go func() {
		workList <- os.Args[1:]
	}()

	// 未探索のリンクを取得するために 20 個のクローラのゴルーチンを生成する。
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {
					workList <- foundLinks
				}()
			}
		}()
	}

	// メインゴルーチンは workList の項目の重複をなくし、未探索の項目をクローラへ送る。
	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
