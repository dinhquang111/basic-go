package jenkins

import (
	"context"

	"github.com/bndr/gojenkins"
)

var JenkinCtx context.Context
var JenkinClient *gojenkins.Jenkins
var defaultBranch = "dev"

func SetupJenkin() {
	JenkinCtx = context.Background()
	JenkinClient := gojenkins.CreateJenkins(nil, "http://113.160.247.103:9000", "nhatcx", "111fc9efe79685c067f923bf3d2ce8545e")
	_, err := JenkinClient.Init(JenkinCtx)
	if err != nil {
		panic("Invalid jenkins client")
	}
}
