// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package common

// Datadog env var names
const (
	DatadogHost                                     = "DATADOG_HOST"
	DDAdmissionControllerEnabled                    = "DD_ADMISSION_CONTROLLER_ENABLED"
	DDAdmissionControllerInjectConfig               = "DD_ADMISSION_CONTROLLER_INJECT_CONFIG_ENABLED"
	DDAdmissionControllerInjectConfigMode           = "DD_ADMISSION_CONTROLLER_INJECT_CONFIG_MODE"
	DDAdmissionControllerInjectTags                 = "DD_ADMISSION_CONTROLLER_INJECT_TAGS_ENABLED"
	DDAdmissionControllerLocalServiceName           = "DD_ADMISSION_CONTROLLER_INJECT_CONFIG_LOCAL_SERVICE_NAME"
	DDAdmissionControllerMutateUnlabelled           = "DD_ADMISSION_CONTROLLER_MUTATE_UNLABELLED"
	DDAdmissionControllerServiceName                = "DD_ADMISSION_CONTROLLER_SERVICE_NAME"
	DDAPIKey                                        = "DD_API_KEY"
	DDAPMEnabled                                    = "DD_APM_ENABLED"
	DDAPMNonLocalTraffic                            = "DD_APM_NON_LOCAL_TRAFFIC"
	DDAPMReceiverPort                               = "DD_APM_RECEIVER_PORT"
	DDAPMReceiverSocket                             = "DD_APM_RECEIVER_SOCKET"
	DDAppKey                                        = "DD_APP_KEY"
	DDAuthTokenFilePath                             = "DD_AUTH_TOKEN_FILE_PATH"
	DDClcRunnerEnabled                              = "DD_CLC_RUNNER_ENABLED"
	DDClcRunnerHost                                 = "DD_CLC_RUNNER_HOST"
	DDClcRunnerID                                   = "DD_CLC_RUNNER_ID"
	DDClusterAgentAuthToken                         = "DD_CLUSTER_AGENT_AUTH_TOKEN"
	DDClusterAgentEnabled                           = "DD_CLUSTER_AGENT_ENABLED"
	DDClusterAgentKubeServiceName                   = "DD_CLUSTER_AGENT_KUBERNETES_SERVICE_NAME"
	DDClusterAgentTokenName                         = "DD_CLUSTER_AGENT_TOKEN_NAME"
	DDClusterChecksEnabled                          = "DD_CLUSTER_CHECKS_ENABLED"
	DDClusterName                                   = "DD_CLUSTER_NAME"
	DDCollectKubernetesEvents                       = "DD_COLLECT_KUBERNETES_EVENTS"
	DDComplianceConfigCheckInterval                 = "DD_COMPLIANCE_CONFIG_CHECK_INTERVAL"
	DDComplianceConfigDir                           = "DD_COMPLIANCE_CONFIG_DIR"
	DDComplianceConfigEnabled                       = "DD_COMPLIANCE_CONFIG_ENABLED"
	DDContainerCollectionEnabled                    = "DD_PROCESS_CONFIG_CONTAINER_COLLECTION_ENABLED"
	DDCriSocketPath                                 = "DD_CRI_SOCKET_PATH"
	DDddURL                                         = "DD_DD_URL"
	DDDogstatsdEnabled                              = "DD_USE_DOGSTATSD"
	DDDogstatsdMapperProfiles                       = "DD_DOGSTATSD_MAPPER_PROFILES"
	DDDogstatsdNonLocalTraffic                      = "DD_DOGSTATSD_NON_LOCAL_TRAFFIC"
	DDDogstatsdOriginDetection                      = "DD_DOGSTATSD_ORIGIN_DETECTION"
	DDDogstatsdPort                                 = "DD_DOGSTATSD_PORT"
	DDDogstatsdSocket                               = "DD_DOGSTATSD_SOCKET"
	DDEnableMetadataCollection                      = "DD_ENABLE_METADATA_COLLECTION"
	DDEnableOOMKillEnvVar                           = "DD_SYSTEM_PROBE_CONFIG_ENABLE_OOM_KILL"
	DDEnableTCPQueueLengthEnvVar                    = "DD_SYSTEM_PROBE_CONFIG_ENABLE_TCP_QUEUE_LENGTH"
	DDExternalMetricsProviderAPIKey                 = "DD_EXTERNAL_METRICS_PROVIDER_API_KEY"
	DDExternalMetricsProviderAppKey                 = "DD_EXTERNAL_METRICS_PROVIDER_APP_KEY"
	DDExternalMetricsProviderEndpoint               = "DD_EXTERNAL_METRICS_PROVIDER_ENDPOINT"
	DDExtraConfigProviders                          = "DD_EXTRA_CONFIG_PROVIDERS"
	DDExtraListeners                                = "DD_EXTRA_LISTENERS"
	DDHealthPort                                    = "DD_HEALTH_PORT"
	DDHostname                                      = "DD_HOSTNAME"
	DDHostRootEnvVar                                = "HOST_ROOT"
	DDIgnoreAutoConf                                = "DD_IGNORE_AUTOCONF"
	DDKubeletCAPath                                 = "DD_KUBELET_CLIENT_CA"
	DDKubeletHost                                   = "DD_KUBERNETES_KUBELET_HOST"
	DDKubeletTLSVerify                              = "DD_KUBELET_TLS_VERIFY"
	DDKubeStateMetricsCoreConfigMap                 = "DD_KUBE_STATE_METRICS_CORE_CONFIGMAP_NAME"
	DDKubeStateMetricsCoreEnabled                   = "DD_KUBE_STATE_METRICS_CORE_ENABLED"
	DDLeaderElection                                = "DD_LEADER_ELECTION"
	DDLeaderLeaseName                               = "DD_LEADER_LEASE_NAME"
	DDLogLevel                                      = "DD_LOG_LEVEL"
	DDLogsConfigContainerCollectAll                 = "DD_LOGS_CONFIG_CONTAINER_COLLECT_ALL"
	DDLogsConfigOpenFilesLimit                      = "DD_LOGS_CONFIG_OPEN_FILES_LIMIT"
	DDLogsContainerCollectUsingFiles                = "DD_LOGS_CONFIG_K8S_CONTAINER_USE_FILE"
	DDLogsEnabled                                   = "DD_LOGS_ENABLED"
	DDMetricsProviderEnabled                        = "DD_EXTERNAL_METRICS_PROVIDER_ENABLED"
	DDMetricsProviderPort                           = "DD_EXTERNAL_METRICS_PROVIDER_PORT"
	DDMetricsProviderUseDatadogMetric               = "DD_EXTERNAL_METRICS_PROVIDER_USE_DATADOGMETRIC_CRD"
	DDMetricsProviderWPAController                  = "DD_EXTERNAL_METRICS_PROVIDER_WPA_CONTROLLER"
	DDNamespaceLabelsAsTags                         = "DD_KUBERNETES_NAMESPACE_LABELS_AS_TAGS"
	DDNodeLabelsAsTags                              = "DD_KUBERNETES_NODE_LABELS_AS_TAGS"
	DDOrchestratorExplorerEnabled                   = "DD_ORCHESTRATOR_EXPLORER_ENABLED"
	DDOrchestratorExplorerExtraTags                 = "DD_ORCHESTRATOR_EXPLORER_EXTRA_TAGS"
	DDOrchestratorExplorerDDUrl                     = "DD_ORCHESTRATOR_EXPLORER_ORCHESTRATOR_DD_URL"
	DDOrchestratorExplorerAdditionalEndpoints       = "DD_ORCHESTRATOR_ADDITIONAL_ENDPOINTS"
	DDOrchestratorExplorerContainerScrubbingEnabled = "DD_ORCHESTRATOR_EXPLORER_CONTAINER_SCRUBBING_ENABLED"
	DDPodAnnotationsAsTags                          = "DD_KUBERNETES_POD_ANNOTATIONS_AS_TAGS"
	DDPodLabelsAsTags                               = "DD_KUBERNETES_POD_LABELS_AS_TAGS"
	DDPPMReceiverSocket                             = "DD_APM_RECEIVER_SOCKET"
	DDProcessCollectionEnabled                      = "DD_PROCESS_CONFIG_PROCESS_COLLECTION_ENABLED"
	DDProcessConfigScrubArgs                        = "DD_PROCESS_CONFIG_SCRUB_ARGS"
	DDProcessConfigStripArgs                        = "DD_PROCESS_CONFIG_STRIP_PROC_ARGUMENTS"
	DDPrometheusScrapeChecks                        = "DD_PROMETHEUS_SCRAPE_CHECKS"
	DDPrometheusScrapeEnabled                       = "DD_PROMETHEUS_SCRAPE_ENABLED"
	DDPrometheusScrapeServiceEndpoints              = "DD_PROMETHEUS_SCRAPE_SERVICE_ENDPOINTS"
	DDRuntimeSecurityConfigEnabled                  = "DD_RUNTIME_SECURITY_CONFIG_ENABLED"
	DDRuntimeSecurityConfigPoliciesDir              = "DD_RUNTIME_SECURITY_CONFIG_POLICIES_DIR"
	DDRuntimeSecurityConfigRemoteTaggerEnabled      = "DD_RUNTIME_SECURITY_CONFIG_REMOTE_TAGGER"
	DDRuntimeSecurityConfigSocket                   = "DD_RUNTIME_SECURITY_CONFIG_SOCKET"
	DDRuntimeSecurityConfigSyscallMonitorEnabled    = "DD_RUNTIME_SECURITY_CONFIG_SYSCALL_MONITOR_ENABLED"
	DDSecretBackendCommand                          = "DD_SECRET_BACKEND_COMMAND"
	DDSite                                          = "DD_SITE"
	DDSystemProbeAgentEnabled                       = "DD_SYSTEM_PROBE_ENABLED"
	DDSystemProbeBPFDebugEnabled                    = DDSystemProbeEnvPrefix + "BPF_DEBUG"
	DDSystemProbeCollectDNSStatsEnabled             = "DD_COLLECT_DNS_STATS"
	DDSystemProbeConntrackEnabled                   = DDSystemProbeEnvPrefix + "ENABLE_CONNTRACK"
	DDSystemProbeDebugPort                          = DDSystemProbeEnvPrefix + "DEBUG_PORT"
	DDSystemProbeEnabled                            = "DD_SYSTEM_PROBE_ENABLED"
	DDSystemProbeEnvPrefix                          = "DD_SYSTEM_PROBE_CONFIG_"
	DDSystemProbeExternal                           = "DD_SYSTEM_PROBE_EXTERNAL"
	DDSystemProbeNPMEnabled                         = "DD_SYSTEM_PROBE_NETWORK_ENABLED"
	DDSystemProbeServiceMonitoringEnabled           = "DD_SYSTEM_PROBE_SERVICE_MONITORING_ENABLED"
	DDSystemProbeSocket                             = "DD_SYSPROBE_SOCKET"
	DDSystemProbeOOMKillEnabled                     = DDSystemProbeEnvPrefix + "ENABLE_OOM_KILL"
	DDSystemProbeTCPQueueLengthEnabled              = DDSystemProbeEnvPrefix + "ENABLE_TCP_QUEUE_LENGTH"
	DDTags                                          = "DD_TAGS"
	DockerHost                                      = "DOCKER_HOST"

	DDOTLPgRPCEndpoint = "DD_OTLP_CONFIG_RECEIVER_PROTOCOLS_GRPC_ENDPOINT"
	DDOTLPHTTPEndpoint = "DD_OTLP_CONFIG_RECEIVER_PROTOCOLS_HTTP_ENDPOINT"

	// KubernetesEnvvarName Env var used by the Datadog Agent container entrypoint
	// to add kubelet config provider and listener
	KubernetesEnvVar = "KUBERNETES"

	ClusterChecksConfigProvider = "clusterchecks"
)
