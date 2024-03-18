package v1beta1

import esmeta "github.com/external-secrets/external-secrets/apis/meta/v1"

type BeyondTrustProviderSecretRef struct {

	// Value can be specified directly to set a value without using a secret.
	// +optional
	Value string `json:"value,omitempty"`

	// SecretRef references a key in a secret that will be used as value.
	// +optional
	SecretRef *esmeta.SecretKeySelector `json:"secretRef,omitempty"`
}

// TODOBT
type BeyondtrustProvider struct {
	Host     string `json:"host"`
	Clientid string `json:"clientid"`
	// ClientSecret is the secret part of the credential.
	Clientsecret  string `json:"clientsecret"`
	Retrievaltype string `json:"retrievaltype"`
}
