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
package v1

import (
	"github.com/openshift/api/image/v1"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
)

// LatestTaggedImage returns the most recent TagEvent for the specified image
// repository and tag. Will resolve lookups for the empty tag. Returns nil
// if tag isn't present in stream.status.tags.
func LatestTaggedImage(stream *v1.ImageStream, tag string) *v1.TagEvent {
	if len(tag) == 0 {
		tag = imageapi.DefaultImageTag
	}
	// find the most recent tag event with an image reference
	if stream.Status.Tags != nil {
		for _, t := range stream.Status.Tags {
			if t.Tag == tag {
				if len(t.Items) == 0 {
					return nil
				}
				return &t.Items[0]
			}
		}
	}

	return nil
}
