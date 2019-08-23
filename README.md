# WurlWind
An open source GO library for interfacing with Highwinds / Striketracker CDN

This library provides no business logic, and operates more like an SDK

### To Do
* Literally everything
* Fix up URL management to resolve references etc
* Env var configuration (config uniformity too)
* The rest of the Origin service

# Usage
You will need your authorizationHeaderToken from Highwinds as well as manage your own accountHashes.

There is a plan to implement username/password based authentication if you do not want to use a permanent API token.

### Client
You must first configure and maintain a client in your application. As of right now this looks like the example below, however it's going to accept EnvVars ultimately for secrets (not implemented yet)

`import "github.com/openwurl/wurlwind/striketracker"`

```
c := striketracker.NewClient(debug (bool), authorizationHeaderToken (string), application-name (string))
```

ToDo:

* `NewClientFromConfig`
  * config-struct functional driven configuration
  * (ex. `NewClientFromConfig(WithAuthorizationToken(token)))`)
* `NewClientFromEnv`
  * env-var driven configuration


# Services
A brief overview of the Highwinds services exposed in this API Client Library

### Origin
The origin service at Highwinds has a few simple interactions

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
