# golog
golog provides loglevel setup, output setup, json/txt formwat and more important logrotate!

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
