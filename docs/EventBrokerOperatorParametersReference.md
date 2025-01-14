# PubSub+ Event Broker Operator API Parameters Reference

Packages:

- [pubsubplus.solace.com/v1beta1](#pubsubplussolacecomv1beta1)

# pubsubplus.solace.com/v1beta1

Resource Types:

- [PubSubPlusEventBroker](#pubsubpluseventbroker)




## PubSubPlusEventBroker
<sup><sup>[↩ Parent](#pubsubplussolacecomv1beta1 )</sup></sup>






PubSub+ Event Broker

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>pubsubplus.solace.com/v1beta1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>PubSubPlusEventBroker</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspec">spec</a></b></td>
        <td>object</td>
        <td>
          EventBrokerSpec defines the desired state of PubSubPlusEventBroker<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerstatus">status</a></b></td>
        <td>object</td>
        <td>
          EventBrokerStatus defines the observed state of the PubSubPlusEventBroker<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec
<sup><sup>[↩ Parent](#pubsubpluseventbroker)</sup></sup>



EventBrokerSpec defines the desired state of PubSubPlusEventBroker

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>adminCredentialsSecret</b></td>
        <td>string</td>
        <td>
          Defines the password for PubSubPlusEventBroker if provided. Random one will be generated if not provided.
When provided, ensure the secret key name is `username_admin_password`. For valid values refer to the Solace documentation https://docs.solace.com/Admin/Configuring-Internal-CLI-User-Accounts.htm.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecbrokercontainersecurity">brokerContainerSecurity</a></b></td>
        <td>object</td>
        <td>
          ContainerSecurityContext defines the container security context for the PubSubPlusEventBroker.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>developer</b></td>
        <td>boolean</td>
        <td>
          Developer true specifies a minimum footprint scaled-down deployment, not for production use.
If set to true it overrides SystemScaling parameters.<br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>enableServiceLinks</b></td>
        <td>boolean</td>
        <td>
          EnableServiceLinks indicates whether information about services should be injected into pod's environment
variables, matching the syntax of Docker links. Optional: Defaults to false.<br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecextraenvvarsindex">extraEnvVars</a></b></td>
        <td>[]object</td>
        <td>
          List of extra environment variables to be added to the PubSubPlusEventBroker container. Note: Do not configure Timezone or SystemScaling parameters here as it could cause unintended consequences.
A primary use case is to specify configuration keys, although the variables defined here will not override the ones defined in ConfigMap<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>extraEnvVarsCM</b></td>
        <td>string</td>
        <td>
          List of extra environment variables to be added to the PubSubPlusEventBroker container from an existing ConfigMap. Note: Do not configure Timezone or SystemScaling parameters here as it could cause unintended consequences.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>extraEnvVarsSecret</b></td>
        <td>string</td>
        <td>
          List of extra environment variables to be added to the PubSubPlusEventBroker container from an existing Secret<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecimage">image</a></b></td>
        <td>object</td>
        <td>
          Image defines container image parameters for the event broker.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecmonitoring">monitoring</a></b></td>
        <td>object</td>
        <td>
          Monitoring specifies a Prometheus monitoring endpoint for the event broker<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>monitoringCredentialsSecret</b></td>
        <td>string</td>
        <td>
          Defines the password for PubSubPlusEventBroker to be used by the Exporter for monitoring.
When provided, ensure the secret key name is `username_monitor_password`. For valid values refer to the Solace documentation https://docs.solace.com/Admin/Configuring-Internal-CLI-User-Accounts.htm.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindex">nodeAssignment</a></b></td>
        <td>[]object</td>
        <td>
          NodeAssignment defines labels to constrain PubSubPlusEventBroker nodes to run on particular node(s), or to prefer to run on particular nodes.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>podAnnotations</b></td>
        <td>map[string]string</td>
        <td>
          PodAnnotations allows adding provider-specific pod annotations to PubSubPlusEventBroker pods<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>podDisruptionBudgetForHA</b></td>
        <td>boolean</td>
        <td>
          PodDisruptionBudgetForHA enables setting up PodDisruptionBudget for the broker pods in HA deployment.
This parameter is ignored for non-HA deployments (if redundancy is false).<br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>podLabels</b></td>
        <td>map[string]string</td>
        <td>
          PodLabels allows adding provider-specific pod labels to PubSubPlusEventBroker pods<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>preSharedAuthKeySecret</b></td>
        <td>string</td>
        <td>
          PreSharedAuthKeySecret defines the PreSharedAuthKey Secret for PubSubPlusEventBroker. Random one will be generated if not provided.
When provided, ensure the secret key name is `preshared_auth_key`. For valid values refer to the Solace documentation https://docs.solace.com/Features/HA-Redundancy/Pre-Shared-Keys-SMB.htm?Highlight=pre%20shared.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>redundancy</b></td>
        <td>boolean</td>
        <td>
          Redundancy true specifies HA deployment, false specifies Non-HA.<br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecsecuritycontext">securityContext</a></b></td>
        <td>object</td>
        <td>
          SecurityContext defines the pod security context for the event broker.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecservice">service</a></b></td>
        <td>object</td>
        <td>
          Service defines broker service details.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecserviceaccount">serviceAccount</a></b></td>
        <td>object</td>
        <td>
          ServiceAccount defines a ServiceAccount dedicated to the PubSubPlusEventBroker<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecstorage">storage</a></b></td>
        <td>object</td>
        <td>
          Storage defines storage details for the broker.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecsystemscaling">systemScaling</a></b></td>
        <td>object</td>
        <td>
          SystemScaling provides exact fine-grained specification of the event broker scaling parameters
and the assigned CPU / memory resources to the Pod.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>timezone</b></td>
        <td>string</td>
        <td>
          Defines the timezone for the event broker container, if undefined default is UTC. Valid values are tz database time zone names.<br/>
          <br/>
            <i>Default</i>: UTC<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspectls">tls</a></b></td>
        <td>object</td>
        <td>
          TLS provides TLS configuration for the event broker.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>updateStrategy</b></td>
        <td>enum</td>
        <td>
          UpdateStrategy specifies how to update an existing deployment. manualPodRestart waits for user intervention.<br/>
          <br/>
            <i>Enum</i>: automatedRolling, manualPodRestart<br/>
            <i>Default</i>: automatedRolling<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.brokerContainerSecurity
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspec)</sup></sup>



ContainerSecurityContext defines the container security context for the PubSubPlusEventBroker.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>runAsGroup</b></td>
        <td>number</td>
        <td>
          Specifies runAsGroup in container security context. 0 or unset defaults either to 1000002, or if OpenShift detected to unspecified (see documentation)<br/>
          <br/>
            <i>Format</i>: int64<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>runAsUser</b></td>
        <td>number</td>
        <td>
          Specifies runAsUser in container security context. 0 or unset defaults either to 1000001, or if OpenShift detected to unspecified (see documentation)<br/>
          <br/>
            <i>Format</i>: int64<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.extraEnvVars[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspec)</sup></sup>



ExtraEnvVar defines environment variables to be added to the PubSubPlusEventBroker container

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Specifies the Name of an environment variable to be added to the PubSubPlusEventBroker container<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>value</b></td>
        <td>string</td>
        <td>
          Specifies the Value of an environment variable to be added to the PubSubPlusEventBroker container<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.image
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspec)</sup></sup>



Image defines container image parameters for the event broker.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>pullPolicy</b></td>
        <td>string</td>
        <td>
          Specifies ImagePullPolicy of the container image for the event broker.<br/>
          <br/>
            <i>Default</i>: IfNotPresent<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecimagepullsecretsindex">pullSecrets</a></b></td>
        <td>[]object</td>
        <td>
          pullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>repository</b></td>
        <td>string</td>
        <td>
          Defines the container image repo where the event broker image is pulled from<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>tag</b></td>
        <td>string</td>
        <td>
          Specifies the tag of the container image to be used for the event broker.<br/>
          <br/>
            <i>Default</i>: latest<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.image.pullSecrets[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecimage)</sup></sup>



LocalObjectReference contains enough information to let you locate the
referenced object inside the same namespace.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the referent.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
TODO: Add other useful fields. apiVersion, kind, uid?<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.monitoring
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspec)</sup></sup>



Monitoring specifies a Prometheus monitoring endpoint for the event broker

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>enabled</b></td>
        <td>boolean</td>
        <td>
          Enabled true enables the setup of the Prometheus Exporter.<br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
       </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecmonitoringextraenvvarsindex">extraEnvVars</a></b></td>
        <td>[]object</td>
        <td>
          List of extra environment variables to be added to the Prometheus Exporter container.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecmonitoringimage">image</a></b></td>
        <td>object</td>
        <td>
          Image defines container image parameters for the Prometheus Exporter.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>includeRates</b></td>
        <td>boolean</td>
        <td>
          Defines if Prometheus Exporter should include rates<br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecmonitoringmetricsendpoint">metricsEndpoint</a></b></td>
        <td>object</td>
        <td>
          MetricsEndpoint defines parameters to configure monitoring for the Prometheus Exporter.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecmonitoringresources">resources</a></b></td>
        <td>object</td>
        <td>
          Defines container resources for the Promethes Exporter.<br/>
          <br/>
            <i>Default</i>: {}<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>sslVerify</b></td>
        <td>boolean</td>
        <td>
          Defines if Prometheus Exporter verifies SSL<br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>timeOut</b></td>
        <td>number</td>
        <td>
          Timeout configuration for Prometheus Exporter scrapper<br/>
          <br/>
            <i>Format</i>: int32<br/>
            <i>Default</i>: 5<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>

### PubSubPlusEventBroker.spec.monitoring.extraEnvVars[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecmonitoring)</sup></sup>



MonitoringExtraEnvVar defines environment variables to be added to the Prometheus Exporter container for Monitoring

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Specifies the Name of an environment variable to be added to the Prometheus Exporter container for Monitoring<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>value</b></td>
        <td>string</td>
        <td>
          Specifies the Value of an environment variable to be added to the Prometheus Exporter container for Monitoring<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>



### PubSubPlusEventBroker.spec.monitoring.image
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecmonitoring)</sup></sup>



Image defines container image parameters for the Prometheus Exporter.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>pullPolicy</b></td>
        <td>string</td>
        <td>
          Specifies ImagePullPolicy of the container image for the Prometheus Exporter.<br/>
          <br/>
            <i>Default</i>: IfNotPresent<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecmonitoringimagepullsecretsindex">pullSecrets</a></b></td>
        <td>[]object</td>
        <td>
          pullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>repository</b></td>
        <td>string</td>
        <td>
          Defines the container image repo where the Prometheus Exporter image is pulled from<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>tag</b></td>
        <td>string</td>
        <td>
          Specifies the tag of the container image to be used for the Prometheus Exporter.<br/>
          <br/>
            <i>Default</i>: latest<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.monitoring.image.pullSecrets[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecmonitoringimage)</sup></sup>



LocalObjectReference contains enough information to let you locate the
referenced object inside the same namespace.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the referent.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
TODO: Add other useful fields. apiVersion, kind, uid?<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.monitoring.metricsEndpoint
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecmonitoring)</sup></sup>



MetricsEndpoint defines parameters to configure monitoring for the Prometheus Exporter.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>containerPort</b></td>
        <td>number</td>
        <td>
          ContainerPort is the port number to expose on the Prometheus Exporter pod.<br/>
          <br/>
            <i>Format</i>: int32<br/>
            <i>Default</i>: 9628<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>endpointTlsConfigPrivateKeyName</b></td>
        <td>string</td>
        <td>
          EndpointTlsConfigPrivateKeyName is the file name of the Private Key used to set up TLS configuration<br/>
          <br/>
            <i>Default</i>: tls.key<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>endpointTlsConfigSecret</b></td>
        <td>string</td>
        <td>
          EndpointTLSConfigSecret defines TLS secret name to set up TLS configuration<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>endpointTlsConfigServerCertName</b></td>
        <td>string</td>
        <td>
          EndpointTlsConfigServerCertName is the file name of the Server Certificate used to set up TLS configuration<br/>
          <br/>
            <i>Default</i>: tls.crt<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>listenTLS</b></td>
        <td>boolean</td>
        <td>
          Defines if Metrics Service Endpoint uses TLS configuration<br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name is a unique name for the port that can be referred to by services.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>protocol</b></td>
        <td>enum</td>
        <td>
          Protocol for port. Must be UDP, TCP, or SCTP.<br/>
          <br/>
            <i>Enum</i>: TCP, UDP, SCTP<br/>
            <i>Default</i>: TCP<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>servicePort</b></td>
        <td>number</td>
        <td>
          ServicePort is the port number to expose on the service<br/>
          <br/>
            <i>Format</i>: int32<br/>
            <i>Default</i>: 9628<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>serviceType</b></td>
        <td>string</td>
        <td>
          Defines the service type for the Metrics Service Endpoint<br/>
          <br/>
            <i>Default</i>: ClusterIP<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.monitoring.resources
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecmonitoring)</sup></sup>



Resources defines parameters to configure resources for the Prometheus Exporter.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#pubsubpluseventbrokerspecmonitoringresourceslimits">limits</a></b></td>
        <td>object</td>
        <td>
          Define the container resource limits<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecmonitoringresourcesrequests">requests</a></b></td>
        <td>object</td>
        <td>
          Define the container resource requests<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.monitoring.resources.limits
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecmonitoringresources)</sup></sup>



Limits defines parameters to configure resource limits for the Prometheus Exporter.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>cpu</b></td>
        <td>string</td>
        <td>
          Define the container cpu resource limit<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>memory</b></td>
        <td>string</td>
        <td>
          Define the container memory resource limit<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.monitoring.resources.requests
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecmonitoringresources)</sup></sup>



requests defines parameters to configure resource requests for the Prometheus Exporter.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>cpu</b></td>
        <td>string</td>
        <td>
          Define the container cpu resource request<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>memory</b></td>
        <td>string</td>
        <td>
          Define the container memory resource request<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspec)</sup></sup>



NodeAssignment defines labels to constrain PubSubPlusEventBroker nodes to specific nodes

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>name</b></td>
        <td>enum</td>
        <td>
          Defines the name of broker node type that has the nodeAssignment spec defined<br/>
          <br/>
            <i>Enum</i>: Primary, Backup, Monitor<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspec">spec</a></b></td>
        <td>object</td>
        <td>
          If provided defines the labels to constrain the PubSubPlusEventBroker node to specific nodes<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindex)</sup></sup>



If provided defines the labels to constrain the PubSubPlusEventBroker node to specific nodes

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinity">affinity</a></b></td>
        <td>object</td>
        <td>
          Affinity if provided defines the conditional approach to assign PubSubPlusEventBroker nodes to specific nodes to which they can be scheduled<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>nodeSelector</b></td>
        <td>map[string]string</td>
        <td>
          NodeSelector if provided defines the exact labels of nodes to which PubSubPlusEventBroker nodes can be scheduled<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspectolerationsindex">tolerations</a></b></td>
        <td>[]object</td>
        <td>
          Toleration if provided defines the exact properties of the PubSubPlusEventBroker nodes can be scheduled on nodes with d matching taint.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspec)</sup></sup>



Affinity if provided defines the conditional approach to assign PubSubPlusEventBroker nodes to specific nodes to which they can be scheduled

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitynodeaffinity">nodeAffinity</a></b></td>
        <td>object</td>
        <td>
          Describes node affinity scheduling rules for the pod.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinity">podAffinity</a></b></td>
        <td>object</td>
        <td>
          Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinity">podAntiAffinity</a></b></td>
        <td>object</td>
        <td>
          Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.nodeAffinity
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinity)</sup></sup>



Describes node affinity scheduling rules for the pod.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindex">preferredDuringSchedulingIgnoredDuringExecution</a></b></td>
        <td>[]object</td>
        <td>
          The scheduler will prefer to schedule pods to nodes that satisfy
the affinity expressions specified by this field, but it may choose
a node that violates one or more of the expressions. The node that is
most preferred is the one with the greatest sum of weights, i.e.
for each node that meets all of the scheduling requirements (resource
request, requiredDuringScheduling affinity expressions, etc.),
compute a sum by iterating through the elements of this field and adding
"weight" to the sum if the node matches the corresponding matchExpressions; the
node(s) with the highest sum are the most preferred.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitynodeaffinityrequiredduringschedulingignoredduringexecution">requiredDuringSchedulingIgnoredDuringExecution</a></b></td>
        <td>object</td>
        <td>
          If the affinity requirements specified by this field are not met at
scheduling time, the pod will not be scheduled onto the node.
If the affinity requirements specified by this field cease to be met
at some point during pod execution (e.g. due to an update), the system
may or may not try to eventually evict the pod from its node.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitynodeaffinity)</sup></sup>



An empty preferred scheduling term matches all objects with implicit weight 0
(i.e. it's a no-op). A null preferred scheduling term matches no objects (i.e. is also a no-op).

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindexpreference">preference</a></b></td>
        <td>object</td>
        <td>
          A node selector term, associated with the corresponding weight.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>weight</b></td>
        <td>integer</td>
        <td>
          Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.<br/>
          <br/>
            <i>Format</i>: int32<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].preference
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindex)</sup></sup>



A node selector term, associated with the corresponding weight.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindexpreferencematchexpressionsindex">matchExpressions</a></b></td>
        <td>[]object</td>
        <td>
          A list of node selector requirements by node's labels.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindexpreferencematchfieldsindex">matchFields</a></b></td>
        <td>[]object</td>
        <td>
          A list of node selector requirements by node's fields.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].preference.matchExpressions[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindexpreference)</sup></sup>



A node selector requirement is a selector that contains values, a key, and an operator
that relates the key and values.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          The label key that the selector applies to.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>operator</b></td>
        <td>string</td>
        <td>
          Represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>values</b></td>
        <td>[]string</td>
        <td>
          An array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. If the operator is Gt or Lt, the values
array must have a single element, which will be interpreted as an integer.
This array is replaced during a strategic merge patch.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].preference.matchFields[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindexpreference)</sup></sup>



A node selector requirement is a selector that contains values, a key, and an operator
that relates the key and values.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          The label key that the selector applies to.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>operator</b></td>
        <td>string</td>
        <td>
          Represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>values</b></td>
        <td>[]string</td>
        <td>
          An array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. If the operator is Gt or Lt, the values
array must have a single element, which will be interpreted as an integer.
This array is replaced during a strategic merge patch.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitynodeaffinity)</sup></sup>



If the affinity requirements specified by this field are not met at
scheduling time, the pod will not be scheduled onto the node.
If the affinity requirements specified by this field cease to be met
at some point during pod execution (e.g. due to an update), the system
may or may not try to eventually evict the pod from its node.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitynodeaffinityrequiredduringschedulingignoredduringexecutionnodeselectortermsindex">nodeSelectorTerms</a></b></td>
        <td>[]object</td>
        <td>
          Required. A list of node selector terms. The terms are ORed.<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitynodeaffinityrequiredduringschedulingignoredduringexecution)</sup></sup>



A null or empty node selector term matches no objects. The requirements of
them are ANDed.
The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitynodeaffinityrequiredduringschedulingignoredduringexecutionnodeselectortermsindexmatchexpressionsindex">matchExpressions</a></b></td>
        <td>[]object</td>
        <td>
          A list of node selector requirements by node's labels.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitynodeaffinityrequiredduringschedulingignoredduringexecutionnodeselectortermsindexmatchfieldsindex">matchFields</a></b></td>
        <td>[]object</td>
        <td>
          A list of node selector requirements by node's fields.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[index].matchExpressions[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitynodeaffinityrequiredduringschedulingignoredduringexecutionnodeselectortermsindex)</sup></sup>



A node selector requirement is a selector that contains values, a key, and an operator
that relates the key and values.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          The label key that the selector applies to.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>operator</b></td>
        <td>string</td>
        <td>
          Represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>values</b></td>
        <td>[]string</td>
        <td>
          An array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. If the operator is Gt or Lt, the values
array must have a single element, which will be interpreted as an integer.
This array is replaced during a strategic merge patch.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[index].matchFields[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitynodeaffinityrequiredduringschedulingignoredduringexecutionnodeselectortermsindex)</sup></sup>



A node selector requirement is a selector that contains values, a key, and an operator
that relates the key and values.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          The label key that the selector applies to.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>operator</b></td>
        <td>string</td>
        <td>
          Represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>values</b></td>
        <td>[]string</td>
        <td>
          An array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. If the operator is Gt or Lt, the values
array must have a single element, which will be interpreted as an integer.
This array is replaced during a strategic merge patch.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAffinity
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinity)</sup></sup>



Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindex">preferredDuringSchedulingIgnoredDuringExecution</a></b></td>
        <td>[]object</td>
        <td>
          The scheduler will prefer to schedule pods to nodes that satisfy
the affinity expressions specified by this field, but it may choose
a node that violates one or more of the expressions. The node that is
most preferred is the one with the greatest sum of weights, i.e.
for each node that meets all of the scheduling requirements (resource
request, requiredDuringScheduling affinity expressions, etc.),
compute a sum by iterating through the elements of this field and adding
"weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the
node(s) with the highest sum are the most preferred.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindex">requiredDuringSchedulingIgnoredDuringExecution</a></b></td>
        <td>[]object</td>
        <td>
          If the affinity requirements specified by this field are not met at
scheduling time, the pod will not be scheduled onto the node.
If the affinity requirements specified by this field cease to be met
at some point during pod execution (e.g. due to a pod label update), the
system may or may not try to eventually evict the pod from its node.
When there are multiple elements, the lists of nodes corresponding to each
podAffinityTerm are intersected, i.e. all terms must be satisfied.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinity)</sup></sup>



The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinityterm">podAffinityTerm</a></b></td>
        <td>object</td>
        <td>
          Required. A pod affinity term, associated with the corresponding weight.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>weight</b></td>
        <td>integer</td>
        <td>
          weight associated with matching the corresponding podAffinityTerm,
in the range 1-100.<br/>
          <br/>
            <i>Format</i>: int32<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindex)</sup></sup>



Required. A pod affinity term, associated with the corresponding weight.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>topologyKey</b></td>
        <td>string</td>
        <td>
          This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching
the labelSelector in the specified namespaces, where co-located is defined as running on a node
whose value of the label with key topologyKey matches that of any node on which any of the
selected pods is running.
Empty topologyKey is not allowed.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermlabelselector">labelSelector</a></b></td>
        <td>object</td>
        <td>
          A label query over a set of resources, in this case pods.
If it's null, this PodAffinityTerm matches with no Pods.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>matchLabelKeys</b></td>
        <td>[]string</td>
        <td>
          MatchLabelKeys is a set of pod label keys to select which pods will
be taken into consideration. The keys are used to lookup values from the
incoming pod labels, those key-value labels are merged with `LabelSelector` as `key in (value)`
to select the group of existing pods which pods will be taken into consideration
for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming
pod labels will be ignored. The default value is empty.
The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.
Also, MatchLabelKeys cannot be set when LabelSelector isn't set.
This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>mismatchLabelKeys</b></td>
        <td>[]string</td>
        <td>
          MismatchLabelKeys is a set of pod label keys to select which pods will
be taken into consideration. The keys are used to lookup values from the
incoming pod labels, those key-value labels are merged with `LabelSelector` as `key notin (value)`
to select the group of existing pods which pods will be taken into consideration
for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming
pod labels will be ignored. The default value is empty.
The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.
Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.
This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermnamespaceselector">namespaceSelector</a></b></td>
        <td>object</td>
        <td>
          A label query over the set of namespaces that the term applies to.
The term is applied to the union of the namespaces selected by this field
and the ones listed in the namespaces field.
null selector and null or empty namespaces list means "this pod's namespace".
An empty selector ({}) matches all namespaces.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>namespaces</b></td>
        <td>[]string</td>
        <td>
          namespaces specifies a static list of namespace names that the term applies to.
The term is applied to the union of the namespaces listed in this field
and the ones selected by namespaceSelector.
null or empty namespaces list and null namespaceSelector means "this pod's namespace".<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.labelSelector
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinityterm)</sup></sup>



A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermlabelselectormatchexpressionsindex">matchExpressions</a></b></td>
        <td>[]object</td>
        <td>
          matchExpressions is a list of label selector requirements. The requirements are ANDed.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>matchLabels</b></td>
        <td>map[string]string</td>
        <td>
          matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.labelSelector.matchExpressions[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermlabelselector)</sup></sup>



A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          key is the label key that the selector applies to.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>operator</b></td>
        <td>string</td>
        <td>
          operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>values</b></td>
        <td>[]string</td>
        <td>
          values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.namespaceSelector
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinityterm)</sup></sup>



A label query over the set of namespaces that the term applies to.
The term is applied to the union of the namespaces selected by this field
and the ones listed in the namespaces field.
null selector and null or empty namespaces list means "this pod's namespace".
An empty selector ({}) matches all namespaces.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermnamespaceselectormatchexpressionsindex">matchExpressions</a></b></td>
        <td>[]object</td>
        <td>
          matchExpressions is a list of label selector requirements. The requirements are ANDed.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>matchLabels</b></td>
        <td>map[string]string</td>
        <td>
          matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.namespaceSelector.matchExpressions[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermnamespaceselector)</sup></sup>



A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          key is the label key that the selector applies to.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>operator</b></td>
        <td>string</td>
        <td>
          operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>values</b></td>
        <td>[]string</td>
        <td>
          values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinity)</sup></sup>



Defines a set of pods (namely those matching the labelSelector
relative to the given namespace(s)) that this pod should be
co-located (affinity) or not co-located (anti-affinity) with,
where co-located is defined as running on a node whose value of
the label with key <topologyKey> matches that of any node on which
a pod of the set of pods is running

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>topologyKey</b></td>
        <td>string</td>
        <td>
          This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching
the labelSelector in the specified namespaces, where co-located is defined as running on a node
whose value of the label with key topologyKey matches that of any node on which any of the
selected pods is running.
Empty topologyKey is not allowed.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindexlabelselector">labelSelector</a></b></td>
        <td>object</td>
        <td>
          A label query over a set of resources, in this case pods.
If it's null, this PodAffinityTerm matches with no Pods.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>matchLabelKeys</b></td>
        <td>[]string</td>
        <td>
          MatchLabelKeys is a set of pod label keys to select which pods will
be taken into consideration. The keys are used to lookup values from the
incoming pod labels, those key-value labels are merged with `LabelSelector` as `key in (value)`
to select the group of existing pods which pods will be taken into consideration
for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming
pod labels will be ignored. The default value is empty.
The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.
Also, MatchLabelKeys cannot be set when LabelSelector isn't set.
This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>mismatchLabelKeys</b></td>
        <td>[]string</td>
        <td>
          MismatchLabelKeys is a set of pod label keys to select which pods will
be taken into consideration. The keys are used to lookup values from the
incoming pod labels, those key-value labels are merged with `LabelSelector` as `key notin (value)`
to select the group of existing pods which pods will be taken into consideration
for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming
pod labels will be ignored. The default value is empty.
The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.
Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.
This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindexnamespaceselector">namespaceSelector</a></b></td>
        <td>object</td>
        <td>
          A label query over the set of namespaces that the term applies to.
The term is applied to the union of the namespaces selected by this field
and the ones listed in the namespaces field.
null selector and null or empty namespaces list means "this pod's namespace".
An empty selector ({}) matches all namespaces.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>namespaces</b></td>
        <td>[]string</td>
        <td>
          namespaces specifies a static list of namespace names that the term applies to.
The term is applied to the union of the namespaces listed in this field
and the ones selected by namespaceSelector.
null or empty namespaces list and null namespaceSelector means "this pod's namespace".<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].labelSelector
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindex)</sup></sup>




A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindexlabelselectormatchexpressionsindex">matchExpressions</a></b></td>
        <td>[]object</td>
        <td>
          matchExpressions is a list of label selector requirements. The requirements are ANDed.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>matchLabels</b></td>
        <td>map[string]string</td>
        <td>
          matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].labelSelector.matchExpressions[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindexlabelselector)</sup></sup>



A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          key is the label key that the selector applies to.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>operator</b></td>
        <td>string</td>
        <td>
          operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>values</b></td>
        <td>[]string</td>
        <td>
          values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].namespaceSelector
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindex)</sup></sup>



A label query over the set of namespaces that the term applies to.
The term is applied to the union of the namespaces selected by this field
and the ones listed in the namespaces field.
null selector and null or empty namespaces list means "this pod's namespace".
An empty selector ({}) matches all namespaces.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindexnamespaceselectormatchexpressionsindex">matchExpressions</a></b></td>
        <td>[]object</td>
        <td>
          matchExpressions is a list of label selector requirements. The requirements are ANDed.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>matchLabels</b></td>
        <td>map[string]string</td>
        <td>
          matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].namespaceSelector.matchExpressions[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindexnamespaceselector)</sup></sup>



A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          key is the label key that the selector applies to.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>operator</b></td>
        <td>string</td>
        <td>
          operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>values</b></td>
        <td>[]string</td>
        <td>
          values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAntiAffinity
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinity)</sup></sup>



Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindex">preferredDuringSchedulingIgnoredDuringExecution</a></b></td>
        <td>[]object</td>
        <td>
          The scheduler will prefer to schedule pods to nodes that satisfy
the anti-affinity expressions specified by this field, but it may choose
a node that violates one or more of the expressions. The node that is
most preferred is the one with the greatest sum of weights, i.e.
for each node that meets all of the scheduling requirements (resource
request, requiredDuringScheduling anti-affinity expressions, etc.),
compute a sum by iterating through the elements of this field and adding
"weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the
node(s) with the highest sum are the most preferred.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindex">requiredDuringSchedulingIgnoredDuringExecution</a></b></td>
        <td>[]object</td>
        <td>
          If the anti-affinity requirements specified by this field are not met at
scheduling time, the pod will not be scheduled onto the node.
If the anti-affinity requirements specified by this field cease to be met
at some point during pod execution (e.g. due to a pod label update), the
system may or may not try to eventually evict the pod from its node.
When there are multiple elements, the lists of nodes corresponding to each
podAffinityTerm are intersected, i.e. all terms must be satisfied.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinity)</sup></sup>



The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinityterm">podAffinityTerm</a></b></td>
        <td>object</td>
        <td>
          Required. A pod affinity term, associated with the corresponding weight.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>weight</b></td>
        <td>integer</td>
        <td>
          weight associated with matching the corresponding podAffinityTerm,
in the range 1-100.<br/>
          <br/>
            <i>Format</i>: int32<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindex)</sup></sup>



Required. A pod affinity term, associated with the corresponding weight.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>topologyKey</b></td>
        <td>string</td>
        <td>
          This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching
the labelSelector in the specified namespaces, where co-located is defined as running on a node
whose value of the label with key topologyKey matches that of any node on which any of the
selected pods is running.
Empty topologyKey is not allowed.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermlabelselector">labelSelector</a></b></td>
        <td>object</td>
        <td>
          A label query over a set of resources, in this case pods.
If it's null, this PodAffinityTerm matches with no Pods.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>matchLabelKeys</b></td>
        <td>[]string</td>
        <td>
          MatchLabelKeys is a set of pod label keys to select which pods will
be taken into consideration. The keys are used to lookup values from the
incoming pod labels, those key-value labels are merged with `LabelSelector` as `key in (value)`
to select the group of existing pods which pods will be taken into consideration
for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming
pod labels will be ignored. The default value is empty.
The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.
Also, MatchLabelKeys cannot be set when LabelSelector isn't set.
This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>mismatchLabelKeys</b></td>
        <td>[]string</td>
        <td>
          MismatchLabelKeys is a set of pod label keys to select which pods will
be taken into consideration. The keys are used to lookup values from the
incoming pod labels, those key-value labels are merged with `LabelSelector` as `key notin (value)`
to select the group of existing pods which pods will be taken into consideration
for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming
pod labels will be ignored. The default value is empty.
The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.
Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.
This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermnamespaceselector">namespaceSelector</a></b></td>
        <td>object</td>
        <td>
          A label query over the set of namespaces that the term applies to.
The term is applied to the union of the namespaces selected by this field
and the ones listed in the namespaces field.
null selector and null or empty namespaces list means "this pod's namespace".
An empty selector ({}) matches all namespaces.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>namespaces</b></td>
        <td>[]string</td>
        <td>
          namespaces specifies a static list of namespace names that the term applies to.
The term is applied to the union of the namespaces listed in this field
and the ones selected by namespaceSelector.
null or empty namespaces list and null namespaceSelector means "this pod's namespace".<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.labelSelector
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinityterm)</sup></sup>




A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermlabelselectormatchexpressionsindex">matchExpressions</a></b></td>
        <td>[]object</td>
        <td>
          matchExpressions is a list of label selector requirements. The requirements are ANDed.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>matchLabels</b></td>
        <td>map[string]string</td>
        <td>
          matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.labelSelector.matchExpressions[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermlabelselector)</sup></sup>



A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          key is the label key that the selector applies to.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>operator</b></td>
        <td>string</td>
        <td>
          operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>values</b></td>
        <td>[]string</td>
        <td>
          values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.namespaceSelector
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinityterm)</sup></sup>



A label query over the set of namespaces that the term applies to.
The term is applied to the union of the namespaces selected by this field
and the ones listed in the namespaces field.
null selector and null or empty namespaces list means "this pod's namespace".
An empty selector ({}) matches all namespaces.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermnamespaceselectormatchexpressionsindex">matchExpressions</a></b></td>
        <td>[]object</td>
        <td>
          matchExpressions is a list of label selector requirements. The requirements are ANDed.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>matchLabels</b></td>
        <td>map[string]string</td>
        <td>
          matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.namespaceSelector.matchExpressions[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermnamespaceselector)</sup></sup>



A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          key is the label key that the selector applies to.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>operator</b></td>
        <td>string</td>
        <td>
          operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>values</b></td>
        <td>[]string</td>
        <td>
          values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinity)</sup></sup>



Defines a set of pods (namely those matching the labelSelector
relative to the given namespace(s)) that this pod should be
co-located (affinity) or not co-located (anti-affinity) with,
where co-located is defined as running on a node whose value of
the label with key <topologyKey> matches that of any node on which
a pod of the set of pods is running

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>topologyKey</b></td>
        <td>string</td>
        <td>
          This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching
the labelSelector in the specified namespaces, where co-located is defined as running on a node
whose value of the label with key topologyKey matches that of any node on which any of the
selected pods is running.
Empty topologyKey is not allowed.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindexlabelselector">labelSelector</a></b></td>
        <td>object</td>
        <td>
          A label query over a set of resources, in this case pods.
If it's null, this PodAffinityTerm matches with no Pods.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>matchLabelKeys</b></td>
        <td>[]string</td>
        <td>
          MatchLabelKeys is a set of pod label keys to select which pods will
be taken into consideration. The keys are used to lookup values from the
incoming pod labels, those key-value labels are merged with `LabelSelector` as `key in (value)`
to select the group of existing pods which pods will be taken into consideration
for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming
pod labels will be ignored. The default value is empty.
The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.
Also, MatchLabelKeys cannot be set when LabelSelector isn't set.
This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>mismatchLabelKeys</b></td>
        <td>[]string</td>
        <td>
          MismatchLabelKeys is a set of pod label keys to select which pods will
be taken into consideration. The keys are used to lookup values from the
incoming pod labels, those key-value labels are merged with `LabelSelector` as `key notin (value)`
to select the group of existing pods which pods will be taken into consideration
for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming
pod labels will be ignored. The default value is empty.
The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.
Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.
This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindexnamespaceselector">namespaceSelector</a></b></td>
        <td>object</td>
        <td>
          A label query over the set of namespaces that the term applies to.
The term is applied to the union of the namespaces selected by this field
and the ones listed in the namespaces field.
null selector and null or empty namespaces list means "this pod's namespace".
An empty selector ({}) matches all namespaces.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>namespaces</b></td>
        <td>[]string</td>
        <td>
          namespaces specifies a static list of namespace names that the term applies to.
The term is applied to the union of the namespaces listed in this field
and the ones selected by namespaceSelector.
null or empty namespaces list and null namespaceSelector means "this pod's namespace".<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].labelSelector
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindex)</sup></sup>



A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindexlabelselectormatchexpressionsindex">matchExpressions</a></b></td>
        <td>[]object</td>
        <td>
          matchExpressions is a list of label selector requirements. The requirements are ANDed.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>matchLabels</b></td>
        <td>map[string]string</td>
        <td>
          matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].labelSelector.matchExpressions[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindexlabelselector)</sup></sup>



A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          key is the label key that the selector applies to.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>operator</b></td>
        <td>string</td>
        <td>
          operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>values</b></td>
        <td>[]string</td>
        <td>
          values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].namespaceSelector
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindex)</sup></sup>



A label query over the set of namespaces that the term applies to.
The term is applied to the union of the namespaces selected by this field
and the ones listed in the namespaces field.
null selector and null or empty namespaces list means "this pod's namespace".
An empty selector ({}) matches all namespaces.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindexnamespaceselectormatchexpressionsindex">matchExpressions</a></b></td>
        <td>[]object</td>
        <td>
          matchExpressions is a list of label selector requirements. The requirements are ANDed.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>matchLabels</b></td>
        <td>map[string]string</td>
        <td>
          matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].namespaceSelector.matchExpressions[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspecaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindexnamespaceselector)</sup></sup>



A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          key is the label key that the selector applies to.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>operator</b></td>
        <td>string</td>
        <td>
          operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>values</b></td>
        <td>[]string</td>
        <td>
          values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.nodeAssignment[index].spec.tolerations[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecnodeassignmentindexspec)</sup></sup>



The pod this Toleration is attached to tolerates any taint that matches the triple <key,value,effect> using the matching operator <operator>.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>effect</b></td>
        <td>string</td>
        <td>
          Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>
          Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>operator</b></td>
        <td>string</td>
        <td>
          Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>tolerationSeconds</b></td>
        <td>integer</td>
        <td>
          TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.<br/>
          <br/>
            <i>Format</i>: int64<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>value</b></td>
        <td>string</td>
        <td>
          Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.securityContext
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspec)</sup></sup>



SecurityContext defines the pod security context for the event broker.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>fsGroup</b></td>
        <td>number</td>
        <td>
          Specifies fsGroup in pod security context. 0 or unset defaults either to 1000002, or if OpenShift detected to unspecified (see documentation)<br/>
          <br/>
            <i>Format</i>: int64<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>runAsUser</b></td>
        <td>number</td>
        <td>
          Specifies runAsUser in pod security context. 0 or unset defaults either to 1000001, or if OpenShift detected to unspecified (see documentation)<br/>
          <br/>
            <i>Format</i>: int64<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.service
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspec)</sup></sup>



Service defines broker service details.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>annotations</b></td>
        <td>map[string]string</td>
        <td>
          Annotations allows adding provider-specific service annotations<br/>
          <br/>
            <i>Default</i>: map[]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecserviceportsindex">ports</a></b></td>
        <td>[]object</td>
        <td>
          Ports specifies the ports to expose PubSubPlusEventBroker services.<br/>
          <br/>
            <i>Default</i>: [map[containerPort:2222 name:tcp-ssh protocol:TCP servicePort:2222] map[containerPort:8080 name:tcp-semp protocol:TCP servicePort:8080] map[containerPort:1943 name:tls-semp protocol:TCP servicePort:1943] map[containerPort:55555 name:tcp-smf protocol:TCP servicePort:55555] map[containerPort:55003 name:tcp-smfcomp protocol:TCP servicePort:55003] map[containerPort:55443 name:tls-smf protocol:TCP servicePort:55443] map[containerPort:55556 name:tcp-smfroute protocol:TCP servicePort:55556] map[containerPort:8008 name:tcp-web protocol:TCP servicePort:8008] map[containerPort:1443 name:tls-web protocol:TCP servicePort:1443] map[containerPort:9000 name:tcp-rest protocol:TCP servicePort:9000] map[containerPort:9443 name:tls-rest protocol:TCP servicePort:9443] map[containerPort:5672 name:tcp-amqp protocol:TCP servicePort:5672] map[containerPort:5671 name:tls-amqp protocol:TCP servicePort:5671] map[containerPort:1883 name:tcp-mqtt protocol:TCP servicePort:1883] map[containerPort:8883 name:tls-mqtt protocol:TCP servicePort:8883] map[containerPort:8000 name:tcp-mqttweb protocol:TCP servicePort:8000] map[containerPort:8443 name:tls-mqttweb protocol:TCP servicePort:8443]]<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>string</td>
        <td>
          ServiceType specifies how to expose the broker services. Options include ClusterIP, NodePort, LoadBalancer (default).<br/>
          <br/>
            <i>Default</i>: LoadBalancer<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.service.ports[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecservice)</sup></sup>



Port defines parameters configure Service details for the Broker

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>containerPort</b></td>
        <td>number</td>
        <td>
          Port number to expose on the pod.<br/>
          <br/>
            <i>Format</i>: int32<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Unique name for the port that can be referred to by services.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>protocol</b></td>
        <td>enum</td>
        <td>
          Protocol for port. Must be UDP, TCP, or SCTP.<br/>
          <br/>
            <i>Enum</i>: TCP, UDP, SCTP<br/>
            <i>Default</i>: TCP<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>servicePort</b></td>
        <td>number</td>
        <td>
          Port number to expose on the service<br/>
          <br/>
            <i>Format</i>: int32<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.serviceAccount
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspec)</sup></sup>



ServiceAccount defines a ServiceAccount dedicated to the PubSubPlusEventBroker

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name specifies the name of an existing ServiceAccount dedicated to the PubSubPlusEventBroker.
If this value is missing a new ServiceAccount will be created.<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.storage
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspec)</sup></sup>



Storage defines storage details for the broker.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#pubsubpluseventbrokerspecstoragecustomvolumemountindex">customVolumeMount</a></b></td>
        <td>[]object</td>
        <td>
          CustomVolumeMount can be used to show the data volume should be mounted instead of using a storage class.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>messagingNodeStorageSize</b></td>
        <td>string</td>
        <td>
          MessagingNodeStorageSize if provided will assign the minimum persistent storage to be used by the message nodes.<br/>
          <br/>
            <i>Default</i>: 30Gi<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>monitorNodeStorageSize</b></td>
        <td>string</td>
        <td>
          MonitorNodeStorageSize if provided this will create and assign the minimum recommended storage to Monitor pods.<br/>
          <br/>
            <i>Default</i>: 3Gi<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>slow</b></td>
        <td>boolean</td>
        <td>
          Slow indicate slow storage is in use, an example is NFS.<br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>useStorageClass</b></td>
        <td>string</td>
        <td>
          UseStrorageClass Name of the StorageClass to be used to request persistent storage volumes. If undefined, the "default" StorageClass will be used.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.storage.customVolumeMount[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecstorage)</sup></sup>



StorageCustomVolumeMount defines Image details and pulling configurations

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>name</b></td>
        <td>enum</td>
        <td>
          Defines the name of PubSubPlusEventBroker node type that has the customVolumeMount spec defined<br/>
          <br/>
            <i>Enum</i>: Primary, Backup, Monitor<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerspecstoragecustomvolumemountindexpersistentvolumeclaim">persistentVolumeClaim</a></b></td>
        <td>object</td>
        <td>
          Defines the customVolumeMount that can be used mount the data volume instead of using a storage class<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.storage.customVolumeMount[index].persistentVolumeClaim
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspecstoragecustomvolumemountindex)</sup></sup>



Defines the customVolumeMount that can be used mount the data volume instead of using a storage class

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>claimName</b></td>
        <td>string</td>
        <td>
          Defines the claimName of a custom PersistentVolumeClaim to be used instead<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.systemScaling
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspec)</sup></sup>



SystemScaling provides exact fine-grained specification of the event broker scaling parameters
and the assigned CPU / memory resources to the Pod.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>maxConnections</b></td>
        <td>integer</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: 100<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>maxQueueMessages</b></td>
        <td>integer</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: 100<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>maxSpoolUsage</b></td>
        <td>integer</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: 1000<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>messagingNodeCpu</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: 2<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>messagingNodeMemory</b></td>
        <td>string</td>
        <td>
          <br/>
          <br/>
            <i>Default</i>: 4025Mi<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.spec.tls
<sup><sup>[↩ Parent](#pubsubpluseventbrokerspec)</sup></sup>



TLS provides TLS configuration for the event broker.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>certFilename</b></td>
        <td>string</td>
        <td>
          Name of the Certificate file in the `serverCertificatesSecret`<br/>
          <br/>
            <i>Default</i>: tls.key<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>certKeyFilename</b></td>
        <td>string</td>
        <td>
          Name of the Key file in the `serverCertificatesSecret`<br/>
          <br/>
            <i>Default</i>: tls.crt<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>enabled</b></td>
        <td>boolean</td>
        <td>
          Enabled true enables TLS for the broker.<br/>
          <br/>
            <i>Default</i>: false<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>serverTlsConfigSecret</b></td>
        <td>string</td>
        <td>
          Specifies the tls configuration secret to be used for the broker<br/>
          <br/>
            <i>Default</i>: example-tls-secret<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.status
<sup><sup>[↩ Parent](#pubsubpluseventbroker)</sup></sup>



EventBrokerStatus defines the observed state of the PubSubPlusEventBroker

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#pubsubpluseventbrokerstatusbroker">broker</a></b></td>
        <td>object</td>
        <td>
          Broker section provides the broker status<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerstatusconditionsindex">conditions</a></b></td>
        <td>[]object</td>
        <td>
          Conditions provide information about the observed status of the deployment<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>podsList</b></td>
        <td>[]string</td>
        <td>
          PodsList are the names of the eventbroker and optionally the monitoring pods<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#pubsubpluseventbrokerstatusprometheusmonitoring">prometheusMonitoring</a></b></td>
        <td>object</td>
        <td>
          Monitoring sectionprovides monitoring support status<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.status.broker
<sup><sup>[↩ Parent](#pubsubpluseventbrokerstatus)</sup></sup>



Broker section provides the broker status

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>adminCredentialsSecret</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>brokerImage</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>haDeployment</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>serviceName</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>serviceType</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>statefulSets</b></td>
        <td>[]string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>tlsSecret</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>tlsSupport</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.status.conditions[index]
<sup><sup>[↩ Parent](#pubsubpluseventbrokerstatus)</sup></sup>



Condition contains details for one aspect of the current state of this API Resource.
---
This struct is intended for direct use as an array at the field path .status.conditions.  For example,


	type FooStatus struct{
	    // Represents the observations of a foo's current state.
	    // Known .status.conditions.type are: "Available", "Progressing", and "Degraded"
	    // +patchMergeKey=type
	    // +patchStrategy=merge
	    // +listType=map
	    // +listMapKey=type
	    Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`


	    // other fields
	}

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>lastTransitionTime</b></td>
        <td>string</td>
        <td>
          lastTransitionTime is the last time the condition transitioned from one status to another.
This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.<br/>
          <br/>
            <i>Format</i>: date-time<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>message</b></td>
        <td>string</td>
        <td>
          message is a human readable message indicating details about the transition.
This may be an empty string.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>reason</b></td>
        <td>string</td>
        <td>
          reason contains a programmatic identifier indicating the reason for the condition's last transition.
Producers of specific condition types may define expected values and meanings for this field,
and whether the values are considered a guaranteed API.
The value should be a CamelCase string.
This field may not be empty.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>status</b></td>
        <td>enum</td>
        <td>
          status of the condition, one of True, False, Unknown.<br/>
          <br/>
            <i>Enum</i>: True, False, Unknown<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>string</td>
        <td>
          type of condition in CamelCase or in foo.example.com/CamelCase.
---
Many .condition.type values are consistent across resources like Available, but because arbitrary conditions can be
useful (see .node.status.conditions), the ability to deconflict is important.
The regex it matches is (dns1123SubdomainFmt/)?(qualifiedNameFmt)<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>observedGeneration</b></td>
        <td>integer</td>
        <td>
          observedGeneration represents the .metadata.generation that the condition was set based upon.
For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date
with respect to the current state of the instance.<br/>
          <br/>
            <i>Format</i>: int64<br/>
            <i>Minimum</i>: 0<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PubSubPlusEventBroker.status.prometheusMonitoring
<sup><sup>[↩ Parent](#pubsubpluseventbrokerstatus)</sup></sup>



Monitoring sectionprovides monitoring support status

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>enabled</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>exporterImage</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>serviceName</b></td>
        <td>string</td>
        <td>
          <br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>
