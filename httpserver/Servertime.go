package httpserver

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func Time(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("你好，王召！"))
	//服务器的时间，年月日，时分秒 ，格式化时间
	now := time.Now()
	newNow := now.Format("2006-01-02 15:04:05")
	//获取id+端口
	addr := r.RemoteAddr
	//获取浏览器
	userAgent := r.Header.Get("User-Agent")
	browser := getBrowserName(userAgent)
	//整合打印
	fmt.Println("客户端id+端口+时间+浏览器:", addr, newNow, browser)
}

func getBrowserName(userAgent string) string {
	// 转小写，避免大小写匹配问题
	ua := strings.ToLower(userAgent)
	// 按优先级匹配浏览器关键词（覆盖主流浏览器）
	switch {
	case strings.Contains(ua, "chrome") && !strings.Contains(ua, "edg"):
		return "Chrome（谷歌浏览器）"
	case strings.Contains(ua, "edg"):
		return "Edge（微软浏览器）"
	case strings.Contains(ua, "firefox"):
		return "Firefox（火狐浏览器）"
	case strings.Contains(ua, "safari") && !strings.Contains(ua, "chrome"):
		return "Safari（苹果浏览器）"
	case strings.Contains(ua, "qqbrowser"):
		return "QQ浏览器"
	default:
		return "未知浏览器"
	}
}
