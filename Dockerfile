###############################################################################

FROM golang:1.19 as builder

WORKDIR /go/src/polaris-exporter

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /go/bin/polaris-exporter

###############################################################################

FROM ubuntu:latest

# Required to allow init function to run => https://github.com/fairwindsops/polaris/blob/7.1.5/pkg/config/checks.go#L73
RUN mkdir -p /checks && for f in automountServiceAccountToken.yaml clusterrolePodExecAttach.yaml clusterrolebindingClusterAdmin.yaml clusterrolebindingPodExecAttach.yaml cpuLimitsMissing.yaml cpuRequestsMissing.yaml dangerousCapabilities.yaml deploymentMissingReplicas.yaml hostIPCSet.yaml hostNetworkSet.yaml hostPIDSet.yaml hostPortSet.yaml insecureCapabilities.yaml linuxHardening.yaml livenessProbeMissing.yaml memoryLimitsMissing.yaml memoryRequestsMissing.yaml metadataAndNameMismatched.yaml missingNetworkPolicy.yaml missingPodDisruptionBudget.yaml notReadOnlyRootFilesystem.yaml pdbDisruptionsIsZero.yaml priorityClassNotSet.yaml privilegeEscalationAllowed.yaml pullPolicyNotAlways.yaml readinessProbeMissing.yaml rolePodExecAttach.yaml rolebindingClusterAdminClusterRole.yaml rolebindingClusterAdminRole.yaml rolebindingClusterRolePodExecAttach.yaml rolebindingRolePodExecAttach.yaml runAsPrivileged.yaml runAsRootAllowed.yaml sensitiveConfigmapContent.yaml sensitiveContainerEnvVar.yaml tagNotSpecified.yaml tlsSettingsMissing.yaml; do touch "/checks/${f}"; done

COPY --from=builder /go/bin/polaris-exporter /usr/local/polaris-exporter

CMD ["/usr/local/polaris-exporter"]

###############################################################################