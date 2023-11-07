// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package chime

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	endpoints_sdkv1 "github.com/aws/aws-sdk-go/aws/endpoints"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	chime_sdkv1 "github.com/aws/aws-sdk-go/service/chime"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
	"log"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceVoiceConnector,
			TypeName: "aws_chime_voice_connector",
			Name:     "Voice Connector",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceVoiceConnectorGroup,
			TypeName: "aws_chime_voice_connector_group",
		},
		{
			Factory:  ResourceVoiceConnectorLogging,
			TypeName: "aws_chime_voice_connector_logging",
		},
		{
			Factory:  ResourceVoiceConnectorOrigination,
			TypeName: "aws_chime_voice_connector_origination",
		},
		{
			Factory:  ResourceVoiceConnectorStreaming,
			TypeName: "aws_chime_voice_connector_streaming",
		},
		{
			Factory:  ResourceVoiceConnectorTermination,
			TypeName: "aws_chime_voice_connector_termination",
		},
		{
			Factory:  ResourceVoiceConnectorTerminationCredentials,
			TypeName: "aws_chime_voice_connector_termination_credentials",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Chime
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*chime_sdkv1.Chime, error) {
	sess := config[names.AttrSession].(*session_sdkv1.Session)

	if endpoint := config[names.AttrEndpoint].(string); endpoint != "" && sess.Config.UseFIPSEndpoint == endpoints_sdkv1.FIPSEndpointStateEnabled {
		// The SDK doesn't allow setting a custom non-FIPS endpoint *and* enabling UseFIPSEndpoint.
		// However there are a few cases where this is necessary; some services don't have FIPS endpoints,
		// and for some services (e.g. CloudFront) the SDK generates the wrong fips endpoint.
		// While forcing this to disabled may result in the end-user not using a FIPS endpoint as specified
		// by setting UseFIPSEndpoint=true in the provider, the user also explicitly changed the endpoint, so
		// here we need to assume the user knows what they're doing.
		log.Printf("[WARN] UseFIPSEndpoint is enabled but a custom endpoint (%s) is configured, ignoring UseFIPSEndpoint.", endpoint)
		sess.Config.UseFIPSEndpoint = endpoints_sdkv1.FIPSEndpointStateDisabled
	}

	return chime_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config[names.AttrEndpoint].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
