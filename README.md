# mikane
mf api server

## Configuration

```yaml
# appengine/secret.yaml
env_variables:
  MF_SESSION: "<mf session id>"
```

## Deploy
deploying on CircleCI

### Environment Variables

|name|value|
|:---|:----|
|GCP_PROJECT_ID| project id on Google Cloud Platform|
|GCP_SECRET_KEY| base64-encoded secret.json content on Google Cloud Platform|
|MF_SESSION| mf session id|
