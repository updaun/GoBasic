// 최종 실습 예제
// 대상 사이트 : 루리웹 (ruliweb.com)

package main

import (
	_ "bufio"
	"fmt"
	_ "fmt"
	"net/http"
	_ "os"
	_ "strings"
	"sync"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// 스크래핑 대상 URL
const (
	urlRoot = "http://ruliweb.com"
)

// 첫 번째 방문(메인페이지) 대상으로 원하는 url을 파싱 후 반환하는 함수
func parseMainNodes(n *html.Node) bool {
	if n.DataAtom == atom.A && n.Parent != nil {
		return scrape.Attr(n.Parent, "class") == "row"
	}
	return false
}

// 에러 체크 공통 함수
func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// 동기화를 위한 작업 그룹 선언
var wg sync.WaitGroup

func main() {
	// 메인 페이지 Get 방식 요청
	resp, err := http.Get(urlRoot)
	errCheck(err)

	// 요청 Body 닫기
	defer resp.Body.Close()

	// 응답 데이터(HTML)
	root, err := html.Parse(resp.Body)
	errCheck(err)

	// ParseMainNodes 메소드를 크로링(스크래핑) 대상 URL 추출
	urlList := scrape.FindAll(root, parseMainNodes)
	fmt.Println("Extract Data Length : ", len(urlList))
}
