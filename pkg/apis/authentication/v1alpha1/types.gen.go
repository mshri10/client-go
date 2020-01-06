// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by kubetype-gen. DO NOT EDIT.

package v1alpha1

import (
	authenticationv1alpha1 "istio.io/api/authentication/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Policy defines what authentication methods can be accepted on workload(s),
// and if authenticated, which method/certificate will set the request principal
// (i.e request.auth.principal attribute).
//
// Authentication policy is composed of 2-part authentication:
// - peer: verify caller service credentials. This part will set source.user
// (peer identity).
// - origin: verify the origin credentials. This part will set request.auth.user
// (origin identity), as well as other attributes like request.auth.presenter,
// request.auth.audiences and raw claims. Note that the identity could be
// end-user, service account, device etc.
//
// Last but not least, the principal binding rule defines which identity (peer
// or origin) should be used as principal. By default, it uses peer.
//
// Examples:
//
// Policy to enable mTLS for all services in namespace frod. The policy name must be
// `default`, and it contains no rule for `targets`.
//
// ```yaml
// apiVersion: authentication.istio.io/v1alpha1
// kind: Policy
// metadata:
//   name: default
//   namespace: frod
// spec:
//   peers:
//   - mtls:
// ```
// Policy to disable mTLS for "productpage" service
//
// ```yaml
// apiVersion: authentication.istio.io/v1alpha1
// kind: Policy
// metadata:
//   name: productpage-mTLS-disable
//   namespace: frod
// spec:
//   targets:
//   - name: productpage
// ```
// Policy to require mTLS for peer authentication, and JWT for origin authentication
// for productpage:9000 except the path '/health_check' . Principal is set from origin identity.
//
// ```yaml
// apiVersion: authentication.istio.io/v1alpha1
// kind: Policy
// metadata:
//   name: productpage-mTLS-with-JWT
//   namespace: frod
// spec:
//   targets:
//   - name: productpage
//     ports:
//     - number: 9000
//   peers:
//   - mtls:
//   origins:
//   - jwt:
//       issuer: "https://securetoken.google.com"
//       audiences:
//       - "productpage"
//       jwksUri: "https://www.googleapis.com/oauth2/v1/certs"
//       jwtHeaders:
//       - "x-goog-iap-jwt-assertion"
//       triggerRules:
//       - excludedPaths:
//         - exact: /health_check
//   principalBinding: USE_ORIGIN
// ```
//
// <!-- crd generation tags
// +cue-gen:Policy:groupName:authentication.istio.io
// +cue-gen:Policy:version:v1alpha1
// +cue-gen:Policy:storageVersion
// +cue-gen:Policy:annotations:helm.sh/resource-policy=keep
// +cue-gen:Policy:labels:app=istio-citadel,chart=istio,heritage=Tiller,release=istio
// +cue-gen:Policy:subresource:status
// +cue-gen:Policy:scope:Namespaced
// +cue-gen:Policy:resource:categories=istio-io,authentication-istio-io,plural=policies
// -->
//
// <!-- crd generation tags
// +cue-gen:MeshPolicy:groupName:authentication.istio.io
// +cue-gen:MeshPolicy:version:v1alpha1
// +cue-gen:MeshPolicy:storageVersion
// +cue-gen:MeshPolicy:annotations:helm.sh/resource-policy=keep
// +cue-gen:MeshPolicy:labels:app=istio-citadel,chart=istio,heritage=Tiller,release=istio
// +cue-gen:MeshPolicy:subresource:status
// +cue-gen:MeshPolicy:scope:Cluster
// +cue-gen:MeshPolicy:resource:categories=istio-io,authentication-istio-io,plural=meshpolicies
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=authentication.istio.io/v1alpha1
// +kubetype-gen:kubeType=Policy
// +kubetype-gen:kubeType=MeshPolicy
// +kubetype-gen:MeshPolicy:tag=genclient:nonNamespaced
// +genclient
// +k8s:deepcopy-gen=true
// -->
type Policy struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec authenticationv1alpha1.Policy `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PolicyList is a collection of Policies.
type PolicyList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []Policy `json:"items" protobuf:"bytes,2,rep,name=items"`
}

//
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Policy defines what authentication methods can be accepted on workload(s),
// and if authenticated, which method/certificate will set the request principal
// (i.e request.auth.principal attribute).
//
// Authentication policy is composed of 2-part authentication:
// - peer: verify caller service credentials. This part will set source.user
// (peer identity).
// - origin: verify the origin credentials. This part will set request.auth.user
// (origin identity), as well as other attributes like request.auth.presenter,
// request.auth.audiences and raw claims. Note that the identity could be
// end-user, service account, device etc.
//
// Last but not least, the principal binding rule defines which identity (peer
// or origin) should be used as principal. By default, it uses peer.
//
// Examples:
//
// Policy to enable mTLS for all services in namespace frod. The policy name must be
// `default`, and it contains no rule for `targets`.
//
// ```yaml
// apiVersion: authentication.istio.io/v1alpha1
// kind: Policy
// metadata:
//   name: default
//   namespace: frod
// spec:
//   peers:
//   - mtls:
// ```
// Policy to disable mTLS for "productpage" service
//
// ```yaml
// apiVersion: authentication.istio.io/v1alpha1
// kind: Policy
// metadata:
//   name: productpage-mTLS-disable
//   namespace: frod
// spec:
//   targets:
//   - name: productpage
// ```
// Policy to require mTLS for peer authentication, and JWT for origin authentication
// for productpage:9000 except the path '/health_check' . Principal is set from origin identity.
//
// ```yaml
// apiVersion: authentication.istio.io/v1alpha1
// kind: Policy
// metadata:
//   name: productpage-mTLS-with-JWT
//   namespace: frod
// spec:
//   targets:
//   - name: productpage
//     ports:
//     - number: 9000
//   peers:
//   - mtls:
//   origins:
//   - jwt:
//       issuer: "https://securetoken.google.com"
//       audiences:
//       - "productpage"
//       jwksUri: "https://www.googleapis.com/oauth2/v1/certs"
//       jwtHeaders:
//       - "x-goog-iap-jwt-assertion"
//       triggerRules:
//       - excludedPaths:
//         - exact: /health_check
//   principalBinding: USE_ORIGIN
// ```
//
// <!-- crd generation tags
// +cue-gen:Policy:groupName:authentication.istio.io
// +cue-gen:Policy:version:v1alpha1
// +cue-gen:Policy:storageVersion
// +cue-gen:Policy:annotations:helm.sh/resource-policy=keep
// +cue-gen:Policy:labels:app=istio-citadel,chart=istio,heritage=Tiller,release=istio
// +cue-gen:Policy:subresource:status
// +cue-gen:Policy:scope:Namespaced
// +cue-gen:Policy:resource:categories=istio-io,authentication-istio-io,plural=policies
// -->
//
// <!-- crd generation tags
// +cue-gen:MeshPolicy:groupName:authentication.istio.io
// +cue-gen:MeshPolicy:version:v1alpha1
// +cue-gen:MeshPolicy:storageVersion
// +cue-gen:MeshPolicy:annotations:helm.sh/resource-policy=keep
// +cue-gen:MeshPolicy:labels:app=istio-citadel,chart=istio,heritage=Tiller,release=istio
// +cue-gen:MeshPolicy:subresource:status
// +cue-gen:MeshPolicy:scope:Cluster
// +cue-gen:MeshPolicy:resource:categories=istio-io,authentication-istio-io,plural=meshpolicies
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=authentication.istio.io/v1alpha1
// +kubetype-gen:kubeType=Policy
// +kubetype-gen:kubeType=MeshPolicy
// +kubetype-gen:MeshPolicy:tag=genclient:nonNamespaced
// +genclient
// +k8s:deepcopy-gen=true
// -->
type MeshPolicy struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec authenticationv1alpha1.Policy `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MeshPolicyList is a collection of MeshPolicies.
type MeshPolicyList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []MeshPolicy `json:"items" protobuf:"bytes,2,rep,name=items"`
}
