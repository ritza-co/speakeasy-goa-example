<div align="center">

<a href="[Speakeasy](https://speakeasyapi.dev/)">
  <img src="https://github.com/speakeasy-api/speakeasy/assets/68016351/e959f81a-b250-4003-8c5c-a45b9463fc95" alt="Speakeasy Logo" width="400">
<h2>Speakeasy Goa Example</h2>
</a>

</div>

This example Goa app demonstrates Speakeasy - recommended practices for generating clear OpenAPI specifications and SDKs.

It's part of a complete tutorial available on the [Speakeasy documentation site](https://www.speakeasyapi.dev/docs).

## Prerequisites

You need to have Docker version 20 or later to run this project.

## Run the app

Clone, or download and unzip, the repository:

```bash
git clone https://github.com/ritza-co/speakeasy-goa-example.git
```

Run it:

```bash
cd speakeasy-goa-example
cd completed_app;
docker run --name gobox --volume .:/go/src/app -it golang:1.21.2 bash; # start a docker container
cd /go/src/app; # now inside the container
go install goa.design/goa/v3/cmd/goa@v3; # install Goa
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest; # install grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest;
export PATH=$PATH:/go/src/app/lib; # add protoc to your path
go get app/cmd/club; # download dependencies
./club; # start the server
```

The Club server is running inside Docker.

Open another terminal on your host machine, and log into a new Docker terminal:

```bash
docker exec -it gobox bash;
cd /go/src/app;
./club-cli --help; # see if you can call the server from a CLI client
./club-cli order tea --body '{"includeMilk": false, "isGreen": false, "numberSugars": 1 }';
./club-cli band play --body '{"style": "Bebop" }';
```

While the CLI won't receive a response from the server because the implementation is just a placeholder, you can see in the server terminal that it has been successfully called.

Speakeasy is an online-only service. Please register before continuing, at https://app.speakeasyapi.dev. Once you've registered, create a workspace named `club`. Browse to API keys. Click `New Api Key`. Name it `club`. Copy and save the key content to use later.

Run the commands below in the second terminal to gobox that you were just working in. Use your API key that you saved earlier.

```bash
apt update;
apt install -y curl unzip sudo nodejs npm; # install dependencies
curl -fsSL https://raw.githubusercontent.com/speakeasy-api/speakeasy/main/install.sh | sh; # install Speakeasy
export SPEAKEASY_API_KEY=your_api_key_here; # <-- overwrite this with your key
cd /go/src/app/sdk;
node test.js;
```

You should see in first terminal, where your `club` server is still running, that the server receives a request. In the second terminal you should see it receive `A nice cup of tea`.

To give yourself permissions to the files on your host machine, run:

```bash
sudo chown -R $(id -u):$(id -g) .
```

## License

This project is licensed under the terms of the Apache 2.0 license.
