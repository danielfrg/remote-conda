package conda

import (
  "testing"
)

func TestList1(t *testing.T) {
  args := make(map[string]string)
  cmd := List(args)

  cmd_ex := "/opt/anaconda/bin/conda list"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}

func TestList2(t *testing.T) {
  args := make(map[string]string)
  args["name"] = "env1"
  cmd := List(args)

  cmd_ex := "/opt/anaconda/bin/conda list -n env1"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}

func TestList3(t *testing.T) {
  args := make(map[string]string)
  args["path"] = "/opt/myenv"
  cmd := List(args)

  cmd_ex := "/opt/anaconda/bin/conda list -p /opt/myenv"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}

func TestList5(t *testing.T) {
  args := make(map[string]string)
  args["no-pip"] = "true"
  cmd := List(args)

  cmd_ex := "/opt/anaconda/bin/conda list --no-pip"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}

func TestList6(t *testing.T) {
  args := make(map[string]string)
  args["canonical"] = "true"
  cmd := List(args)

  cmd_ex := "/opt/anaconda/bin/conda list --canonical"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}

func TestList7(t *testing.T) {
  args := make(map[string]string)
  args["explicit"] = "true"
  cmd := List(args)

  cmd_ex := "/opt/anaconda/bin/conda list --explicit"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}
