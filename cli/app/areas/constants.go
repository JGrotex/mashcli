package areas

const IODOCS_EXPORT_FIELDS = `definition, created, ServiceId, defaultApi`

const SERVICES_COLLECTION_FIELDS = `id, name`

const SERVICE_EXPORT_FIELDS = `id, name, created, updated,editorHandle, revisionNumber, robotsPolicy
, crossdomainPolicy,description, errorSets, errorSets.name, errorSets.type, errorSets.jsonp
, errorSets.jsonpType, errorSets.errorMessages, qpsLimitOverall
, rfc3986Encode, securityProfile, version, cache, roles, roles.id
, roles.created, roles.updates, roles.name, roles.action
, endpoints.allowMissingApiKey, endpoints.apiKeyValueLocationKey
, endpoints.apiKeyValueLocations, endpoints.apiMethodDetectionKey
, endpoints.apiMethodDetectionLocations
, endpoints.cache.clientSurrogateControlEnabled
, endpoints.cache.contentCacheKeyHeaders
, endpoints.connectionTimeoutForSystemDomainRequest
, endpoints.connectionTimeoutForSystemDomainResponse
, endpoints.cookiesDuringHttpRedirectsEnabled, endpoints.cors
, endpoints.cors.allDomainsEnabled, endpoints.cors.maxAge
, endpoints.customRequestAuthenticationAdapter
, endpoints.dropApiKeyFromIncomingCall, endpoints.forceGzipOfBackendCallid
, name, created, updated, editorHandle, revisionNumber, robotsPolicy
, crossdomainPolicy, description, errorSets, errorSets.name, errorSets.type
, errorSets.jsonp, errorSets.jsonpType, errorSets.errorMessages
, qpsLimitOverall, rfc3986Encode, securityProfile, version, cache, roles
, roles.id, roles.created, roles.updates, roles.name, roles.action
, endpoints.allowMissingApiKey, endpoints.apiKeyValueLocationKey
, endpoints.apiKeyValueLocations, endpoints.apiMethodDetectionKey
, endpoints.apiMethodDetectionLocations
, endpoints.cache.clientSurrogateControlEnabled
, endpoints.cache.contentCacheKeyHeaders
, endpoints.connectionTimeoutForSystemDomainRequest
, endpoints.connectionTimeoutForSystemDomainResponse
, endpoints.cookiesDuringHttpRedirectsEnabled, endpoints.cors
, endpoints.cors.allDomainsEnabled, endpoints.cors.maxAge
, endpoints.customRequestAuthenticationAdapter
, endpoints.dropApiKeyFromIncomingCall, endpoints.forceGzipOfBackendCall
, endpoints.gzipPassthroughSupportEnabled
, endpoints.headersToExcludeFromIncomingCall, endpoints.highSecurity
, endpoints.hostPassthroughIncludedInBackendCallHeader
, endpoints.inboundSslRequired, endpoints.jsonpCallbackParameter
, endpoints.jsonpCallbackParameterValue, endpoints.scheduledMaintenanceEvent
, endpoints.forwardedHeaders, endpoints.returnedHeaders, endpoints.methods
, endpoints.methods.name, endpoints.methods.sampleJsonResponse
, endpoints.methods.sampleXmlResponse, endpoints.methods.responseFilters
, endpoints.methods.responseFilters.id, endpoints.methods.responseFilters.name
, endpoints.methods.responseFilters.created
, endpoints.methods.responseFilters.updated
, endpoints.methods.responseFilters.notes
, endpoints.methods.responseFilters.xmlFilterFields
, endpoints.methods.responseFilters.jsonFilterFields, endpoints.name
, endpoints.numberOfHttpRedirectsToFollow, endpoints.outboundRequestTargetPath
, endpoints.outboundRequestTargetQueryParameters
, endpoints.outboundTransportProtocol, endpoints.processor
, endpoints.publicDomains, endpoints.requestAuthenticationType
, endpoints.scheduledMaintenanceEvent, endpoints.scheduledMaintenanceEvent.id
, endpoints.scheduledMaintenanceEvent.name
, endpoints.scheduledMaintenanceEvent.startDateTime
, endpoints.scheduledMaintenanceEvent.endDateTime
, endpoints.scheduledMaintenanceEvent.endpoints, endpoints.requestPathAlias
, endpoints.requestProtocol, endpoints.oauthGrantTypes
, endpoints.stringsToTrimFromApiKey, endpoints.supportedHttpMethods
, endpoints.systemDomainAuthentication
, endpoints.systemDomainAuthentication.type
, endpoints.systemDomainAuthentication.username
, endpoints.systemDomainAuthentication.certificate
, endpoints.systemDomainAuthentication.password, endpoints.systemDomains
, endpoints.trafficManagerDomain, endpoints.useSystemDomainCredentials`
