# cloud run job example

# やること
Cloud Run Jobの動作確認をする
https://cloud.google.com/run/docs/create-jobs


# Cloud Run Jobsだけで何ができるの？
## ドキュメント
- https://cloud.google.com/run/docs/create-jobs
- https://cloud.google.com/run/docs/execute/jobs
- https://cloud.google.com/run/docs/configuring/containers#configure-entrypoint
- https://cloud.google.com/run/docs/configuring/environment-variables#setting

## 環境変数を用意
```
export GOOGLE_PROJECT=stock-data-dev
export REGION=europe-west9
export IMAGE_URL=asia.gcr.io/stock-data-dev/cloud-run-job-example:latest
export JOB_NAME=sample-job
```

## イメージを用意
```
docker image build .  --platform amd64 -t $IMAGE_URL
docker container run --rm $IMAGE_URL
docker image push $IMAGE_URL
```


## コマンドフラグなし -> 動作する
```
gcloud beta run jobs create $JOB_NAME \
    --image $IMAGE_URL \
    --project $GOOGLE_PROJECT \
    --region $REGION
gcloud beta run jobs execute $JOB_NAME \
    --project $GOOGLE_PROJECT \
    --region $REGION
gcloud beta run jobs delete $JOB_NAME \
    --project $GOOGLE_PROJECT \
    --region $REGION
```

## コマンドと引数は指定できる
```
gcloud beta run jobs create $JOB_NAME \
    --image $IMAGE_URL \
    --command "/app" \
    --args a,b,c,100 \
    --project $GOOGLE_PROJECT \
    --region $REGION
gcloud beta run jobs execute $JOB_NAME \
    --project $GOOGLE_PROJECT \
    --region $REGION
gcloud beta run jobs delete $JOB_NAME \
    --project $GOOGLE_PROJECT \
    --region $REGION
```

## フラグっぽい引数は指定する -> エラー
```
gcloud beta run jobs create $JOB_NAME \
    --image $IMAGE_URL \
    --command "/app" \
    --args --user,hoge \
    --project $GOOGLE_PROJECT \
    --region $REGION
```

##  yamlで無理やり更新を試みる -> servicesと違ってjobsはreplaceコマンド使えない
なので、今のところフラグっぽい引数を指定することは不可能っぽい
```
gcloud beta run jobs describe $JOB_NAME --format export \
    --project $GOOGLE_PROJECT \
    --region $REGION
gcloud beta run jobs replace sample_job.yaml \
    --project $GOOGLE_PROJECT \
    --region $REGION
```

## 環境変数を指定できる
```
gcloud beta run jobs create $JOB_NAME \
    --image $IMAGE_URL \
    --set-env-vars FOO=foo,BAR=bar \
    --project $GOOGLE_PROJECT \
    --region $REGION
gcloud beta run jobs execute $JOB_NAME \
    --project $GOOGLE_PROJECT \
    --region $REGION
gcloud beta run jobs delete $JOB_NAME \
    --project $GOOGLE_PROJECT \
    --region $REGION
```

## tasksを指定する -> 並列に実行される
```
gcloud beta run jobs create $JOB_NAME \
    --tasks 3 \
    --image $IMAGE_URL \
    --project $GOOGLE_PROJECT \
    --region $REGION
gcloud beta run jobs execute $JOB_NAME \
    --project $GOOGLE_PROJECT \
    --region $REGION
gcloud beta run jobs delete $JOB_NAME \
    --project $GOOGLE_PROJECT \
    --region $REGION
```

## tasksとparallelismを指定する -> 並列数を制御できる
```
gcloud beta run jobs create $JOB_NAME \
    --tasks 3 \
    --parallelism 1 \
    --image $IMAGE_URL \
    --project $GOOGLE_PROJECT \
    --region $REGION
gcloud beta run jobs execute $JOB_NAME \
    --project $GOOGLE_PROJECT \
    --region $REGION
gcloud beta run jobs delete $JOB_NAME \
    --project $GOOGLE_PROJECT \
    --region $REGION
```


## その他使ったコマンド
```
gcloud beta run jobs execute $JOB_NAME \
    --project $GOOGLE_PROJECT \
    --region $REGION

gcloud beta run jobs execute $JOB_NAME \
    --project $GOOGLE_PROJECT \
    --region $REGION

gcloud beta run jobs list \
    --project $GOOGLE_PROJECT \
    --region $REGION

gcloud beta run jobs executions list\
    --project $GOOGLE_PROJECT \
    --region $REGION

gcloud beta run jobs executions describe sample-job-m9xkv \
    --project $GOOGLE_PROJECT \
    --region $REGION


gcloud beta run jobs delete $JOB_NAME \
    --project $GOOGLE_PROJECT \
    --region $REGION
```

# 他のサービスと組み合わせで何ができるの？
## Cloud Scheduler
Cloud Schedulerとの組み合わせで定期実行ができるらしい
https://cloud.google.com/run/docs/execute/jobs-on-schedule#console

## Cloud Workflow
https://cloud.google.com/run/docs/create-jobs

>After you create or update a job, you can execute the job as a one-off, on a schedule or as part of a workflow. You can manage individual job executions and view the execution logs.

Cloud Workflowで実行できるとの記載があるが、ドキュメントはまだないようで、yamlの書き方が不明だったので試せなかった
