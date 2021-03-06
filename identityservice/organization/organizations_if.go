package organization

//This file is auto-generated by go-raml
//Do not edit this file by hand since it will be overwritten during the next generation

import (
	"github.com/gorilla/mux"
	"net/http"
)

type OrganizationsInterface interface {
	// Get organizations. Authorization limits are applied to requesting user.
	// It is handler for GET /organizations
	Get(http.ResponseWriter, *http.Request)

	// Create a new organization. 1 user should be in the owners list. Validation is performed
	// to check if the securityScheme allows management on this user.
	// It is handler for POST /organizations
	Post(http.ResponseWriter, *http.Request)

	// Get organization info
	// It is handler for GET /organizations/{globalid}
	globalidGet(http.ResponseWriter, *http.Request)

	// Update organization info
	// It is handler for PUT /organizations/{globalid}
	globalidPut(http.ResponseWriter, *http.Request)

	// Assign a member to organization.
	// It is handler for POST /organizations/{globalid}/members
	globalidmembersPost(http.ResponseWriter, *http.Request)

	// Remove a member from organization
	// It is handler for DELETE /organizations/{globalid}/members/{username}
	globalidmembersusernameDelete(http.ResponseWriter, *http.Request)

	// Invite a user to become owner of an organization.
	// It is handler for POST /organizations/{globalid}/owners
	globalidownersPost(http.ResponseWriter, *http.Request)

	// Remove an owner from organization
	// It is handler for DELETE /organizations/{globalid}/owners/{username}
	globalidownersusernameDelete(http.ResponseWriter, *http.Request)

	// Get the contracts where the organization is 1 of the parties. Order descending by
	// date.
	// It is handler for GET /organizations/{globalid}/contracts
	GetContracts(http.ResponseWriter, *http.Request)

	// Get the list of pending invitations for users to join this organization.
	// It is handler for GET /organizations/{globalid}/invitations
	GetPendingInvitations(http.ResponseWriter, *http.Request)

	// Cancel a pending invitation.
	// It is handler for DELETE /organizations/{globalid}/invitations/{username}
	RemovePendingInvitation(http.ResponseWriter, *http.Request)
}

func OrganizationsInterfaceRoutes(r *mux.Router, i OrganizationsInterface) {
	r.HandleFunc("/organizations", i.Get).Methods("GET")
	r.HandleFunc("/organizations", i.Post).Methods("POST")
	r.HandleFunc("/organizations/{globalid}", i.globalidGet).Methods("GET")
	r.HandleFunc("/organizations/{globalid}", i.globalidPut).Methods("PUT")
	r.HandleFunc("/organizations/{globalid}/members", i.globalidmembersPost).Methods("POST")
	r.HandleFunc("/organizations/{globalid}/members/{username}", i.globalidmembersusernameDelete).Methods("DELETE")
	r.HandleFunc("/organizations/{globalid}/owners", i.globalidownersPost).Methods("POST")
	r.HandleFunc("/organizations/{globalid}/owners/{username}", i.globalidownersusernameDelete).Methods("DELETE")
	r.HandleFunc("/organizations/{globalid}/contracts", i.GetContracts).Methods("GET")
	r.HandleFunc("/organizations/{globalid}/invitations", i.GetPendingInvitations).Methods("GET")
	r.HandleFunc("/organizations/{globalid}/invitations/{username}", i.RemovePendingInvitation).Methods("DELETE")
}
