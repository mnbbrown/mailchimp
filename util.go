package mailchimp

import (
	"crypto/md5"
	"encoding/hex"
)

// From http://developer.mailchimp.com/documentation/mailchimp/guides/manage-subscribers-with-the-mailchimp-api/?_ga=1.198202223.1371869516.1465805004
// "In previous versions of the API, we exposed internal database IDs eid and leid
// for emails and list/email combinations. In API 3.0, we no longer use or expose
// either of these IDs. Instead, we identify your subscribers by the MD5 hash of the
// lowercase version of their email address so you can easily predict the API URL
// of a subscriber’s data."
func memberIDFromEmail(email string) string {
	idData := []byte(email)
	id := md5.Sum(idData)
	return hex.EncodeToString(id[:])
}
