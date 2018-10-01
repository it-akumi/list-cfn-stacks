list-cfn-stacks
====

A slack slash command to list stacks of AWS CloudFormation

## Requirements

* Go
* SAM CLI or AWS CLI

## Usage

After a deployment, set a slash command to your slack.

Then you can use the command. If you execute `/yourSlashCommand`
in your slack channel, this function shows stack lists.

## How to deploy

### Install dependencies and Build

```
$ go get github.com/aws/aws-lambda-go/lambda \
         github.com/aws/aws-lambda-go/events \
         github.com/aws/aws-sdk-go
$ go build -o src/list-cfn-stacks src/list-cfn-stacks.go
```

### Packaging and Deployment

```
$ sam package \
     --template-file ./template.yml \
     --s3-bucket your-s3-bucket-name \
     --output-template-file your-output-template.yml
$ sam deploy \
     --template-file your-output-template.yml \
     --stack-name your-stack-name
     --capabilities CAPABILITY_IAM
```

You can use `aws cloudformation` instead of `sam`.

See [AWS document](https://docs.aws.amazon.com/lambda/latest/dg/serverless-deploy-wt.html#serverless-deploy) for more details.

## Author
[Takumi Ishii](https://github.com/it-akumi)

## License
[MIT](https://github.com/it-akumi/list-cfn-stacks/blob/master/LICENSE)
