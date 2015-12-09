package conda

import (
  "testing"
)

func TestRemoveSingle(t *testing.T) {
  args := make(map[string]string)
  cmd := Remove(args, "numpy")

  cmd_ex := "/opt/anaconda/bin/conda remove numpy -y -q"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}

func TestRemoveMultiple(t *testing.T) {
  args := make(map[string]string)
  cmd := Remove(args, "numpy", "scipy", "scikit-learn")

  cmd_ex := "/opt/anaconda/bin/conda remove numpy scipy scikit-learn -y -q"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}

func TestRemoveWithName(t *testing.T) {
  args := make(map[string]string)
  args["name"] = "py35"
  cmd := Remove(args, "requests", "boto")

  cmd_ex := "/opt/anaconda/bin/conda remove requests boto -n py35 -y -q"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}

func TestRemoveWithPath(t *testing.T) {
  args := make(map[string]string)
  args["path"] = "/opt/anaconda/envs/py35"
  cmd := Remove(args, "pip", "boto")

  cmd_ex := "/opt/anaconda/bin/conda remove pip boto -p /opt/anaconda/envs/py35 -y -q"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}

func TestRemoveChannel(t *testing.T) {
  args := make(map[string]string)
  args["channels"] = "custom_channel"
  cmd := Remove(args, "requests", "boto")

  cmd_ex := "/opt/anaconda/bin/conda remove requests boto -c custom_channel -y -q"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}

func TestRemoveMultipleChannels(t *testing.T) {
  args := make(map[string]string)
  args["channels"] = "custom_channel,mychannel"
  cmd := Remove(args, "requests", "boto")

  cmd_ex := "/opt/anaconda/bin/conda remove requests boto -c custom_channel -c mychannel -y -q"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}

func TestRemoveOverriteChannels(t *testing.T) {
  args := make(map[string]string)
  args["channels"] = "custom_channel,mychannel"
  args["override-channels"] = "true"
  cmd := Remove(args, "requests", "boto")

  cmd_ex := "/opt/anaconda/bin/conda remove requests boto -c custom_channel -c mychannel --override-channels -y -q"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}

func TestRemoveDRY(t *testing.T) {
  args := make(map[string]string)
  args["name"] = "py35"
  args["dry"] = "true"
  cmd := Remove(args, "requests", "boto")

  cmd_ex := "/opt/anaconda/bin/conda remove requests boto -n py35 --dry-run -y -q"
  if cmd != cmd_ex {
    t.Logf("Commands '%s' != '%s'", cmd, cmd_ex)
    t.Fail()
  }
}
