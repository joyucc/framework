package framework

type Setting struct {
	framework Framework
}

func (l *Setting) SetLog() BuildProcess {
	l.framework.Orm = nil
	return l
}

func (l *Setting) SetOrm() BuildProcess {
	l.framework.Orm = nil
	return l
}

func (l *Setting) SetTrace() BuildProcess {
	l.framework.Orm = nil
	return l
}

func (l *Setting) SetStat() BuildProcess {
	l.framework.Orm = nil
	return l
}

func (l *Setting) GetFramework() Framework {

	return l.framework
}
