# Changelog v1.59

## Know before update


 - Application traffic redirection setup method in the istio module deligated to a special CNI plugin. Cilium agent pods will restart, and L7-based policies will flap.
 - Deckhouse will not update if `linstor` module is enabled. [Migrate](https://deckhouse.io/modules/sds-replicated-volume/stable/faq.html#migrating-from-the-deckhouse-kubernetes-platform-linstorhttpsdeckhouseiodocumentationv157modules041-linstor--built-in-module-to-sds-replicated-volume) from `linstor` module to `sds-replicated-volume`.
 - `sds-drbd` module is renamed. Please switch to [sds-replicated-volume](https://deckhouse.io/modules/sds-replicated-volume/stable/faq.html#migrating-from-sds-drbd-module-to-sds-replicated-volume) module ASAP. `sds-drbd` module cannot be enabled but will continue to work if it was already enabled before.

## Features


 - **[admission-policy-engine]** Add new columns to Grafana representing policy name and policy type. [#7926](https://github.com/deckhouse/deckhouse/pull/7926)
 - **[admission-policy-engine]** Add the ability to specify restrictions on applying cluster roles to users. [#7660](https://github.com/deckhouse/deckhouse/pull/7660)
 - **[admission-policy-engine]** Add a security policy to control the mounting of service account tokens in pods. This rule will prohibit running pods with `automountServiceAccountToken` for existing _SecurityPolicy_. To restore the old behavior, add `automountServiceAccountToken: true` to the policy. [#7571](https://github.com/deckhouse/deckhouse/pull/7571)
 - **[candi]** Add Deckhouse Kubernetes Platform Standart Edition (SE). [#7828](https://github.com/deckhouse/deckhouse/pull/7828)
 - **[candi]** Add support for Debian 12, remove support for Debian 9. [#7631](https://github.com/deckhouse/deckhouse/pull/7631)
 - **[candi]** Add OpenStack server group support for master nodes. [#7312](https://github.com/deckhouse/deckhouse/pull/7312)
 - **[candi]** Use FQDN for a node name (for the new nodes). [#7117](https://github.com/deckhouse/deckhouse/pull/7117)
 - **[cloud-provider-openstack]** Allow to install clusters with http proxy in OpenStack [#7924](https://github.com/deckhouse/deckhouse/pull/7924)
 - **[cloud-provider-openstack]** Add alert about orphaned disks (without PV) in the cloud. [#7458](https://github.com/deckhouse/deckhouse/pull/7458)
 - **[cloud-provider-vsphere]** Add a feature to discover orphaned disks for vSphere. [#7676](https://github.com/deckhouse/deckhouse/pull/7676)
 - **[cloud-provider-vsphere]** Added a parameter that allows you to use an existing VM folder. [#7667](https://github.com/deckhouse/deckhouse/pull/7667)
 - **[cloud-provider-yandex]** Add a feature to discover orphaned disks for YandexCloud. [#7603](https://github.com/deckhouse/deckhouse/pull/7603)
 - **[cni-cilium]** Add workaround for the stale DNS connections to the node-local-dns in the host network applications. [#7632](https://github.com/deckhouse/deckhouse/pull/7632)
 - **[deckhouse]** Fail `deckhouse-controller` if module restoration failed on startup. [#7915](https://github.com/deckhouse/deckhouse/pull/7915)
 - **[deckhouse]** Create embedded default `ModuleUpdatePolicy`, if no others are set. [#7703](https://github.com/deckhouse/deckhouse/pull/7703)
 - **[deckhouse]** Add `Ignore` mode for `ModuleUpdatePolicy` to exclude desired modules. [#7703](https://github.com/deckhouse/deckhouse/pull/7703)
 - **[deckhouse-controller]** Ability to set and display the current life cycle stage for modules. [#7807](https://github.com/deckhouse/deckhouse/pull/7807)
 - **[deckhouse-controller]** Rework the process of updating modules' statuses. [#7741](https://github.com/deckhouse/deckhouse/pull/7741)
 - **[deckhouse-controller]** Provides Deckhouse HA mode. [#7634](https://github.com/deckhouse/deckhouse/pull/7634)
 - **[deckhouse-controller]** Added module loading metrics. [#7424](https://github.com/deckhouse/deckhouse/pull/7424)
 - **[dhctl]** Support for creating chunked images bundles in mirror. [#8000](https://github.com/deckhouse/deckhouse/pull/8000)
 - **[extended-monitoring]** Support force image check with `image-availability-exporter`, even when workloads are disabled or suspended. [#7606](https://github.com/deckhouse/deckhouse/pull/7606)
 - **[external-module-manager]** Disable module hooks on restart. [#7582](https://github.com/deckhouse/deckhouse/pull/7582)
 - **[ingress-nginx]** Add the IP address and hostname to the `status` field of the `IngressNginxController` resource. [#7785](https://github.com/deckhouse/deckhouse/pull/7785)
 - **[ingress-nginx]** Add VPA support to OpenKruise advance daemonsets. [#7557](https://github.com/deckhouse/deckhouse/pull/7557)
    Vpa-admission-controller/recommender/updater pods will be recreated. OpenKruise controller manager pod will restart.
 - **[istio]** The new traffic redirection mode "IstioCNI" for the Istio module is to replace the old one, "InitContainer". [#7488](https://github.com/deckhouse/deckhouse/pull/7488)
    Application traffic redirection setup method in the istio module deligated to a special CNI plugin. Cilium agent pods will restart, and L7-based policies will flap.
 - **[l2-load-balancer]** Show public IP in the status field of the `L2LoadBalancer` CustomResource. [#7937](https://github.com/deckhouse/deckhouse/pull/7937)
 - **[l2-load-balancer]** The new module for redundant L2 load-balancing and a new IngressNginxController inlet L2LoadBalancer. [#7752](https://github.com/deckhouse/deckhouse/pull/7752)
 - **[log-shipper]** Validate the TLS certificate of the remote host for Loki. [#6911](https://github.com/deckhouse/deckhouse/pull/6911)
    Loki and Grafana will restart.
 - **[loki]** Add `kube-rbac-proxy` to prevent access to Loki from unauthorized clients. [#6911](https://github.com/deckhouse/deckhouse/pull/6911)
    Loki and Grafana will restart.
 - **[monitoring-kubernetes]** `ebpf_exporter` image is based on a distroless image. Bump version to v2.3.0. [#7769](https://github.com/deckhouse/deckhouse/pull/7769)
 - **[node-local-dns]** Add a workaround for the stale DNS connections to the `node-local-dns` in the host network applications. [#7632](https://github.com/deckhouse/deckhouse/pull/7632)
 - **[node-manager]** Scale all NodeGroups (with low priority) if any of them have a priority. [#7772](https://github.com/deckhouse/deckhouse/pull/7772)
 - **[node-manager]** Check the `StaticInstance` address by establishing a TCP connection. [#7640](https://github.com/deckhouse/deckhouse/pull/7640)
 - **[node-manager]** Add check for minimal Debian version. [#7631](https://github.com/deckhouse/deckhouse/pull/7631)
 - **[prometheus]** - Equal alerts with different severity are deduplicated (only lowest severity alert is displayed in the cluster).
    - Alerts queue size changed from 100 to 500.
    - Removed `DeadMansSwitch` alert and added `MissingDeadMansSwitch` alert in the case of `DeadMansSwitch` alert does not send by cluster Prometheus.
    - Added `ClusterHasTooManyAlerts` alert to indicate queue fullness. [#7809](https://github.com/deckhouse/deckhouse/pull/7809)
    `alerts-receiver` will restart.
 - **[prometheus]** Aggregating proxy for Prometheus in HA mode. [#7757](https://github.com/deckhouse/deckhouse/pull/7757)
    Now there is one Grafana datasource to see combined metrics from the Main and the Longterm Prometheus.
 - **[prometheus]** Add separate transition deployment for Grafana v10. [#7018](https://github.com/deckhouse/deckhouse/pull/7018)
 - **[registrypackages]** Add the `d8` tools (`deckhouse-cli`) to registry packages. [#7792](https://github.com/deckhouse/deckhouse/pull/7792)
 - **[vertical-pod-autoscaler]** Add VPA support to OpenKruise advance daemonsets. [#7557](https://github.com/deckhouse/deckhouse/pull/7557)
    Vpa-admission-controller/recommender/updater pods will be recreated. OpenKruise controller manager pod will be recreated.

## Fixes


 - **[admission-policy-engine]** Fix storageClass fake alert. [#7912](https://github.com/deckhouse/deckhouse/pull/7912)
 - **[admission-policy-engine]** Prevent rendering of unavailable components. [#7663](https://github.com/deckhouse/deckhouse/pull/7663)
 - **[candi]** Disable updating tfadm for bootstraping. [#7930](https://github.com/deckhouse/deckhouse/pull/7930)
 - **[candi]** Fix bootstrap by replace bb-error checks with trap. [#7647](https://github.com/deckhouse/deckhouse/pull/7647)
 - **[candi]** Remove `yum versionlock` and `apt-mark hold`. [#7280](https://github.com/deckhouse/deckhouse/pull/7280)
 - **[cloud-provider-aws]** Update `aws-ebs-csi-driver` version to `v1.28.0`. [#7691](https://github.com/deckhouse/deckhouse/pull/7691)
 - **[cloud-provider-aws]** Prevent rendering of unavailable components. [#7663](https://github.com/deckhouse/deckhouse/pull/7663)
 - **[cloud-provider-azure]** Prevent rendering of unavailable components. [#7663](https://github.com/deckhouse/deckhouse/pull/7663)
 - **[cloud-provider-gcp]** Prevent rendering of unavailable components. [#7663](https://github.com/deckhouse/deckhouse/pull/7663)
 - **[cloud-provider-yandex]** Prevent rendering of unavailable components. [#7663](https://github.com/deckhouse/deckhouse/pull/7663)
 - **[cloud-provider-yandex]** Fix a long node name. [#6718](https://github.com/deckhouse/deckhouse/pull/6718)
 - **[dashboard]** Add `kube-rbac-proxy-ca.crt` configMap to `d8-dashboard` namespace. [#7766](https://github.com/deckhouse/deckhouse/pull/7766)
 - **[deckhouse]** Use the same logic as minor/patch for `apply-now` release, except the time settings. [#7988](https://github.com/deckhouse/deckhouse/pull/7988)
 - **[deckhouse]** Сhange the way the `deckhouse` pod readiness is determined during the minor version update. [#7770](https://github.com/deckhouse/deckhouse/pull/7770)
 - **[deckhouse]** Validation configs of disabled modules is disabled. [#7744](https://github.com/deckhouse/deckhouse/pull/7744)
 - **[descheduler]** Remove incorrect inclusions of the `removeDuplicates` field. [#7242](https://github.com/deckhouse/deckhouse/pull/7242)
 - **[dhctl]** Fix meta config deep copy method does not return the copy. [#7854](https://github.com/deckhouse/deckhouse/pull/7854)
 - **[external-module-manager]** Fix panic at a module status update. [#7648](https://github.com/deckhouse/deckhouse/pull/7648)
 - **[global-hooks]** Fix ability to downgrade Kubernetes by more than 1 minor version. [#7279](https://github.com/deckhouse/deckhouse/pull/7279)
 - **[ingress-nginx]** Digital ocean Kubernetes upgrade, update `timeoutSeconds`. [#7933](https://github.com/deckhouse/deckhouse/pull/7933)
 - **[ingress-nginx]** Improve description of `NginxIngressSslWillExpire` and `NginxIngressSslExpired` alerts. [#7830](https://github.com/deckhouse/deckhouse/pull/7830)
 - **[ingress-nginx]** Fix HTTPS port validation for HostPort inlet. [#7813](https://github.com/deckhouse/deckhouse/pull/7813)
 - **[ingress-nginx]** Prevent rendering of unavailable components. [#7663](https://github.com/deckhouse/deckhouse/pull/7663)
 - **[l2-load-balancer]** Fix missed `externalTrafficPolicy` option for L2LoadBalancer. [#7968](https://github.com/deckhouse/deckhouse/pull/7968)
 - **[monitoring-kubernetes]** Reducing the number of scans of release secrets in the cluster. [#7558](https://github.com/deckhouse/deckhouse/pull/7558)
 - **[node-manager]** Simplify ng zones check expression. [#8022](https://github.com/deckhouse/deckhouse/pull/8022)
 - **[node-manager]** Prevent rendering of unavailable components. [#7663](https://github.com/deckhouse/deckhouse/pull/7663)
 - **[node-manager]** Add the `PasswordAuthentication=no` option to the OpenSSH client if public key authentication is used. [#7621](https://github.com/deckhouse/deckhouse/pull/7621)
 - **[node-manager]** Check for `SSHCredentials` in the `StaticInstance` webhook. [#7619](https://github.com/deckhouse/deckhouse/pull/7619)
 - **[node-manager]** Fix the cleanup phase in the 'caps-controller-manager' if the cleanup script does not exist. [#7615](https://github.com/deckhouse/deckhouse/pull/7615)
 - **[node-manager]** Add NodeGroup validation to prevent adding more than one taint with the same key and effect. [#7474](https://github.com/deckhouse/deckhouse/pull/7474)
 - **[node-manager]** Fix a race in instance controller if NodeGroup has been deleted. [#7454](https://github.com/deckhouse/deckhouse/pull/7454)
 - **[node-manager]** Add the ability to disable StaticInstance bootstrapping by adding the `"node.deckhouse.io/allow-bootstrap": "false"` label. [#7244](https://github.com/deckhouse/deckhouse/pull/7244)
 - **[node-manager]** Add the alert about an impossible node drain. [#6190](https://github.com/deckhouse/deckhouse/pull/6190)
 - **[operator-prometheus]** Prevent rendering of unavailable components. [#7663](https://github.com/deckhouse/deckhouse/pull/7663)
 - **[prometheus]** Fix aggregation with longterm Prometheus. [#8056](https://github.com/deckhouse/deckhouse/pull/8056)
 - **[prometheus]** Migrate from `statusmap` and `status-history` to `state-timeline` plugin. [#8040](https://github.com/deckhouse/deckhouse/pull/8040)
 - **[prometheus]** Increase `sampleLimit` threshold to 100000. [#7858](https://github.com/deckhouse/deckhouse/pull/7858)

## Chore


 - **[candi]** Bump patch versions of Kubernetes images: `v1.26.15`, `v1.27.12`, `v1.28.8`, `v1.29.3` [#7834](https://github.com/deckhouse/deckhouse/pull/7834)
    Kubernetes control-plane components will restart, kubelet will restart.
 - **[candi]** Change the terraform version to `0.14.8`. [#7513](https://github.com/deckhouse/deckhouse/pull/7513)
    All cloud envirinment(clusters in AWS,GCP,Azure,Yandex-Cloud,OpenStack,VSphere)
 - **[cni-cilium]** Improvement of the cilium-agent version update process. [#7658](https://github.com/deckhouse/deckhouse/pull/7658)
    `cilium-agent` pods can restart.
 - **[cni-cilium]** cilium-agent is now based on a distroless images. [#7225](https://github.com/deckhouse/deckhouse/pull/7225)
    `cilium-agent` pods will restart.
 - **[deckhouse]** Add a validating webhook to prevent the `sds-drbd` module from being enabled. [#7918](https://github.com/deckhouse/deckhouse/pull/7918)
    `sds-drbd` module is renamed. Please switch to [sds-replicated-volume](https://deckhouse.io/modules/sds-replicated-volume/stable/faq.html#migrating-from-sds-drbd-module-to-sds-replicated-volume) module ASAP. `sds-drbd` module cannot be enabled but will continue to work if it was already enabled before.
 - **[deckhouse-controller]** Remove linstor module [#7091](https://github.com/deckhouse/deckhouse/pull/7091)
    Deckhouse will not update if `linstor` module is enabled. [Migrate](https://deckhouse.io/modules/sds-replicated-volume/stable/faq.html#migrating-from-the-deckhouse-kubernetes-platform-linstorhttpsdeckhouseiodocumentationv157modules041-linstor--built-in-module-to-sds-replicated-volume) from `linstor` module to `sds-replicated-volume`.
 - **[dhctl]** Support Deckhouse controller high availability mode during bootstrap. [#7919](https://github.com/deckhouse/deckhouse/pull/7919)
 - **[prometheus]** Move memcached to distroless. [#7945](https://github.com/deckhouse/deckhouse/pull/7945)
    Memcached (Prometheus) will restart.
 - **[prometheus]** Update metrics rotation value. [#7560](https://github.com/deckhouse/deckhouse/pull/7560)
 - **[prometheus]** Grafana v10.x migration-related deprecation alerts. [#7340](https://github.com/deckhouse/deckhouse/pull/7340)

