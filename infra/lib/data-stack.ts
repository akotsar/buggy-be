import * as cdk from '@aws-cdk/core';
import * as dynamo from "@aws-cdk/aws-dynamodb";

export class DataStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const table = new dynamo.Table(this, "DynamodbTable", {
      tableName: "buggy-data",
      billingMode: dynamo.BillingMode.PAY_PER_REQUEST,
      partitionKey: {
        name: "RecordID",
        type: dynamo.AttributeType.STRING,
      },
      sortKey: {
        name: "TypeAndID",
        type: dynamo.AttributeType.STRING,
      },
      removalPolicy: cdk.RemovalPolicy.RETAIN,
    });

    table.addGlobalSecondaryIndex({
      indexName: "TypeAndID",
      partitionKey: {
        name: "ShardID",
        type: dynamo.AttributeType.NUMBER,
      },
      sortKey: {
        name: "TypeAndID",
        type: dynamo.AttributeType.STRING,
      },
    });

    table.addGlobalSecondaryIndex({
      indexName: "Votes",
      partitionKey: {
        name: "EntityType",
        type: dynamo.AttributeType.STRING,
      },
      sortKey: {
        name: "Votes",
        type: dynamo.AttributeType.NUMBER,
      },
    });

    new cdk.CfnOutput(this, "DataTableArn", {
      value: table.tableArn,
      exportName: `${this.stackName}-data-table-arn`,
    });

    new cdk.CfnOutput(this, "DataTableName", {
      value: table.tableName,
      exportName: `${this.stackName}-data-table-name`,
    });
  }
}
