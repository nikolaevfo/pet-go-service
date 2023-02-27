package api

import (
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type Config struct {
	URL   string
	Token string
}

type Api struct {
	config *Config
}

func (a Api) GetUsers() string {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}

	req, _ := http.NewRequest(http.MethodGet, a.config.URL, nil)

	resp, err := client.Do(req)
	if err != nil {
		return "error happend" + string(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}

func NewApi(config *Config) *Api {
	return &Api{
		config: config,
	}
}



// простой запрос================================================
// resp, err := http.Get(url)

// if err != nil {
// 	fmt.Println("error happend", err)
// 	return "error happend" + string(err.Error())
// }

// defer resp.Body.Close()

// respBody, err := ioutil.ReadAll(resp.Body)

// более сложный запрос тоже лучше не использовать ============================================
// req := &http.Request{
// 	Method: http.MethodGet,
// }

// req.URL, _ = url.Parse(urlString)

// resp, err := http.DefaultClient.Do(req)
// if err != nil {
// 	return "error happend" + string(err.Error())
// }

// defer resp.Body.Close()

// body, err := ioutil.ReadAll(resp.Body)