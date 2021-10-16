package enums

const planningAndStrategy = "Planning and Strategy"
const policy = "Policy"
const riskManagement = "Risk Management"

const businessEnvironment = "Business Environment"
const assetManagement = "Asset Management"
const riskAssessment = "Risk Assessment"

const identityManagementAndAccessControl = "Identity Management and Access Control"
const networkSecurity = "Network Security"
const dataSecurity = "Data Security"
const vulnerabilityManagement = "Vulnerability Management"
const changeAndConfigurationManagement = "Change and Configuration Management"

const monitoringAndLogging = "Monitoring and Logging"
const benchmarkingReviewsAndAssessments = "Benchmarking, Reviews and Assessments"
const secureSoftwareDevelopment = "Secure Software Development"

const incidentManagement = "Incident Management"

const testingAndPlanning = "Testing and Planning"

const continuousImprovement = "Continuous Improvement"
const securityEducation = "Security Education"

const governanceAndManagement = "Governance and Management"
const cloudServiceProviders = "Cloud Service Providers"

type categories struct {
	PlanningAndStrategy string
	Policy              string
	RiskManagement      string

	BusinessEnvironment string
	AssetManagement     string
	RiskAssessment      string

	IdentityManagementAndAccessControl string
	NetworkSecurity                    string
	DataSecurity                       string
	VulnerabilityManagement            string
	ChangeAndConfigurationManagement   string

	MonitoringAndLogging              string
	BenchmarkingReviewsAndAssessments string
	SecureSoftwareDevelopment         string

	IncidentManagement string

	TestingAndPlanning string

	ContinuousImprovement string
	SecurityEducation     string

	GovernanceAndManagement string
	CloudServiceProviders   string
}

var Categories categories = categories{
	PlanningAndStrategy: planningAndStrategy,
	Policy:              policy,
	RiskManagement:      riskManagement,

	BusinessEnvironment: businessEnvironment,
	AssetManagement:     assetManagement,
	RiskAssessment:      riskAssessment,

	IdentityManagementAndAccessControl: identityManagementAndAccessControl,
	NetworkSecurity:                    networkSecurity,
	DataSecurity:                       dataSecurity,
	VulnerabilityManagement:            vulnerabilityManagement,
	ChangeAndConfigurationManagement:   changeAndConfigurationManagement,

	MonitoringAndLogging:              monitoringAndLogging,
	BenchmarkingReviewsAndAssessments: benchmarkingReviewsAndAssessments,
	SecureSoftwareDevelopment:         secureSoftwareDevelopment,

	IncidentManagement: incidentManagement,

	TestingAndPlanning: testingAndPlanning,

	ContinuousImprovement: continuousImprovement,
	SecurityEducation:     securityEducation,

	GovernanceAndManagement: governanceAndManagement,
	CloudServiceProviders:   cloudServiceProviders,
}
