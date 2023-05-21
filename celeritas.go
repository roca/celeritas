package celeritas

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const version = "1.0.0"

type Celeritas struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RooPath  string
	config   config
}

type config struct {
	port     string
	renderer string // The rendering engine Go or Jet
}

func (c *Celeritas) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath: rootPath,
		folderNames: []string{
			"handlers",
			"migrations",
			"views",
			"data",
			"public",
			"tmp",
			"logs",
			"middleware",
		},
	}
	err := c.Init(pathConfig)
	if err != nil {
		return err
	}

	err = c.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	// read .env file
	err = godotenv.Load(fmt.Sprintf("%s/.env", rootPath))
	if err != nil {
		return err
	}

	// create loggers
	infoLog, errorLog := c.startLogggers()
	c.InfoLog = infoLog
	c.ErrorLog = errorLog
	c.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	c.Version = version
	c.RooPath = rootPath

	c.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
	}

	return nil
}

func (c *Celeritas) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		// create folder if it doesn't exist
		err := c.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Celeritas) checkDotEnv(path string) error {
	err := c.CreateFileIfNotExist(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}
	return nil
}

func (c *Celeritas) startLogggers() (*log.Logger, *log.Logger) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return infoLog, errorLog
}
