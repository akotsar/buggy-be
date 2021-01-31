import * as cdk from '@aws-cdk/core';
import * as cognito from '@aws-cdk/aws-cognito';

export class AuthStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const pool = new cognito.UserPool(this, "UserPool", {
      selfSignUpEnabled: true,
      signInAliases: {
        username: true,
      },
      standardAttributes: {
        familyName: { required: true, mutable: true },
        givenName: { required: true, mutable: true }
      },
      customAttributes: {
        'is_admin': new cognito.BooleanAttribute({ mutable: true }),
      }
    });

    const client = pool.addClient('buggy-app', {
      authFlows: {
        adminUserPassword: true,
        userPassword: true,
      }
    });

    new cdk.CfnOutput(this, "PoolId", {
      value: pool.userPoolId,
      exportName: `${this.stackName}-user-pool-id`,
    });

    new cdk.CfnOutput(this, "PoolClientId", {
      value: client.userPoolClientId,
      exportName: `${this.stackName}-user-pool-client-id`,
    });
  }
}
