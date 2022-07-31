import { Stack, StackProps } from "aws-cdk-lib";
import * as Ecr from "aws-cdk-lib/aws-ecr";
import { Construct } from "constructs";

export class EcrStack extends Stack {
  public repository: Ecr.Repository;

  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    this.repository = new Ecr.Repository(this, "GoApiSqliteApiRepository", {
      repositoryName: `go-api-sqlite`,
    });
  }
}
