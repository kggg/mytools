package config

import (
	"strconv"

	"github.com/Unknwon/goconfig"
)

type Hostinfo struct {
	Ipaddr string
	User   string
	Pass   string
	Port   int
	Skey   bool
}

const configpath = "./config/remote.ini"

func Readconfig(sectionname string) (Hostinfo, error) {
	cfg, err := goconfig.LoadConfigFile(configpath)
	if err != nil {
		return Hostinfo{}, err
	}
	sec, err := cfg.GetSection(sectionname)
	if err != nil {
		return Hostinfo{}, err
	}
	var hostinfo Hostinfo
	for v, k := range sec {
		if v == "ipaddr" {
			hostinfo.Ipaddr = k
		}
		if v == "user" {
			hostinfo.User = k
		}
		if v == "pass" {
			hostinfo.Pass = k
		}
		if v == "port" {
			hostinfo.Port, err = strconv.Atoi(k)
			if err != nil {
				return Hostinfo{}, err
			}
		}
		if v == "skey" {
			if k == "true" || k == "yes" {
				hostinfo.Skey = true
			} else {
				hostinfo.Skey = false
			}
		}
	}
	return hostinfo, nil
}
