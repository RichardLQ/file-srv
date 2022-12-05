package auth

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

//SendHttpRequest 发送HTTP请求
func SendHttpRequest(url, method, body string, cookies []http.Cookie, headers []map[string]string) (string, error) {
	if method == "" {
		method = "GET"
	}
	var client = &http.Client{
		Transport: &http.Transport{
			Proxy: nil,
			DialContext: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2: true,
		},
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       5 * time.Second,
	}
	var req *http.Request

	//实例化一个request对象
	req, _ = http.NewRequest(method, url, strings.NewReader(body))

	//设置Cookie
	if nil != cookies && len(cookies) == 0 {
		for i := range cookies {
			req.AddCookie(&cookies[i])
		}
	}

	//设置Header
	if nil != headers && len(headers) > 0 {
		for i := range headers {
			for k, v := range headers[i] {
				req.Header.Add(k, v)
			}
		}
	}

	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer req.Body.Close()
	buf := bytes.Buffer{}
	buf.ReadFrom(resp.Body)
	return buf.String(), nil
}

//RedisConn redis的连接
func RedisConn() redis.Conn {
	c, err := redis.Dial("tcp", "localhost:2000")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return nil
	}
	return c
}

//FileMD5 文件md5名称
func FileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func FileOpenMD5(files *multipart.FileHeader) (string, error) {
	//file, err := os.Open(filePath)
	//if err != nil {
	//	return "", err
	//}
	file, err := files.Open()
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}