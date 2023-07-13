package ginoapi3

import (
	_ "embed"
)

const (
	DefaultSchemaVersion = "3.0.0"         // DefaultSchemaVersion default schema version. Only "3.0.0" is supported now.
	DefaultSchemaPath    = "/openapi.json" // DefaultSchemaPath default schema path
	DefaultSchemaUIPath  = "/openapi"      // DefaultSchemaUIPath default schema ui path
)

//go:embed ui/redoc/index.html
var redocUI string

// RedocUIOption redoc ui option.
// Ref: https://redocly.com/docs/api-reference-docs/configuration/functionality/
type RedocUIOption struct {
	DisableSearch                  bool   `json:"disableSearch,omitempty"`
	MinCharacterLengthToInitSearch uint   `json:"minCharacterLengthToInitSearch,omitempty"`
	ExpandDefaultServerVariables   bool   `json:"expandDefaultServerVariables,omitempty"`
	ExpandResponses                string `json:"expandResponses,omitempty"`
	ExpandSingleSchemaField        bool   `json:"expandSingleSchemaField,omitempty"`
	HideDownloadButton             bool   `json:"hideDownloadButton,omitempty"`
	HideHostname                   bool   `json:"hideHostname,omitempty"`
	HideLoading                    bool   `json:"hideLoading,omitempty"`
	HideRequestPayloadSample       bool   `json:"hideRequestPayloadSample,omitempty"`
	HideSchemaPattern              bool   `json:"hideSchemaPattern,omitempty"`
	HideOneOfDescription           bool   `json:"hideOneOfDescription,omitempty"`
	HideSchemaTitles               bool   `json:"hideSchemaTitles,omitempty"`
	HideSingleRequestSampleTab     bool   `json:"hideSingleRequestSampleTab,omitempty"`
	ShowObjectSchemaExamples       bool   `json:"showObjectSchemaExamples,omitempty"`
	HtmlTemplate                   string `json:"htmlTemplate,omitempty"`
	MaxDisplayedEnumValues         uint   `json:"maxDisplayedEnumValues,omitempty"`
	MenuToggle                     *bool  `json:"menuToggle,omitempty"`
	NativeScrollbars               bool   `json:"nativeScrollbars,omitempty"`
	OnlyRequiredInSamples          bool   `json:"onlyRequiredInSamples,omitempty"`
	PathInMiddlePanel              bool   `json:"pathInMiddlePanel,omitempty"`
	PayloadSampleIdx               uint   `json:"payloadSampleIdx,omitempty"`
	RequiredPropsFirst             bool   `json:"requiredPropsFirst,omitempty"`
	ShowWebhookVerbose             bool   `json:"showWebhookVerbose,omitempty"`
	HideSecuritySection            bool   `json:"hideSecuritySection,omitempty"`
	SimpleOneOfTypeLabel           bool   `json:"simpleOneOfTypeLabel,omitempty"`
	SortPropsAlphabetically        bool   `json:"sortPropsAlphabetically,omitempty"`
	UntrustedDefinition            bool   `json:"untrustedDefinition,omitempty"`
}

func FillPtr[T any](v T) *T {
	return &v
}
