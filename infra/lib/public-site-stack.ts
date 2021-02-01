import * as cdk from '@aws-cdk/core';
import * as s3 from '@aws-cdk/aws-s3';
import * as cloudfront from '@aws-cdk/aws-cloudfront';
import * as origins from '@aws-cdk/aws-cloudfront-origins';
import * as acm from '@aws-cdk/aws-certificatemanager';

export class PublicSiteStack extends cdk.Stack {
  bucket: s3.Bucket;
  distribution: cloudfront.Distribution;

  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    this.bucket = new s3.Bucket(this, 'PublicSiteBucket', {
      versioned: false
    });

    const cert = acm.Certificate.fromCertificateArn(this, "Certificate", "arn:aws:acm:us-east-1:739075733351:certificate/a0d84763-4813-4623-a435-218ff2d6258a")

    this.distribution = new cloudfront.Distribution(this, "PublicSiteDistribution", {
      defaultBehavior: { origin: new origins.S3Origin(this.bucket) },
      defaultRootObject: "index.html",
      certificate: cert,
      domainNames: [
        "buggy.justtestit.org"
      ],
      errorResponses: [
        {
          httpStatus: 404,
          responseHttpStatus: 200,
          responsePagePath: "/index.html"
        }
      ]
    })
  }
}
