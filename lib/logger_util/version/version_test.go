package version_test

import (
	"amfLoadBalancer/lib/logger_util/version"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	assert.Equal(t, "2020-03-31-01", version.GetVersion())
}
