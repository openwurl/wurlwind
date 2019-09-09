// Package striketracker provides a library for interfacing with Highwinds CDN
// without business logic
// A client is acquired by
//  c, err := striketracker.NewClientWithOptions(
//  	striketracker.WithApplicationID("DescriptiveApplicationName"),
//  	striketracker.WithDebug(true),
//  	striketracker.WithAuthorizationHeaderToken(authToken),
//  )
// Or you can configure it manually via
//  c, err := striketracker.NewClientFromConfiguration(&striketracker.Configuration{
//  	Debug: false,
//  	AuthorizationHeaderToken: yourtoken,
//  	Timeout: TimeOutInSeconds,
//  	ApplicationID: YourApplicationName,
//  })
package striketracker
