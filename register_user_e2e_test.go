package e2e

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
)

type RegisterUserTestsuite struct {
	suite.Suite
	compose tc.ComposeStack
}

func (c *RegisterUserTestsuite) SetupSuite() {
	composeFilePaths := []string{"resources/docker-compose.yml"}

	compose, err := tc.NewDockerCompose(composeFilePaths...)
	if err != nil {
		log.Fatalf("failed to init docker compose")
	}
	c.compose = compose

	ctx := context.Background()

	log.Println("Starting Docker Compose...")
	var emt []tc.StackUpOption
	upErr := compose.Up(ctx, emt...)
	if upErr != nil {
		log.Fatalf("failed to bring up docker compose: %v", upErr)
	}
	time.Sleep(5 * time.Second)
}

func (c *RegisterUserTestsuite) TearDownSuite() {
	ctx := context.Background()
	var emt []tc.StackDownOption
	execError := c.compose.Down(ctx, emt...)
	if execError != nil {
		log.Fatalf("Could not shutdown compose stack: %v", execError)
	}
}

func TestCreateOrderTestSuite(t *testing.T) {
	suite.Run(t, new(RegisterUserTestsuite))
}
