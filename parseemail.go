package ParseEmail

import (
	"errors"
	"regexp"
	"strings"
)

// ErrMissingAt if the email doesn't have an @ at all
var ErrMissingAt = errors.New("Missing @ character to separate out Domain from username")

// ErrBadUsername if the part before the @ is bad
var ErrBadUsername = errors.New("Username portion of the email must only contain alphanumeric and certain punctuation characters")

// ErrBadDomain if the part after the @ is bad
var ErrBadDomain = errors.New("Domain portion of an email contains invalid characters")

// ErrBadEmailSyntax if the email is really bad (multiple @ or other error)
var ErrBadEmailSyntax = errors.New("Unable to parse email at all (two many @ in the address)")

// Address takes a string which should be of the form
//  username@domain
// and splits it into the username/domain portion checking for any errors
// It also splits off any tags from the email address
//
func Address(emailAddr string) (username string, domain string, err error) {
	usernameRegex := regexp.MustCompile(`^[A-Za-z0-9._%\-]+$`) // Note + for tags has been removed because we handle it special
	domainRegex := regexp.MustCompile(`^[A-Za-z0-9.\-]+\.[a-z]{2,6}$`)
	// Set default return values
	username = ""
	domain = ""
	err = nil
	// Start out by splitting on the @ (assuming it is in there)
	pieces := strings.Split(emailAddr, "@")
	switch len(pieces) {
	case 2:
		// This is the normal case
		username = pieces[0]
		domain = pieces[1]
		// We can split off any tags from the username (they have a +) and throw them away
		pieces = strings.Split(username, "+")
		username = pieces[0]
		// Validate that the username is valid syntactically
		if !usernameRegex.MatchString(username) {
			err = ErrBadUsername
			return
		}
		// same with checking the domain portion syntactically
		if !domainRegex.MatchString(domain) {
			err = ErrBadDomain
			return
		}
	case 1:
		// They gave us an email address without a domain
		err = ErrMissingAt
		return
	default:
		err = ErrBadEmailSyntax
	}
	return
}
