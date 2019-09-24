# kube-sec-decoder

## Usage

### Basic

```shell
$ cat test.yaml
apiVersion: v1
kind: Secret
metadata:
  name: testsecret
  namespace: default
data:
  FOO: Zm9v
  BAR: YmFy
```

```shell
$ cat test.yaml| kube-sec-decoder
apiVersion: v1
kind: Secret
metadata:
  name: testsecret
  namespace: default
data:
  FOO: foo
  BAR: bar
```

### Hide .data.*

```shell
$ cat test.yaml| kube-sec-decoder --hide-data
apiVersion: v1
kind: Secret
metadata:
  name: testsecret
  namespace: default
data:
  FOO: <secret>
  BAR: <secret>

$ cat test.yaml| kube-sec-decoder --hide-data --replace-data xxx
apiVersion: v1
kind: Secret
metadata:
  name: testsecret
  namespace: default
data:
  FOO: xxx
  BAR: xxx
```
