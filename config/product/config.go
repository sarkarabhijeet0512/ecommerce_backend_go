package product

// Module is config module
// var Module = fx.Options(
// 	fx.Provide(
// 		New,
// 	),
// )

// type argvMeta struct {
// 	desc       string
// 	defaultVal string
// }

// // New returns a viper object.
// // This object is used to read environment variables or command line arguments.
// func New() (config *viper.Viper) {
// 	config = viper.New()

// 	confList := map[string]argvMeta{
// 		"env": {
// 			defaultVal: "development",
// 			desc:       "Environment",
// 		},
// 		"mysql_db": {
// 			defaultVal: "Test",
// 			desc:       "mysql db name",
// 		},
// 		"mysql_host": {
// 			defaultVal: "localhost",
// 			desc:       "mysql host",
// 		},
// 		"mysql_port": {
// 			defaultVal: "3306",
// 			desc:       "mysql port",
// 		},
// 		"mysql_user": {
// 			defaultVal: "root",
// 			desc:       "mysql username",
// 		},
// 		"mysql_password": {
// 			defaultVal: "",
// 			desc:       "mysql password",
// 		},
// 		"port": {
// 			defaultVal: "8765",
// 			desc:       "Port number of user API server",
// 		},
// 		"mode": {
// 			defaultVal: "server",
// 			desc:       "App mode eg. consumer, server, worker",
// 		},
// 		"log_level": {
// 			defaultVal: "debug",
// 			desc:       "Log level to be printed. List of log level by Priority - debug, info, warn, error, dpanic, panic, fatal",
// 		},
// 		"aws_region": {
// 			defaultVal: "ap-south-1",
// 			desc:       "AWS region",
// 		},
// 		"aws_bucket_name": {
// 			defaultVal: "akstestss3",
// 			desc:       "AWS s3 bucket name",
// 		},
// 		"aws_key_id": {
// 			defaultVal: "AKIA34XKABSDXUEVSKW5",
// 			desc:       "AWS access key",
// 		},
// 		"aws_secret_key": {
// 			defaultVal: "g9k2KqSu/QezqT+oqAAGl0yEHeu9oT2hj0Po9q2H",
// 			desc:       "AWS secret access key",
// 		},
// 		"zoho_api_key": {
// 			defaultVal: "1000.EQJ3183NUD5IZANWD3DGBHUMCPW4UF",
// 			desc:       "zoho api key id",
// 		},
// 		"zoho_api_secret": {
// 			defaultVal: "e663672395555e0c8edea1e612476077e62f386d9e",
// 			desc:       "zoho api key secret",
// 		},
// 		"zoho_auth_code": {
// 			defaultVal: "1000.463437b7a3238da8698cc4277a04caa9.6d0aaecd379431c6c5fbdae14970beb4",
// 			desc:       "zoho self-generated auth code",
// 		},
// 		"zoho_redirect_url": {
// 			defaultVal: "localhost:8765",
// 			desc:       "zoho redirect url",
// 		},
// 		"zoho_access_token": {
// 			defaultVal: "1000.463437b7a3238da8698cc4277a04caa9.6d0aaecd379431c6c5fbdae14970beb4",
// 			desc:       "zoho access token",
// 		},
// 		"zoho_refresh_token": {
// 			defaultVal: "1000.629874b7b8a7f4cb23754e7a07d0797e.3bd9a5f699b50e72ccedfc02b41f7abf",
// 			desc:       "zoho refresh token",
// 		},
// 	}

// 	for key, meta := range confList {
// 		// automatic conversion of environment var key to `UPPER_CASE` will happen.
// 		config.BindEnv(key)

// 		// read command-line arguments
// 		pflag.String(key, meta.defaultVal, meta.desc)
// 	}

// 	pflag.Parse()
// 	config.BindPFlags(pflag.CommandLine)
// 	return
// }
