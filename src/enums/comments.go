package enums

type Comment struct {
	Bad  string
	Good string
}

var Comments map[string]Comment = map[string]Comment{
	"Governance": {
		Bad:  "Sua organização precisa se atentar às questões de governança, às quais envolvem políticas da empresa, gerenciamento de risco, planejamentos e estratégias.",
		Good: "Sua organização possui uma boa governança, abrangendo tanto as áreas e planejamento estratégico, quanto às políticas da empresa e gerenciamento de risco.",
	},
	"Identify": {
		Bad:  "Sua organização precisa melhorar sua identificação de riscos de ciber segurança, tanto no ambiente de negócio quanto administração de recursos e avaliação de riscos.",
		Good: "Sua organização identifica com precisão possíveis riscos na avaliação de risco, no ambiente de negócio e na administração de recursos.",
	},
	"Defend": {
		Bad:  "Sua organização carece de defesas relativas à controle de acesso, gerenciamento de identidade, segurança de rede, segurança de dados, gerenciamento de vulnerabilidades e administração de configurações e mudanças.",
		Good: "Sua organização se defende satisfatoriamente quanto a ataques de redes, tentativas de acessos, hosts não identificados e vulnerabilidades.",
	},
	"Detect": {
		Bad:  "Falta uma maior atenção aos registros, benchmarking, avaliações e desenvolvimento de softwares em sua organização.",
		Good: "Sua organização é muito boa em detectar problemas de ciber segurança por meio de avaliações, registro e desenvolvimento de software seguro.",
	},
	"Respond": {
		Bad:  "Sua organização não é eficiente em responder a possíveis incidentes.",
		Good: "Sua organização é eficiente em responder e gerenciar incidentes.",
	},
	"Recover": {
		Bad:  "Sua oganização precisa melhorar na sua recuperação de danos oriundos de falhas ou ataques cibernéticos.",
		Good: "Sua organização se recupera muito bem contra ataques cibernéticos.",
	},
	"Learn": {
		Bad:  "Sua organização não aprende com seus erros e experiências.",
		Good: "Sua organização tem uma boa cultura de aprendizado, melhoramento contínuo e educação em segurança.",
	},
	"Third Party Providers": {
		Bad:  "Sua organização não lida bem com recursos fornecidos por terceiros, como provedores de serviços em nuvem.",
		Good: "Sua organização lida bem com recursos fornecidos por terceiros, como provedores de serviços em nuvem.",
	},
}
