package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strconv"
)

const LimitDefault string = "-1"
const OrderDefault string = "asc"

type GetRoutesResponse struct {
	Source string  `json:"source"`
	Routes []Route `json:"routes"`
}

func GetRoutes(c *gin.Context) {
	src := c.Query("src")
	dsts := c.QueryArray("dst")
	limit := c.DefaultQuery("limit", LimitDefault)
	order := c.DefaultQuery("order", OrderDefault)

	routes := FetchAllRoutes(src, dsts)

	if order == "desc" {
		sort.Sort(sort.Reverse(ByDurationAndDistance(routes)))
	} else {
		sort.Sort(ByDurationAndDistance(routes))
	}

	if limit != LimitDefault {
		routes = headRoutes(routes, limit)
	}

	r := GetRoutesResponse{
		Source: src,
		Routes: routes,
	}

	c.JSON(http.StatusOK, r)
}

func headRoutes(routes []Route, limitParam string) []Route {
	limitInt, err := strconv.Atoi(limitParam)
	if err != nil {
		fmt.Println(err.Error())
		return routes
	} else if limitInt < 0 {
		return routes
	}

	return routes[:min(limitInt, len(routes))]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
