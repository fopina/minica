# minica

[![goreference](https://pkg.go.dev/badge/github.com/fopina/minica.svg)](https://pkg.go.dev/github.com/fopina/minica)
[![release](https://img.shields.io/github/v/release/fopina/golang-template)](https://github.com/fopina/minica/releases)
[![downloads](https://img.shields.io/github/downloads/fopina/golang-template/total.svg)](https://github.com/fopina/minica/releases)
[![ci](https://github.com/fopina/minica/actions/workflows/publish-main.yml/badge.svg)](https://github.com/fopina/minica/actions/workflows/publish-main.yml)
[![test](https://github.com/fopina/minica/actions/workflows/test.yml/badge.svg)](https://github.com/fopina/minica/actions/workflows/test.yml)
[![codecov](https://codecov.io/github/fopina/minica/graph/badge.svg)](https://codecov.io/github/fopina/minica)

> **NOTE**  
> This is a fork from [jsha/minica](https://github.com/jsha/minica).  
> This was only detached as PRs were not getting any reactions - but all credits go to [@jsha](https://github.com/jsha)

Minica is a simple CA intended for use in situations where the CA operator
also operates each host where a certificate will be used. It automatically
generates both a key and a certificate when asked to produce a certificate.
It does not offer OCSP or CRL services. Minica is appropriate, for instance,
for generating certificates for RPC systems or microservices.

On first run, minica will generate a keypair and a root certificate in the
current directory, and will reuse that same keypair and root certificate
unless they are deleted.

On each run, minica will generate a new keypair and sign an end-entity (leaf)
certificate for that keypair. The certificate will contain a list of DNS names
and/or IP addresses from the command line flags. The key and certificate are
placed in a new directory whose name is chosen as the first domain name from
the certificate, or the first IP address if no domain names are present. It
will not overwrite existing keys or certificates.

The certificate will have a validity of 2 years and 30 days.

# Installation

First, install the [Go tools](https://golang.org/dl/) and set up your `$GOPATH`.
Then, run:

`go install github.com/jsha/minica@latest`

When using Go 1.11 or newer you don't need a $GOPATH and can instead do the
following:

```
cd /ANY/PATH
git clone https://github.com/jsha/minica.git
go build
## or
# go install
```

Mac OS users could alternatively use Homebrew: `brew install minica`

# Example usage

```
# Generate a root key and cert in minica-key.pem, and minica.pem, then
# generate and sign an end-entity key and cert, storing them in ./foo.com/
$ minica --domains foo.com

# Wildcard
$ minica --domains '*.foo.com'
```
