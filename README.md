# libre-izly

Libre and open source Izly client implementation

See the [changelog](/CHANGELOG.md) for the latest updates.

## Table of content

-   [**Features**](#features)
-   [**Installation**](#installation)
-   [**Compiling from source**](#compiling-from-source)
-   [**Configuring the client**](#configuring-the-client)
-   [**How to use**](#how-to-use)
-   [**Copyright**](#copyright)

## Features

-   Login to Izly with 2FA
-   Payment QR code generation

## Installation

Here's how to install the client:

-   Download [go](https://go.dev/dl/) (go 1.20 required).
-   Download or clone the project.
-   Download the binary from the [Releases](../../releases) or [build it](#compiling-from-source) yourself.

## Compiling from source

-   Use the `make` command or use `go build cmd/libre-izly/main.go` in this folder

## Configuring the client

For now you can copy or rename the [`.env.example`](.env.example) file to `.env`.

You do not need to edit anything except maybe the `CLIENT_VERSION` if Izly gets updated.

## How to use

-   First login to Izly by running `./main login -u YOUR_LOGIN -p YOUR_PASSWORD`
-   Then if everything worked correctly you should receive an SMS with a 2FA URL like this `https://mon-espace.izly.fr/tools/Activation/YOUR_PHONE_NUMBER/SOME_TEXT`
-   Now run `./main login -u YOUR_LOGIN -c SOME_TEXT`
-   Again, if everything is okay, you should be logged in.
-   You can now generate payment QR codes by running `./main qr`

## Copyright

See the [license](/LICENSE).