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

In order to run the service locally, create the file `.env` and add the following env vars:

```
MSG_RECEIVER_SECRET_KEY="NasSecretKey"
MSG_RECEIVER_ISSUER="Nathali"
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

To test the rate limiter, you can run the following command:
```
ab -n 20 -c 5 -p payload.json -T application/json http://localhost:8080/token
```
Where:

`-n 20`: Total number of requests.

`-c 5`: Number of concurrent requests.

`-p payload.json`: Specifies the file containing the POST body.

`-T application/json`: Sets the Content-Type header.

That means: this sends 20 requests with a concurrency of 5 using the tool ab (Apache Benchmark). You should see that the first 5 requests are successful, and the rest are blocked by the rate limiter.

Example of the output:
![ratelimit](/internal/docs/images/ratelimit.png)

## Deployment
Once you create a PR to `main` a GitHub action will run the testing suite, if it succeeds, then you will be able to merge the PR.


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
   [GitHub Workflow](https://github.com/nathaliaguayos/msg-receiver/actions) to see the progress - `UNDER CONSTRUCTION`
4. After a successful pipeline execution, a new Docker image for this service will be created -
   `UNDER CONSTRUCTION`

## Monitoring
`TO BE CONSIDERED`
## Who do I talk to?
* Nathali Aguayo - **nathaliaguayo@gmail.com**