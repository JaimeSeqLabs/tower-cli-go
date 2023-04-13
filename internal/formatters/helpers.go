package formatters

import (
	"fmt"
)

func FormatOrgID(orgID int64, serverUrl string, orgName string) string {
	return FormatLink(
		fmt.Sprintf("%d", orgID),
		fmt.Sprintf("%s/orgs/%s/workspaces", serverUrl, orgName),
	)
}

// FormatLink makes title as a clickable link if ANSI is enabled
func FormatLink(title, link string) string {
	// TODO
	/* if !ansiEnabled {
		return title
	} */
	return "\u001b]8;;" + link + "\u001b\\" + title + "\u001b]8;;\u001b\\"
}
