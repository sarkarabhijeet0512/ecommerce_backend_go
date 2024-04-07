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
