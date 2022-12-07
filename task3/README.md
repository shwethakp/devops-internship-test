# Kubernetes

## 1. Create a Kubernetes cluster

Create a Kubernetes cluster in your local machine using [kind](https://kind.sigs.k8s.io/), [minikube](https://minikube.sigs.k8s.io/docs/start/), or whatever other k8s flavor you prefer.

## 2. Deploy the application

1. Create a Kubernetes manifest file that deploys the application to the cluster.
   1. Create the manifest file in the `task3` directory.
   2. The manifest should contain a `Deployment` and a `Service` for the application.
   3. For the sake of simplicity, a `Deployment` and `Service` for the database are **not** required.
   4. If you didn't complete [Task 1.5](../task1/README.md#5-push-the-image-to-docker-hub), you can use the [nginx](https://hub.docker.com/_/nginx) image instead of the application image.
2. Deploy the application to the cluster using `kubectl`.
   1. Deploy the application using `kubectl apply -f <manifest-file>` command and verify that the application is running. You can use `kubectl get pods` and `kubectl get services` to verify that the application is running.
   2. Provide the output of the `kubectl get pods` and `kubectl get services` commands in the `task3/output.txt` file.

## 3. ConfigMaps

1. Add a `ConfigMap` to the manifest file created in the previous step. The `ConfigMap` should contain the following key-value pairs:
   1. `APP_COLOR` - `blue`
   2. `APP_MODIFICATION_TIME` - `2021-01-01T00:00:00Z`
2. Mount the `ConfigMap` as an environment variable in the application container.
   1. The environment variable should be named `APP_COLOR`.
   2. The environment variable should contain the value of the `APP_COLOR` key.
3. *EXTRA*: Mount the `ConfigMap` as a volume in the application container.
   1. The volume should be mounted to the `/etc/config` directory.
   2. The volume should contain the `APP_COLOR` and `APP_MODIFICATION_TIME` keys.

## 4. Secrets

1. Add a `Secret` to the manifest file created in the previous step. The `Secret` should contain the following key-value pairs:
   1. `SUPER_SECRET_VALUE` - `1234567890`
2. Mount the `Secret` as an environment variable in the application container.
   1. The environment variable should be named `SUPER_SECRET_ENV_VAR`.
   2. The environment variable should contain the value of the `SUPER_SECRET_VALUE` key.

## 5. Helm (Optional)

This task is optional in case you want to go the extra mile. Don't worry if you don't have time to complete it, or if you don't know Helm. You can skip to [Task 6](../task6/README.md).

1. Create a Helm chart for the application.
   1. The chart should be created in the `task3` directory.
   2. The chart should contain a `Deployment` and a `Service` for the application.
   3. The chart should contain a `ConfigMap` and a `Secret` for the application.
   4. The chart should contain a `values.yaml` file that contains the following values:
      1. `image.repository` - (The image repository of the application, or `nginx` if you didn't complete [Task 1.5](../task1/README.md#5-push-the-image-to-docker-hub))
      2. `image.tag` - `latest` (or the image tag of the application, if you completed [Task 1.5](../task1/README.md#5-push-the-image-to-docker-hub))
      3. `image.pullPolicy` - `IfNotPresent`
      4. `service.type` - `ClusterIP`
      5. `service.port` - `80`
      6. `config.APP_COLOR` - `blue`
      7. `config.APP_MODIFICATION_TIME` - `2021-01-01T00:00:00Z`
      8. `secret.SUPER_SECRET_VALUE` - `1234567890`
