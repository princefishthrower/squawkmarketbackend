package supabase

// enum for the different types of user subscriptions
type SubscriptionType int

const (
	// SubscriptionTypeFree is the free subscription type
	SubscriptionTypeFree SubscriptionType = iota
	// SubscriptionTypePremium is the premium subscription type
	SubscriptionTypePremium
)
