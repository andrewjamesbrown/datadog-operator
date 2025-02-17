// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package clusterchecks

import (
	"fmt"
	"testing"

	v2alpha1test "github.com/DataDog/datadog-operator/apis/datadoghq/v2alpha1/test"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	apicommon "github.com/DataDog/datadog-operator/apis/datadoghq/common"
	apicommonv1 "github.com/DataDog/datadog-operator/apis/datadoghq/common/v1"
	"github.com/DataDog/datadog-operator/apis/datadoghq/v1alpha1"
	"github.com/DataDog/datadog-operator/apis/datadoghq/v2alpha1"
	apiutils "github.com/DataDog/datadog-operator/apis/utils"
	"github.com/DataDog/datadog-operator/controllers/datadogagent/feature"
	"github.com/DataDog/datadog-operator/controllers/datadogagent/feature/fake"
	"github.com/DataDog/datadog-operator/controllers/datadogagent/feature/test"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

func TestClusterChecksFeature(t *testing.T) {
	tests := test.FeatureTestSuite{
		//////////////////////////
		// v1Alpha1.DatadogAgent
		//////////////////////////
		{
			Name:          "v1alpha1 cluster checks not enabled and runners not enabled",
			DDAv1:         newV1Agent(false, false),
			WantConfigure: false,
		},
		{
			Name:          "v1alpha1 cluster checks not enabled and runners enabled",
			DDAv1:         newV1Agent(false, true),
			WantConfigure: false,
		},
		{
			Name:          "v1alpha1 cluster checks enabled and runners not enabled",
			DDAv1:         newV1Agent(true, false),
			WantConfigure: true,
			ClusterAgent:  test.NewDefaultComponentTest().WithWantFunc(wantClusterAgentHasExpectedEnvs),
			Agent:         testAgentHasExpectedEnvsWithNoRunners(apicommonv1.CoreAgentContainerName),
		},
		{
			Name:                "v1alpha1 cluster checks enabled and runners enabled",
			DDAv1:               newV1Agent(true, true),
			WantConfigure:       true,
			ClusterAgent:        test.NewDefaultComponentTest().WithWantFunc(wantClusterAgentHasExpectedEnvs),
			ClusterChecksRunner: testClusterChecksRunnerHasExpectedEnvs(),
			Agent:               testAgentHasExpectedEnvsWithRunners(apicommonv1.CoreAgentContainerName),
		},

		//////////////////////////
		// v2Alpha1.DatadogAgent
		//////////////////////////
		{
			Name: "v2alpha1 cluster checks empty, checksum set",
			DDAv2: &v2alpha1.DatadogAgent{
				Spec: v2alpha1.DatadogAgentSpec{
					Features: &v2alpha1.DatadogFeatures{
						ClusterChecks: &v2alpha1.ClusterChecksFeatureConfig{},
					},
				},
			},
			ClusterAgent:  test.NewDefaultComponentTest().WithWantFunc(wantClusterAgentHasNonEmptyChecksumAnnotation),
			WantConfigure: false,
		},
		{
			Name: "v2alpha1 cluster checks not enabled and runners not enabled",
			DDAv2: v2alpha1test.NewDatadogAgentBuilder().
				WithClusterChecksEnabled(false).
				WithClusterChecksUseCLCEnabled(false).
				Build(),
			ClusterAgent:  test.NewDefaultComponentTest().WithWantFunc(wantClusterAgentHasNonEmptyChecksumAnnotation),
			WantConfigure: false,
		},
		{
			Name: "v2alpha1 cluster checks not enabled and runners enabled",
			DDAv2: v2alpha1test.NewDatadogAgentBuilder().
				WithClusterChecksEnabled(false).
				WithClusterChecksUseCLCEnabled(true).
				Build(),
			ClusterAgent:  test.NewDefaultComponentTest().WithWantFunc(wantClusterAgentHasNonEmptyChecksumAnnotation),
			WantConfigure: false,
		},
		{
			Name: "v2alpha1 cluster checks enabled and runners not enabled",
			DDAv2: v2alpha1test.NewDatadogAgentBuilder().
				WithClusterChecksEnabled(true).
				WithClusterChecksUseCLCEnabled(false).
				Build(),
			WantConfigure: true,
			ClusterAgent:  test.NewDefaultComponentTest().WithWantFunc(wantClusterAgentHasExpectedEnvsAndChecksum),
			Agent:         testAgentHasExpectedEnvsWithNoRunners(apicommonv1.CoreAgentContainerName),
		},
		{
			Name: "v2alpha1 cluster checks enabled and runners not enabled with multi-process container",
			DDAv2: v2alpha1test.NewDatadogAgentBuilder().
				WithClusterChecksEnabled(true).
				WithClusterChecksUseCLCEnabled(false).
				WithMultiProcessContainer(true).
				Build(),
			WantConfigure: true,
			ClusterAgent:  test.NewDefaultComponentTest().WithWantFunc(wantClusterAgentHasExpectedEnvsAndChecksum),
			Agent:         testAgentHasExpectedEnvsWithNoRunners(apicommonv1.UnprivilegedMultiProcessAgentContainerName),
		},
		{
			Name: "v2alpha1 cluster checks enabled and runners enabled",
			DDAv2: v2alpha1test.NewDatadogAgentBuilder().
				WithClusterChecksEnabled(true).
				WithClusterChecksUseCLCEnabled(true).
				Build(),
			WantConfigure:       true,
			ClusterAgent:        test.NewDefaultComponentTest().WithWantFunc(wantClusterAgentHasExpectedEnvsAndChecksum),
			ClusterChecksRunner: testClusterChecksRunnerHasExpectedEnvs(),
			Agent:               testAgentHasExpectedEnvsWithRunners(apicommonv1.CoreAgentContainerName),
		},
		{
			Name: "v2alpha1 cluster checks enabled and runners enabled with multi-process container",
			DDAv2: v2alpha1test.NewDatadogAgentBuilder().
				WithClusterChecksEnabled(true).
				WithClusterChecksUseCLCEnabled(true).
				WithMultiProcessContainer(true).
				Build(),
			WantConfigure:       true,
			ClusterAgent:        test.NewDefaultComponentTest().WithWantFunc(wantClusterAgentHasExpectedEnvsAndChecksum),
			ClusterChecksRunner: testClusterChecksRunnerHasExpectedEnvs(),
			Agent:               testAgentHasExpectedEnvsWithRunners(apicommonv1.UnprivilegedMultiProcessAgentContainerName),
		},
	}

	tests.Run(t, buildClusterChecksFeature)
}

func TestClusterAgentChecksumsDifferentForDifferentConfig(t *testing.T) {
	logf.SetLogger(zap.New(zap.UseDevMode(true)))
	logger := logf.Log.WithName("checksum unique")

	annotationKey := fmt.Sprintf(apicommon.MD5ChecksumAnnotationKey, feature.ClusterChecksIDType)
	feature := buildClusterChecksFeature(&feature.Options{
		Logger: logger,
	})

	podTemplateManager := fake.NewPodTemplateManagers(t, corev1.PodTemplateSpec{})
	md5Values := map[string]string{}

	datadogAgents := []*v2alpha1.DatadogAgent{
		{
			Spec: v2alpha1.DatadogAgentSpec{
				Features: &v2alpha1.DatadogFeatures{
					ClusterChecks: &v2alpha1.ClusterChecksFeatureConfig{},
				},
			},
		},
		v2alpha1test.NewDatadogAgentBuilder().
			WithClusterChecksEnabled(false).
			WithClusterChecksUseCLCEnabled(false).
			Build(),
		v2alpha1test.NewDatadogAgentBuilder().
			WithClusterChecksEnabled(false).
			WithClusterChecksUseCLCEnabled(true).
			Build(),
		v2alpha1test.NewDatadogAgentBuilder().
			WithClusterChecksEnabled(true).
			WithClusterChecksUseCLCEnabled(false).
			Build(),
		v2alpha1test.NewDatadogAgentBuilder().
			WithClusterChecksEnabled(true).
			WithClusterChecksUseCLCEnabled(true).
			Build(),
	}

	for _, datadogAgent := range datadogAgents {
		feature.Configure(datadogAgent)
		feature.ManageClusterAgent(podTemplateManager)
		md5 := podTemplateManager.AnnotationMgr.Annotations[annotationKey]
		md5Values[md5] = ""
	}

	// First three cases, when cluster checks is disabled md5 is empty string
	assert.Equal(t, 3, len(md5Values))
}

func newV1Agent(enableClusterChecks bool, enableClusterCheckRunners bool) *v1alpha1.DatadogAgent {
	return &v1alpha1.DatadogAgent{
		Spec: v1alpha1.DatadogAgentSpec{
			ClusterAgent: v1alpha1.DatadogAgentSpecClusterAgentSpec{
				Config: &v1alpha1.ClusterAgentConfig{
					ClusterChecksEnabled: apiutils.NewBoolPointer(enableClusterChecks),
				},
			},
			ClusterChecksRunner: v1alpha1.DatadogAgentSpecClusterChecksRunnerSpec{
				Enabled: apiutils.NewBoolPointer(enableClusterCheckRunners),
			},
		},
	}
}

func wantClusterAgentHasExpectedEnvsAndChecksum(t testing.TB, mgrInterface feature.PodTemplateManagers) {
	wantClusterAgentHasExpectedEnvs(t, mgrInterface)
	wantClusterAgentHasNonEmptyChecksumAnnotation(t, mgrInterface)
}

func wantClusterAgentHasExpectedEnvs(t testing.TB, mgrInterface feature.PodTemplateManagers) {
	mgr := mgrInterface.(*fake.PodTemplateManagers)

	clusterAgentEnvs := mgr.EnvVarMgr.EnvVarsByC[apicommonv1.ClusterAgentContainerName]
	expectedClusterAgentEnvs := []*corev1.EnvVar{
		{
			Name:  apicommon.DDClusterChecksEnabled,
			Value: "true",
		},
		{
			Name:  apicommon.DDExtraConfigProviders,
			Value: apicommon.KubeServicesAndEndpointsConfigProviders,
		},
		{
			Name:  apicommon.DDExtraListeners,
			Value: apicommon.KubeServicesAndEndpointsListeners,
		},
	}

	assert.True(
		t,
		apiutils.IsEqualStruct(clusterAgentEnvs, expectedClusterAgentEnvs),
		"Cluster Agent ENVs \ndiff = %s", cmp.Diff(clusterAgentEnvs, expectedClusterAgentEnvs),
	)
}

func wantClusterAgentHasNonEmptyChecksumAnnotation(t testing.TB, mgrInterface feature.PodTemplateManagers) {
	mgr := mgrInterface.(*fake.PodTemplateManagers)
	annotationKey := fmt.Sprintf(apicommon.MD5ChecksumAnnotationKey, feature.ClusterChecksIDType)
	annotations := mgr.AnnotationMgr.Annotations
	assert.NotEmpty(t, annotations[annotationKey])
}

func testClusterChecksRunnerHasExpectedEnvs() *test.ComponentTest {
	return test.NewDefaultComponentTest().WithWantFunc(
		func(t testing.TB, mgrInterface feature.PodTemplateManagers) {
			mgr := mgrInterface.(*fake.PodTemplateManagers)

			clusterRunnerEnvs := mgr.EnvVarMgr.EnvVarsByC[apicommonv1.ClusterChecksRunnersContainerName]
			expectedClusterRunnerEnvs := []*corev1.EnvVar{
				{
					Name:  apicommon.DDClusterChecksEnabled,
					Value: "true",
				},
				{
					Name:  apicommon.DDExtraConfigProviders,
					Value: apicommon.ClusterChecksConfigProvider,
				},
			}

			assert.True(
				t,
				apiutils.IsEqualStruct(clusterRunnerEnvs, expectedClusterRunnerEnvs),
				"Cluster Runner ENVs \ndiff = %s", cmp.Diff(clusterRunnerEnvs, expectedClusterRunnerEnvs),
			)
		},
	)
}

func testAgentHasExpectedEnvsWithRunners(agentContainerName apicommonv1.AgentContainerName) *test.ComponentTest {
	return test.NewDefaultComponentTest().WithWantFunc(
		func(t testing.TB, mgrInterface feature.PodTemplateManagers) {
			mgr := mgrInterface.(*fake.PodTemplateManagers)

			agentEnvs := mgr.EnvVarMgr.EnvVarsByC[agentContainerName]
			expectedAgentEnvs := []*corev1.EnvVar{
				{
					Name:  apicommon.DDExtraConfigProviders,
					Value: apicommon.EndpointsChecksConfigProvider,
				},
			}

			assert.True(
				t,
				apiutils.IsEqualStruct(agentEnvs, expectedAgentEnvs),
				"Cluster Runner ENVs \ndiff = %s", cmp.Diff(agentEnvs, expectedAgentEnvs),
			)
		},
	)
}

func testAgentHasExpectedEnvsWithNoRunners(agentContainerName apicommonv1.AgentContainerName) *test.ComponentTest {
	return test.NewDefaultComponentTest().WithWantFunc(
		func(t testing.TB, mgrInterface feature.PodTemplateManagers) {
			mgr := mgrInterface.(*fake.PodTemplateManagers)

			agentEnvs := mgr.EnvVarMgr.EnvVarsByC[agentContainerName]
			expectedAgentEnvs := []*corev1.EnvVar{
				{
					Name:  apicommon.DDExtraConfigProviders,
					Value: apicommon.ClusterAndEndpointsConfigProviders,
				},
			}

			assert.True(
				t,
				apiutils.IsEqualStruct(agentEnvs, expectedAgentEnvs),
				"Cluster Runner ENVs \ndiff = %s", cmp.Diff(agentEnvs, expectedAgentEnvs),
			)
		},
	)
}
