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

// GetRoutes godoc
// @Summary Retrieves routes from OSRM 3rd party service based on provided source and destination latitude and longitude
// @Produce json
// @Param src query string true "Comma-delimited latitude and longitude values of source location in decimal format"
// @Param dst query []string true "Comma-delimited latitude and longitude values of destination location in decimal format"
// @Param limit query int false "Limit of returned values"
// @Param order query string false "Ordering of routes in the response (asc or desc) - default is asc (ascending)"
// @Success 200 {object} GetRoutesResponse
// @Router /routes [get]
func GetRoutes(c *gin.Context) {
	src := c.Query("src")
	dsts := c.QueryArray("dst")
	limit := c.DefaultQuery("limit", LimitDefault)
	order := c.DefaultQuery("order", OrderDefault)

	if src == "" || len(dsts) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "No src or dst parameter in the query string"})
		return
	}

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
