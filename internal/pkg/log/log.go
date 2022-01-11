package log
import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
)
type myHook struct {
	baseurl string
	appname string
}
func (h *myHook) Levels() []log.Level {
	return []log.Level{
		log.InfoLevel,
		log.ErrorLevel,
	}
}
func (h *myHook) Fire(entry *log.Entry) error {
	dst := "info/"
	switch entry.Level.String() {
	case "error":
		dst = "error/"
		break
	case "info":
		dst  = "info/"
		break
	default:
		dst = "info/"
		break
	}
	req, err := http.NewRequest("POST", h.baseurl + dst + h.appname, bytes.NewBufferString(entry.Message))
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	if err != nil {
		fmt.Println(err)
	}
	defer req.Body.Close()
	//发送请求
	client := &http.Client{}
	response, _ := client.Do(req)
	if response!= nil && response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}
	defer func() {
		if response != nil{
			response.Body.Close()
		}
	}()
	return nil
}
func NewLogger(ip string, port int,appname string)*log.Logger{
	logger := log.New()
	logger.Hooks.Add(&myHook{baseurl:"http://" + ip + ":" + strconv.Itoa(port) +  "/" ,appname:appname})
	return logger
}