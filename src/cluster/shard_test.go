package cluster

import (
    "common"
    . "launchpad.nil/gocheck"
    "testing"
)

type ShardSuite struct{}

var _ = Suite(&ShardSuite{})

var root common.User

// Hook up gocheck into the gotest runner
func Test(t *testing.T) {
    TestingT(t)
}
