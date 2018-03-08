package deployment

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"testing"
	"k8s.io/api/apps/v1"
	"github.com/stretchr/testify/assert"

	"github.com/TimWoolford/podrick/pkg/config"
)

func TestReturnsNameWithNoLabels(t *testing.T) {
	deployment := v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{Name: "foo"},
	}

	dep := &K8sDeployment{
		deployment: deployment,
	}

	assert.Equal(t, "foo", dep.Name())
}

func TestReturnsNameWithLabelsButNoValues(t *testing.T) {
	deployment := v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{Name: "foo"},
	}

	dep := &K8sDeployment{
		deployment: deployment,
		config:     config.Config{VersionLabels: []string{"label1"}},
	}

	assert.Equal(t, "foo", dep.Name())
}

func TestReturnsNameWithLabels(t *testing.T) {
	deployment := v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{Name: "foo", Labels: map[string]string{"label1": "bar"}},
	}

	dep := &K8sDeployment{
		deployment: deployment,
		config:     config.Config{AppNameLabel: "label1"},
	}

	assert.Equal(t, "bar", dep.Name())
}

func TestReturnsVersion(t *testing.T) {
	deployment := v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{Name: "foo", Labels: map[string]string{"app_ver": "1.778"}},
	}

	dep := &K8sDeployment{
		deployment: deployment,
		config:     config.Config{VersionLabels: []string{"version", "app_ver", "conf_ver"}},
	}

	assert.Equal(t, "1.778", dep.Version())
}

func TestReturnsCompositeVersion(t *testing.T) {
	deployment := v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{Name: "foo", Labels: map[string]string{"app_ver": "1.778", "conf_ver": "143"}},
	}

	dep := &K8sDeployment{
		deployment: deployment,
		config:     config.Config{VersionLabels: []string{"version", "app_ver", "conf_ver"}},
	}

	assert.Equal(t, "1.778-143", dep.Version())
}

func TestReturnsDefaultVersion(t *testing.T) {
	deployment := v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{Name: "foo", Labels: map[string]string{"version": "v1.0.20"}},
	}

	dep := &K8sDeployment{
		deployment: deployment,
		config:     config.Config{VersionLabels: []string{"version", "app_ver", "conf_ver"}},
	}

	assert.Equal(t, "v1.0.20", dep.Version())
}

func TestReturnsUnknownVersion(t *testing.T) {
	deployment := v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{Name: "foo", Labels: map[string]string{}},
	}

	dep := &K8sDeployment{
		deployment: deployment,
		config:     config.Config{VersionLabels: []string{"version", "app_ver", "conf_ver"}},
	}

	assert.Equal(t, "Unknown", dep.Version())
}
