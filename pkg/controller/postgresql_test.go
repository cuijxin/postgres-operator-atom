package controller

import (
	"reflect"
	"testing"

	acidv1 "github.com/cuijxin/postgres-operator-atom/pkg/apis/acid.zalan.do/v1"
	"github.com/cuijxin/postgres-operator-atom/pkg/spec"
)

var (
	True  = true
	False = false
)

func TestMergeDeprecatedPostgreSQLSpecParameters(t *testing.T) {
	c := NewController(&spec.ControllerConfig{})

	tests := []struct {
		name  string
		in    *acidv1.PostgresSpec
		out   *acidv1.PostgresSpec
		error string
	}{
		{
			"Check that old parameters propagate values to the new ones",
			&acidv1.PostgresSpec{UseLoadBalancer: &True, ReplicaLoadBalancer: &True},
			&acidv1.PostgresSpec{UseLoadBalancer: nil, ReplicaLoadBalancer: nil,
				EnableMasterLoadBalancer: &True, EnableReplicaLoadBalancer: &True},
			"New parameters should be set from the values of old ones",
		},
		{
			"Check that new parameters are not set when both old and new ones are present",
			&acidv1.PostgresSpec{UseLoadBalancer: &True, EnableMasterLoadBalancer: &False},
			&acidv1.PostgresSpec{UseLoadBalancer: nil, EnableMasterLoadBalancer: &False},
			"New parameters should remain unchanged when both old and new are present",
		},
	}
	for _, tt := range tests {
		result := c.mergeDeprecatedPostgreSQLSpecParameters(tt.in)
		if !reflect.DeepEqual(result, tt.out) {
			t.Errorf("%s: %v", tt.name, tt.error)
		}
	}
}
