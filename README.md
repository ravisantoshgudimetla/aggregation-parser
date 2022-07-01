# aggregation-parser
Aggregation parser to parse output from openshift aggregation jobs. As of now, it just parses the output
from the aggregation job. The plan is to extend this to generate html files or update existing spreadsheet
in future.

```
$ make build
$ ./ap <aggregation_url>
For example:
$ ./ap https://pr-payload-tests.ci.openshift.org/runs/ci/150ace20-ebf4-11ec-88c9-c77511494be1-2

Would generate the following output:

Visiting https://pr-payload-tests.ci.openshift.org/runs/ci/150ace20-ebf4-11ec-88c9-c77511494be1-2
Got a response from https://pr-payload-tests.ci.openshift.org/runs/ci/150ace20-ebf4-11ec-88c9-c77511494be1-2
periodic-ci-openshift-release-master-nightly-4.11-e2e-metal-ipi-serial-ovn-dualstack {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-metal-ipi-serial-ovn-dualstack/1536727778810925056 failed}
periodic-ci-openshift-release-master-nightly-4.11-e2e-metal-ipi-upgrade-ovn-ipv6 {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-metal-ipi-upgrade-ovn-ipv6/1536727783001034752 failed}
periodic-ci-openshift-release-master-nightly-4.11-e2e-metal-single-node-live-iso {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-metal-single-node-live-iso/1536727788063559680 Success}
periodic-ci-openshift-release-master-nightly-4.11-e2e-aws {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-aws/1536727755318628352 failed}
periodic-ci-openshift-release-master-ci-4.11-e2e-aws-techpreview {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-ci-4.11-e2e-aws-techpreview/1536727761194848256 failed}
periodic-ci-openshift-release-master-nightly-4.11-upgrade-from-stable-4.10-e2e-metal-ipi-upgrade {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-upgrade-from-stable-4.10-e2e-metal-ipi-upgrade/1536727781340090368 failed}
periodic-ci-openshift-release-master-nightly-4.11-upgrade-from-stable-4.10-e2e-metal-ipi-upgrade-ovn-ipv6 {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-upgrade-from-stable-4.10-e2e-metal-ipi-upgrade-ovn-ipv6/1536727783844089856 failed}
periodic-ci-openshift-release-master-nightly-4.11-e2e-gcp-rt {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-gcp-rt/1536727773769371648 Success}
periodic-ci-openshift-release-master-ci-4.11-e2e-gcp-techpreview {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-ci-4.11-e2e-gcp-techpreview/1536727775036051456 Success}
periodic-ci-openshift-release-master-nightly-4.11-e2e-metal-assisted {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-metal-assisted/1536727776290148352 Success}
periodic-ci-openshift-release-master-nightly-4.11-e2e-metal-ipi-upgrade {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-metal-ipi-upgrade/1536727780509618176 Success}
periodic-ci-openshift-release-master-ci-4.11-upgrade-from-stable-4.10-e2e-gcp-ovn-rt-upgrade {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/aggregator-periodic-ci-openshift-release-master-ci-4.11-upgrade-from-stable-4.10-e2e-gcp-ovn-rt-upgrade/1536727753699627008 failed}
periodic-ci-openshift-release-master-ci-4.11-e2e-aws-ovn {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-ci-4.11-e2e-aws-ovn/1536727758674071552 failed}
periodic-ci-openshift-release-master-nightly-4.11-e2e-vsphere-upi-serial {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-vsphere-upi-serial/1536727795588141056 Success}
periodic-ci-openshift-release-master-nightly-4.11-e2e-aws-proxy {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-aws-proxy/1536727787249864704 failed}
periodic-ci-openshift-release-master-ci-4.11-e2e-gcp-ovn {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-ci-4.11-e2e-gcp-ovn/1536727772951482368 Success}
periodic-ci-openshift-release-master-nightly-4.11-e2e-vsphere {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-vsphere/1536727790978600960 Success}
periodic-ci-openshift-release-master-nightly-4.11-e2e-azure {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-azure/1536727762897735680 failed}
periodic-ci-openshift-release-master-nightly-4.11-e2e-metal-ipi-ovn-dualstack {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-metal-ipi-ovn-dualstack/1536727777129009152 failed}
periodic-ci-openshift-release-master-ci-4.11-e2e-azure-techpreview-serial {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-ci-4.11-e2e-azure-techpreview-serial/1536727767066873856 failed}
periodic-ci-openshift-release-master-nightly-4.11-console-aws {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-console-aws/1536727756165877760 failed}
periodic-ci-openshift-release-master-nightly-4.11-e2e-aws-single-node {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-aws-single-node/1536727759533903872 failed}
periodic-ci-openshift-release-master-nightly-4.11-e2e-metal-ipi-serial-ovn-ipv6 {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-metal-ipi-serial-ovn-ipv6/1536727777976258560 failed}
periodic-ci-openshift-release-master-nightly-4.11-upgrade-from-stable-4.10-e2e-aws-upgrade {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-upgrade-from-stable-4.10-e2e-aws-upgrade/1536727789711921152 failed}
periodic-ci-openshift-release-master-ci-4.11-e2e-gcp-techpreview-serial {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-ci-4.11-e2e-gcp-techpreview-serial/1536727775627448320 Success}
periodic-ci-openshift-release-master-nightly-4.11-e2e-vsphere-techpreview-serial {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-vsphere-techpreview-serial/1536727793977528320 Success}
periodic-ci-openshift-release-master-nightly-4.11-e2e-aws-csi {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-aws-csi/1536727756996349952 failed}
periodic-ci-openshift-release-master-ci-4.11-e2e-azure-ovn {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-ci-4.11-e2e-azure-ovn/1536727764558680064 failed}
periodic-ci-openshift-release-master-nightly-4.11-e2e-aws-driver-toolkit {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-aws-driver-toolkit/1536727770451677184 failed}
periodic-ci-openshift-release-master-nightly-4.11-e2e-aws-upgrade {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-aws-upgrade/1536727788956946432 failed}
periodic-ci-openshift-release-master-nightly-4.11-e2e-alibaba {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-alibaba/1536727754479767552 Success}
periodic-ci-openshift-release-master-nightly-4.11-e2e-metal-ipi-virtualmedia {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-metal-ipi-virtualmedia/1536727784792002560 Success}
periodic-ci-openshift-release-master-nightly-4.11-e2e-ovirt {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-ovirt/1536727785689583616 Success}
periodic-ci-openshift-release-master-nightly-4.11-e2e-aws-fips {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-aws-fips/1536727757839405056 failed}
periodic-ci-openshift-release-master-ci-4.11-e2e-aws-techpreview-serial {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-ci-4.11-e2e-aws-techpreview-serial/1536727762184704000 failed}
periodic-ci-openshift-release-master-nightly-4.11-e2e-gcp-csi {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-gcp-csi/1536727772100038656 failed}
periodic-ci-openshift-release-master-nightly-4.11-e2e-azure-csi {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-azure-csi/1536727763719819264 Success}
periodic-ci-openshift-release-master-ci-4.11-e2e-azure-techpreview {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-ci-4.11-e2e-azure-techpreview/1536727765384957952 Success}
periodic-ci-openshift-release-master-nightly-4.11-e2e-vsphere-techpreview {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-vsphere-techpreview/1536727793201582080 Success}
periodic-ci-openshift-release-master-nightly-4.11-e2e-vsphere-upi {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-vsphere-upi/1536727794753474560 Success}
periodic-ci-openshift-release-master-nightly-4.11-e2e-azure-deploy-cnv {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-azure-deploy-cnv/1536727768748789760 failed}
periodic-ci-openshift-release-master-nightly-4.11-e2e-gcp {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-gcp/1536727771265372160 failed}
periodic-ci-openshift-release-master-nightly-4.11-e2e-metal-ipi-serial-virtualmedia {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-metal-ipi-serial-virtualmedia/1536727779649785856 failed}
periodic-ci-openshift-release-master-nightly-4.11-e2e-ovirt-csi {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-ovirt-csi/1536727786364866560 failed}
periodic-ci-openshift-release-master-nightly-4.11-e2e-azure-upgrade-cnv {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-azure-upgrade-cnv/1536727769583456256 Success}
periodic-ci-openshift-release-master-nightly-4.11-e2e-vsphere-csi {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-vsphere-csi/1536727791439974400 Success}
periodic-ci-openshift-release-master-nightly-4.11-e2e-vsphere-serial {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-vsphere-serial/1536727792232697856 Success}
periodic-ci-openshift-release-master-nightly-4.11-e2e-aws-single-node-serial {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-nightly-4.11-e2e-aws-single-node-serial/1536727760351793152 failed}
periodic-ci-openshift-release-master-ci-4.11-e2e-azure-upgrade-single-node {https://prow.ci.openshift.org/view/gs/origin-ci-test/logs/openshift-origin-27244-ci-4.11-e2e-azure-upgrade-single-node/1536727767909928960 failed
```