# kube-sec-decode-data

## Usage

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
$ cat test.yaml| ./kube-sec-decode-data
apiVersion: v1
data:
  BAR: bar
  FOO: foo
kind: Secret
metadata:
  name: testsecret
  namespace: default

```
