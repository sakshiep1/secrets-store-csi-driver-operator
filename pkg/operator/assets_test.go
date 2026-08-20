package operator

import (
	"testing"

	"github.com/openshift/secrets-store-csi-driver-operator/assets"
	appsv1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/yaml"
)

const clusterAutoscalerEnableDsEvictionAnnotation = "cluster-autoscaler.kubernetes.io/enable-ds-eviction"

func TestNodeDaemonSetHasClusterAutoscalerEvictionOptOut(t *testing.T) {
	data, err := assets.ReadFile("node.yaml")
	if err != nil {
		t.Fatalf("failed to read node.yaml: %v", err)
	}

	var ds appsv1.DaemonSet
	if err := yaml.Unmarshal(data, &ds); err != nil {
		t.Fatalf("failed to unmarshal node.yaml: %v", err)
	}

	val, ok := ds.Spec.Template.Annotations[clusterAutoscalerEnableDsEvictionAnnotation]
	if !ok {
		t.Fatalf("annotation %q not found on DaemonSet pod template", clusterAutoscalerEnableDsEvictionAnnotation)
	}
	if val != "false" {
		t.Fatalf("annotation %q = %q, want %q", clusterAutoscalerEnableDsEvictionAnnotation, val, "false")
	}
}
