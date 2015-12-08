package servant

import (
	"appengine"
	"appengine/urlfetch"
	"encoding/base64"
	"encoding/binary"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

func ToEng(v string) string {
	switch {
	case "黒" == v:
		return "black"
	case "青" == v:
		return "blue"
	case "無" == v:
		return "colorless"
	case "赤" == v:
		return "red"
	case "緑" == v:
		return "green"
	case "白" == v:
		return "white"
	}
	return ""
}

func ToInt(v string) int {
	if v == "-" {
		return -1
	}
	i, _ := strconv.Atoi(v)
	return i
}

func TrimHyphen(v string) string {
	if v == "-" {
		return ""
	}
	return v
}

func TrimLinefeed(v string) string {
	return strings.TrimRight(strings.TrimLeft(strings.TrimSpace(v), "\n"), "\n")
}

func Contains(str string, list []string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}
	return false
}

/**
 * Getリクエストを送る
 * @function
 * @param {appengine.Context} c コンテキスト
 * @param {string} targetUrl 送信先
 * @param {map[string]string} query クエリリスト
 * @param {map[string]string} header ヘッダー
 * @returns {*http.Response} レスポンス
 */
func Get(c appengine.Context, targetUrl string, query map[string]string, header map[string]string) *http.Response {
	var request *http.Request
	var err error

	// クエリ埋め込み
	if query != nil || len(query) > 0 {
		paramStrings := make([]string, 0)
		for key, value := range query {
			param := strings.Join([]string{key, value}, "=")
			paramStrings = append(paramStrings, param)
		}
		paramString := ""
		if len(query) == 1 {
			paramString = paramStrings[0]
		} else {
			paramString = strings.Join(paramStrings, "&")
		}
		targetUrl = strings.Join([]string{targetUrl, paramString}, "?")
	}

	// リクエスト作成
	request, err = http.NewRequest("GET", targetUrl, nil)
	Check(c, err)

	// ヘッダー設定
	if header != nil || len(header) > 0 {
		for key, value := range header {
			request.Header.Add(key, value)
		}
	}

	// 送受信
	client := urlfetch.Client(c)
	response, err := client.Do(request)
	Check(c, err)

	return response
}

/**
 * HTTP リクエストを送信してレスポンスを返す
 * @function
 * @param {appengine.Context} c コンテキスト
 * @param {string} method POST または GET
 * @param {string} targetUrl 送信先のURL
 * @param {map[string]string} params パラーメタリスト 指定しない場合は nil または空マップ
 * @param {string} body リクエストボディ GET の場合は無視される
 * @param {*http.Response} レスポンス
 */
func Request(c appengine.Context, method string, targetUrl string, params map[string]string, body string) *http.Response {
	var request *http.Request
	var err error

	// methodのチェック
	if method != "GET" && method != "POST" {
		c.Infof("request(): method must set GET or POST only.")
		return nil
	}

	// GET なら URL にクエリ埋め込み
	if method == "GET" && (params != nil || len(params) > 0) {
		paramStrings := make([]string, 0)
		for key, value := range params {
			param := strings.Join([]string{key, value}, "=")
			paramStrings = append(paramStrings, param)
		}
		paramString := ""
		if len(params) == 1 {
			paramString = paramStrings[0]
		} else {
			paramString = strings.Join(paramStrings, "&")
		}
		targetUrl = strings.Join([]string{targetUrl, paramString}, "?")
	}

	// リクエスト作成
	if method == "GET" || body == "" {
		request, err = http.NewRequest(method, targetUrl, nil)
	} else {
		request, err = http.NewRequest(method, targetUrl, NewReader(body))
	}
	Check(c, err)

	// POST なら Header にパラメータ設定
	if method == "POST" && (params != nil || len(params) > 0) {
		for key, value := range params {
			request.Header.Add(key, value)
		}
	}

	// 送受信
	client := urlfetch.Client(c)
	response, err := client.Do(request)
	Check(c, err)

	return response
}

/**
 * エラーチェック
 * エラーがあればコンソールに出力する
 * @function
 * @param {appengine.Context} c コンテキスト
 * @param {error} err チェックするエラーオブジェクト
 */
func Check(c appengine.Context, err error) {
	if err != nil {
		c.Errorf(err.Error())
	}
}

/**
 * リクエストボディ用のリーダー
 * request() で body を送信するために使う
 * @class
 * @member {[]byte} body 本文
 * @member {int} pointer 何バイト目まで読み込んだか表すポインタ
 */
type Reader struct {
	io.Reader
	body    []byte
	pointer int
}

/**
 * Reader のインスタンスを作成する
 * @param {string} body 本文
 * @returns {*Reader} 作成したインスタンス
 */
func NewReader(body string) *Reader {
	reader := new(Reader)
	reader.body = []byte(body)
	reader.pointer = 0
	return reader
}

/**
 * ランダムな文字列を取得する
 * 64bit のランダムデータを Base64 エンコードして記号を抜いたもの
 * @function
 * @returns {string} ランダムな文字列
 */
func GetRandomizedString() string {
	r := rand.Int63()
	b := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(b, int64(r))
	e := base64.StdEncoding.EncodeToString(b)
	e = strings.Replace(e, "+", "", -1)
	e = strings.Replace(e, "/", "", -1)
	e = strings.Replace(e, "=", "", -1)
	return e
}
