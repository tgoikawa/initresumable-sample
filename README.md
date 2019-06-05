# initresumable-sample

## Prepare

1. 以下のcorsをGCSのバケットに設定します.
	```js
		[{"maxAgeSeconds": 3600, "method": ["GET", "POST", "PUT", "DELETE"], "origin": ["*"], "responseHeader": ["Content-Type", "x-goog-resumable"]}]
	```

## How to deploy

1. `dep ensure` します

2. `gcloud app deploy` します

## How to run local dev_appserver.py

1. 環境変数 `GOOGLE_CLOUD_PROJECT` を設定します

2. `dev_appserver app.yaml` を実行します
