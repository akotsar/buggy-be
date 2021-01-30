import * as cdk from '@aws-cdk/core';
import * as s3 from '@aws-cdk/aws-s3';
import * as cloudfront from '@aws-cdk/aws-cloudfront';
import * as origins from '@aws-cdk/aws-cloudfront-origins';

export class PublicSiteStack extends cdk.Stack {
  bucket: s3.Bucket;
  distribution: cloudfront.Distribution;

  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    this.bucket = new s3.Bucket(this, 'PublicSiteBucket', {
      versioned: false,
      websiteIndexDocument: "index.html"
    });

    this.distribution = new cloudfront.Distribution(this, "PublicSiteDistribution", {
      defaultBehavior: { origin: new origins.S3Origin(this.bucket) },
    })
  }
}
