package types

// Minting module event types
const (
	EventTypeMint = ModuleName

	AttributeKeyDayProvisions    = "day_provisions"
	AttributeKeyPeriodProvisions = "period_provisions"

	EventTypeTransfer = "burn"

	AttributeKeyRecipient = "recipient"
	AttributeKeySender    = "burner"

	AttributeValueCategory = ModuleName
)
