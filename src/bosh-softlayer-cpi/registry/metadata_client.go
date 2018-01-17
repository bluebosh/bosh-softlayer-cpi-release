package registry

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/softlayer/softlayer-go/sl"

	"bosh-softlayer-cpi/api"
	"bosh-softlayer-cpi/logger"
	"bosh-softlayer-cpi/softlayer/client"
)

const metadataClientLogTag = "RegistryMetadataClient"

// HTTPClient represents a BOSH Registry Client.
type MetadataClient struct {
	options         ClientOptions
	logger          logger.Logger
	softlayerClient client.Client
}

// NewHTTPClient creates a new BOSH Registry Client.
func NewMetadataClient(
	options ClientOptions,
	logger logger.Logger,
	softlayerClient client.Client,
) MetadataClient {
	return MetadataClient{
		options:         options,
		logger:          logger,
		softlayerClient: softlayerClient,
	}
}

// Delete deletes the instance settings for a given instance ID. This is
// deleted when the VM is terminated.
func (c MetadataClient) Delete(instanceID string) error {
	return nil
}

// Fetch gets the agent settings for a given instance ID.
func (c MetadataClient) Fetch(instanceID string) (AgentSettings, error) {
	var agentSettings AgentSettings

	cid, err := strconv.Atoi(instanceID)
	if err != nil {
		return agentSettings, bosherr.WrapErrorf(err, "Convert id %q to type int.", agentSettings)
	}
	settingsContents, err := c.softlayerClient.GetInstanceMetadata(cid)
	if err != nil {
		return AgentSettings{}, err
	}

	settingsBytesWithoutQuotes := strings.Replace(string(settingsContents), `"`, ``, -1)
	decodedSettings, err := base64.RawURLEncoding.DecodeString(settingsBytesWithoutQuotes)
	if err != nil {
		return agentSettings, bosherr.WrapError(err, "Decoding url encoded user data")
	}

	err = json.Unmarshal([]byte(decodedSettings), &agentSettings)
	if err != nil {
		return agentSettings, bosherr.WrapErrorf(err, "Parsing metadata settings from %q", decodedSettings)
	}

	return agentSettings, nil
}

// Update updates the agent settings for a given instance ID. If there are not already agent settings for the instance, it will create ones.
func (c MetadataClient) Update(instanceID string, agentSettings AgentSettings) error {
	settingsJSON, err := json.Marshal(agentSettings)
	if err != nil {
		return bosherr.WrapErrorf(err, "Marshalling agent settings, contents: '%#v", agentSettings)
	}
	c.logger.Debug(metadataClientLogTag, "Updating instance metadata for %q with agent settings %q", instanceID, settingsJSON)

	cid, err := strconv.Atoi(instanceID)
	if err != nil {
		return bosherr.WrapErrorf(err, "Convert id %q to type int.", agentSettings)
	}

	settingsEncodeString := base64.RawURLEncoding.EncodeToString(settingsJSON)
	found, err := c.softlayerClient.SetInstanceMetadata(cid, sl.String(settingsEncodeString))
	if err != nil {
		return bosherr.WrapErrorf(err, "Set Instance Metadata for %q with contents: '%s'", instanceID, settingsEncodeString)
	}

	if !found {
		return bosherr.WrapErrorf(api.NewVMNotFoundError(instanceID), "Set Instance Metadata for %q with contents: '%#v", instanceID, agentSettings)
	}
	c.logger.Debug(httpClientLogTag, "Updated registry endpoint '%s' with agent settings '%s'", instanceID, settingsJSON)
	return nil
}
