# golog
golog package provide log rotate 
This is a wrapper of logrus and lumberjack. It provide loglevel, output, json formwat and more import logrotate. 

Example: 

    Option 1. Simplely import golog and write log with golog.Infoln("This is golog package"). 
    Option 2. Create log instance with NewLogger. 
    	          logger = golog.NewLogger()
	              logger.SetLevel(log.InfoLevel)
	              logger.SetFormatter("json")
	              logger.SetOutPutAndRotate(log.LogRotate{
                    Filename:  <filename>,
                    MaxSize:   <Megabyte>,
                    MaxBackup: <Copies>,
                    MaxAge:    <Age>,
                    DupToStd:  <STD>,
                })
