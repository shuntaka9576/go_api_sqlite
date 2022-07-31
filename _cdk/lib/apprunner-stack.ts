import { CfnOutput, Stack, StackProps } from "aws-cdk-lib";
import * as Ecr from "aws-cdk-lib/aws-ecr";
import * as Iam from "aws-cdk-lib/aws-iam";
import * as S3 from "aws-cdk-lib/aws-s3";
import * as AppRunner from "@aws-cdk/aws-apprunner-alpha";
import { Construct } from "constructs";

type Props = {
  repository: Ecr.Repository;
  replicateBucket: S3.Bucket;
} & StackProps;

const API_LISTEN_PORT = 8080;

export class AppRunnerStack extends Stack {
  constructor(scope: Construct, id: string, props: Props) {
    super(scope, id, props);

    const instanceRole = new Iam.Role(this, "AppRunnerInstanceRole", {
      assumedBy: new Iam.ServicePrincipal("tasks.apprunner.amazonaws.com"),
    });
    instanceRole.addToPolicy(
      new Iam.PolicyStatement({
        actions: [
          "s3:GetBucketLocation",
          "s3:ListBucket",
          "s3:PutObject",
          "s3:DeleteObject",
          "s3:GetObject",
        ],
        effect: Iam.Effect.ALLOW,
        resources: [
          `arn:aws:s3:::${props.replicateBucket.bucketName}`,
          `arn:aws:s3:::${props.replicateBucket.bucketName}/*`,
        ],
      })
    );

    const accessRole = new Iam.Role(this, "AppRunnerAccessRole", {
      assumedBy: new Iam.ServicePrincipal("build.apprunner.amazonaws.com"),
    });

    const service = new AppRunner.Service(this, "AppRunnerExampleService", {
      source: AppRunner.Source.fromEcr({
        imageConfiguration: {
          port: API_LISTEN_PORT, // REST API Listen port
          environment: {
            PORT: String(API_LISTEN_PORT),
            DB_PATH: "./todo.db",
            REPLICATE_BUCKET_NAME: props.replicateBucket.bucketName,
          },
        },
        repository: props.repository,
        tagOrDigest: "latest",
      }),
      instanceRole: instanceRole,
      accessRole: accessRole,
    });

    new CfnOutput(this, "AppRunnerOutput", {
      value: `https://${service.serviceUrl}
curl -XGET https://${service.serviceUrl}/tasks
curl -XPOST https://${service.serviceUrl}/tasks -d '{"title": "test1"}'`,
      description: `Output AppRunner endpoint and API exec commands`,
    });
  }
}
