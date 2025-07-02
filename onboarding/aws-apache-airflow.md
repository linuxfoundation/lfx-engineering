# Apache Airflow on AWS

## Overview

The following is a guide to gaining access to the AWS Airflow Instance running in the AWS account.

## Login Steps

1. Navigate to our [AWS LF SSO portal link](https://lfx.awsapps.com/start)
2. Authenticate via LF SSO.
3. From the AWS access portal landing page, select the “LFX Product
  (Production)” -> PowerUser option to assume the role in the appropriate account.
  See [screenshot #1](./screenshots/aws-access-portal-account-selection.png)
4. Once into the AWS console home/products/catalog page -> ensure you have the
  US West (Oregon) region selected in the top-right corner. This is not the
  default region - I always need to change this. See [screenshot
  #2](./screenshots/aws-access-portal-region-selection.png).
5. From the console home page, search for Amazon MWAA (Managed Apache Airflow)
  and select the product - see [screenshot
  #3](./screenshots/aws-access-portal-managed-airflow-search.png).
6. From the menu, select Environments on the left. You should see lfx-airstream.
  Select this. See [screenshot
  #4](./screenshots/aws-access-portal-mwaa-environment-selection.png).
7. Once on the lfx-airstream deployment page, select the Airflow UI link or
  button to FINALLY navigate to the Airflow console. See [screenshot
  #5](./screenshots/aws-access-portal-mwaa-airflow-ui-link.png).

## Screenshots

### Screenshot #1

![Screenshot #1](./screenshots/aws-access-portal-account-selection.png)

### Screenshot #2

![Screenshot #2](./screenshots/aws-access-portal-region-selection.png)

### Screenshot #3

![Screenshot #3](./screenshots/aws-access-portal-managed-airflow-search.png)

### Screenshot #4

![Screenshot #4](./screenshots/aws-access-portal-mwaa-environment-selection.png)

### Screenshot #5

![Screenshot #5](./screenshots/aws-access-portal-mwaa-airflow-ui-link.png)
