# Test lambda

This function assumes that SNS message contains JSON and adds platform:farmroad to the root level of said payload regardless of its contents.

Tests do not mock enviroment variables, however that could be achieved in the similar fashion as AWS API mocks.
