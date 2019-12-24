package web

import (
	"mytools/web/routes"
)

func Service(addr string) {
	route := routes.Initroute()
	route.Run(addr)
}
