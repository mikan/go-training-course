package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mikan/go-training-course/ch05/ex13/links"
)

var tokens = make(chan struct{}, 20)

func crawl(url link) []link {
	fmt.Println(url)
	tokens <- struct{}{} // トークンを獲得
	list, err := links.Extract(url.url)
	<-tokens // トークンを開放
	if err != nil {
		log.Print(err)
	}
	var listWithDepth []link
	for _, l := range list {
		listWithDepth = append(listWithDepth, link{l, url.depth + 1})
	}
	return listWithDepth
}

type link struct {
	url   string
	depth int
}

func (l link) String() string {
	return fmt.Sprintf("depth=%d url=%s", l.depth, l.url)
}

func main() {
	depth := flag.Int("depth", 0, "深さ制限, 0で無制限")
	flag.Parse()
	workList := make(chan []link)  // URL のリスト、重複を含む
	unseenLinks := make(chan link) // 重複していない URL

	// コマンドライン引数を workList へ追加する
	go func() {
		var list []link
		for _, arg := range os.Args[1:] {
			if strings.HasPrefix(arg, "http") {
				list = append(list, link{arg, 1})
			}
		}
		workList <- list
	}()

	// 未探索のリンクを取得するために 20 個のクローラのゴルーチンを生成する。
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				if *depth > 0 && link.depth > *depth {
					continue
				}
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
			if !seen[link.url] {
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
}
