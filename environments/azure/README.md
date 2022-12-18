## How to get the most updated azure-pipeline-schema?

There's yamlschema endpoint in Azure DevOps REST API that returns schema for YAML pipeline: 
```
GET https://dev.azure.com/{organization}/_apis/distributedtask/yamlschema?api-version=6
```

change the api-version for getting the last version of yamlschema.
