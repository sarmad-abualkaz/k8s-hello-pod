package config

import "os"

type PodMetaData struct {
	ClusterID, NodeName, Namespace, PodName string
}

type Config struct {
	Pod PodMetaData
	Port string
}

// Return a new config struct:
func New() *Config {
	return &Config{
		Pod: PodMetaData{
			ClusterID: getEnv("MY_CLUSTER_ID", ""),
			NodeName: getEnv("MY_NODE_NAME", ""),
			Namespace: getEnv("MY_POD_NAMESPACE", ""),
			PodName: getEnv("MY_POD_NAME", ""),
		},
		Port: getEnv("PORT", "8080"),
	}
}

// Key function to return environment variable if exists, otherwise return default
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// Initialize config for environmnent
var ConfigEnv = New()