package satisphp

import (
    "log"
    "os/exec"
    "satis-go/satis/satisphp/db"
)

var _ = log.Print

type Generator interface {
    Generate() error
    UpdateRepo(repo string) error
}

type StaticWebGenerator struct {
    DbPath  string
    WebPath string
}

func (s *StaticWebGenerator) UpdateRepo(repo string) error {

    log.Print("UpdateRepo Generating...")

    out, err := exec.
    Command("satis", "--no-interaction", "build", "--repository-url", repo, s.DbPath+db.StagingFile, s.WebPath).
        CombinedOutput()

    if err != nil {
        log.Printf("Satis Generation Error: %s", string(out[:]))
    }

    return err
}

func (s *StaticWebGenerator) Generate() error {
    log.Print("Generating...")

    out, err := exec.
    Command("satis", "--no-interaction", "build", s.DbPath+db.StagingFile, s.WebPath).
        CombinedOutput()
    if err != nil {
        log.Printf("Satis Generation Error: %s", string(out[:]))
    }
    return err
}
