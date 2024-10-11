// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/AMAndres/iskaypetChallenge/restapi/operations/pet"
)

// NewBackendAPI creates a new Backend instance
func NewBackendAPI(spec *loads.Document) *BackendAPI {
	return &BackendAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer:    runtime.JSONConsumer(),
		UrlformConsumer: runtime.DiscardConsumer,
		XMLConsumer:     runtime.XMLConsumer(),

		JSONProducer: runtime.JSONProducer(),
		XMLProducer:  runtime.XMLProducer(),

		PetAddPetHandler: pet.AddPetHandlerFunc(func(params pet.AddPetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pet.AddPet has not yet been implemented")
		}),
		PetDeletePetHandler: pet.DeletePetHandlerFunc(func(params pet.DeletePetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pet.DeletePet has not yet been implemented")
		}),
		PetFindPetsByStatusHandler: pet.FindPetsByStatusHandlerFunc(func(params pet.FindPetsByStatusParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pet.FindPetsByStatus has not yet been implemented")
		}),
		PetFindPetsByTagsHandler: pet.FindPetsByTagsHandlerFunc(func(params pet.FindPetsByTagsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pet.FindPetsByTags has not yet been implemented")
		}),
		PetGetPetByIDHandler: pet.GetPetByIDHandlerFunc(func(params pet.GetPetByIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pet.GetPetByID has not yet been implemented")
		}),
		PetUpdatePetHandler: pet.UpdatePetHandlerFunc(func(params pet.UpdatePetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pet.UpdatePet has not yet been implemented")
		}),
		PetUpdatePetWithFormHandler: pet.UpdatePetWithFormHandlerFunc(func(params pet.UpdatePetWithFormParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pet.UpdatePetWithForm has not yet been implemented")
		}),

		// Applies when the "api_key" header is set
		APIKeyAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (api_key) api_key from header param [api_key] has not yet been implemented")
		},
		PetstoreAuthAuth: func(token string, scopes []string) (interface{}, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (petstore_auth) has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*
BackendAPI API definition of the code challenge proposed.
Useful links - [GitHub project](https://github.com/AMAndres/iskaypetchallenge) - [Swagger PetStore editor](https://editor.swagger.io/?url=https://petstore.swagger.io/v2/swagger.yaml)
*/
type BackendAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer
	// UrlformConsumer registers a consumer for the following mime types:
	//   - application/x-www-form-urlencoded
	UrlformConsumer runtime.Consumer
	// XMLConsumer registers a consumer for the following mime types:
	//   - application/xml
	XMLConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer
	// XMLProducer registers a producer for the following mime types:
	//   - application/xml
	XMLProducer runtime.Producer

	// APIKeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key api_key provided in the header
	APIKeyAuth func(string) (interface{}, error)

	// PetstoreAuthAuth registers a function that takes an access token and a collection of required scopes and returns a principal
	// it performs authentication based on an oauth2 bearer token provided in the request
	PetstoreAuthAuth func(string, []string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// PetAddPetHandler sets the operation handler for the add pet operation
	PetAddPetHandler pet.AddPetHandler
	// PetDeletePetHandler sets the operation handler for the delete pet operation
	PetDeletePetHandler pet.DeletePetHandler
	// PetFindPetsByStatusHandler sets the operation handler for the find pets by status operation
	PetFindPetsByStatusHandler pet.FindPetsByStatusHandler
	// PetFindPetsByTagsHandler sets the operation handler for the find pets by tags operation
	PetFindPetsByTagsHandler pet.FindPetsByTagsHandler
	// PetGetPetByIDHandler sets the operation handler for the get pet by Id operation
	PetGetPetByIDHandler pet.GetPetByIDHandler
	// PetUpdatePetHandler sets the operation handler for the update pet operation
	PetUpdatePetHandler pet.UpdatePetHandler
	// PetUpdatePetWithFormHandler sets the operation handler for the update pet with form operation
	PetUpdatePetWithFormHandler pet.UpdatePetWithFormHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *BackendAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *BackendAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *BackendAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *BackendAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *BackendAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *BackendAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *BackendAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *BackendAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *BackendAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the BackendAPI
func (o *BackendAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}
	if o.UrlformConsumer == nil {
		unregistered = append(unregistered, "UrlformConsumer")
	}
	if o.XMLConsumer == nil {
		unregistered = append(unregistered, "XMLConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}
	if o.XMLProducer == nil {
		unregistered = append(unregistered, "XMLProducer")
	}

	if o.APIKeyAuth == nil {
		unregistered = append(unregistered, "APIKeyAuth")
	}
	if o.PetstoreAuthAuth == nil {
		unregistered = append(unregistered, "PetstoreAuthAuth")
	}

	if o.PetAddPetHandler == nil {
		unregistered = append(unregistered, "pet.AddPetHandler")
	}
	if o.PetDeletePetHandler == nil {
		unregistered = append(unregistered, "pet.DeletePetHandler")
	}
	if o.PetFindPetsByStatusHandler == nil {
		unregistered = append(unregistered, "pet.FindPetsByStatusHandler")
	}
	if o.PetFindPetsByTagsHandler == nil {
		unregistered = append(unregistered, "pet.FindPetsByTagsHandler")
	}
	if o.PetGetPetByIDHandler == nil {
		unregistered = append(unregistered, "pet.GetPetByIDHandler")
	}
	if o.PetUpdatePetHandler == nil {
		unregistered = append(unregistered, "pet.UpdatePetHandler")
	}
	if o.PetUpdatePetWithFormHandler == nil {
		unregistered = append(unregistered, "pet.UpdatePetWithFormHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *BackendAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *BackendAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "api_key":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.APIKeyAuth)

		case "petstore_auth":
			result[name] = o.BearerAuthenticator(name, o.PetstoreAuthAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *BackendAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *BackendAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		case "application/x-www-form-urlencoded":
			result["application/x-www-form-urlencoded"] = o.UrlformConsumer
		case "application/xml":
			result["application/xml"] = o.XMLConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *BackendAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		case "application/xml":
			result["application/xml"] = o.XMLProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *BackendAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the backend API
func (o *BackendAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *BackendAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/pet"] = pet.NewAddPet(o.context, o.PetAddPetHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/pet/{petId}"] = pet.NewDeletePet(o.context, o.PetDeletePetHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/pet/findByStatus"] = pet.NewFindPetsByStatus(o.context, o.PetFindPetsByStatusHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/pet/findByTags"] = pet.NewFindPetsByTags(o.context, o.PetFindPetsByTagsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/pet/{petId}"] = pet.NewGetPetByID(o.context, o.PetGetPetByIDHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/pet"] = pet.NewUpdatePet(o.context, o.PetUpdatePetHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/pet/{petId}"] = pet.NewUpdatePetWithForm(o.context, o.PetUpdatePetWithFormHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *BackendAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *BackendAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *BackendAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *BackendAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *BackendAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[um][path] = builder(h)
	}
}
