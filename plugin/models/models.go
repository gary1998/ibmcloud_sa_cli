package models

import "time"

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type ApiEmpty struct {
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type ApiProvider struct {
	Name string `json:"name"`

	Id string `json:"id"`
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type ApiListProvidersResponse struct {
	Providers []InlineResponse200Providers `json:"providers,omitempty"`
}

type InlineResponse200Providers struct {
	Name string `json:"name"`

	Id string `json:"id"`
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type ApiListNotesResponse struct {
	Notes []InlineResponse2001Notes `json:"notes,omitempty"`
	// The next pagination token in the list response. It should be used as page_token for the following request. An empty value means no more result.
	NextPageToken string `json:"next_page_token,omitempty"`
}

type InlineResponse2001Notes struct {
	Name string `json:"name,omitempty"`

	// A one sentence description of this `Note`.
	ShortDescription string `json:"short_description"`

	// A detailed description of this `Note`.
	LongDescription string `json:"long_description"`

	// This must be 1&#58;1 with members of our oneofs, it can be used for filtering Note and Occurrence on their kind.  - FINDING&#58; The note and occurrence represent a finding.  - KPI&#58; The note and occurrence represent a KPI value.  - CARD&#58; The note represents a card showing findings and related metric values.  - CARD_CONFIGURED&#58; The note represents a card configured for a user account.  - SECTION&#58; The note represents a section in a dashboard.
	Kind string `json:"kind"`

	RelatedUrl []InlineResponse2001RelatedUrl `json:"related_url,omitempty"`

	// Time of expiration for this note, null if note does not expire.
	ExpirationTime time.Time `json:"expiration_time,omitempty"`

	// Output only. The time this note was created. This field can be used as a filter in list requests.
	CreateTime time.Time `json:"create_time,omitempty"`

	// Output only. The time this note was last updated. This field can be used as a filter in list requests.
	UpdateTime time.Time `json:"update_time,omitempty"`

	ProviderId string `json:"provider_id,omitempty"`

	Id string `json:"id"`

	// True if this `Note` can be shared by multiple accounts.
	Shared bool `json:"shared,omitempty"`

	ReportedBy InlineResponse2001ReportedBy `json:"reported_by,omitempty"`

	Finding InlineResponse2001Finding `json:"finding,omitempty"`

	Kpi InlineResponse2001Kpi `json:"kpi,omitempty"`

	Card InlineResponse2001Card `json:"card,omitempty"`

	Section InlineResponse2001Section `json:"section,omitempty"`
}

type ApiNoteRelatedUrl struct {
	Label string `json:"label,omitempty"`

	Url string `json:"url,omitempty"`
}

type ApiNote struct {
	Name string `json:"name,omitempty"`

	// A one sentence description of this `Note`.
	ShortDescription string `json:"short_description"`

	// A detailed description of this `Note`.
	LongDescription string `json:"long_description"`

	// This must be 1&#58;1 with members of our oneofs, it can be used for filtering Note and Occurrence on their kind.  - FINDING&#58; The note and occurrence represent a finding.  - KPI&#58; The note and occurrence represent a KPI value.  - CARD&#58; The note represents a card showing findings and related metric values.  - CARD_CONFIGURED&#58; The note represents a card configured for a user account.  - SECTION&#58; The note represents a section in a dashboard.
	Kind string `json:"kind"`

	RelatedUrl []InlineResponse2001RelatedUrl `json:"related_url,omitempty"`

	// Time of expiration for this note, null if note does not expire.
	ExpirationTime time.Time `json:"expiration_time,omitempty"`

	// Output only. The time this note was created. This field can be used as a filter in list requests.
	CreateTime time.Time `json:"create_time,omitempty"`

	// Output only. The time this note was last updated. This field can be used as a filter in list requests.
	UpdateTime time.Time `json:"update_time,omitempty"`

	ProviderId string `json:"provider_id,omitempty"`

	Id string `json:"id"`

	// True if this `Note` can be shared by multiple accounts.
	Shared bool `json:"shared,omitempty"`

	ReportedBy InlineResponse2001ReportedBy `json:"reported_by,omitempty"`

	Finding InlineResponse2001Finding `json:"finding,omitempty"`

	Kpi InlineResponse2001Kpi `json:"kpi,omitempty"`

	Card InlineResponse2001Card `json:"card,omitempty"`

	Section InlineResponse2001Section `json:"section,omitempty"`
}

type InlineResponse2001RelatedUrl struct {
	Label string `json:"label,omitempty"`

	Url string `json:"url,omitempty"`
}

type InlineResponse2001ReportedBy struct {

	// The id of this reporter
	Id string `json:"id"`

	// The title of this reporter
	Title string `json:"title"`

	// The url of this reporter
	Url string `json:"url,omitempty"`
}

type InlineResponse2001Finding struct {

	// Note provider-assigned severity/impact ranking - LOW&#58; Low Impact - MEDIUM&#58; Medium Impact - HIGH&#58; High Impact
	Severity string `json:"severity"`

	// Common remediation steps for the finding of this type
	NextSteps []InlineResponse2001FindingNextSteps `json:"next_steps,omitempty"`
}
type InlineResponse2001Kpi struct {

	// The aggregation type of the KPI values. - SUM&#58; A single-value metrics aggregation type that sums up numeric values   that are extracted from KPI occurrences.
	AggregationType string `json:"aggregation_type"`
}

type InlineResponse2001Card struct {

	// The section this card belongs to
	Section string `json:"section"`

	// The title of this card
	Title string `json:"title"`

	// The subtitle of this card
	Subtitle string `json:"subtitle"`

	// The order of the card in which it will appear on SA dashboard in the mentioned section
	Order int32 `json:"order,omitempty"`

	// The finding note names associated to this card
	FindingNoteNames []string `json:"finding_note_names"`

	RequiresConfiguration bool `json:"requires_configuration,omitempty"`

	// The text associated to the card's badge
	BadgeText string `json:"badge_text,omitempty"`

	// The base64 content of the image associated to the card's badge
	BadgeImage string `json:"badge_image,omitempty"`

	// The elements of this card
	Elements []InlineResponse2001CardElements `json:"elements"`
}

type InlineResponse2001Section struct {

	// The title of this section
	Title string `json:"title"`

	// The image of this section
	Image string `json:"image"`
}

type InlineResponse2001FindingNextSteps struct {

	// Title of this next step
	Title string `json:"title,omitempty"`

	// The URL associated to this next steps
	Url string `json:"url,omitempty"`
}

type InlineResponse2001CardElements struct {

	// Kind of element - NUMERIC&#58; Single numeric value - BREAKDOWN&#58; Breakdown of numeric values - TIME_SERIES&#58; Time-series of numeric values
	Kind string `json:"kind"`

	// The default time range of this card element
	DefaultTimeRange string `json:"default_time_range,omitempty"`
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type GetNotificationChannel struct {
	Channel NotificationChannel `json:"channel,omitempty"`
}

type ChannelList struct {
	Channels []NotificationChannel `json:"channels,omitempty"`
}

type NotificationChannel struct {
	Name string `json:"name"`

	// A one sentence description of this `Channel`.
	Description string `json:"description,omitempty"`

	// Type of callback URL
	Type string `json:"type"`

	// ID of Channel
	ID string `json:"channel_id"`

	// Frequency of channel
	Frequency string `json:"frequency,omitempty"`

	// Severity of the notification to be received.
	Severity Severity `json:"severity,omitempty"`

	// The callback URL which receives the notification
	Endpoint string `json:"endpoint"`

	// Channel is enabled or not. Default is disabled
	Enabled bool `json:"enabled,omitempty"`

	AlertSource []V1accountIdnotificationschannelsAlertSource `json:"alertSource,omitempty"`
}

type Severity struct {
	Low    bool `json:"low"`
	Medium bool `json:"medium"`
	High   bool `json:"high"`
}

type V1accountIdnotificationschannelsAlertSource struct {

	// Below is a list of builtin providers that you can select in addition to the ones you obtain by calling Findings API /v1/{account_id}/providers :  | provider_name | The source they represnt |  |-----|-----|  | VA  | Vulnerable image findings|  | NA  | Network Insights findings|  | ATA | Activity Insights findings|  | CERT | Certificate Manager findings|  | ALL | Special provider name to represent all the providers. Its mutually exclusive with other providers meaning either you choose ALL or you don't|
	ProviderName string `json:"provider_name,omitempty"`
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type PublicKey struct {
	Key string `json:"publicKey,omitempty"`
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type TestChannelResponse struct {
	Message string `json:"test"`
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
