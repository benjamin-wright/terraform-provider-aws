// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package cognitoidp

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	endpoints_sdkv1 "github.com/aws/aws-sdk-go/aws/endpoints"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	cognitoidentityprovider_sdkv1 "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
	"log"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory: newUserGroupDataSource,
			Name:    "User Group",
		},
		{
			Factory: newUserGroupsDataSource,
			Name:    "User Groups",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newManagedUserPoolClientResource,
			Name:    "Managed User Pool Client",
		},
		{
			Factory: newUserPoolClientResource,
			Name:    "User Pool Client",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceUserPoolClient,
			TypeName: "aws_cognito_user_pool_client",
			Name:     "User Pool Client",
		},
		{
			Factory:  dataSourceUserPoolClients,
			TypeName: "aws_cognito_user_pool_clients",
			Name:     "User Pool Clients",
		},
		{
			Factory:  dataSourceUserPoolSigningCertificate,
			TypeName: "aws_cognito_user_pool_signing_certificate",
			Name:     "User Pool Signing Certificate",
		},
		{
			Factory:  dataSourceUserPools,
			TypeName: "aws_cognito_user_pools",
			Name:     "User Pools",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceIdentityProvider,
			TypeName: "aws_cognito_identity_provider",
			Name:     "Identity Provider",
		},
		{
			Factory:  resourceResourceServer,
			TypeName: "aws_cognito_resource_server",
			Name:     "Resource Server",
		},
		{
			Factory:  resourceRiskConfiguration,
			TypeName: "aws_cognito_risk_configuration",
			Name:     "Risk Configuration",
		},
		{
			Factory:  resourceUser,
			TypeName: "aws_cognito_user",
			Name:     "User",
		},
		{
			Factory:  resourceUserGroup,
			TypeName: "aws_cognito_user_group",
			Name:     "User Group",
		},
		{
			Factory:  resourceUserInGroup,
			TypeName: "aws_cognito_user_in_group",
			Name:     "Group User",
		},
		{
			Factory:  resourceUserPool,
			TypeName: "aws_cognito_user_pool",
			Name:     "User Pool",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  resourceUserPoolDomain,
			TypeName: "aws_cognito_user_pool_domain",
			Name:     "User Pool Domain",
		},
		{
			Factory:  resourceUserPoolUICustomization,
			TypeName: "aws_cognito_user_pool_ui_customization",
			Name:     "User Pool UI Customization",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.CognitoIDP
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*cognitoidentityprovider_sdkv1.CognitoIdentityProvider, error) {
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

	return cognitoidentityprovider_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config[names.AttrEndpoint].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
