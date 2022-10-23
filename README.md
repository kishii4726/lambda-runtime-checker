# lambda_runtime_checker
Tool to investigate runtime of Lambda functions.

## Install
### Mac(amd64)
```
$ LAMBDA_RUNTIME_CHECKER_VERSION=0.0.1
$ curl -OL https://github.com/kishii4726/lambda-runtime-checker/releases/download/v${LAMBDA_RUNTIME_CHECKER_VERSION}/lambda-runtime-checker_v${LAMBDA_RUNTIME_CHECKER_VERSION}_darwin_amd64.zip

$ unzip lambda-runtime-checker_v${LAMBDA_RUNTIME_CHECKER_VERSION}_darwin_amd64.zip lambda-runtime-checker

$ sudo cp lambda-runtime-checker /usr/local/bin
```

### Mac(arm64)
```
$ LAMBDA_RUNTIME_CHECKER_VERSION=0.0.1
$ curl -OL https://github.com/kishii4726/lambda-runtime-checker/releases/download/v${LAMBDA_RUNTIME_CHECKER_VERSION}/lambda-runtime-checker_v${LAMBDA_RUNTIME_CHECKER_VERSION}_darwin_arm64.zip

$ unzip lambda-runtime-checker_v${LAMBDA_RUNTIME_CHECKER_VERSION}_darwin_arm64.zip lambda-runtime-checker

$ sudo cp lambda-runtime-checker /usr/local/bin
```

### Linux
```
$ LAMBDA_RUNTIME_CHECKER_VERSION=0.0.1
$ curl -OL https://github.com/kishii4726/lambda-runtime-checker/releases/download/v${LAMBDA_RUNTIME_CHECKER_VERSION}/lambda-runtime-checker_v${LAMBDA_RUNTIME_CHECKER_VERSION}_linux_amd64.zip

$ unzip lambda-runtime-checker_v${LAMBDA_RUNTIME_CHECKER_VERSION}_linux_amd64.zip lambda-runtime-checker

$ sudo cp lambda-runtime-checker /usr/local/bin
```

## Usage
### Show runtimes for all Lambda functions
```
$ lambda-runtime-checker all
+---------------+------------+
| FUNCTION NAME |  RUNTIME   |
+---------------+------------+
| test1         | nodejs12.x |
+---------------+------------+
| test2         | ruby2.5    |
+---------------+------------+
| test3         | python3.9  |
+---------------+------------+
| test4         | ruby2.5    |
+---------------+------------+
| test5         | python3.9  |
+---------------+------------+
| test6         | python3.8  |
+---------------+------------+

```

### Search for Lambda functions that use the specified runtime
```
$ lambda-runtime-checker runtime python3.9
+---------------+------------+
| FUNCTION NAME |  RUNTIME   |
+---------------+------------+
| test3         | python3.9  |
+---------------+------------+
| test5         | python3.9  |
+---------------+------------+
```

### Search for Lambda functions using deprecated runtimes
```
$ lambda-runtime-checker deprecated
+---------------+------------+
| FUNCTION NAME |  RUNTIME   |
+---------------+------------+
| test1         | nodejs12.x |
+---------------+------------+
| test2         | ruby2.5    |
+---------------+------------+
| test4         | ruby2.5    |
+---------------+------------+
```
