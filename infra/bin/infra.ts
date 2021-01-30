#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from '@aws-cdk/core';
import { PublicSiteStack } from '../lib/public-site-stack';

const app = new cdk.App();
new PublicSiteStack(app, 'PublicSiteStack');
