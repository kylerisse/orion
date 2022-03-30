package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gliderlabs/ssh"
	"gitlab.com/nebulaworks/eng/bazaar/prj/term-apply/pkg/applicant"
)

type TeaManager struct {
	Appmgr *applicant.ApplicantManager
}

func NewTeaManager(applicantManager *applicant.ApplicantManager) *TeaManager {
	return &TeaManager{
		Appmgr: applicantManager,
	}
}

func (t *TeaManager) TeaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	_, _, active := s.Pty()
	if !active {
		fmt.Println("no active terminal, skipping")
		return nil, nil
	}
	m := InitialModel(t.Appmgr, s.User())
	return m, []tea.ProgramOption{tea.WithAltScreen()}
}
