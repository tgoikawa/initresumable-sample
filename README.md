# initresumable-sample

## Prepare

1. 以下のcorsをGCSのバケット `resumable_test` に設定します.
	```js
	[{"maxAgeSeconds": 3600, "method": ["GET", "POST", "PUT", "DELETE"], "origin": ["*"], "responseHeader": ["Content-Type", "x-goog-resumable"]}]
	```

2. `dep ensure` します

## How to deploy

1. `gcloud app deploy` します

## How to run local dev_appserver.py

1. 環境変数 `GOOGLE_CLOUD_PROJECT` に自分のプロジェクトIDを設定します

2. `dev_appserver app.yaml` を実行します
