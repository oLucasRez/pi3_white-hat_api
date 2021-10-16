package enums

const governance = "Governance"
const identify = "Identify"
const defend = "Defend"
const detect = "Detect"
const respond = "Respond"
const recover = "Recover"
const learn = "Learn"
const thirdPartyProviders = "Third Party Providers"

type focuses struct {
	Governance          string
	Identify            string
	Defend              string
	Detect              string
	Respond             string
	Recover             string
	Learn               string
	ThirdPartyProviders string
}

var Focuses focuses = focuses{
	Governance:          governance,
	Identify:            identify,
	Defend:              defend,
	Detect:              detect,
	Respond:             respond,
	Recover:             recover,
	Learn:               learn,
	ThirdPartyProviders: thirdPartyProviders,
}
