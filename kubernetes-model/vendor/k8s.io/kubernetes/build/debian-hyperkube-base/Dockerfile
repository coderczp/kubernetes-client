#
# Copyright (C) 2015 Red Hat, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#         http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

FROM BASEIMAGE

RUN echo CACHEBUST>/dev/null && clean-install \
    bash

# The samba-common, cifs-utils, and nfs-common packages depend on
# ucf, which itself depends on /bin/bash.
RUN echo "dash dash/sh boolean false" | debconf-set-selections
RUN DEBIAN_FRONTEND=noninteractive dpkg-reconfigure dash

RUN echo CACHEBUST>/dev/null && clean-install \
    ca-certificates \
    ceph-common \
    cifs-utils \
    conntrack \
    e2fsprogs \
    ebtables \
    ethtool \
    git \
    glusterfs-client \
    iptables \
    jq \
    kmod \
    openssh-client \
    nfs-common \
    socat \
    util-linux

COPY cni-bin/bin /opt/cni/bin
