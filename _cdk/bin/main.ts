import * as cdk from "aws-cdk-lib";
import { AppRunnerStack } from "../lib/apprunner-stack";
import { BucketStack } from "../lib/bucket-stack";
import { PROJECT_NAME } from "../lib/const";
import { EcrStack } from "../lib/ecr-stack";

const app = new cdk.App();
const stageName = app.node.tryGetContext("stageName");

const ecrStack = new EcrStack(app, `${stageName}-${PROJECT_NAME}-ecr`);
const bucketStack = new BucketStack(
  app,
  `${stageName}-${PROJECT_NAME}-bucket`
);
new AppRunnerStack(app, `${stageName}-${PROJECT_NAME}-app-runner`, {
  repository: ecrStack.repository,
  replicateBucket: bucketStack.replicateBucket,
});
