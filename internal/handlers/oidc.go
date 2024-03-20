package handlers

import (
	oauthelia2 "authelia.com/provider/oauth2"
	"github.com/authelia/authelia/v4/internal/authentication"
	"github.com/authelia/authelia/v4/internal/middlewares"
	"github.com/authelia/authelia/v4/internal/model"
	"github.com/authelia/authelia/v4/internal/oidc"
	"github.com/authelia/authelia/v4/internal/session"
	"github.com/authelia/authelia/v4/internal/utils"
	"github.com/google/uuid"
)

func oidcGrantRequests(ar oauthelia2.AuthorizeRequester, consent *model.OAuth2ConsentSession, userSession *session.UserSession) (extraClaims map[string]any) {
	extraClaims = map[string]any{}

	for _, scope := range consent.GrantedScopes {
		if ar != nil {
			ar.GrantScope(scope)
		}

		switch scope {
		case oidc.ScopeGroups:
			extraClaims[oidc.ClaimGroups] = userSession.Groups
		case oidc.ScopeProfile:
			extraClaims[oidc.ClaimPreferredUsername] = userSession.Username
			extraClaims[oidc.ClaimFullName] = userSession.DisplayName
		case oidc.ScopeEmail:
			if len(userSession.Emails) != 0 {
				extraClaims[oidc.ClaimPreferredEmail] = userSession.Emails[0]
				if len(userSession.Emails) > 1 {
					extraClaims[oidc.ClaimEmailAlts] = userSession.Emails[1:]
				}

				// TODO (james-d-elliott): actually verify emails and record that information.
				extraClaims[oidc.ClaimEmailVerified] = true
			}
		}
	}

	if ar != nil {
		for _, audience := range consent.GrantedAudience {
			ar.GrantAudience(audience)
		}
	}

	return extraClaims
}

func oidcApplyUserInfoClaims(ctx *middlewares.AutheliaCtx, client oidc.Client, clientID string, claims map[string]any) {
	delete(claims, oidc.ClaimJWTID)
	delete(claims, oidc.ClaimSessionID)
	delete(claims, oidc.ClaimAccessTokenHash)
	delete(claims, oidc.ClaimCodeHash)
	delete(claims, oidc.ClaimExpirationTime)
	delete(claims, oidc.ClaimNonce)

	audience, ok := claims[oidc.ClaimAudience].([]string)

	if !ok || len(audience) == 0 {
		audience = []string{client.GetID()}
	} else if !utils.IsStringInSlice(clientID, audience) {
		audience = append(audience, clientID)
	}

	claims[oidc.ClaimAudience] = audience

	oidcUpdateClaims(ctx, claims)
}

func oidcUpdateClaims(ctx *middlewares.AutheliaCtx, claims map[string]any) {
	var (
		claim      any
		identifier *model.UserOpaqueIdentifier
		details    *authentication.UserDetails
		subject    uuid.UUID
		subjectStr string
		ok         bool
		err        error
	)

	if claim, ok = claims[oidc.ClaimSubject]; !ok {
		return
	}

	if !oidcHasClaims(claims, oidc.ClaimPreferredUsername, oidc.ClaimGroups, oidc.ClaimPreferredEmail) {
		return
	}

	if subjectStr, ok = claim.(string); !ok {
		return
	}

	if subject, err = uuid.Parse(subjectStr); err != nil {
		return
	}

	if identifier, err = ctx.Providers.StorageProvider.LoadUserOpaqueIdentifier(ctx, subject); err != nil {
		return
	}

	if details, err = ctx.Providers.UserProvider.GetDetails(identifier.Username); err != nil {
		return
	}

	oidcUpdateClaimsWithUserDetails(details, claims)
}

func oidcUpdateClaimsWithUserDetails(details *authentication.UserDetails, claims map[string]any) {
	var ok bool

	if _, ok = claims[oidc.ClaimPreferredUsername]; ok {
		claims[oidc.ClaimPreferredUsername] = details.Username
	}

	if _, ok = claims[oidc.ClaimGroups]; ok {
		claims[oidc.ClaimGroups] = details.Groups
	}
	if _, ok = claims[oidc.ClaimPreferredEmail]; ok {
		if n := len(details.Emails); n == 0 {
			delete(claims, oidc.ClaimPreferredEmail)
			delete(claims, oidc.ClaimEmailAlts)
			delete(claims, oidc.ClaimEmailVerified)
		} else {
			claims[oidc.ClaimPreferredEmail] = details.Emails[0]

			if n > 1 {
				claims[oidc.ClaimEmailAlts] = details.Emails[:1]
			} else {
				delete(claims, oidc.ClaimEmailAlts)
			}
		}
	}
}

func oidcHasClaims(claims map[string]any, names ...string) bool {
	for _, name := range names {
		if _, ok := claims[name]; ok {
			return true
		}
	}

	return false
}
