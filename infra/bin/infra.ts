#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from '@aws-cdk/core';
import { PublicSiteStack } from '../lib/public-site-stack';
import { DataStack } from '../lib/data-stack';
import { AuthStack } from '../lib/auth-stack';

const app = new cdk.App();
new PublicSiteStack(app, 'PublicSiteStack');
new DataStack(app, 'DataStack');
new AuthStack(app, 'AuthStack');
