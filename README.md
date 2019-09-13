# WurlWind [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/openwurl/wurlwind/striketracker?status.svg)](https://godoc.org/github.com/openwurl/wurlwind/striketracker) [![Build Status](https://travis-ci.org/openwurl/wurlwind.svg?branch=master)](https://travis-ci.org/openwurl/wurlwind)

![wurlwind](static/wurlwind.png)

An open source GO library for interfacing with Highwinds / Striketracker CDN

This library provides no business logic, and operates more like an SDK

### To Do
* Literally everything
* List is ever changing

# Testing
The basic tests are surfaced via the Makefile

* `make test`
  * Runs Unit tests and ignores Integration suite
  * Requires no Environment variables
  * Ex. `make test`
* `make cover`
  * Load go cover details in your browser
* `make integration`
  * Runs integration suite against Striketracker API
  * Requires `INTEGRATIONACCOUNTHASH` & `AUTHORIZATIONHEADERKEY` environment variables
    * INTEGRATIONACCOUNTHASH
      * The account hash of the subaccount you are running integration tests against
    * AUTHORIZATIONHEADERKEY
      * The authorization header key for authenticated API access
  * Ex. `INTEGRATIONACCOUNTHASH=f98fsj32k AUTHORIZATIONHEADERKEY=fj32jk43kj32kj3rkhj make integration`

# Usage
You will need your authorizationHeaderToken from Highwinds as well as manage your own accountHashes.

There is no plan to implement username/password based authentication. While it is supported by the Striketracker API, this library is focused on using a permanent API token.

### Client
You must first configure and maintain a client in your application. There are a few ways to do this.

`import "github.com/openwurl/wurlwind/striketracker"`

Via Functional Parameters
```
c, err := striketracker.NewClientWithOptions(
    striketracker.WithApplicationID("DescriptiveApplicationName"),
    striketracker.WithDebug(true),
    striketracker.WithAuthorizationHeaderToken(stringAuthToken),
    striketracker.WithRequestTimeout(intSeconds),
)
```

Via user-defined configuration
```
c, err := striketracker.NewClientFromConfiguration(&striketracker.Configuration{
	Debug: false,
	AuthorizationHeaderToken: yourtoken,
	Timeout: TimeOutInSeconds,
	ApplicationID: YourApplicationName,
})
```

# Services
A brief overview of the Highwinds services exposed in this API Client Library

### Context
Context can be passed into most, if not all service methods - to be used for early cancellation or to configure a timeout per operation.

```
ctx := context.Background()
ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
defer cancel()

resp, err := service.Action(ctx, ...)
```

### Origin
The origin service at highwinds defines the upstream origins used as the cache basis / source for your edge distributions.

It has a few simple interactions

`import "github.com/openwurl/wurlwind/striketracker/services/origin"`

##### Instantiation
```
o := origin.New(*striketracker.Client)
```

##### Surfaced Operations
* Create New Origin
  * `POST /api/v1/accounts/{account_hash}/origins`
  * `origin.Create(accountHash, Origin)`
* List All Origins
  * `GET /api/v1/accounts/{account_hash}/origins`
  * `origin.List(accountHash)`
* Delete Origin
  * `DELETE /api/v1/accounts/{account_hash}/origins/{origin_id}`
  * `origin.Delete(accountHash, originID)`
* Get Individual Origin
  * `GET /api/v1/accounts/{account_hash}/origins/{origin_id}`
  * `origin.Get(accountHash, originID)`
* Update Individual Origin
  * `PUT /api/v1/accounts/{account_hash}/origins/{origin_id}`
  * `origin.Update(accountHash, originID, Origin)`

##### Examples
Setup
```
c, err := striketracker.NewClientWithOptions(
    striketracker.WithApplicationID("SomeApplication"),
    striketracker.WithDebug(true),
    striketracker.WithAuthorizationHeaderToken(stringAuthToken),
    striketracker.WithRequestTimeout(3),
)

ctx := context.Background()
ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
defer cancel()
```

Create
```
response, err := o.Create(ctx, accountHash, &models.Origin{
    Name: "Some Origin",
    Hostname: "some.origin.com",
    Port: 8080,
})
if err != nil {
    // deal with error
}
fmt.Printf("Origin %s created with ID: %s", response.Name, response.ID)
```

Get & Delete
```
orig, err := o.Get(ctx, accountHash, 8675309)
if err != nil {
    // handle err
}
response, err := o.Delete(ctx, accountHash, orig)
if err != nil {
    // deal with error
}
fmt.Printf("Origin %s deleted with ID: %s", orig.Name, orig.ID)
```

### Certificates
* TODO

### Configuration
* TODO

### Hosts
* TODO

### Search
* TODO

### Purge
* TODO

### Accounts
* TODO

### Authentication (Token management)
* TODO

### Users
* TODO