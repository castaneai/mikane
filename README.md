# mikane
mf api server

## Configuration

```yaml
# appengine/secret.yaml
env_variables:
  MF_SESSION: "<mf session id>"
```

## Deploy

```
goapp deploy -application <gcp_project_id> -version <version> appengine/app.yaml
```
