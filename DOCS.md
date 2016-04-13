Use this plugin to post deployment notifications to New Relic.
The following parameters are available:

- `license_key` - **required** Your New Relic license key
- `application_id` - The ID # of the application
- `app_name` - The name of the application, found in the newrelic.yml file

> One of `application_id` or `app_name` is **required** as per New Relic's API specs.

- `description` - Text annotation for the deployment — notes for you
- `user` - The name of the user/process that triggered this deployment.  
Defaults to `drone-ci`
- `debug` - Prints the request payload and response code

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
notify:
  newrelic:
    license_key: PXCmiqaFhhHILKr6tAFStVDCGsdsQ+41mE2W6ocpT34
    app_name: Acme
    description: Continuous deployment via Drone CI
    when:
      branch: master
```
