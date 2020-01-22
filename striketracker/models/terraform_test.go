package models

/*
func TestMapFromStruct(t *testing.T) {
	var mfsTests = []struct {
		in  interface{}
		out map[string]interface{}
	}{
		{
			&Compression{
				Enabled: true,
				GZIP:    "txt,js,html,css",
				Level:   createIntPointer(5),
				Mime:    "test/*,application/x-mpegUR",
			},
			map[string]interface{}{
				"ttTestName": "Compression Model",
				"enabled":    true,
				"gzip":       "txt,js,html,css",
				"level":      5,
				"mime":       "test/*,application/x-mpegUR",
			},
		},
		{
			&OriginPullPolicy{
				Enabled:              true,
				ExpirePolicy:         "CACHE_CONTROL",
				ExpireSeconds:        createIntPointer(5),
				DefaultCacheBehavior: "ttl",
			},
			map[string]interface{}{
				"ttTestName":                          "Origin Pull Policy (pointer int)",
				"enabled":                             true,
				"expire_policy":                       "CACHE_CONTROL",
				"expire_seconds":                      5,
				"default_cache_behavior":              "ttl",
				"force_bypass_cache":                  false,
				"honor_must_revalidate":               false,
				"honor_no_cache":                      false,
				"honor_no_store":                      false,
				"honor_private":                       false,
				"honor_smax_age":                      false,
				"http_headers":                        "",
				"must_revalidate_to_no_cache":         false,
				"no_cache_behavior":                   "",
				"update_http_headers_on_304_response": false,
				"max_age_zero_to_no_cache":            false,
				"bypass_cache_identifier":             "",
				"content_type_filter":                 "",
				"header_filter":                       "",
				"method_filter":                       "",
				"path_filter":                         "",
				"status_code_match":                   "",
			},
		},
		{
			&StaticHeader{
				Enabled:                  true,
				HTTP:                     "GET",
				ClientResponseCodeFilter: "2*",
			},
			map[string]interface{}{
				"ttTestName":                  "Static Header",
				"enabled":                     true,
				"http":                        "GET",
				"client_response_code_filter": "2*",
				"origin_pull":                 "",
				"client_request":              "",
				"method_filter":               "",
				"path_filter":                 "",
				"header_filter":               "",
			},
		},
	}

	for _, tt := range mfsTests {
		t.Run(tt.out["ttTestName"].(string), func(t *testing.T) {
			c := MapFromStruct(tt.in)

			c["ttTestName"] = tt.out["ttTestName"]

			if !reflect.DeepEqual(tt.out, c) {
				t.Errorf("Expected: %v | Got: %v\n", tt.out, c)
			}
		})
	}

}

func TestStructFromMap(t *testing.T) {

	var sfmTests = []struct {
		in  map[string]interface{}
		out interface{}
	}{
		{
			map[string]interface{}{
				"ttTestName": "Compression Model",
				"ttTestType": 0,
				"enabled":    true,
				"gzip":       "txt,js,html,css",
				"level":      5,
				"mime":       "test/*,application/x-mpegUR",
			},
			&Compression{
				Enabled: true,
				GZIP:    "txt,js,html,css",
				Level:   createIntPointer(5),
				Mime:    "test/*,application/x-mpegUR",
			},
		},
		{
			map[string]interface{}{
				"ttTestName":             "Origin Pull Policy",
				"ttTestType":             1,
				"enabled":                true,
				"expire_policy":          "CACHE_CONTROL",
				"expire_seconds":         5,
				"default_cache_behavior": "ttl",
			},
			&OriginPullPolicy{
				Enabled:              true,
				ExpirePolicy:         "CACHE_CONTROL",
				ExpireSeconds:        createIntPointer(5),
				DefaultCacheBehavior: "ttl",
			},
		},
		{
			map[string]interface{}{
				"ttTestName":                  "Static Header",
				"ttTestType":                  2,
				"enabled":                     true,
				"http":                        "GET",
				"client_response_code_filter": "2*",
			},
			&StaticHeader{
				Enabled:                  true,
				HTTP:                     "GET",
				ClientResponseCodeFilter: "2*",
			},
		},
	}

	for _, tt := range sfmTests {
		t.Run(tt.in["ttTestName"].(string), func(t *testing.T) {
			switch tt.in["ttTestType"].(int) {
			case 0:
				c := &Compression{}
				d := StructFromMap(c, tt.in)

				if d == nil {
					t.Error("Did not attempt to pack map into struct, map may be nil")
				}

				d = d.(*Compression)

				if !reflect.DeepEqual(tt.out, d) {
					t.Errorf("Failed to pack map into struct. | Map: [%v] | Got: [%v]", tt.in, spew.Sprint(d))
				}
			case 1:
				c := &OriginPullPolicy{}
				d := StructFromMap(c, tt.in)

				spew.Dump(d)

				if d == nil {
					t.Error("Did not attempt to pack map into struct, map may be nil")
				}

				d = d.(*OriginPullPolicy)

				if !reflect.DeepEqual(tt.out, d) {
					t.Errorf("Failed to pack map into struct. | Map: [%v] | Got: [%v]", tt.in, spew.Sprint(d))
				}
			case 2:
				c := &StaticHeader{}
				d := StructFromMap(c, tt.in)

				spew.Dump(d)

				if d == nil {
					t.Error("Did not attempt to pack map into struct, map may be nil")
				}

				d = d.(*StaticHeader)

				if !reflect.DeepEqual(tt.out, d) {
					t.Errorf("Failed to pack map into struct. | Map: [%v] | Got: [%v]", tt.in, spew.Sprint(d))
				}
			}

		})
	}
}

func createIntPointer(x int) *int {
	return &x
}
*/
