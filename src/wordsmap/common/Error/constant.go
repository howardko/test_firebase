package Error

import "net/http"

// StatusContinue           = 100
// StatusSwitchingProtocols = 101

// StatusOK                   = 200
// StatusCreated              = 201
// StatusAccepted             = 202
// StatusNonAuthoritativeInfo = 203
// StatusNoContent            = 204
// StatusResetContent         = 205
// StatusPartialContent       = 206

// StatusMultipleChoices   = 300
// StatusMovedPermanently  = 301
// StatusFound             = 302
// StatusSeeOther          = 303
// StatusNotModified       = 304
// StatusUseProxy          = 305
// StatusTemporaryRedirect = 307

// StatusBadRequest                   = 400
// StatusUnauthorized                 = 401
// StatusPaymentRequired              = 402
// StatusForbidden                    = 403
// StatusNotFound                     = 404
// StatusMethodNotAllowed             = 405
// StatusNotAcceptable                = 406
// StatusProxyAuthRequired            = 407
// StatusRequestTimeout               = 408
// StatusConflict                     = 409
// StatusGone                         = 410
// StatusLengthRequired               = 411
// StatusPreconditionFailed           = 412
// StatusRequestEntityTooLarge        = 413
// StatusRequestURITooLong            = 414
// StatusUnsupportedMediaType         = 415
// StatusRequestedRangeNotSatisfiable = 416
// StatusExpectationFailed            = 417
// StatusTeapot                       = 418

// StatusInternalServerError     = 500
// StatusNotImplemented          = 501
// StatusBadGateway              = 502
// StatusServiceUnavailable      = 503
// StatusGatewayTimeout          = 504
// StatusHTTPVersionNotSupported = 505

var (
	// common
	UnexpectedException      = Error{"common.unexpected.error", http.StatusInternalServerError, "Unknown Error."}
	UnknownResource          = Error{"common.unknown.resource", http.StatusBadRequest, "Unknown Resource."}
	ParameterValidatedFailed = Error{"common.parameter.validated.error", http.StatusBadRequest, "Parameter validated error."}
	JSONParsingFailed        = Error{"common.json.parsed.error", http.StatusBadRequest, "Parsing json string failed."}

	// auth
	AccessTokenEmpty        = Error{"auth.accesstoken.empty.error", http.StatusUnauthorized, "Authorization header should be in the format 'Bearer {access_token}'."}
	OAuthInvalidAccessToken = Error{"auth.invalid.accesstoken.error", http.StatusUnauthorized, "Invalid access token."}
	OAuthPermissionDenied   = Error{"auth.oauth.permission.denied.error", http.StatusForbidden, "Access token scope unmatched."}

	// notify
	PairNotFound = Error{"notify.pair.notfound.error", http.StatusBadRequest, "Pair is not found."}
)
