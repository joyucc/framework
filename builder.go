package framework

import "github.com/jinzhu/gorm"

type Framework struct {
	Orm *gorm.DB
}

type BuildProcess interface {
	SetLog() BuildProcess
	SetOrm() BuildProcess
	SetTrace() BuildProcess
	SetStat() BuildProcess
	GetFramework() Framework
}

type FMDirector struct {
	builder BuildProcess
}

func (m *FMDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}

func (m *FMDirector) BuildWeb() Framework {
	m.builder.SetOrm()
	return m.builder.GetFramework()
}

func (m *FMDirector) BuildSrv() Framework {
	m.builder.SetOrm()
	return m.builder.GetFramework()
}
