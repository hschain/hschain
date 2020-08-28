package types

// Minting module event types
const (
	EventTypeMint = ModuleName

	AttributeKeyCurrentDayProvisions    = "current_day_provisions"
	AttributeKeyNextPeriodDayProvisions = "next_period_day_provisions"

	EventTypeTransfer = "burn"

	AttributeKeyRecipient = "recipient"
	AttributeKeySender    = "burner"

	AttributeValueCategory = ModuleName
)
