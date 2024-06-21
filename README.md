## Google Cloud Services Infrastructure

_The cloud services used_

<p style="text-align: center; background-color: #eee; display: inline-block; padding: 14px 20px; border-radius: 15px;">
<img src="https://github.com/C241-PS120/.github/blob/main/profile/image/cloud-infrastructure.png?raw=true" width="800"/>
</p>

_CI/CD Implementation on GitHub Actions with Keyless Authentication_
<p style="text-align: center; background-color: #eee; display: inline-block; padding: 14px 20px; border-radius: 15px;">
<img src="https://github.com/C241-PS120/bangkit-cloud-computing/assets/87903309/37e1c3c3-9d84-43f3-af07-e7a455d50d8f?raw=true" width="800"/>
</p>


**Powered by:**

<p style="text-align: center; background-color: #eee; display: inline-block; padding: 14px 20px; border-radius: 15px;">
<img src="https://upload.wikimedia.org/wikipedia/commons/5/51/Google_Cloud_logo.svg" width="250"/>
</p>

Google Cloud Platform (GCP) is a Google-provided set of cloud computing services. It is a platform that offers computing infrastructure and services for running applications, storing and managing data, and developing software solutions.

## API Documentation
[Bump.sh - Coptas Backend API Documentation](https://bump.sh/coptas/doc/backend-api)

## Technology Used for Model Service

[Model Repository](https://github.com/C241-PS120/bangkit-cloud-computing/tree/model)

### Cloud Run

<img src="https://www.svgrepo.com/show/375383/cloud-run.svg" width="100" height="70"/>
Cloud Run: For deploying Model API.

Service details:

```YAML
Location          : asia-southeast2 (Jakarta)
Min-Instance      : 1
VCpu              : 2
Memory            : 4GB
Runtime           : python:3.12-slim
```

### Cloud Storage

<img src="https://symbols.getvecta.com/stencil_4/47_google-cloud-storage.fee263d33a.svg" width="100" height="50"/>

Cloud Storage: For storing the Image that is predicted and the AI Model.

```YAML
Location Type   : Single-Region
Location        : asia-southeast2 (Jakarta)
Storage Class   : Standard
```

## Technology Used for Article Service

[Article Repository](https://github.com/C241-PS120/bangkit-cloud-computing/tree/article)

### Cloud Run

<img src="https://www.svgrepo.com/show/375383/cloud-run.svg" width="100" height="70"/>
Cloud Run: For deploying Article API.

Service details:

```YAML
Location          : asia-southeast2 (Jakarta)
Min-Instance      : 1
VCpu              : 1
Memory            : 512MB
Runtime           : golang:1.22.3
```

### Cloud Storage

<img src="https://symbols.getvecta.com/stencil_4/47_google-cloud-storage.fee263d33a.svg" width="100" height="50"/>

Cloud Storage: For storing the article Image.

```YAML
Location Type   : Single-Region
Location        : asia-southeast2 (Jakarta)
Storage Class   : Standard
```

### Cloud SQL

<img src="https://www.svgrepo.com/show/375389/cloud-sql.svg" width="120" height="100"/>

Cloud SQL: For storing the Article

Service details:

```YAML
Database Engine   : Mysql
Location          : asia-southeast2 (Jakarta)
```
