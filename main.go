package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"ingrid-coding-assignment/route"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
)

type RouteResponse struct {
	Source string        `json:"source"`
	Routes []route.Route `json:"routes"`
}

type osrmResponse struct {
	Code   string `json:"code"`
	Routes []struct {
		Duration float64 `json:"duration"`
		Distance float64 `json:"distance"`
	} `json:"routes"`
}

const OSRM_DRIVING_BASE_URL string = "http://router.project-osrm.org/route/v1/driving/"
const LIMIT_DEFAULT string = "-1"
const ORDER_DEFAULT string = "asc"

func main() {
	// test URL: http://localhost:8080/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219
	r := gin.Default()
	r.GET("/routes", getRoutes)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func getRoutes(c *gin.Context) {
	src := c.Query("src")
	dsts := c.QueryArray("dst")
	limit := c.DefaultQuery("limit", LIMIT_DEFAULT)
	order := c.DefaultQuery("order", ORDER_DEFAULT)

	routes := fetchAllRoutes(src, dsts)

	if order == "desc" {
		sort.Sort(sort.Reverse(route.ByDurationAndDistance(routes)))
	} else {
		sort.Sort(route.ByDurationAndDistance(routes))
	}

	if limit != LIMIT_DEFAULT {
		routes = headRoutes(routes, limit)
	}

	r := RouteResponse{
		Source: src,
		Routes: routes,
	}

	c.JSON(http.StatusOK, r)
}

func headRoutes(routes []route.Route, limitParam string) []route.Route {
	limitInt, err := strconv.Atoi(limitParam)
	if err != nil {
		fmt.Println(err.Error())
		return routes
	}

	return routes[:(limitInt + 1)]
}

func fetchAllRoutes(src string, dsts []string) []route.Route {
	var out []route.Route
	for _, dst := range dsts {
		ch := fetchOsrmRoute(src, dst)
		osrmResponse := <-ch

		if isRouteValid(osrmResponse) {
			out = append(out, route.Route{
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

		response, err := http.Get(OSRM_DRIVING_BASE_URL + src + ";" + dst + "?overview=false")

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err.Error())
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
