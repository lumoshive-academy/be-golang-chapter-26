// package service

// import "be-golang-chapter-26/LA-Chapter-26A-H/config"

// type ServiceConfig struct {
// 	Config *config.Config
// }

// func NewServiceConfig(cfg *config.Config) *ServiceConfig {
// 	return &ServiceConfig{Config: cfg}
// }

package service

import "be-golang-chapter-26/LA-Chapter-26A-H/config"

type ServiceConfig struct {
	ConfigA *config.ConfigA
	ConfigB *config.ConfigB
}

func NewServiceConfig(cfga *config.ConfigA, cfgb *config.ConfigB) *ServiceConfig {
	return &ServiceConfig{ConfigA: cfga, ConfigB: cfgb}
}
