import { expect as expectCDK, matchTemplate, MatchStyle } from '@aws-cdk/assert';
import * as cdk from '@aws-cdk/core';
import { PublicSiteStack } from '../lib/public-site-stack';

test('Empty Stack', () => {
    const app = new cdk.App();
    // WHEN
    const stack = new PublicSiteStack(app, 'MyTestStack');
    // THEN
    expectCDK(stack).to(matchTemplate({
      "Resources": {}
    }, MatchStyle.EXACT))
});
