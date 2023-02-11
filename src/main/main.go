package main

import (
	"config"
)

func main() {

	c := new(config.RedisConfig)
	c.Host = "10.216.233.41:34999"
	c.Pass = "Walsinyonyou1201"
	c.Zero()
	c.Flush()

}
