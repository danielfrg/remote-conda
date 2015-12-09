package main

import (
  "os"
  "fmt"
  "time"
  "strings"

  "golang.org/x/crypto/ssh"
  "github.com/codegangsta/cli"
  "github.com/danielfrg/remote-conda/conda"
  "github.com/danielfrg/remote-conda/ssh"
)

func install(c *cli.Context) {
	cliArgs := c.Args()
	if len(cliArgs) == 0 {
		fmt.Fprintln(os.Stderr, "Error: too few arguments, must supply package specs")
    os.Exit(1)
	}

	args := make(map[string]string)
	args["conda"] = c.String("conda")
	args["name"] = c.String("name")
	args["path"] = c.String("path")
	args["channels"] = strings.Join(c.StringSlice("channel"), ",")
	args["override-channels"] = c.String("override-channels")
	args["dry"] = c.String("dry")
	args["copy"] = c.String("copy")
	cmd := conda.Install(args, cliArgs...)
	fmt.Println("Executing command:", cmd)

	clients := makeClients(cmd, c)
	results := make(chan bool, len(clients))

	execClients(clients, results)
	waitResults(clients, results, c)
	printResults(clients, c)
}

func remove(c *cli.Context) {
	cliArgs := c.Args()
	if len(cliArgs) == 0 {
		fmt.Fprintln(os.Stderr, "Error: too few arguments, must supply package specs")
    os.Exit(1)
	}

	args := make(map[string]string)
	args["conda"] = c.String("conda")
	args["name"] = c.String("name")
	args["path"] = c.String("path")
	args["channels"] = strings.Join(c.StringSlice("channel"), ",")
	args["override-channels"] = c.String("override-channels")
	args["dry"] = c.String("dry")
	cmd := conda.Remove(args)
	fmt.Println("Executing command:", cmd)

	clients := makeClients(cmd, c)
	results := make(chan bool, len(clients))

	execClients(clients, results)
	waitResults(clients, results, c)
	printResults(clients, c)
}

func list(c *cli.Context) {
	args := make(map[string]string)
	args["conda"] = c.String("conda")
	args["name"] = c.String("name")
	args["path"] = c.String("path")
	args["no-pip"] = c.String("no-pip")
	args["canonical"] = c.String("canonical")
	args["explicit"] = c.String("explicit")
	cmd := conda.List(args)
	fmt.Println("Executing command:", cmd)

	clients := makeClients(cmd, c)
	results := make(chan bool, len(clients))

	execClients(clients, results)
	waitResults(clients, results, c)
	printResults(clients, c)
}

// Get an Slice of SSHClient.SSHClient based on cli args
func makeClients(cmd string, c *cli.Context) []*SSHClient.SSHClient {
	if c.String("user") == "" {
		fmt.Fprintln(os.Stderr, "Error: No --user/-u provided")
    os.Exit(1)
	}

	if c.String("pkey") == "" {
		fmt.Fprintln(os.Stderr, "Error: No --pkey/-k provided")
    os.Exit(1)
	}

	sshConfig := &ssh.ClientConfig{
		User: c.String("user"),
		Auth: []ssh.AuthMethod{
			SSHClient.PublicKeyFile(c.String("pkey")),
		},
	}

	hosts := deleteEmpty(c.StringSlice("host"))
	if len(hosts) == 0 {
		fmt.Fprintln(os.Stderr, "Error: No --hosts/-x provided")
    os.Exit(1)
	}

	clients := make([]*SSHClient.SSHClient, len(hosts))
	for i, host := range hosts {
		_cmd :=  &SSHClient.SSHCommand{
			Cmd: cmd,
		}
		client := &SSHClient.SSHClient{
			Config: sshConfig,
			Host:   host,
			Port:   22,
			Cmd:    _cmd,
		}
		clients[i] = client
	}

	return clients
}

// Execute Client.Cmd
func execClients(clients []*SSHClient.SSHClient, results chan bool) {
	for _, client := range clients {
		go func(client *SSHClient.SSHClient) {
			results <- SSHClient.ExecuteCmd(client)
		}(client)
	}
}

// Fill the commands with the results from the channel
func waitResults(clients []*SSHClient.SSHClient, results chan bool, c *cli.Context) {
	timeout := time.After(time.Duration(c.Int("timeout")) * time.Second)

	for i := 0; i < len(clients); i++ {
		select {
		case res := <-results:
			res = !res
		case <-timeout:
			fmt.Println("Timed out! Commands might be still running on the hosts")
			os.Exit(1)
		}
	}
}

// Print the results
func printResults(clients []*SSHClient.SSHClient, c *cli.Context) {
	stdout := func(client *SSHClient.SSHClient) string {
		return client.Cmd.Stdout
	}

	groups := groupBy(clients, stdout)

	for k, _ := range groups {
		fmt.Printf("Response from %d node(s):\n", len(groups[k]))
		fmt.Println(k)
	}
}

func groupBy(clients []*SSHClient.SSHClient, f func(*SSHClient.SSHClient) string) map[string][]string{
	ret := make(map[string][]string)
	for _, client := range clients {
		key := client.Cmd.Stdout
		ret[key] = append(ret[key], client.Host)
	}
	return ret
}

// Delete Empty strings from a Slice or Array
func deleteEmpty(s []string) []string {
    var r []string
    for _, str := range s {
        if str != "" {
            r = append(r, str)
        }
    }
    return r
}

func main() {
	app := cli.NewApp()
  app.Name = "Remote Conda"
  app.Usage = "Install conda packages in remote hosts"
  app.Version = "1.0.0rc1"

	defaultFlags := []cli.Flag {
	  cli.StringSliceFlag{
	    Name: "host, x",
	    Usage: "Host",
	  },
		cli.StringFlag{
	    Name: "user, u",
	    Usage: "Username",
	  },
		cli.StringFlag{
	    Name: "pkey, k",
	    Usage: "Private Key",
	  },
		cli.IntFlag{
	    Name: "timeout, t",
	    Value: 300,
	    Usage: "Timeout",
	  },
		cli.StringFlag{
	    Name: "conda",
	    Value: "/opt/anaconda/bin/conda",
	    Usage: "Conda path",
	  },
		cli.StringFlag{
	    Name: "name, n",
	    Usage: "Name of the environment",
	  },
		cli.StringFlag{
	    Name: "path, p",
	    Usage: "Full path to environment prefix",
	  },
	}

	channelsFlags := []cli.Flag {
		cli.StringSliceFlag{
			Name: "channel, c",
			Usage: "Additional channel to search for packages",
		},
		cli.BoolFlag{
	    Name: "override-channels",
	    Usage: "Do not search default or .condarc channels. Requires --channel.",
	  },
	}

	dryFlag := cli.BoolFlag{
		Name: "dry-run",
		Usage: "Only display what would have been done",
	}

	copyFlag := cli.BoolFlag{
		Name: "copy",
		Usage: "Install all packages using copies instead of hard- or soft-linking",
	}

	installFlags := append(defaultFlags, channelsFlags...)
	installFlags = append(installFlags, dryFlag)
	installFlags = append(installFlags, copyFlag)

	removeFlags := append(defaultFlags, channelsFlags...)
	removeFlags = append(defaultFlags, dryFlag)

	listFlags := defaultFlags

	nopipFlag := cli.BoolFlag{
		Name: "no-pip",
		Usage: "Do not include pip-only installed packages",
	}
	listFlags = append(listFlags, nopipFlag)

	canonicalFlag := cli.BoolFlag{
		Name: "canonical",
		Usage: "Output canonical names of packages only. Implies --no-pip",
	}
	listFlags = append(listFlags, canonicalFlag)

	explicitFlag := cli.BoolFlag{
		Name: "explicit",
		Usage: "List explicitly all installed conda packaged with URL",
	}
	listFlags = append(listFlags, explicitFlag)

	app.Commands = []cli.Command{
		{
			Name:     "install",
			Usage:    "Install conda packages",
			Action:   install,
			Flags:    installFlags,
		},
		{
			Name:     "remove",
			Usage:    "Remove conda packages",
			Action:   remove,
			Flags:    removeFlags,
		},
		{
			Name:     "list",
			Usage:    "List conda packages",
			Action:   list,
			Flags:    listFlags,
		},
	}

  app.Run(os.Args)
}
