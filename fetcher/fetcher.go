package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var rateLimiter = time.Tick(30 * time.Millisecond)

// 获取网页信息
// input: url
// output:
func Fetch(url string) ([]byte, error) {
	// 流量限制：执行到这里，需要隔10毫秒才继续往下执行
	<-rateLimiter
	//resp, err := http.Get(url)
	// 直接用http.Get(url)进行获取信息会报错：Error: status code 403
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	// 查看自己浏览器中的User-Agent信息（检查元素->Network->User-Agent）
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	// 自动识别网页html编码，并转换为utf-8
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	// Peek 返回缓存的一个切片，该切片引用缓存中前 n 字节数据，
	// 该操作不会将数据读出，只是引用，引用的数据在下一次读取操作之前是有效的
	// 如果引用的数据长度小于 n，则返回一个错误信息；如果 n 大于缓存的总大小，则返回 ErrBufferFull
	// 通过 Peek 的返回值，可以修改缓存中的数据，但是不能修改底层 io.Reader 中的数据
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
