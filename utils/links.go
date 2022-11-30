package utils

type Link struct {
	Short string `json:"short_link" bigtable:"shortLink"`
	Long  string `json:"long_link" bigtable:"longLink"`
}

const tableLinks = "team2_Link"
const linkColumnFamilyName = "LinkString"
const shortLinkColumnName = "shortLink"
const longLinkColumnName = "longLink"
