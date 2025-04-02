ARG ARCH=amd64

ARG CHART_DEFAULT_BRANCH=dev-v2.11
ARG PARTNER_CHART_DEFAULT_BRANCH=main
ARG RKE2_CHART_DEFAULT_BRANCH=main

ARG CATTLE_K3S_VERSION=v1.32.1+k3s1

FROM registry.suse.com/bci/bci-base:15.6 AS tools
RUN zypper modifyrepo --keep-packages --all
RUN --mount=type=cache,target=/var/cache/zypp,sharing=locked,id=demo-zypper zypper refresh && \
    zypper -n install git

FROM tools as charts
ARG CHART_DEFAULT_BRANCH
RUN mkdir -p /var/lib/rancher-data/local-catalogs/v2 && \
    git config --global url."https://github.com/rancher/".insteadOf https://git.rancher.io/ && \
    git clone --no-checkout -b $CHART_DEFAULT_BRANCH --depth 1 https://git.rancher.io/charts /var/lib/rancher-data/local-catalogs/v2/rancher-charts/4b40cac650031b74776e87c1a726b0484d0877c3ec137da0872547ff9b73a721

FROM tools as partner-charts
ARG PARTNER_CHART_DEFAULT_BRANCH
RUN mkdir -p /var/lib/rancher-data/local-catalogs/v2 && \
    git config --global url."https://github.com/rancher/".insteadOf https://git.rancher.io/ && \
    git clone --no-checkout -b $PARTNER_CHART_DEFAULT_BRANCH --depth 1 https://git.rancher.io/partner-charts /var/lib/rancher-data/local-catalogs/v2/rancher-partner-charts/8f17acdce9bffd6e05a58a3798840e408c4ea71783381ecd2e9af30baad65974

FROM tools as rke2-charts
ARG RKE2_CHART_DEFAULT_BRANCH
RUN mkdir -p /var/lib/rancher-data/local-catalogs/v2 && \
    git config --global url."https://github.com/rancher/".insteadOf https://git.rancher.io/ && \
    git clone --no-checkout -b $RKE2_CHART_DEFAULT_BRANCH --depth 1 https://git.rancher.io/rke2-charts /var/lib/rancher-data/local-catalogs/v2/rancher-rke2-charts/675f1b63a0a83905972dcab2794479ed599a6f41b86cd6193d69472d0fa889c9

FROM tools as tini
ARG ARCH
ENV TINI_VERSION v0.18.0
ENV TINI_URL_amd64=https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini \
    TINI_URL_arm64=https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini-arm64 \
    TINI_URL_s390x=https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini-s390x \
    TINI_URL=TINI_URL_${ARCH}
RUN curl -sLf ${!TINI_URL} > /usr/bin/tini && \
    chmod +x /usr/bin/tini

FROM tools as etcd
ARG ARCH
ENV CATTLE_ETCD_VERSION v3.5.16
ENV ETCD_URL=https://github.com/etcd-io/etcd/releases/download/${CATTLE_ETCD_VERSION}/etcd-${CATTLE_ETCD_VERSION}-linux-${ARCH}.tar.gz
RUN mkdir -p /etcd && curl -sfL ${ETCD_URL} | tar xvzf - --strip-components=1 --no-same-owner -C /etcd etcd-${CATTLE_ETCD_VERSION}-linux-${ARCH}/etcdctl

FROM --platform=$BUILDPLATFORM registry.suse.com/bci/golang:1.23 AS go-tools
ENV GOMODCACHE=/root/.cache/go/mod-cache
ENV GOCACHE=/root/.cache/go/cache
RUN zypper modifyrepo --keep-packages --all
RUN --mount=type=cache,target=/var/cache/zypp,sharing=locked,id=demo-zypper zypper refresh && \
    zypper -n install libbtrfs-devel libgpgme-devel
WORKDIR /src

FROM go-tools as server-builder
ARG CATTLE_K3S_VERSION
COPY server/airgap/go.mod server/airgap/go.sum /src/
RUN --mount=type=cache,target=/root/.cache,id=demo go mod download -x
COPY server/airgap/ /src/
RUN --mount=type=cache,target=/root/.cache,id=demo go build -tags k3s_export -o export-images ./...
RUN ./export-images -k3s-version ${CATTLE_K3S_VERSION} -output /src/k3s-airgap-images.tar

FROM go-tools as agent-builder
ARG CATTLE_K3S_VERSION
COPY agent/airgap/go.mod agent/airgap/go.sum /src/
RUN --mount=type=cache,target=/root/.cache,id=demo go mod download -x
COPY agent/airgap/ /src/
RUN --mount=type=cache,target=/root/.cache,id=demo go build -tags k3s_export -o export-images ./...
RUN ./export-images -k3s-version ${CATTLE_K3S_VERSION} -output /src/k3s-airgap-images.tar

# Final image for server
FROM registry.suse.com/bci/bci-base:15.6 as server

ENV CATTLE_CHART_DEFAULT_BRANCH=$CHART_DEFAULT_BRANCH
ENV CATTLE_PARTNER_CHART_DEFAULT_BRANCH=$PARTNER_CHART_DEFAULT_BRANCH
ENV CATTLE_RKE2_CHART_DEFAULT_BRANCH=$RKE2_CHART_DEFAULT_BRANCH

RUN mkdir -p /var/lib/rancher/k3s/agent/images/ && \
    mkdir -p /var/lib/rancher-data/driver-metadata

COPY --from=etcd /etcd/* /usr/bin/
COPY --from=charts /var/lib/rancher-data/local-catalogs/v2/rancher-charts/4b40cac650031b74776e87c1a726b0484d0877c3ec137da0872547ff9b73a721 /var/lib/rancher-data/local-catalogs/v2/rancher-charts/4b40cac650031b74776e87c1a726b0484d0877c3ec137da0872547ff9b73a721
COPY --from=partner-charts /var/lib/rancher-data/local-catalogs/v2/rancher-partner-charts/8f17acdce9bffd6e05a58a3798840e408c4ea71783381ecd2e9af30baad65974 /var/lib/rancher-data/local-catalogs/v2/rancher-partner-charts/8f17acdce9bffd6e05a58a3798840e408c4ea71783381ecd2e9af30baad65974
COPY --from=rke2-charts /var/lib/rancher-data/local-catalogs/v2/rancher-rke2-charts/675f1b63a0a83905972dcab2794479ed599a6f41b86cd6193d69472d0fa889c9 /var/lib/rancher-data/local-catalogs/v2/rancher-rke2-charts/675f1b63a0a83905972dcab2794479ed599a6f41b86cd6193d69472d0fa889c9
COPY --from=tini /usr/bin/tini /usr/bin/tini
COPY --from=server-builder /src/k3s-airgap-images.tar /server.tar


# The final image for agent
FROM registry.suse.com/bci/bci-base:15.6 as agent
COPY --from=charts /var/lib/rancher-data/local-catalogs/v2/rancher-charts/4b40cac650031b74776e87c1a726b0484d0877c3ec137da0872547ff9b73a721 /var/lib/rancher-data/local-catalogs/v2/rancher-charts/4b40cac650031b74776e87c1a726b0484d0877c3ec137da0872547ff9b73a721
COPY --from=partner-charts /var/lib/rancher-data/local-catalogs/v2/rancher-partner-charts/8f17acdce9bffd6e05a58a3798840e408c4ea71783381ecd2e9af30baad65974 /var/lib/rancher-data/local-catalogs/v2/rancher-partner-charts/8f17acdce9bffd6e05a58a3798840e408c4ea71783381ecd2e9af30baad65974
COPY --from=rke2-charts /var/lib/rancher-data/local-catalogs/v2/rancher-rke2-charts/675f1b63a0a83905972dcab2794479ed599a6f41b86cd6193d69472d0fa889c9 /var/lib/rancher-data/local-catalogs/v2/rancher-rke2-charts/675f1b63a0a83905972dcab2794479ed599a6f41b86cd6193d69472d0fa889c9
COPY --from=tini /usr/bin/tini /usr/bin/tini
COPY --from=agent-builder /src/k3s-airgap-images.tar /agent.tar
