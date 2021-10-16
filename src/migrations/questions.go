package migrations

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"white-hat.api/src/entities"
	"white-hat.api/src/enums"
)

func InsertQuestions(database *mongo.Database) {
	questions := []entities.Question{{
		Id:       1,
		Focus:    enums.Focuses.Governance,
		Category: enums.Categories.PlanningAndStrategy,
		Content:  "The FRFI has published a cyber risk strategy that is aligned with the technology and business strategies.",
	}, {
		Id:       2,
		Focus:    enums.Focuses.Governance,
		Category: enums.Categories.PlanningAndStrategy,
		Content:  "The FRFI has an established cyber risk framework (e.g., a complete set of elements including policies, standards, roles and responsibilities, risk management processes, risk taxonomy, risk appetite and emerging threats and technologies) in support of the cyber risk strategy, and ongoing threat, risk and incident management.",
	}, {
		Id:       3,
		Focus:    enums.Focuses.Governance,
		Category: enums.Categories.PlanningAndStrategy,
		Content:  "The FRFI conducts regular reviews of the cyber risk strategy and cyber risk framework, to ensure compliance with legal and regulatory requirements.",
	}, {
		Id:       4,
		Focus:    enums.Focuses.Governance,
		Category: enums.Categories.PlanningAndStrategy,
		Content:  "The FRFI considers cyber risk compliance requirements, identified risks, current and emerging threats, and potential incident related impacts on operations and services, as inputs to planning and prioritizing cyber risk projects, programs and budgets.",
	}, {
		Id:       5,
		Focus:    enums.Focuses.Governance,
		Category: enums.Categories.PlanningAndStrategy,
		Content:  "The FRFI has appointed an executive responsible for the cyber risk strategy, the cyber risk framework and for cyber risk awareness and knowledge at the executive level.",
	}, {
		Id:       6,
		Focus:    enums.Focuses.Governance,
		Category: enums.Categories.Policy,
		Content:  "The FRFI has documented cyber risk policies to explain staff and contractor roles, responsibilities, rules and constraints as well as possible penalties for non-compliance.",
	}, {
		Id:       7,
		Focus:    enums.Focuses.Governance,
		Category: enums.Categories.Policy,
		Content:  "The roles and responsibilities of each of the three lines of defence and other stakeholders are clearly described within the cyber risk framework.",
	}, {
		Id:       8,
		Focus:    enums.Focuses.Governance,
		Category: enums.Categories.RiskManagement,
		Content:  "Key risk and performance indicators as well as thresholds have been established for the FRFI’s key cyber risk and controls. The risk indicators should align with the cyber risk appetite as stated in the cyber risk framework.",
	}, {
		Id:       9,
		Focus:    enums.Focuses.Governance,
		Category: enums.Categories.RiskManagement,
		Content:  "Cyber risks to the organization and its programs or customers are regularly reviewed, prioritized, escalated, explained to the appropriate executives or senior management, and those risks are prioritized for mitigation.",
	}, {
		Id:       10,
		Focus:    enums.Focuses.Governance,
		Category: enums.Categories.RiskManagement,
		Content:  "The second line of defence regularly provides an independent review of the various cyber risk assessments and other control activities conducted by the first line of defence.",
	}, {
		Id:       11,
		Focus:    enums.Focuses.Governance,
		Category: enums.Categories.RiskManagement,
		Content:  "The FRFI ensures that background checks have been implemented for personnel/contractors and at third party providers, commensurate with the sensitivity and cyber risk needs of FRFI assets being managed.",
	}, {
		Id:       12,
		Focus:    enums.Focuses.Governance,
		Category: enums.Categories.RiskManagement,
		Content:  "The FRFI has implemented a formal process for risk acceptance that is measured, tracked and reported.",
	}, {
		Id:       13,
		Focus:    enums.Focuses.Identify,
		Category: enums.Categories.BusinessEnvironment,
		Content:  "The FRFI has allocated sufficient and skilled resources for the sustainment of cyber risk programs, systems, roles and services.  ",
	}, {
		Id:       14,
		Focus:    enums.Focuses.Identify,
		Category: enums.Categories.BusinessEnvironment,
		Content:  "The FRFI has identified its critical technology assets and has implemented appropriate controls to ensure confidentiality, integrity and availability. The controls are regularly reviewed and tested.",
	}, {
		Id:       15,
		Focus:    enums.Focuses.Identify,
		Category: enums.Categories.BusinessEnvironment,
		Content:  "The FRFI ensures that contracts for outsourcing and external services (e.g., third party providers, Cloud Service Providers) include supplier and service provider responsibilities for the security of the FRFI’s information.",
	}, {
		Id:       16,
		Focus:    enums.Focuses.Identify,
		Category: enums.Categories.AssetManagement,
		Content:  "The FRFI maintains a configuration management database (CMDB) or similar utility for documenting and tracking IT component configurations (i.e., hardware, software, network addresses, security systems, dependencies, etc.).",
	}, {
		Id:       17,
		Focus:    enums.Focuses.Identify,
		Category: enums.Categories.AssetManagement,
		Content:  "The FRFI's IT assets and information are classified and managed according to a classification scheme.",
	}, {
		Id:       18,
		Focus:    enums.Focuses.Identify,
		Category: enums.Categories.AssetManagement,
		Content:  "The FRFI has established procedures for the disposal or destruction of IT assets.",
	}, {
		Id:       19,
		Focus:    enums.Focuses.Identify,
		Category: enums.Categories.RiskAssessment,
		Content:  "The FRFI conducts Threat and Risk Assessments in the early stages of new initiatives/projects or prior to changes in existing systems and data, to identify and prioritize threats, risks and remediation options.",
	}, {
		Id:       20,
		Focus:    enums.Focuses.Identify,
		Category: enums.Categories.RiskAssessment,
		Content:  "The FRFI should periodically assess their cyber risks, which will require consideration for and assessment of the robustness, currency and completeness of the cyber risk practices and controls.",
	}, {
		Id:       21,
		Focus:    enums.Focuses.Identify,
		Category: enums.Categories.RiskAssessment,
		Content:  "The FRFI conducts regular penetration testing against the network, Cloud environment and all critical IT systems to identify security gaps and deficiencies, and to affirm strengths.",
	}, {
		Id:       22,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.IdentityManagementAndAccessControl,
		Content:  "The FRFI implements a consistent access control model (e.g., Role Based Access Control) across all critical systems.",
	}, {
		Id:       23,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.IdentityManagementAndAccessControl,
		Content:  "The FRFI requires that all persons, systems or services be identified, authenticated and authorized prior to granting access to FRFI systems, services or data.",
	}, {
		Id:       24,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.IdentityManagementAndAccessControl,
		Content:  "The FRFI consistently applies the principle of 'least privilege', such that the permissions and access granted to an authenticated person, system or service is sufficient to their operational need, and no higher.",
	}, {
		Id:       25,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.IdentityManagementAndAccessControl,
		Content:  "The FRFI ensures that permissions are revoked and accounts or active connections are terminated, when no longer required.",
	}, {
		Id:       26,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.IdentityManagementAndAccessControl,
		Content:  "The FRFI implements  Multi-Factor Authentication for access to critical systems and for remote access to the FRFI network.",
	}, {
		Id:       27,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.IdentityManagementAndAccessControl,
		Content:  "The FRFI encrypts and securely stores identity and access control credentials (e.g. passwords), separate from other data.",
	}, {
		Id:       28,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.IdentityManagementAndAccessControl,
		Content:  "Privileged account credentials are managed, monitored and secured.",
	}, {
		Id:       29,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.NetworkSecurity,
		Content:  "The FRFI follows a positive security model for network security, allowing only pre-defined and authorized traffic (IP addresses, protocols, ports, etc.).",
	}, {
		Id:       30,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.NetworkSecurity,
		Content:  "The FRFI defines logical network zones, and applies controls to segregate and limit or block traffic between those zones, to help track, manage and secure the assets within those zones.",
	}, {
		Id:       31,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.NetworkSecurity,
		Content:  "The FRFI places all internet facing systems and services in a DMZ or similar, segregated and closely monitored network zone with carefully secured and limited connection into the broader environment.",
	}, {
		Id:       32,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.NetworkSecurity,
		Content:  "The FRFI engages in ongoing Threat Hunting (e.g., using manual techniques and machine learning tools) to proactively identify and isolate advanced threats which may not be detected by automated tools.",
	}, {
		Id:       33,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.NetworkSecurity,
		Content:  "The FRFI implements critical network security and traffic management controls to be fault tolerant, and to fail securely, so that security will not be compromised during any fault, outage or security incident.",
	}, {
		Id:       34,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.NetworkSecurity,
		Content:  "The FRFI limits remote access and connection options to authorized personnel, including  third party providers, and secures all remote sessions (e.g., with session encryption, MFA, session timeouts).",
	}, {
		Id:       35,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.DataSecurity,
		Content:  "The FRFI has implemented data loss prevention (DLP) controls across all technology assets for data at rest, data in use and data in transit to identify attempts at unauthorized data exfiltration, and to automatically limit or stop associated data loss.  ",
	}, {
		Id:       36,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.DataSecurity,
		Content:  "The FRFI assesses all external data interfaces (e.g. APIs) to ascertain if implemented security controls are appropriate to the sensitivity of the FRFI's data.",
	}, {
		Id:       37,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.DataSecurity,
		Content:  "The FRFI uses automated tools to examine all data (including source code and configuration data) prior to its introduction into FRFI's systems, to identify and quarantine unauthorized executable code (e.g., malware), and potentially harmful data.",
	}, {
		Id:       38,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.DataSecurity,
		Content:  "The FRFI encrypts all data to be physically transported internally or externally (e.g., on portable/removable storage media), and restricts such data transport to authorized individuals only.",
	}, {
		Id:       39,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.DataSecurity,
		Content:  "FRFI personnel 'work from home' solutions are implemented with strong end-point controls (e.g., in laptops or other mobile devices) to maintain robust security.",
	}, {
		Id:       40,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.DataSecurity,
		Content:  "The FRFI conducts regular, automated back-ups of its data.",
	}, {
		Id:       41,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.VulnerabilityManagement,
		Content:  "The FRFI has published and implemented a Vulnerability and Patch Management Program, providing rules and guidance on roles, responsibilities, the FRFI's vulnerability management life cycle, vulnerability prioritization (e.g., based on risk), remediation timeframes, exception/exemption approvals, monitoring and reporting, and tools to be applied.",
	}, {
		Id:       42,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.VulnerabilityManagement,
		Content:  "The FRFI has identified reputable sources of vulnerability information, and subscribes to recognized and authoritative vulnerability reporting services.",
	}, {
		Id:       43,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.VulnerabilityManagement,
		Content:  "The FRFI conducts regular vulnerability scanning to identify new vulnerabilities.",
	}, {
		Id:       44,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.VulnerabilityManagement,
		Content:  "The FRFI prioritizes identified vulnerabilities for resolution, based on the risk and potential impact represented.",
	}, {
		Id:       45,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.VulnerabilityManagement,
		Content:  "The FRFI has an exception/exemption management process that documents and requires appropriate management approvals, for delays or exceptions to vulnerability remediation (e.g., through application of vendor supplied patches).",
	}, {
		Id:       46,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.VulnerabilityManagement,
		Content:  "The FRFI verifies and tests vulnerability patches, prior to general deployment within the operational environment.",
	}, {
		Id:       47,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.VulnerabilityManagement,
		Content:  "The FRFI identifies contingency options for reversing vulnerability resolution measures (e.g., through roll-back of patches), prior to general deployment.",
	}, {
		Id:       48,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.VulnerabilityManagement,
		Content:  "The FRFI has established timelines for applying patches based on risk.",
	}, {
		Id:       49,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.ChangeAndConfigurationManagement,
		Content:  "The FRFI has created, documented and implemented standardized, secure configurations for all hardware and software (e.g., Operating Systems, VMs, desktop image).",
	}, {
		Id:       50,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.ChangeAndConfigurationManagement,
		Content:  "The FRFI hardens all critical systems and networks.",
	}, {
		Id:       51,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.ChangeAndConfigurationManagement,
		Content:  "The FRFI enforces security policies through the use of automated tools to identify and block use of unauthorized software and hardware across all of its systems.",
	}, {
		Id:       52,
		Focus:    enums.Focuses.Defend,
		Category: enums.Categories.ChangeAndConfigurationManagement,
		Content:  "The FRFI has documented and implemented a Change Management process, to formally identify, assess, approve and document configuration changes.",
	}, {
		Id:       53,
		Focus:    enums.Focuses.Detect,
		Category: enums.Categories.MonitoringAndLogging,
		Content:  "The FRFI monitors all networks, sub-networks, and interfaces to identify information security events such as unauthorized connection attempts, unusual or suspicious traffic patterns or use of unauthorized ports and protocols.",
	}, {
		Id:       54,
		Focus:    enums.Focuses.Detect,
		Category: enums.Categories.MonitoringAndLogging,
		Content:  "The FRFI has established requirements for log collection and retention across all IT assets.",
	}, {
		Id:       55,
		Focus:    enums.Focuses.Detect,
		Category: enums.Categories.MonitoringAndLogging,
		Content:  "The FRFI uses automated tools (e.g., a SIEM or Log Analytics Tool) to collect, aggregate and analyze event data in real time or near to real time (e.g., anomalous activity), and alerts personnel according to established use cases and rules.",
	}, {
		Id:       56,
		Focus:    enums.Focuses.Detect,
		Category: enums.Categories.MonitoringAndLogging,
		Content:  "The FRFI's network monitoring and management processes are integrated with Incident Response processes, for rapid and formal escalations, communications and resolution of priority events.  ",
	}, {
		Id:       57,
		Focus:    enums.Focuses.Detect,
		Category: enums.Categories.MonitoringAndLogging,
		Content:  "FRFI and service provider logs and related records pertaining to security events are encrypted, time stamped and archived for later reference as needed. Event logs are maintained in a secure location.",
	}, {
		Id:       58,
		Focus:    enums.Focuses.Detect,
		Category: enums.Categories.BenchmarkingReviewsAndAssessments,
		Content:  "The FRFI conducts ongoing and periodic assessments (e.g., of cyber risk processes), with reference to external security frameworks, best practices, and emerging vulnerabilities to identify control gaps or deficiencies across the FRFI environment, and to identify opportunities and recommendations for improvement.",
	}, {
		Id:       59,
		Focus:    enums.Focuses.Detect,
		Category: enums.Categories.BenchmarkingReviewsAndAssessments,
		Content:  "The FRFI conducts ongoing reviews to determine policy compliance.",
	}, {
		Id:       60,
		Focus:    enums.Focuses.Detect,
		Category: enums.Categories.BenchmarkingReviewsAndAssessments,
		Content:  "The FRFI conducts regular, automated reviews of IT infrastructure (e.g., endpoints) to verify that security controls are configured and functioning as expected.",
	}, {
		Id:       61,
		Focus:    enums.Focuses.Detect,
		Category: enums.Categories.BenchmarkingReviewsAndAssessments,
		Content:  "The FRFI communicates security assessment and audit results to appropriate internal management, and to the executive(s) responsible for the cyber risk framework.",
	}, {
		Id:       62,
		Focus:    enums.Focuses.Detect,
		Category: enums.Categories.SecureSoftwareDevelopment,
		Content:  "The FRFI treats security and the adoption of security best practices as a priority within the software development life cycle.",
	}, {
		Id:       63,
		Focus:    enums.Focuses.Detect,
		Category: enums.Categories.SecureSoftwareDevelopment,
		Content:  "The FRFI deploys all software, including off the shelf products, in a segregated test environment, and executes relevant testing and security scans, prior to general deployment.",
	}, {
		Id:       64,
		Focus:    enums.Focuses.Detect,
		Category: enums.Categories.SecureSoftwareDevelopment,
		Content:  "The FRFI verifies the code from external sources is from a reputable and recognized source (e.g., by review of digital signature, or hash function).",
	}, {
		Id:       65,
		Focus:    enums.Focuses.Respond,
		Category: enums.Categories.IncidentManagement,
		Content:  "The FRFI’s Incident Management standard is designed to respond rapidly to cyber risk incidents.",
	}, {
		Id:       66,
		Focus:    enums.Focuses.Respond,
		Category: enums.Categories.IncidentManagement,
		Content:  "The FRFI has established a 'whole of organization' response including but not limited to: cyber risk team, IT team, business owner, legal, privacy, and communications (public affairs), and others as required and has developed playbooks and runbooks as needed.",
	}, {
		Id:       67,
		Focus:    enums.Focuses.Respond,
		Category: enums.Categories.IncidentManagement,
		Content:  "The FRFI regularly exercises the Incident Management standard.",
	}, {
		Id:       68,
		Focus:    enums.Focuses.Respond,
		Category: enums.Categories.IncidentManagement,
		Content:  "The FRFI has an established communication plan that includes, but is not limited to, customers/clients, business partners, provincial or federal regulatory or security agencies, law enforcement, internal staff, and others as appropriate.",
	}, {
		Id:       69,
		Focus:    enums.Focuses.Respond,
		Category: enums.Categories.IncidentManagement,
		Content:  "The FRFI conducts post-incident analysis to identify root cause, vulnerabilities, remedies and to document lessons learned for future reference by staff.",
	}, {
		Id:       70,
		Focus:    enums.Focuses.Recover,
		Category: enums.Categories.TestingAndPlanning,
		Content:  "The FRFI regularly tests data back-ups to verify their integrity, and to confirm that restoration of data is feasible in case of need.",
	}, {
		Id:       71,
		Focus:    enums.Focuses.Recover,
		Category: enums.Categories.TestingAndPlanning,
		Content:  "The FRFI develops and tests playbooks to ensure timely restoration of data, systems or services impacted by cyber risk incidents.",
	}, {
		Id:       72,
		Focus:    enums.Focuses.Recover,
		Category: enums.Categories.TestingAndPlanning,
		Content:  "The FRFI has a Disaster Recovery Plan and/or Business Continuity Plan to execute in the event of a material cyber risk incident.",
	}, {
		Id:       73,
		Focus:    enums.Focuses.Learn,
		Category: enums.Categories.ContinuousImprovement,
		Content:  "The FRFI regularly reviews its IT environment and mitigates risks from end of life/support hardware and software.",
	}, {
		Id:       74,
		Focus:    enums.Focuses.Learn,
		Category: enums.Categories.ContinuousImprovement,
		Content:  "The FRFI conducts threat modeling to improve cyber resilience.",
	}, {
		Id:       75,
		Focus:    enums.Focuses.Learn,
		Category: enums.Categories.ContinuousImprovement,
		Content:  "The FRFI conducts regular simulation exercises (e.g. ransomware, DDOS) to validate response plans, and familiarize stakeholders with their roles and responsibilities.",
	}, {
		Id:       76,
		Focus:    enums.Focuses.Learn,
		Category: enums.Categories.ContinuousImprovement,
		Content:  "The FRFI subscribes to reputable information sources for understanding of emerging threats, trends, vulnerabilities, and cyber risk best practices.",
	}, {
		Id:       77,
		Focus:    enums.Focuses.Learn,
		Category: enums.Categories.ContinuousImprovement,
		Content:  "The FRFI keeps abreast of new and emerging technologies and their impact on cyber risk.",
	}, {
		Id:       78,
		Focus:    enums.Focuses.Learn,
		Category: enums.Categories.SecurityEducation,
		Content:  "The FRFI has a cyber risk education and awareness plan for employees, customers and other stakeholders.",
	}, {
		Id:       79,
		Focus:    enums.Focuses.Learn,
		Category: enums.Categories.SecurityEducation,
		Content:  "The FRFI provides for necessary and appropriate training for cyber risk personnel, to maintain current knowledge and skills, in support of their roles and responsibilities.",
	}, {
		Id:       80,
		Focus:    enums.Focuses.Learn,
		Category: enums.Categories.SecurityEducation,
		Content:  "The FRFI provides all staff with ongoing security awareness education to make them aware of their role and responsibilities with respect to cyber risk, to help them identify threats and to explain cyber risk best practices.",
	}, {
		Id:       81,
		Focus:    enums.Focuses.Learn,
		Category: enums.Categories.SecurityEducation,
		Content:  "FRFI executives and senior management are regularly briefed on cyber risk trends, identified risks, incidents, planned cyber risk initiatives and associated, potential impacts on the organization.",
	}, {
		Id:       82,
		Focus:    enums.Focuses.ThirdPartyProviders,
		Category: enums.Categories.GovernanceAndManagement,
		Content:  "The FRFI has identified and assessed cyber risk arising from its third party providers. The risk assessment is regularly refreshed and drives the frequency and intensity of risk management activities (e.g., due diligence, contract obligations, monitoring, reporting and assurance activities).",
	}, {
		Id:       83,
		Focus:    enums.Focuses.ThirdPartyProviders,
		Category: enums.Categories.GovernanceAndManagement,
		Content:  "The FRFI ensures that cyber risk controls implemented by third party providers are appropriate to the sensitivity of FRFI data, and are as robust and comprehensive as those which the FRFI implements on premise.",
	}, {
		Id:       84,
		Focus:    enums.Focuses.ThirdPartyProviders,
		Category: enums.Categories.GovernanceAndManagement,
		Content:  "FRFI has developed exit strategies for critical third party providers that outline possible cyber related scenarios, triggers and alternative solutions developed and assessed for viability. ",
	}, {
		Id:       85,
		Focus:    enums.Focuses.ThirdPartyProviders,
		Category: enums.Categories.GovernanceAndManagement,
		Content:  "The FRFI periodically obtains independent assurance of third party controls using various methods such as audit certifications, internal audit reviews, pooled audits etc.",
	}, {
		Id:       86,
		Focus:    enums.Focuses.ThirdPartyProviders,
		Category: enums.Categories.GovernanceAndManagement,
		Content:  "The FRFI ensures that the third party provider has established incident response playbooks, including procedures as to when and how the FRFI will be informed of any impact on its systems, services or data.",
	}, {
		Id:       87,
		Focus:    enums.Focuses.ThirdPartyProviders,
		Category: enums.Categories.GovernanceAndManagement,
		Content:  "The FRFI verifies that third party providers completely delete all FRFI data including backups, when no longer required.",
	}, {
		Id:       88,
		Focus:    enums.Focuses.ThirdPartyProviders,
		Category: enums.Categories.CloudServiceProviders,
		Content:  "The FRFI has a documented Cloud exit strategy that defines cyber risk processes, roles and responsibilities to be implemented if the FRFI discontinues CSP services (e.g., to migrate to a different CSP).",
	}, {
		Id:       89,
		Focus:    enums.Focuses.ThirdPartyProviders,
		Category: enums.Categories.CloudServiceProviders,
		Content:  "The FRFI ensures that all cyber risk roles and responsibilities (e.g., for implementation and management of controls), are clearly documented and agreed by all parties when implementing Cloud services (IaaS, PaaS, and SaaS).",
	}, {
		Id:       90,
		Focus:    enums.Focuses.ThirdPartyProviders,
		Category: enums.Categories.CloudServiceProviders,
		Content:  "Centralized logging and monitoring processes are implemented across all Cloud assets, with the capability to conduct consolidated analysis and reporting on the security posture across all platforms.",
	}}

	var array []interface{}

	for _, val := range questions {
		array = append(array, val)
	}

	_, err := database.Collection("questions").InsertMany(context.TODO(), array)

	if err != nil {
		panic(err)
	}
}
