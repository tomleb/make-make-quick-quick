# Make `make quick` quick

This is a helper repository for the presentation on making the `make quick`
command for building Rancher faster.

It show cases a few ways to optimize Docker builds

Some steps for the demo are listed in [demo.txt](./demo.txt)

**Using cache mounts**

Demo found here: [mount-cache/go](./mount-cache/go) and [mount-cache/zypper/](./mount-cache/zypper/)

**Using multi-stage build and `COPY --from`

Demo found here: [copy-from](./copy-from/)

**Using a single Dockerfile to share intermediate stages**

Demo found here: [multi-target](./multi-target/)

**Using `--platform=$BUILDPLATFORM` images to avoid rebuild for different architecture**

Not shown during the presentation.

Demo found here: [multi-platform](./multi-platform/)
