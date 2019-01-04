<!-- Code generated by solo-kit. DO NOT EDIT. -->

### Package: `core.solo.io` 
##### Types:


- [Status](#Status)
- [State](#State)
  



##### Source File: [github.com/solo-io/solo-kit/api/v1/status.proto](https://github.com/solo-io/solo-kit/blob/master/api/v1/status.proto)





---
### <a name="Status">Status</a>

 
*
Status indicates whether a resource has been (in)validated by a reporter in the system.
Statuses are meant to be read-only by users

```yaml
"state": .core.solo.io.Status.State
"reason": string
"reported_by": string
"subresource_statuses": map<string, .core.solo.io.Status>

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `state` | [.core.solo.io.Status.State](status.proto.sk.md#State) | State is the enum indicating the state of the resource |  |
| `reason` | `string` | Reason is a description of the error for Rejected resources. If the resource is pending or accepted, this field will be empty |  |
| `reported_by` | `string` | Reference to the reporter who wrote this status |  |
| `subresource_statuses` | `map<string, .core.solo.io.Status>` | Reference to statuses (by resource-ref string: "Kind.Namespace.Name") of subresources of the parent resource |  |




---
### <a name="State">State</a>



| Name | Description |
| ----- | ----------- | 
| `Pending` | Pending status indicates the resource has not yet been validated |
| `Accepted` | Accepted indicates the resource has been validated |
| `Rejected` | Rejected indicates an invalid configuration by the user Rejected resources may be propagated to the xDS server depending on their severity |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->