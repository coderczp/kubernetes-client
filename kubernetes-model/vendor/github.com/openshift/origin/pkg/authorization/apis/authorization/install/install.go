/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package install

import (
	"k8s.io/apimachinery/pkg/apimachinery/announced"
	"k8s.io/apimachinery/pkg/apimachinery/registered"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/kubernetes/pkg/api/legacyscheme"

	"github.com/openshift/origin/pkg/api/legacy"
	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
	"github.com/openshift/origin/pkg/authorization/apis/authorization/rbacconversion"
	authorizationapiv1 "github.com/openshift/origin/pkg/authorization/apis/authorization/v1"
)

func init() {
	legacy.InstallLegacyAuthorization(legacyscheme.Scheme, legacyscheme.Registry)
	Install(legacyscheme.GroupFactoryRegistry, legacyscheme.Registry, legacyscheme.Scheme)
}

// Install registers the API group and adds types to a scheme
func Install(groupFactoryRegistry announced.APIGroupFactoryRegistry, registry *registered.APIRegistrationManager, scheme *runtime.Scheme) {
	if err := announced.NewGroupMetaFactory(
		&announced.GroupMetaFactoryArgs{
			GroupName:                  authorizationapi.GroupName,
			VersionPreferenceOrder:     []string{authorizationapiv1.SchemeGroupVersion.Version},
			AddInternalObjectsToScheme: internalObjectsToScheme,
			RootScopedKinds:            sets.NewString("ClusterRole", "ClusterRoleBinding", "ClusterPolicy", "ClusterPolicyBinding", "SubjectAccessReview", "ResourceAccessReview", "ResourceAccessReviewResponse", "SubjectAccessReviewResponse"),
		},
		announced.VersionToSchemeFunc{
			authorizationapiv1.SchemeGroupVersion.Version: authorizationapiv1.AddToScheme,
		},
	).Announce(groupFactoryRegistry).RegisterAndEnable(registry, scheme); err != nil {
		panic(err)
	}
}

func internalObjectsToScheme(scheme *runtime.Scheme) error {
	if err := authorizationapi.AddToScheme(scheme); err != nil {
		return err
	}
	return rbacconversion.AddToScheme(scheme)
}
