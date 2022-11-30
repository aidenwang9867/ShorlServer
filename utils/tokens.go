package utils

type Token struct {
	Token             string `json:"token" bigtable:"token"`
	AccessLimitPerMin string `json:"access_limit_per_min" bigtable:"accessLimitPerMin"`
}

const tableTokens = "team2_Token"
const tokenColumnFamilyName = "TokenString"
const tokenColumnName = "token"
const accessLimitPerMinColumnName = "accessLimitPerMin"
