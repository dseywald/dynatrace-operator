package status

import (
	dynatracev1alpha1 "github.com/Dynatrace/dynatrace-operator/api/v1alpha1"
	"github.com/Dynatrace/dynatrace-operator/controllers/kubesystem"
	"github.com/Dynatrace/dynatrace-operator/dtclient"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Options struct {
	dtc       dtclient.Client
	apiClient client.Client
}

func SetDynakubeStatus(instance *dynatracev1alpha1.DynaKube, opts Options) error {
	clt := opts.apiClient
	dtc := opts.dtc

	uid, err := kubesystem.GetUID(clt)
	if err != nil {
		return errors.WithStack(err)
	}

	communicationHost, err := dtc.GetCommunicationHostForClient()
	if err != nil {
		return errors.WithStack(err)
	}

	connectionInfo, err := dtc.GetConnectionInfo()
	if err != nil {
		return errors.WithStack(err)
	}

	communicationHostStatus := dynatracev1alpha1.CommunicationHostStatus{
		Protocol: communicationHost.Protocol,
		Host:     communicationHost.Host,
		Port:     communicationHost.Port,
	}

	connectionInfoStatus := dynatracev1alpha1.ConnectionInfoStatus{
		CommunicationHosts: communicationHostsToStatus(connectionInfo.CommunicationHosts),
		TenantUUID:         connectionInfo.TenantUUID,
	}

	instance.Status.KubeSystemUUID = string(uid)
	instance.Status.CommunicationHostForClient = communicationHostStatus
	instance.Status.ConnectionInfo = connectionInfoStatus
	instance.Status.EnvironmentID = connectionInfo.TenantUUID

	return nil
}

func communicationHostsToStatus(communicationHosts []dtclient.CommunicationHost) []dynatracev1alpha1.CommunicationHostStatus {
	var communicationHostStatuses []dynatracev1alpha1.CommunicationHostStatus

	for _, communicationHost := range communicationHosts {
		communicationHostStatuses = append(communicationHostStatuses, dynatracev1alpha1.CommunicationHostStatus{
			Protocol: communicationHost.Protocol,
			Host:     communicationHost.Host,
			Port:     communicationHost.Port,
		})
	}

	return communicationHostStatuses
}
