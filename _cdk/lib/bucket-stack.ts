import { CfnOutput, RemovalPolicy, Stack, StackProps } from "aws-cdk-lib";
import * as S3 from "aws-cdk-lib/aws-s3";
import { Construct } from "constructs";
import { PROJECT_NAME } from "./const";

export class BucketStack extends Stack {
  public replicateBucket: S3.Bucket;

  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);
    const stageName: string = this.node.tryGetContext("stageName");
    const accountId = Stack.of(this).account;

    const s3 = new S3.Bucket(this, "SqliteBackupBucket", {
      bucketName: `${stageName}-${PROJECT_NAME}-replica-${accountId}`,
      removalPolicy: RemovalPolicy.DESTROY,
    });

    this.replicateBucket = s3;

    new CfnOutput(this, "bucketStackOutput", {
      value: s3.bucketName,
      description: "sqlite3 replica bucket",
    });
  }
}
