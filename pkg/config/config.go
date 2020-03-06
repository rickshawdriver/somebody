package config

import (
	"github.com/BurntSushi/toml"
	"github.com/rickshawdriver/somebody/pkg/log"
	"os"
	"sync"
	"sync/atomic"
)

var (
	confLog = log.Get().WithField("prifix", "config")

	global   atomic.Value
	globalMu sync.Mutex

	defaultConfigPath = []string{
		"../somebody.toml",
		"somebody.toml",
		"/etc/rickshawdriver/somebody.toml",
	}
)

type FilePath struct {
	OriginalPath    string `toml:"original_path"`
	PidFileLocation string `toml:"pid_file_location"`
}

type HttpConf struct {
	Addr string `toml:"addr"`
	Port int    `toml:"port"`
}

type Config struct {
	Http     HttpConf `toml:"http"`
	Log      log.Log
	FilePath FilePath `toml:"filepath"`
}

// load my config file
func Load(config *Config) error {
	for _, path := range defaultConfigPath {
		_, err := os.Open(path)
		if err == nil { // success
			config.FilePath.OriginalPath = path
			break
		}
		if os.IsNotExist(err) {
			continue
		}

		return err
	}

	// if file not exist
	if config.FilePath.OriginalPath == "" && nil != createDefaultFile() {
		//create default file
		confLog.Warnf("not found config file, creating")
	}

	// analyse config file
	if _, err := toml.DecodeFile(config.FilePath.OriginalPath, &config); err != nil {
		return err
	}

	return nil
}

func createDefaultFile() error {
	confLog.Println("create default config file")
	f, err := os.Create("somebody.toml")
	defer f.Close()
	if err != nil {
		return err
	}

	_, err = f.Write([]byte(`log_level = "debug"
pid_file_location = ""
[http]
addr = "127.0.0.1"
port = 8080
https_port = 443
default_tls_cert  = "MIIDYDCCAkigAwIBAgIJALSbF0IPueo1MA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwHhcNMTkwOTI3MDcyNzU0WhcNMjAwOTI2MDcyNzU0WjBFMQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxP9CBRyY8EFcsPtNn8CwnHGIsXR3pESLQG0BWvZIpFIO3Cz2WDT7AI0hiBdAC3gOgZag3Jr76aKVNPVG3DhvCTp+n6xK4vNr8TFF4ZVfS95ZVXr6NuABbAm5hOo5/k7UhPu+9qd94UlT/tng2uaIZbyqykJSta5M9C6cmmEpe9POXZn+sFoRbXv21MmLBc5U1/v0PlCDF2d9u1n/ffwI2wtr3IWBU/oPKeaRzuDCXpkwZlAu9NYT3Bznl4cj5LYOtIgfFKARfcnSo2lJwk6/w/69oLzfWTx9DUn4NXX4+z7Qp2bOeao0CmvhOfCUQ1Snx/Mw6zOpV0sr5BZ/SCEH3QIDAQABo1MwUTAdBgNVHQ4EFgQUFRFA5b+OQLi69kDsiBAm045f2xIwHwYDVR0jBBgwFoAUFRFA5b+OQLi69kDsiBAm045f2xIwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAON9DgA1pj0fn3uvHgzDrQwFQIa/afZDScVISSmhiWgMSja8i9yvXIgJX3hKip193Jvd+ekNKlpJOT43Ydwo7oAsp7MFfBYpjNtgRYPYumRmA8SVIMTRnKJ96aPfGgyk3ZJGJswVbvlHMGNAbBZ6XtbNO2+XuuViRCxJbDnQoWkzPbIBOhw5KNn6fXbtsd9aP9UJQ04QGydr60MFaUgQgBIgYlR+XioWxo57V4KrpaTz/ao89kd9Ca3ZlNDrgy4uwYA3gZZtAK0EXMvXsOH3hEFswhojjUWgRkHLxHD+tNKBX2dV7eoA3w9ofTl3MbzaO1ncD7oab6bdypkemBgLuMQ=="
default_tls_key = "MIIEowIBAAKCAQEAxP9CBRyY8EFcsPtNn8CwnHGIsXR3pESLQG0BWvZIpFIO3Cz2WDT7AI0hiBdAC3gOgZag3Jr76aKVNPVG3DhvCTp+n6xK4vNr8TFF4ZVfS95ZVXr6NuABbAm5hOo5/k7UhPu+9qd94UlT/tng2uaIZbyqykJSta5M9C6cmmEpe9POXZn+sFoRbXv21MmLBc5U1/v0PlCDF2d9u1n/ffwI2wtr3IWBU/oPKeaRzuDCXpkwZlAu9NYT3Bznl4cj5LYOtIgfFKARfcnSo2lJwk6/w/69oLzfWTx9DUn4NXX4+z7Qp2bOeao0CmvhOfCUQ1Snx/Mw6zOpV0sr5BZ/SCEH3QIDAQABAoIBAHXTLoucJSVeErCQPkdUms2XYmiw/nYzwQ4RpIPhVmVh5x1tjxIG7jHQN8QME/RIJHUBwMjxscZ1xcRdB7rjzhW49M9P36KKcX9bNy7Lhqn8HXZxDYMQzAjwcBO9fF5Mi/PWFsu0NigvEZwWeNS0mdQv2f8LWCjuTXym/GehwYwbgFAHhyAQCkW0cHG0SEn0wJzeawRQ0MXQIn9v6XfLcrvSPYCHszdcC8i8N1//VdZVzlawJWL0pI8N2vEDaHQyOyhF6mw/v7FxJvxG0oMRu8XOHPUZT1enKN6ENUQLMlVaB8Es7aqvAsr6kWfO+srAmFjoGp5jz9oKgoQHHYhb0UkCgYEA5s8l4F6R/RgGMtsJg7n0tB8PfONzXJIGfCAzbyv14x88Pr4P1CLtfWvWXlNkZU/4gayejYEQMIvV2xITJkK/qDPmmdm8bPUnWteTigtpsACKXGda8JSVxZ1NUnMenbF0MQgWPgn+goQdF0pyPOOqB2axbu5h08GY29hmPOqqLcCgYEA2n9k7teX0EUdbU8a1bY4Jyu+ed7/weNacnoF49plsRHQwtsADkzCGJNCbGDpzA8qYGctweCyJ5O2n6UPglS8cvJ4q920X0I9FCnyd4JOVb0y3pI6UCa7k/KKLe9bB2pivCC9coRqxVUousLVRqF0FJSsOhUkOSp2XMHQOmL6eAsCgYEAmVGVKXvooiUpSPLzWQtXn6baVO2KHj3qDN4fDIQ5LAv9qvf578Lb22qc1b+zexEqVIlkMT9Aj97zyjgxfELuqT8AucZHQmF/KPr2yMZYCE/zmPoXEKTEpWUHqsRBDkCDmvANRMxjUdcJoyKXtSME6mvv+dV/ooA+ECNWw5FpFfUCgYAsWFCu3Ni172ESj0x99VaeNJWa/HRh8Hep5jQN04RyFJCPC42OWWvHOxPvFRg+TxGlsSsrPyRJPgSBkCl+pR3+IlH1Z7C06Kem0QCh2rN1WUnavEjTfZjyZPZAbmTGM4RIdEie1lspI6h5hxNsG1aI2se9ng7U/1Y4aymMwAvfkwKBgF+WmgCFDNIQLxvQ51kJahX1vBsWLd/koJDNjn4WJghnfwtgCWwtYeX6rpm0GcYU9F6rvGBjwfm9KUSDW9W8YVpBzQF8Rozbl7FQ9WSP/77IW22FPwK+TJXPH49dM3L/g5ToYzjPDbA6j/I42fN4YI7qMnkpFp5g2AV4IWCyPlp0"
limit_buffer_read = 2048
limit_buffer_write = 1024
limit_bytes_body = 10485760
[health_check]
enable_health_checks = true
health_check_timeout = 5
`))

	return nil
}

func SetGlobal(conf Config) {
	globalMu.Lock()
	defer globalMu.Unlock()
	global.Store(conf)
}

func Global() Config {
	return global.Load().(Config)
}
