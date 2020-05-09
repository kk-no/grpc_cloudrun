# cloud build を使用して docker image のビルド

gcloud builds submit --tag gcr.io/{project_name}/{registry_name}

# デプロイ

gcloud run deploy --image gcr.io/{project_name}/{registry_name}