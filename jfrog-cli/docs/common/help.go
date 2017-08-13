package common

const GlobalEnvVars string =
`	JFROG_CLI_LOG_LEVEL
		[Default: INFO]
		This variable determines the log level of the JFrog CLI.
		Possible values are: INFO, ERROR, and DEBUG.
		If set to ERROR, JFrog CLI logs error messages only.
		It is useful when you wish to read or parse the JFrog CLI output and do not want any other information logged.

	JFROG_CLI_OFFER_CONFIG
		[Default: true]
		If true, JFrog CLI prompts for product server details and saves them in its config file.
		To avoid having automation scripts interrupted, set this value to false, and instead,
		provide product server details using the config command.

	JFROG_CLI_HOME
		[Default: ~/.jfrog]
		This variable determines the path to the cli home directory.
		The home directory is used to save the cli configuration.
		`
