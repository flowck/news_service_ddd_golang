package components

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/flowck/news_service_ddd_golang/tests/client"
	"github.com/stretchr/testify/require"
)

var (
	ctx          context.Context
	testsTimeout = time.Minute * 2
)

func TestMain(m *testing.M) {
	tmpCtx, cancel := context.WithTimeout(context.Background(), testsTimeout)
	defer cancel()
	ctx = tmpCtx

	os.Exit(m.Run())
}

func getClient(t *testing.T) *client.ClientWithResponses {
	cli, err := client.NewClientWithResponses(fmt.Sprintf("http://localhost:%s", os.Getenv("PORT")))
	require.Nil(t, err)

	return cli
}

func toPtr[V any](v V) *V {
	return &v
}
