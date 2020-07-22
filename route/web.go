package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const OsrmDrivingBaseUrl string = "http://router.project-osrm.org/route/v1/driving/"

type osrmResponse struct {
	Code   string `json:"code"`
	Routes []struct {
		Duration float64 `json:"duration"`
		Distance float64 `json:"distance"`
	} `json:"routes"`
}

func FetchAllRoutes(src string, dsts []string) []Route {
	var out []Route
	for _, dst := range dsts {
		ch := fetchOsrmRoute(src, dst)
		osrmResponse := <-ch

		if isRouteValid(osrmResponse) {
			out = append(out, Route{
				Dest:     dst,
				Duration: osrmResponse.Routes[0].Duration,
				Distance: osrmResponse.Routes[0].Distance,
			})
		}
	}

	return out
}

func isRouteValid(response osrmResponse) bool {
	return response.Code == "Ok"
}

func fetchOsrmRoute(src string, dst string) <-chan osrmResponse {
	r := make(chan osrmResponse)

	go func() {
		defer close(r)

		response, err := http.Get(OsrmDrivingBaseUrl + src + ";" + dst + "?overview=false")

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if response.StatusCode != http.StatusOK {
			errorMessage := struct {
				Message string `json:"message"`
			}{}
			json.Unmarshal(responseData, &errorMessage)
			fmt.Println(errorMessage.Message, "code: ", response.StatusCode)
			return
		}

		osrmResponse := osrmResponse{}
		err = json.Unmarshal(responseData, &osrmResponse)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		r <- osrmResponse
	}()

	return r
}
