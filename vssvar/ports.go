package vssvar

import "os"

func GetAPIPort() string {
	var apiPort string = "8080"
	apiPortValue, apiPortPresent := os.LookupEnv("API_PORT")
	if apiPortPresent {
		apiPort = apiPortValue
	}

	return apiPort
}

func GetPrometheusPort() string {
	var prometheusPort string = "8081"
	prometheusPortValue, prometheusPortPresent := os.LookupEnv("PROM_PORT")
	if prometheusPortPresent {
		prometheusPort = prometheusPortValue
	}

	return prometheusPort
}
