package service

import (
	"berkeley/config"
	"berkeley/utils"
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"
)

var rinconRetries = 0
var rinconHost = "http://rincon" + ":" + config.RinconPort

func RegisterRincon() {
	var portInt, _ = strconv.Atoi(config.Port)
	rinconBody, _ := json.Marshal(map[string]interface{}{
		"name":         "Berkeley",
		"version":      config.Version,
		"url":          "http://berkeley:" + config.Port,
		"port":         portInt,
		"status_email": config.StatusEmail,
	})
	// Azure Container App deployment
	var ContainerAppEnvDNSSuffix = os.Getenv("CONTAINER_APP_ENV_DNS_SUFFIX")
	if ContainerAppEnvDNSSuffix != "" {
		utils.SugarLogger.Infoln("Found Azure Container App environment variables, using internal DNS suffix: " + ContainerAppEnvDNSSuffix)
		rinconHost = "http://rincon.internal." + ContainerAppEnvDNSSuffix
		rinconBody, _ = json.Marshal(map[string]interface{}{
			"name":         "Berkeley",
			"version":      config.Version,
			"url":          "http://berkeley.internal." + ContainerAppEnvDNSSuffix,
			"port":         portInt,
			"status_email": config.StatusEmail,
		})
	}

	responseBody := bytes.NewBuffer(rinconBody)
	res, err := http.Post(rinconHost+"/services", "application/json", responseBody)
	if err != nil {
		utils.SugarLogger.Errorln(err.Error())
		if rinconRetries < 15 {
			rinconRetries++
			if rinconRetries%2 == 0 {
				rinconHost = "http://localhost" + ":" + config.RinconPort
				utils.SugarLogger.Errorln("failed to register with rincon, retrying with \"http://localhost\" in 5s...")
			} else {
				rinconHost = "http://rincon" + ":" + config.RinconPort
				utils.SugarLogger.Errorln("failed to register with rincon, retrying with \"http://rincon\" in 5s...")
			}
			time.Sleep(time.Second * 5)
			RegisterRincon()
		} else {
			utils.SugarLogger.Fatalln("failed to register with rincon after 15 attempts, terminating program...")
		}
	} else {
		defer res.Body.Close()
		if res.StatusCode == 200 {
			json.NewDecoder(res.Body).Decode(&config.Service)
		} else {
			utils.SugarLogger.Errorln("Failed to register with Rincon! Status code: " + strconv.Itoa(res.StatusCode))
		}
		utils.SugarLogger.Infoln("Registered service with Rincon! Service ID: " + strconv.Itoa(config.Service.ID))
		RegisterRinconRoute("/berkeley")
		RegisterRinconRoute("/schools")
	}
}

func RegisterRinconRoute(route string) {
	rinconBody, _ := json.Marshal(map[string]string{
		"route":        route,
		"service_name": "Berkeley",
	})
	responseBody := bytes.NewBuffer(rinconBody)
	_, err := http.Post(rinconHost+"/routes", "application/json", responseBody)
	if err != nil {
		utils.SugarLogger.Errorln(err.Error())
	}
	utils.SugarLogger.Infoln("Registered route " + route)
}
