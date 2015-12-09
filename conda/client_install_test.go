package conda

import (
  "testing"
)

func TestInstallSingle(t *testing.T) {
  args := make(map[string]string)
  cmd := Install(args, "numpy")

  cmd_ex := "/opt/anaconda/bin/conda install numpy -y -q"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}

func TestInstallMultiple(t *testing.T) {
  args := make(map[string]string)
  cmd := Install(args, "numpy", "scipy", "scikit-learn")

  cmd_ex := "/opt/anaconda/bin/conda install numpy scipy scikit-learn -y -q"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}

func TestInstallWithName(t *testing.T) {
  args := make(map[string]string)
  args["name"] = "py35"
  cmd := Install(args, "requests", "boto")

  cmd_ex := "/opt/anaconda/bin/conda install requests boto -n py35 -y -q"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}

func TestInstallWithPath(t *testing.T) {
  args := make(map[string]string)
  args["path"] = "/opt/anaconda/envs/py35"
  cmd := Install(args, "pip", "boto")

  cmd_ex := "/opt/anaconda/bin/conda install pip boto -p /opt/anaconda/envs/py35 -y -q"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}

func TestInstallChannel(t *testing.T) {
  args := make(map[string]string)
  args["channels"] = "custom_channel"
  cmd := Install(args, "requests", "boto")

  cmd_ex := "/opt/anaconda/bin/conda install requests boto -c custom_channel -y -q"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}

func TestInstallMultipleChannels(t *testing.T) {
  args := make(map[string]string)
  args["channels"] = "custom_channel,mychannel"
  cmd := Install(args, "requests", "boto")

  cmd_ex := "/opt/anaconda/bin/conda install requests boto -c custom_channel -c mychannel -y -q"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}

func TestInstallOverriteChannels(t *testing.T) {
  args := make(map[string]string)
  args["channels"] = "custom_channel,mychannel"
  args["override-channels"] = "true"
  cmd := Install(args, "requests", "boto")

  cmd_ex := "/opt/anaconda/bin/conda install requests boto -c custom_channel -c mychannel --override-channels -y -q"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}

func TestInstallDRY(t *testing.T) {
  args := make(map[string]string)
  args["name"] = "py35"
  args["dry"] = "true"
  cmd := Install(args, "requests", "boto")

  cmd_ex := "/opt/anaconda/bin/conda install requests boto -n py35 --dry-run -y -q"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}

func TestInstallCopy(t *testing.T) {
  args := make(map[string]string)
  args["name"] = "py35"
  args["copy"] = "true"
  cmd := Install(args, "requests", "boto")

  cmd_ex := "/opt/anaconda/bin/conda install requests boto -n py35 --copy -y -q"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}
