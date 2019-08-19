package upnp

type logger struct{}

func (me logger) Infof(format string, args ...interface{})  {}
func (me logger) Debugf(format string, args ...interface{}) {}
func (me logger) Errorf(format string, args ...interface{}) {}
