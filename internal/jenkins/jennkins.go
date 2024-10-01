package jenkins

import (
	"context"
	"github.com/rs/zerolog/log"

	"github.com/bndr/gojenkins"
)

var Ctx context.Context
var Client *gojenkins.Jenkins
var defaultBranch = "dev"

func SetupJenkins() {
	Ctx = context.Background()
	JenkinsClient := gojenkins.CreateJenkins(nil, "http://113.160.247.103:9000", "nhatcx", "111fc9efe79685c067f923bf3d2ce8545e")
	_, err := JenkinsClient.Init(Ctx)
	if err != nil {
		log.Fatal().Msg("Failed to init Jenkins client")
		panic("Invalid jenkins client")
	}
}
