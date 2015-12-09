package conda

import (
  "fmt"
  "strings"
)

// Build a valid `conda install` command
func Install(args map[string]string, packages ...string) string {
  conda := "/opt/anaconda/bin/conda"
  if val, ok := args["conda"]; ok {
    conda = val
  }
  cmd := fmt.Sprintf("%s %s", conda, "install")

  for _, pkg := range packages {
    cmd = cmd + fmt.Sprintf(" %s", pkg)
  }

  if val, ok := args["name"]; ok {
    if val != "" {
      cmd = cmd + fmt.Sprintf(" %s %s", "-n", val)
    }
  }

  if val, ok := args["path"]; ok {
    if val != "" {
      cmd = cmd + fmt.Sprintf(" %s %s", "-p", val)
    }
  }

  if val, ok := args["channels"]; ok {
    channels := strings.Split(val, ",")
    for _, channel := range channels {
      cmd = cmd + fmt.Sprintf(" -c %s", channel)
    }
  }

  if val, ok := args["override-channels"]; ok {
    if val == "true" {
      cmd = cmd + fmt.Sprintf(" %s", "--override-channels")
    }
  }

  if val, ok := args["dry"]; ok {
    if val == "true" {
      cmd = cmd + fmt.Sprintf(" %s", "--dry-run")
    }
  }

  if val, ok := args["copy"]; ok {
    if val == "true" {
      cmd = cmd + fmt.Sprintf(" %s", "--copy")
    }
  }

  cmd = cmd + fmt.Sprintf(" %s", "-y -q")
  return cmd
}

// Build a valid `conda install` command
func Remove(args map[string]string, packages ...string) string {
  conda := "/opt/anaconda/bin/conda"
  if val, ok := args["conda"]; ok {
    conda = val
  }
  cmd := fmt.Sprintf("%s %s", conda, "remove")

  for _, pkg := range packages {
    cmd = cmd + fmt.Sprintf(" %s", pkg)
  }

  if val, ok := args["name"]; ok {
    if val != "" {
      cmd = cmd + fmt.Sprintf(" %s %s", "-n", val)
    }
  }

  if val, ok := args["path"]; ok {
    if val != "" {
      cmd = cmd + fmt.Sprintf(" %s %s", "-p", val)
    }
  }

  if val, ok := args["channels"]; ok {
    channels := strings.Split(val, ",")
    for _, channel := range channels {
      cmd = cmd + fmt.Sprintf(" -c %s", channel)
    }
  }

  if val, ok := args["override-channels"]; ok {
    if val == "true" {
      cmd = cmd + fmt.Sprintf(" %s", "--override-channels")
    }
  }

  if val, ok := args["dry"]; ok {
    if val == "true" {
      cmd = cmd + fmt.Sprintf(" %s", "--dry-run")
    }
  }

  cmd = cmd + fmt.Sprintf(" %s", "-y -q")
  return cmd
}

// Build a valid `conda list` command
func List(args map[string]string) string {
  conda := "/opt/anaconda/bin/conda"
  if val, ok := args["conda"]; ok {
    conda = val
  }
  cmd := fmt.Sprintf("%s %s", conda, "list")

  if val, ok := args["name"]; ok {
    if val != "" {
      cmd = cmd + fmt.Sprintf(" %s %s", "-n", val)
    }
  }

  if val, ok := args["path"]; ok {
    if val != "" {
      cmd = cmd + fmt.Sprintf(" %s %s", "-p", val)
    }
  }

  if val, ok := args["no-pip"]; ok {
    if val == "true" {
      cmd = cmd + fmt.Sprintf(" %s", "--no-pip")
    }
  }

  if val, ok := args["canonical"]; ok {
    if val == "true" {
      cmd = cmd + fmt.Sprintf(" %s", "--canonical")
    }
  }

  if val, ok := args["explicit"]; ok {
    if val == "true" {
      cmd = cmd + fmt.Sprintf(" %s", "--explicit")
    }
  }

  return cmd
}
