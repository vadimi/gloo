
---
title: "stitching_info.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `graphql.gloo.solo.io` 
#### Types:


- [GraphQLToolsStitchingInput](#graphqltoolsstitchinginput)
- [Schema](#schema)
- [TypeMergeConfig](#typemergeconfig)
- [GraphQlToolsStitchingOutput](#graphqltoolsstitchingoutput)
  



##### Source File: [github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/graphql/v1alpha1/stitching_info.proto](https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/enterprise/options/graphql/v1alpha1/stitching_info.proto)





---
### GraphQLToolsStitchingInput

 
------------- Graphql Tools JS Input -------------
This is not user-facing and is only used to pass to the graphql-tools js script
This is the message which the graphql-tools js script will consume

```yaml
"subschemas": []graphql.gloo.solo.io.GraphQLToolsStitchingInput.Schema

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `subschemas` | [[]graphql.gloo.solo.io.GraphQLToolsStitchingInput.Schema](../stitching_info.proto.sk/#schema) |  |




---
### Schema



```yaml
"name": string
"schema": string
"typeMergeConfig": map<string, .graphql.gloo.solo.io.GraphQLToolsStitchingInput.Schema.TypeMergeConfig>

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `name` | `string` | name of the subschema, arbitrary name but must be unique in a gateway schema. generally generated from the graphql schema ref. |
| `schema` | `string` | GraphQL schema SDL for the subschema. |
| `typeMergeConfig` | `map<string, .graphql.gloo.solo.io.GraphQLToolsStitchingInput.Schema.TypeMergeConfig>` | Type merge config that the graphql-tools stitching script needs to generate stitching info for the data plane. |




---
### TypeMergeConfig



```yaml
"selectionSet": string
"fieldName": string
"args": map<string, string>

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `selectionSet` | `string` |  |
| `fieldName` | `string` |  |
| `args` | `map<string, string>` |  |




---
### GraphQlToolsStitchingOutput

 
------------- Graphql Tools JS Out ------------------
This is not user-facing and is only used to pass data back from the graphql-tools js script
The message that is the output of the graphql tools stitching info script

```yaml
"fieldNodesByType": map<string, .envoy.config.resolver.stitching.v2.FieldNodes>
"fieldNodesByField": map<string, .envoy.config.resolver.stitching.v2.FieldNodeMap>
"mergedTypes": map<string, .envoy.config.resolver.stitching.v2.MergedTypeConfig>
"stitchedSchema": string

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `fieldNodesByType` | `map<string, .envoy.config.resolver.stitching.v2.FieldNodes>` |  |
| `fieldNodesByField` | `map<string, .envoy.config.resolver.stitching.v2.FieldNodeMap>` |  |
| `mergedTypes` | `map<string, .envoy.config.resolver.stitching.v2.MergedTypeConfig>` |  |
| `stitchedSchema` | `string` |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->
