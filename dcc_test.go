package dcc

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// func TestNewClient(t *testing.T) {}
// func TestSendRequest(t *testing.T) {}

func TestWrapperMethods(t *testing.T) {
	// there's gotta be a better way to reference a set of functions and call them sequentially...?
	// cases := []struct {
	// 	method string
	// 	code int
	// }{
	// 	{"get", 200},
	// }

	t.Parallel()
	ts := httptest.NewTLSServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			f, err := os.Open("fixtures/icanhazdadjoke.json")
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()
			io.Copy(w, f)
		},
	))
	defer ts.Close()
	c := NewClient()
	c.BaseURL = ts.URL
	c.HTTPClient = ts.Client()

	t.Run("Get", func(t *testing.T) {
		response, err := c.Get("/")
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
		response, err := c.Post("/")
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
