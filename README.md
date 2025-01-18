[![Development strategy](https://img.shields.io/static/v1?label=DEVELOPMENT%20STRATEGY&message=GITHUB%20FLOW&color=blue)](https://docs.github.com/en/get-started/quickstart/github-flow)
# msg-receiver

## General Description
API Gateway that will be a REST API that will allow to receive messages from a client and produce them to a kafka queue.

## Design
![diagram](/internal/docs/images/diagram.png)

## Environment variables
msg-receiver uses the following environment variables in order to be up and running:

## Execution
**Run the service locally**

In order to run the service locally, create the file `local.env` or `.env` and add the following env vars:

```
TBD
```

run it by executing the following command:

```
make run
```

You will be able to see the service is up and listening at the port specified at *env* file.

**Execute unit testing**
You can run the unit test simply by running:
```
make test
```
## Deployment
Once you create a PR to `main` a GitHub action will run the testing suite, if it succeed, then you will be able to merge the PR.


1. After merging a PR for this repo, you should tag it with the version of the [`VERSION`](/internal/version/VERSION) file in `main` branch
```sh
$ git checkout main
$ git tag $(cat internal/verison/VERSION)
```
2. Push tag to remote to trigger the build pipeline (see next step)
```sh
$ git push origin $(cat internal/verison/VERSION)
```

3. The GitHub workflow specified in `.github/workflows/build-production.yml` file will run, you can access to
   [GitHub Workflow](https://github.com/msg-receiver/actions) to see the progress - `UNDER CONSTRUCTION`
4. After a successful pipeline execution, a new Docker image for this service will be created -
   `UNDER CONSTRUCTION`

## Monitoring
`TO BE CONSIDERED`
## Who do I talk to?
* Nathali Aguayo - **nathaliaguayo@gmail.com**