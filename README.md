# Kastomize

jlink wrapper that creates custom minimalistic JRE with selected modules only.

Kastomize creates JVM without man pages, headers and debug attributes as they are hardly needed to run application.

## How to use it

### Command line (CLI)

* Download the latest release and make it executable

```shell
wget https://github.com/kyberorg/kastomize/releases/download/1.0/kastomize
chmod +x kastomize
```

* Create you module file

Tip: you can find out required options by running `jdeps your_app.jar`

```shell
cat <<EOF > modulesfile
java.base
java.logging
java.sql
EOF
```

* Finally, run kastomize

```shell
./kastomize --modules-file modulesfile --output myjre
```

* Find your custom-made JVM is ready at `myjre` folder

```shell
ls -al ./myjre
```

#### Additional Options and configuration

* You can set custom JDK folder by defining `--java-home` option or `JAVA_HOME` environment var.

```shell
./kastomize --modules-file modulesfile --output myjre --java-home /path/to/my/jdk
```

* Absolute path are also supported

```shell
./kastomize --modules-file /path/to/modulesfile --output /path/to/myjre
```

## How to build it

### Using make

```shell
make binary
```

### Without make

```shell
 CGO_ENABLED=0 go build github.com/kyberorg/kastomize/cmd/kastomize
```
