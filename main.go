package main

import (
	"bili_video_fetch/model"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"
)

var (
	mixinKeyEncTab = []int{
		46, 47, 18, 2, 53, 8, 23, 32, 15, 50, 10, 31, 58, 3, 45, 35, 27, 43, 5, 49,
		33, 9, 42, 19, 29, 28, 14, 39, 12, 38, 41, 13, 37, 48, 7, 16, 24, 55, 40,
		61, 26, 17, 0, 1, 60, 51, 30, 4, 22, 25, 54, 21, 56, 59, 6, 63, 57, 62, 11,
		36, 20, 34, 44, 52,
	}
	sessdata = ""
)

const (
	user_agent   = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"
	baseUrl      = "https://api.bilibili.com/x/player/wbi/playurl"
	navUrl       = "https://api.bilibili.com/x/web-interface/nav"
	videoInfoUrl = "https://api.bilibili.com/x/web-interface/wbi/view?bvid="
)

// 发送请求的优化
func main() {
	http.HandleFunc("/video", GetVideoUrl)
	http.HandleFunc("/login", UserLogin)
	http.ListenAndServe(":9090", nil)
}
func UserLogin(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	loginMethod := params.Get("method")
	switch loginMethod {
	case "qr":
	default:
	}
}

// GetVideoUrl 处理视频请求
func GetVideoUrl(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	bvid := params.Get("bvid")
	videoInfoUrl := WBI(baseUrl, bvid)
	videoDownloadUrl, audioDownloadUrl := GetVideoDownloadUrl(videoInfoUrl)
	FetchVAData(videoDownloadUrl, "video.m4s")
	FetchVAData(audioDownloadUrl, "audio.m4s")
	MixVA(bvid) //ffmpeg整合视频和音频
	fmt.Fprintf(w, "%s", videoInfoUrl)
}
func MixVA(name string) {
	// 定义要执行的 FFmpeg 命令
	cmd := exec.Command("ffmpeg", "-i", "video.m4s", "-i", "audio.m4s", "-c:v", "copy", "-c:a", "copy", "-f", "mp4", name+".mp4")
	cmd.Run()
	os.Remove("video.m4s")
	os.Remove("audio.m4s")
}
func FetchVAData(videoUrl string, fileName string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", videoUrl, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	req.Header.Set("Referer", "https://www.bilibili.com")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	outFile, _ := os.Create(fileName)
	defer outFile.Close()
	io.Copy(outFile, resp.Body)
	fmt.Println("Download completed successfully!")
}
func GetVideoDownloadUrl(videoInfoUrl string) (string, string) {
	videoStream, _ := execute[model.VideoStream](videoInfoUrl, "GET")
	return videoStream.Data.Dash.Video[0].BaseURL, videoStream.Data.Dash.Audio[0].BaseURL
}

// WBI wbi鉴权
func WBI(composedUrl, bvid string) string {
	//1.获取img_key、sub_key
	wbi_key, _ := execute[model.Root[model.LoginData]](navUrl, "GET")
	// 提取文件名
	imgUrl := wbi_key.Data.WbiImg.ImgUrl
	subUrl := wbi_key.Data.WbiImg.SubUrl
	img_key := getFileNameWithoutExtension(imgUrl)
	sub_key := getFileNameWithoutExtension(subUrl)

	fmt.Println("ImgUrl 文件名:", img_key)
	fmt.Println("SubUrl 文件名:", sub_key)
	//2.sub_key拼接img_key
	raw_wbi_key := img_key + sub_key
	mixin_key := getMixinKey(raw_wbi_key)
	//3.计算w_rid
	wts := time.Now().Unix()
	videInfo, _ := execute[model.VideoInfo](videoInfoUrl+bvid, "GET")
	// 构造请求参数（示例）
	params := map[string]interface{}{
		"bvid":     bvid,
		"fourk":    1,
		"fnver":    0,
		"fnval":    16,
		"platform": "pc",
		"cid":      videInfo.Data.Cid,
		"wts":      wts,
	}
	w_rid := generateWrid(params, mixin_key)
	// 将 w_rid 添加到参数
	query := url.Values{}
	params["w_rid"] = w_rid
	for k, v := range params {
		query.Add(k, fmt.Sprintf("%v", v))
	}
	finalUrl := fmt.Sprintf("%s?%s", composedUrl, query.Encode())
	return finalUrl
}

// 生成签名 w_rid
func generateWrid(params map[string]interface{}, mixin_key string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var queryStrings []string
	for _, key := range keys {
		value := params[key]
		encodedValue := encodeURIComponent(fmt.Sprintf("%v", value))
		queryStrings = append(queryStrings, fmt.Sprintf("%s=%s", key, encodedValue))
	}
	queryString := strings.Join(queryStrings, "&") + mixin_key
	hash := md5.Sum([]byte(queryString))
	return hex.EncodeToString(hash[:])
}

// URL 编码处理
func encodeURIComponent(str string) string {
	var result string
	for i := 0; i < len(str); i++ {
		ch := str[i]
		if ch == '%' || ch == '+' || (ch >= '0' && ch <= '9') || (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || ch == '-' || ch == '.' || ch == '_' || ch == '~' {
			result += string(ch)
		} else {
			result += fmt.Sprintf("%%%02X", ch)
		}
	}
	return result
}
func getMixinKey(orig string) string {
	var str strings.Builder
	for _, v := range mixinKeyEncTab {
		if v < len(orig) {
			str.WriteByte(orig[v])
		}
	}
	return str.String()[:32]
}

func getFileNameWithoutExtension(wbiUrl string) string {
	//https://i0.hdslb.com/bfs/wbi/7cd084941338484aae1ad9425b84077c.png
	// 获取 URL 中最后一个 "/" 后的文件名部分
	fileName := wbiUrl[strings.LastIndex(wbiUrl, "/")+1:][:len(wbiUrl[strings.LastIndex(wbiUrl, "/")+1:])-4]
	// 去掉后缀 ".png"
	return fileName
}

func execute[T any](webUrl, method string) (T, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, webUrl, nil)
	req.Header.Add("User-Agent", user_agent)
	req.Header.Add("Cookie", fmt.Sprintf("SESSDATA=%s", sessdata))
	req.Header.Add("Referer", "https://www.bilibili.com")
	resp, _ := client.Do(req)
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	var v T
	err := json.Unmarshal(body, &v)
	return v, err
}
