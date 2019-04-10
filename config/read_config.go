package config

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"../mylog"
	"../tools"
)

const server_conf string = "./config/conf/server.conf"

//the config variaible  and their default value
var (
	Listen_addr    = "localhost:8080"
	Read_time_out  = 5 * time.Second
	Write_time_out = 5 * time.Second
	Public_addr    = "localhost:8080"
)

func init() {
	read_server_config()
	list_config()
}

//read the config from file into a map
func read_server_config() {
	confmap := make(map[string]string)
	f, err := os.Open(server_conf)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 { //can't find '='
			continue
		}
		key := strings.TrimSpace(s[:index])
		value := strings.TrimSpace(s[index+1:])
		if len(key) == 0 || len(value) == 0 {
			continue
		}
		confmap[key] = value
	}
	// fmt.Println(confmap)
	set_variaible(confmap)
}

//read the config from map and write in variaible
func set_variaible(confmap map[string]string) {

	//if there is value in config file then use new one,else use default
	Listen_addr = tools.NotNullVal(confmap["Listen_addr"], Listen_addr)
	Public_addr = tools.NotNullVal(confmap["Public_addr"], Public_addr)

	//change the config-variable if they have specised in the config file
	//else use the default value
	if confmap["Read_time_out"] != "" {
		newint, err := strconv.Atoi(confmap["Read_time_out"])
		if err != nil {
			log.Fatal("worng at reading config Read_time_out : ", err)
		}
		Read_time_out = time.Duration(newint) * time.Second
	}

	if confmap["Write_time_out"] != "" {
		newint, err := strconv.Atoi(confmap["Write_time_out"])
		if err != nil {
			log.Fatal("worng at reading config Write_time_out : ", err)
		}
		Write_time_out = time.Duration(newint) * time.Second
	}

}

//list config
func list_config() {
	mylog.Log("Listen_addr  " + Listen_addr)
	mylog.Log("Public_addr  " + Public_addr)
	mylog.Log("Read_time_out  " + Read_time_out.String())
	mylog.Log("Write_time_out  " + Write_time_out.String())
}
