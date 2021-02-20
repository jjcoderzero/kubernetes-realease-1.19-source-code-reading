package v1beta1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	clientauthentication "k8s.io/client-go/pkg/apis/clientauthentication"
)

func Convert_clientauthentication_ExecCredentialSpec_To_v1beta1_ExecCredentialSpec(in *clientauthentication.ExecCredentialSpec, out *ExecCredentialSpec, s conversion.Scope) error {
	return nil
}
