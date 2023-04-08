package simplehttp

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// some good test examples at https://github.com/imroc/req/blob/master/req_test.go

var fixturePath string

func init() {
	pwd, _ := os.Getwd()
	fixturePath = filepath.Join(pwd, "fixtures")
}

func TestWrapperMethods(t *testing.T) {
	// there's gotta be a better way to reference a set of functions and call them sequentially...?
	// cases := []struct {
	// 	method string
	// 	code int
	// }{
	// 	{"get", 200},
	// }

	t.Parallel()
	ts := httptest.NewTLSServer(http.HandlerFunc(handleHTTP))
	defer ts.Close()
	c := New(ts.URL)
	c.HTTPClient = ts.Client()

	t.Run("Get", func(t *testing.T) {
		response, err := c.Get("/icanhazdadjoke")
		if err != nil {
			log.Panicln("error:", err)
		}

		got := response
		want := HttpResponse{
			Body: "",
			Code: 200,
		}

		if !cmp.Equal(want.Code, got.Code) {
			t.Error(cmp.Diff(want.Code, got.Code))
		}
	})

	t.Run("Post", func(t *testing.T) {
		c.Data["key"] = "value"
		response, err := c.Post("/echo")
		if err != nil {
			log.Panicln("error:", err)
		}

		got := response
		want := HttpResponse{
			Body: "{\"header\":{\"Accept-Encoding\":[\"gzip\"],\"Content-Length\":[\"15\"],\"User-Agent\":[\"Go-http-client/1.1\"]},\"body\":\"{\\\"key\\\":\\\"value\\\"}\"}",
			Code: 200,
		}

		if !cmp.Equal(want.Code, got.Code) {
			t.Error(cmp.Diff(want.Code, got.Code))
		}
		if !cmp.Equal(want.Body, got.Body) {
			t.Error(cmp.Diff(want.Body, got.Body))
		}
	})

	t.Run("Patch", func(t *testing.T) {
		response, err := c.Patch("/")
		if err != nil {
			log.Panicln("error:", err)
		}

		got := response
		want := HttpResponse{
			Body: "",
			Code: 200,
		}

		if !cmp.Equal(want.Code, got.Code) {
			t.Error(cmp.Diff(want.Code, got.Code))
		}
	})

	t.Run("Put", func(t *testing.T) {
		response, err := c.Put("/")
		if err != nil {
			log.Panicln("error:", err)
		}

		got := response
		want := HttpResponse{
			Body: "",
			Code: 200,
		}

		if !cmp.Equal(want.Code, got.Code) {
			t.Error(cmp.Diff(want.Code, got.Code))
		}
	})

	t.Run("Delete", func(t *testing.T) {
		response, err := c.Delete("/")
		if err != nil {
			log.Panicln("error:", err)
		}

		got := response
		want := HttpResponse{
			Body: "",
			Code: 200,
		}

		if !cmp.Equal(want.Code, got.Code) {
			t.Error(cmp.Diff(want.Code, got.Code))
		}
	})
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Method", r.Method)
	switch r.Method {
	case http.MethodGet:
		handleGet(w, r)
	case http.MethodPost:
		handlePost(w, r)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	// 	case "/":
	// 		w.Write([]byte("TestGet: text response"))
	case "/icanhazdadjoke":
		f, err := ioutil.ReadFile(fmt.Sprintf("%s/icanhazdadjoke.json", fixturePath))
		if err != nil {
			log.Panicf("Error. %+v", err)
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(f)
		// 	case "/bad-request":
		// 		w.WriteHeader(http.StatusBadRequest)
		// 	case "/too-many":
		// 		w.WriteHeader(http.StatusTooManyRequests)
		// 		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 		w.Write([]byte(`{"errMsg":"too many requests"}`))
		// 	case "/chunked":
		// 		w.Header().Add("Trailer", "Expires")
		// 		w.Write([]byte(`This is a chunked body`))
		// 	case "/host-header":
		// 		w.Write([]byte(r.Host))
		// 	case "/json":
		// 		r.ParseForm()
		// 		if r.FormValue("type") != "no" {
		// 			w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 		}
		// 		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 		if r.FormValue("error") == "yes" {
		// 			w.WriteHeader(http.StatusBadRequest)
		// 			w.Write([]byte(`{"message": "not allowed"}`))
		// 		} else {
		// 			w.Write([]byte(`{"name": "roc"}`))
		// 		}
		// 	case "/xml":
		// 		r.ParseForm()
		// 		if r.FormValue("type") != "no" {
		// 			w.Header().Set("Content-Type", header.XmlContentType)
		// 		}
		// 		w.Write([]byte(`<user><name>roc</name></user>`))
		// 	case "/unlimited-redirect":
		// 		w.Header().Set("Location", "/unlimited-redirect")
		// 		w.WriteHeader(http.StatusMovedPermanently)
		// 	case "/redirect-to-other":
		// 		w.Header().Set("Location", "http://dummy.local/test")
		// 		w.WriteHeader(http.StatusMovedPermanently)
		// 	case "/pragma":
		// 		w.Header().Add("Pragma", "no-cache")
		// 	case "/payload":
		// 		b, _ := ioutil.ReadAll(r.Body)
		// 		w.Write(b)
		// 	case "/gbk":
		// 		w.Header().Set("Content-Type", "text/plain; charset=gbk")
		// 		w.Write(toGbk("我是roc"))
		// 	case "/gbk-no-charset":
		// 		b, err := ioutil.ReadFile(tests.GetTestFilePath("sample-gbk.html"))
		// 		if err != nil {
		// 			panic(err)
		// 		}
		// 		w.Header().Set("Content-Type", "text/html")
		// 		w.Write(b)
		// 	case "/header":
		// 		b, _ := json.Marshal(r.Header)
		// 		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 		w.Write(b)
		// 	case "/user-agent":
		// 		w.Write([]byte(r.Header.Get(header.UserAgent)))
		// 	case "/content-type":
		// 		w.Write([]byte(r.Header.Get("Content-Type")))
		// 	case "/query-parameter":
		// 		w.Write([]byte(r.URL.RawQuery))
		// 	case "/search":
		// 		handleSearch(w, r)
		// 	case "/download":
		// 		size := 100 * 1024 * 1024
		// 		w.Header().Set("Content-Length", strconv.Itoa(size))
		// 		buf := make([]byte, 1024)
		// 		for i := 0; i < 1024; i++ {
		// 			buf[i] = 'h'
		// 		}
		// 		for i := 0; i < size; {
		// 			wbuf := buf
		// 			if size-i < 1024 {
		// 				wbuf = buf[:size-i]
		// 			}
		// 			n, err := w.Write(wbuf)
		// 			if err != nil {
		// 				break
		// 			}
		// 			i += n
		// 		}
		// 	case "/protected":
		// 		auth := r.Header.Get("Authorization")
		// 		if auth == "Bearer goodtoken" {
		// 			w.Write([]byte("good"))
		// 		} else {
		// 			w.WriteHeader(http.StatusUnauthorized)
		// 			w.Write([]byte(`bad`))
		// 		}
		// 	default:
		// 		if strings.HasPrefix(r.URL.Path, "/user") {
		// 			handleGetUserProfile(w, r)
		// 		}
	}
}

// Echo is used in "/echo" API.
type Echo struct {
	Header http.Header `json:"header" xml:"header"`
	Body   string      `json:"body" xml:"body"`
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		io.Copy(ioutil.Discard, r.Body)
		w.Write([]byte("TestPost: text response"))
	case "/raw-upload":
		io.Copy(ioutil.Discard, r.Body)
	case "/file-text":
		r.ParseMultipartForm(10e6)
		files := r.MultipartForm.File["file"]
		file, _ := files[0].Open()
		b, _ := ioutil.ReadAll(file)
		r.ParseForm()
		if a := r.FormValue("attempt"); a != "" && a != "2" {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(b)
	case "/form":
		r.ParseForm()
		ret, _ := json.Marshal(&r.Form)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(ret)
	case "/multipart":
		r.ParseMultipartForm(10e6)
		m := make(map[string]interface{})
		m["values"] = r.MultipartForm.Value
		m["files"] = r.MultipartForm.File
		ret, _ := json.Marshal(&m)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(ret)
	// case "/search":
	// 	handleSearch(w, r)
	case "/redirect":
		io.Copy(ioutil.Discard, r.Body)
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusMovedPermanently)
	case "/content-type":
		io.Copy(ioutil.Discard, r.Body)
		w.Write([]byte(r.Header.Get("Content-Type")))
	case "/echo":
		b, _ := ioutil.ReadAll(r.Body)
		e := Echo{
			Header: r.Header,
			Body:   string(b),
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		result, _ := json.Marshal(&e)
		w.Write(result)
	}
}

// online-only test; might use this eventually if I can mock out the correct return values
// func TestWeather(t *testing.T) {
// 	type WeatherResponse struct {
// 		CurrentCondition []struct {
// 			FeelsLikeC string `json:"FeelsLikeC,omitempty"`
// 			FeelsLikeF string `json:"FeelsLikeF,omitempty"`
// 			Cloudcover string `json:"cloudcover,omitempty"`
// 			Humidity   string `json:"humidity,omitempty"`
// 			LangFr     []struct {
// 				Value string `json:"value,omitempty"`
// 			} `json:"lang_fr,omitempty"`
// 			LocalObsDateTime string `json:"localObsDateTime,omitempty"`
// 			ObservationTime  string `json:"observation_time,omitempty"`
// 			PrecipInches     string `json:"precipInches,omitempty"`
// 			PrecipMM         string `json:"precipMM,omitempty"`
// 			Pressure         string `json:"pressure,omitempty"`
// 			PressureInches   string `json:"pressureInches,omitempty"`
// 			TempC            string `json:"temp_C,omitempty"`
// 			TempF            string `json:"temp_F,omitempty"`
// 			UvIndex          string `json:"uvIndex,omitempty"`
// 			Visibility       string `json:"visibility,omitempty"`
// 			VisibilityMiles  string `json:"visibilityMiles,omitempty"`
// 			WeatherCode      string `json:"weatherCode,omitempty"`
// 			WeatherDesc      []struct {
// 				Value string `json:"value,omitempty"`
// 			} `json:"weatherDesc,omitempty"`
// 			WeatherIconURL []struct {
// 				Value string `json:"value,omitempty"`
// 			} `json:"weatherIconUrl,omitempty"`
// 			Winddir16Point string `json:"winddir16Point,omitempty"`
// 			WinddirDegree  string `json:"winddirDegree,omitempty"`
// 			WindspeedKmph  string `json:"windspeedKmph,omitempty"`
// 			WindspeedMiles string `json:"windspeedMiles,omitempty"`
// 		} `json:"current_condition,omitempty"`
// 		NearestArea []struct {
// 			AreaName []struct {
// 				Value string `json:"value,omitempty"`
// 			} `json:"areaName,omitempty"`
// 			Country []struct {
// 				Value string `json:"value,omitempty"`
// 			} `json:"country,omitempty"`
// 			Latitude   string `json:"latitude,omitempty"`
// 			Longitude  string `json:"longitude,omitempty"`
// 			Population string `json:"population,omitempty"`
// 			Region     []struct {
// 				Value string `json:"value,omitempty"`
// 			} `json:"region,omitempty"`
// 			WeatherURL []struct {
// 				Value string `json:"value,omitempty"`
// 			} `json:"weatherUrl,omitempty"`
// 		} `json:"nearest_area,omitempty"`
// 		Request []struct {
// 			Query string `json:"query,omitempty"`
// 			Type  string `json:"type,omitempty"`
// 		} `json:"request,omitempty"`
// 		Weather []struct {
// 			Astronomy []struct {
// 				MoonIllumination string `json:"moon_illumination,omitempty"`
// 				MoonPhase        string `json:"moon_phase,omitempty"`
// 				Moonrise         string `json:"moonrise,omitempty"`
// 				Moonset          string `json:"moonset,omitempty"`
// 				Sunrise          string `json:"sunrise,omitempty"`
// 				Sunset           string `json:"sunset,omitempty"`
// 			} `json:"astronomy,omitempty"`
// 			AvgtempC string `json:"avgtempC,omitempty"`
// 			AvgtempF string `json:"avgtempF,omitempty"`
// 			Date     string `json:"date,omitempty"`
// 			Hourly   []struct {
// 				DewPointC        string `json:"DewPointC,omitempty"`
// 				DewPointF        string `json:"DewPointF,omitempty"`
// 				FeelsLikeC       string `json:"FeelsLikeC,omitempty"`
// 				FeelsLikeF       string `json:"FeelsLikeF,omitempty"`
// 				HeatIndexC       string `json:"HeatIndexC,omitempty"`
// 				HeatIndexF       string `json:"HeatIndexF,omitempty"`
// 				WindChillC       string `json:"WindChillC,omitempty"`
// 				WindChillF       string `json:"WindChillF,omitempty"`
// 				WindGustKmph     string `json:"WindGustKmph,omitempty"`
// 				WindGustMiles    string `json:"WindGustMiles,omitempty"`
// 				Chanceoffog      string `json:"chanceoffog,omitempty"`
// 				Chanceoffrost    string `json:"chanceoffrost,omitempty"`
// 				Chanceofhightemp string `json:"chanceofhightemp,omitempty"`
// 				Chanceofovercast string `json:"chanceofovercast,omitempty"`
// 				Chanceofrain     string `json:"chanceofrain,omitempty"`
// 				Chanceofremdry   string `json:"chanceofremdry,omitempty"`
// 				Chanceofsnow     string `json:"chanceofsnow,omitempty"`
// 				Chanceofsunshine string `json:"chanceofsunshine,omitempty"`
// 				Chanceofthunder  string `json:"chanceofthunder,omitempty"`
// 				Chanceofwindy    string `json:"chanceofwindy,omitempty"`
// 				Cloudcover       string `json:"cloudcover,omitempty"`
// 				Humidity         string `json:"humidity,omitempty"`
// 				LangFr           []struct {
// 					Value string `json:"value,omitempty"`
// 				} `json:"lang_fr,omitempty"`
// 				PrecipInches    string `json:"precipInches,omitempty"`
// 				PrecipMM        string `json:"precipMM,omitempty"`
// 				Pressure        string `json:"pressure,omitempty"`
// 				PressureInches  string `json:"pressureInches,omitempty"`
// 				TempC           string `json:"tempC,omitempty"`
// 				TempF           string `json:"tempF,omitempty"`
// 				Time            string `json:"time,omitempty"`
// 				UvIndex         string `json:"uvIndex,omitempty"`
// 				Visibility      string `json:"visibility,omitempty"`
// 				VisibilityMiles string `json:"visibilityMiles,omitempty"`
// 				WeatherCode     string `json:"weatherCode,omitempty"`
// 				WeatherDesc     []struct {
// 					Value string `json:"value,omitempty"`
// 				} `json:"weatherDesc,omitempty"`
// 				WeatherIconURL []struct {
// 					Value string `json:"value,omitempty"`
// 				} `json:"weatherIconUrl,omitempty"`
// 				Winddir16Point string `json:"winddir16Point,omitempty"`
// 				WinddirDegree  string `json:"winddirDegree,omitempty"`
// 				WindspeedKmph  string `json:"windspeedKmph,omitempty"`
// 				WindspeedMiles string `json:"windspeedMiles,omitempty"`
// 			} `json:"hourly,omitempty"`
// 			MaxtempC    string `json:"maxtempC,omitempty"`
// 			MaxtempF    string `json:"maxtempF,omitempty"`
// 			MintempC    string `json:"mintempC,omitempty"`
// 			MintempF    string `json:"mintempF,omitempty"`
// 			SunHour     string `json:"sunHour,omitempty"`
// 			TotalSnowCm string `json:"totalSnow_cm,omitempty"`
// 			UvIndex     string `json:"uvIndex,omitempty"`
// 		} `json:"weather,omitempty"`
// 	}

// 	c := NewClient()
// 	// Get weather in French JSON
// 	c.BaseURL = "http://wttr.in"
// 	c.Headers["Accept-Language"] = "fr"
// 	c.Params["format"] = "j1"
// 	response, err := c.Get("/55328")
// 	if err != nil {
// 		log.Panicln("error:", err)
// 	}
// 	got := response
// 	want := HttpResponse{
// 		Body: "",
// 		Code: 200,
// 	}
// 	if !cmp.Equal(want.Code, got.Code) {
// 		t.Error(cmp.Diff(want, got))
// 	}

// 	var gotbody WeatherResponse
// 	unmarshall_error := json.Unmarshal([]byte(got.Body), &gotbody)
// 	if unmarshall_error != nil {
// 		log.Panicln("error:", err)
// 	}
// 	log.Println("type:", gotbody.Request)
// }
