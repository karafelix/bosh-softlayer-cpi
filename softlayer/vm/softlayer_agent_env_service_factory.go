package vm

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

type SoftLayerAgentEnvServiceFactory struct {
	agentEnvService string
	registryOptions RegistryOptions
	logger          boshlog.Logger
}

func NewSoftLayerAgentEnvServiceFactory(
	agentEnvService string,
	registryOptions RegistryOptions,
	logger boshlog.Logger,
) SoftLayerAgentEnvServiceFactory {
	return SoftLayerAgentEnvServiceFactory{
		logger:          logger,
		agentEnvService: agentEnvService,
		registryOptions: registryOptions,
	}
}

func (f SoftLayerAgentEnvServiceFactory) New(
	softlayerFileService SoftlayerFileService,
	instanceID string,
) AgentEnvService {
	if f.agentEnvService == "registry" {
		return NewRegistryAgentEnvService(f.registryOptions, instanceID, f.logger)
	}
	return NewFSAgentEnvService(softlayerFileService, f.logger)
}
