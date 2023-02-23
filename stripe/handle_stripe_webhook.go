package stripe

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"squawkmarketbackend/supabase"
	"squawkmarketbackend/utils"

	"log"
	"net/http"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/customer"
	"github.com/stripe/stripe-go/v72/webhook"
)

// we need this custom type because apparently the stripe.Subscription type
// has a Customer key of type *Customer instead of string
// incoming webhook request has a "customer" key of type string,
// see https://stripe.com/docs/api/subscriptions/object#subscription_object-customer
type StripeWebhookSubscription struct {
	ID         string                    `json:"id"`
	CustomerID string                    `json:"customer"`
	Status     stripe.SubscriptionStatus `json:"status"`
}

func generateErrorResponse(w http.ResponseWriter, httpCode int, message string) {
	log.Printf("error: HandleStripeWebhook: %s\n", message)
	// return http error code and message
	w.WriteHeader(httpCode)
	w.Write([]byte(message))
}

// subscription handler for stripe webhooks, originally from official stripe webhook docs:
// https://stripe.com/docs/billing/quickstart
// it is a standard HTTP-like handler because Stripe can't call a graphql endpoint
func HandleStripeWebhook(w http.ResponseWriter, r *http.Request) {

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		generateErrorResponse(w, http.StatusServiceUnavailable, fmt.Sprintf("request body read failed: %+v\n", err))
		return
	}

	// Replace this endpoint secret with your endpoint's unique secret
	// If you are testing with the CLI, find the secret by running 'stripe listen'
	// If you are using an endpoint defined with the API or dashboard, look in your webhook settings
	// at https://dashboard.stripe.com/webhooks
	endpointSecret := os.Getenv("STRIPE_WEBHOOK_SECRET_PRODUCTION")

	// if the origin includes "staging", use the staging secret
	if r.Header.Get("Origin") == "https://staging.squawkmarket.com" {
		endpointSecret = os.Getenv("STRIPE_WEBHOOK_SECRET_STAGING")
	}

	signatureHeader := r.Header.Get("Stripe-Signature")
	event, err := webhook.ConstructEvent(payload, signatureHeader, endpointSecret)
	if err != nil {
		generateErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("⚠️ webhook signature verification failed: %v\n", err))
		return
	}

	// unmarshal stripe subscription object
	var subscription StripeWebhookSubscription
	err = json.Unmarshal(event.Data.Raw, &subscription)
	if err != nil {
		generateErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("webhook JSON parse failed: %+v\n", err))
		return
	}

	stripeCustomer, err := customer.Get(subscription.CustomerID, nil)
	if err != nil {
		generateErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("customer.subscription.deleted: customer.Get: failed: %+v\n", err))
		return
	}

	// check if user exists by email
	email := stripeCustomer.Email
	userID, err := supabase.GetUserIdByEmail(email)
	if err != nil {
		generateErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("customer.subscription.deleted: DoesUserExistWithEmail: failed: %+v\n", err))
		return
	}

	// create user and send magic link if user does not exist
	if userID == "" {
		userID, err = supabase.CreateUserWithEmail(email)
		if err != nil {
			generateErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("customer.subscription.deleted: CreateUserWithEmail: failed: %+v\n", err))
			return
		}
		supabase.SendUserMagicLinkToEmail(email)
	}

	// handle stripe event types
	switch event.Type {
	case "customer.subscription.deleted":
		// this fires when a subscription is canceled
		log.Printf("Subscription DELETED for %s (user id: %s).", subscription.ID, userID)
		// _, err = graphqlclient.GraphQLClient.SetSubscriberSubscription(c, userID, graphqlclient.SubscriptionTiersEnumFree)

		err = supabase.SetUserSubscription(userID, false, "monthly")

		if err != nil {
			generateErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("customer.subscription.deleted: SetSubscriberSubscriptionFree: failed: %+v\n", err))
			return
		}

	case "customer.subscription.updated":
		// this fires when a user creates, updates or deletes a subscription
		// stripe emits this event type when a payment is accepted and the subscription is set to active

		log.Printf("Subscription updated for %s (user id: %s).", subscription.ID, userID)

		if subscription.Status == stripe.SubscriptionStatusActive {
			// set user to premium
			// _, err = graphqlclient.GraphQLClient.SetSubscriberSubscription(c, userID, graphqlclient.SubscriptionTiersEnumPremium)
			err = supabase.SetUserSubscription(userID, true, "monthly")
			if err != nil {
				generateErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("customer.subscription.updated: SetSubscriberSubscriptionPremium: failed: %+v\n", err))
				return
			}

			log.Printf("Subscription ACTIVATED for %s (user id: %s).", subscription.ID, userID)

			// log the event on slack
			utils.SendSlackMessage(fmt.Sprintf("💲💲💲Subscription ACTIVATED for user with ID %s!!!", userID))
		}

	case "customer.subscription.created":
		// we don't care about this event because it's not fired if the payment is successful, but just when a subscription is initialized by the user
		log.Printf("Subscription CREATED for %s (user id: %s).", subscription.ID, userID)
	default:
		log.Printf("error: HandleStripeWebhook: Unhandled event type: %s\n", event.Type)
	}

	// w.String(http.StatusOK, "OK")
	// return a 200 status code to stripe
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
