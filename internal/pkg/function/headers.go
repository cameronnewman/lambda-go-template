package function

const (
	//MIMEApplicationJSON is the http header for a JSON type
	MIMEApplicationJSON = "application/json"

	//MIMEApplicationJSONCharsetUTF8 is the http header for a JSON type with a charset of UTF8
	MIMEApplicationJSONCharsetUTF8 = "application/json; charset=utf-8"

	//MIMEApplicationJavaScript is the http header for a JSON type
	MIMEApplicationJavaScript = "application/javascript"

	//MIMEApplicationJavaScriptCharsetUTF8 is the http header for a JSON type with a charset of UTF8
	MIMEApplicationJavaScriptCharsetUTF8 = "application/javascript; charset=utf-8"

	//MIMEApplicationXML is the http header for a XML type
	MIMEApplicationXML = "application/xml"

	//MIMEApplicationXMLCharsetUTF8 is the http header for a XML type with a charset of UTF8
	MIMEApplicationXMLCharsetUTF8 = "application/xml; charset=utf-8"

	//MIMEApplicationForm is the http header for a form type
	MIMEApplicationForm = "application/x-www-form-urlencoded"

	//MIMEApplicationProtobuf is the http header for a protobuf type
	MIMEApplicationProtobuf = "application/protobuf"

	//MIMEApplicationMsgpack is the http header for a msgpack type
	MIMEApplicationMsgpack = "application/msgpack"

	//MIMETextHTML is the http header for a HTML type
	MIMETextHTML = "text/html"

	//MIMETextHTMLCharsetUTF8 is the http header for a HTML type with a charset of UTF8
	MIMETextHTMLCharsetUTF8 = "text/html; charset=utf-8"

	//MIMETextPlain is the http header for a plain text type
	MIMETextPlain = "text/plain"

	//MIMETextPlainCharsetUTF8 is the http header for a plain text type with a charset of UTF8
	MIMETextPlainCharsetUTF8 = "text/plain; charset=utf-8"

	//MIMEMultipartForm is the http header for a multi part form data type
	MIMEMultipartForm = "multipart/form-data"

	//MIMEOctetStream is the http header for a octet stream type
	MIMEOctetStream = "application/octet-stream"
)

const (
	//RequestContextName is a variable name
	RequestContextName = "REQUEST_CONTEXT"

	//RequestContextID is a variable name
	RequestContextID = "REQUEST_ID"

	//HeaderAuthenticateRealm is a http header auth realm
	HeaderAuthenticateRealm = `Basic realm="api.google.com"`

	//HeaderTrue is a http header variable name true
	HeaderTrue = "true"

	//ContentTypeIsXML is a variable when the payload is xml
	ContentTypeIsXML = "xml"

	//ContentTypeIsJSON is a variable when the payload is json
	ContentTypeIsJSON = "json"

	//ProtocalHTTP is a variable name when the request is over unsecure transport
	ProtocalHTTP = "http"

	//ProtocalHTTPS is a variable name when the request is over secure transport
	ProtocalHTTPS = "https"

	//HeaderAcceptEncoding is a http header variable name for encoding
	HeaderAcceptEncoding = "Accept-Encoding"

	//HeaderAllow is a http header variable name for allowing
	HeaderAllow = "Allow"

	//HeaderAuthorization is a http header variable name
	HeaderAuthorization = "Authorization"

	//HeaderContentDisposition is a http header variable name
	HeaderContentDisposition = "Content-Disposition"

	//HeaderContentEncoding is a http header variable name
	HeaderContentEncoding = "Content-Encoding"

	//HeaderContentLength is a http header variable name
	HeaderContentLength = "Content-Length"

	//HeaderContentType is a http header variable name
	HeaderContentType = "Content-Type"

	//HeaderCookie is a http header variable name
	HeaderCookie = "Cookie"

	//HeaderSetCookie is a http header variable name
	HeaderSetCookie = "Set-Cookie"

	//HeaderIfModifiedSince is a http header variable name
	HeaderIfModifiedSince = "If-Modified-Since"

	//HeaderLastModified is a http header variable name
	HeaderLastModified = "Last-Modified"

	//HeaderLocation is a http header variable name
	HeaderLocation = "Location"

	//HeaderUpgrade is a http header variable name
	HeaderUpgrade = "Upgrade"

	//HeaderUserAgent is a http header variable name
	HeaderUserAgent = "User-Agent"

	//HeaderVary is a http header variable name
	HeaderVary = "Vary"

	//HeaderWWWAuthenticate is a http header variable name
	HeaderWWWAuthenticate = "WWW-Authenticate"

	//HeaderXForwardedProto is a http header variable name
	HeaderXForwardedProto = "X-Forwarded-Proto"

	//HeaderXHTTPMethodOverride is a http header variable name
	HeaderXHTTPMethodOverride = "X-HTTP-Method-Override"

	//HeaderXForwardedFor is a http header variable name
	HeaderXForwardedFor = "X-Forwarded-For"

	//HeaderXRealIP is a http header variable name
	HeaderXRealIP = "X-Real-IP"

	//HeaderXRateLimitLimit is a http header variable name
	HeaderXRateLimitLimit = "X-Ratelimit-Limit"

	//HeaderXRateLimitRemainingRequests is a http header variable name
	HeaderXRateLimitRemainingRequests = "X-Ratelimit-Remaining"

	//HeaderXRateLimitResetTime is a http header variable name
	HeaderXRateLimitResetTime = "X-Ratelimit-Reset"

	//HeaderServer is a http header variable name
	HeaderServer = "Server"

	//HeaderOrigin is a http header variable name
	HeaderOrigin = "Origin"

	//HeaderAccessControlRequestMethod is a http header variable name
	HeaderAccessControlRequestMethod = "Access-Control-Request-Method"

	//HeaderAccessControlRequestHeaders is a http header variable name
	HeaderAccessControlRequestHeaders = "Access-Control-Request-Headers"

	//HeaderAccessControlAllowOrigin is a http header variable name
	HeaderAccessControlAllowOrigin = "Access-Control-Allow-Origin"

	//HeaderAccessControlAllowMethods is a http header variable name
	HeaderAccessControlAllowMethods = "Access-Control-Allow-Methods"

	//HeaderAccessControlAllowHeaders is a http header variable name
	HeaderAccessControlAllowHeaders = "Access-Control-Allow-Headers"

	//HeaderAccessControlAllowCredentials is a http header variable name
	HeaderAccessControlAllowCredentials = "Access-Control-Allow-Credentials"

	//HeaderAccessControlExposeHeaders is a http header variable name
	HeaderAccessControlExposeHeaders = "Access-Control-Expose-Headers"

	//HeaderAccessControlMaxAge is a http header variable name
	HeaderAccessControlMaxAge = "Access-Control-Max-Age"

	//HeaderPlatformPrivacyPreferences is a http header variable name
	HeaderPlatformPrivacyPreferences = "P3P"

	//PlatformPrivacyPreferencesPolicy is a http header variable name
	PlatformPrivacyPreferencesPolicy = "CP=\"NON DSP LAW CUR ADM DEV TAI PSA PSD HIS OUR DEL IND UNI PUR COM NAV INT DEM CNT STA POL HEA PRE LOC IVD SAM IVA OTC\""

	//allowedCORSHeaders is a http header variable name
	allowedCORSHeaders = "X-Requested-With,X-Prototype-Version,Authorization,Content-Type,Cache-Control,Pragma,Origin,Accept,Cookie"

	// Security related

	//HeaderStrictTransportSecurity is a http header variable name
	HeaderStrictTransportSecurity = "Strict-Transport-Security"

	//HeaderXContentTypeOptions is a http header variable name
	HeaderXContentTypeOptions = "X-Content-Type-Options"

	//HeaderXXSSProtection is a http header variable name
	HeaderXXSSProtection = "X-XSS-Protection"

	//HeaderXFrameOptions is a http header variable name
	HeaderXFrameOptions = "X-Frame-Options"

	//HeaderContentSecurityPolicy is a http header variable name
	HeaderContentSecurityPolicy = "Content-Security-Policy"

	//HeaderXCSRFToken is a http header variable name
	HeaderXCSRFToken = "X-CSRF-Token"

	//httpStrictTransportSecurity is a http header variable name
	httpStrictTransportSecurity = "max-age=31536000; includeSubDomains; preload"

	//httpUpgradeInsecureRequests is a http header variable name
	httpUpgradeInsecureRequests = "upgrade-insecure-requests"
)
