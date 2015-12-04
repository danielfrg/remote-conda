package main

import "os"
import "fmt"
import "time"
// import "sort"

import "golang.org/x/crypto/ssh"
import "github.com/codegangsta/cli"
import "github.com/danielfrg/remote-conda/ssh"

func list(c *cli.Context) {
	condaPath := c.String("conda")
	cmd := fmt.Sprintf("%s %s", condaPath, "list")

	clients := makeClients(cmd, c)
	results := make(chan bool, len(clients))

	execClients(clients, results)
	waitResults(clients, results, c)
	printResults(clients, c)
}

func install(c *cli.Context) {
	condaPath := c.String("conda")
	cmd := fmt.Sprintf("%s %s %s", condaPath, "install", "-y -q")

	args := c.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Error: No packages provided")
    os.Exit(1)
	}

	for _, arg := range args {
		cmd = cmd + " " + arg
	}

	clients := makeClients(cmd, c)
	results := make(chan bool, len(clients))

	execClients(clients, results)
	waitResults(clients, results, c)
	printResults(clients, c)
}

func remove(c *cli.Context) {
	condaPath := c.String("conda")
	cmd := fmt.Sprintf("%s %s %s", condaPath, "remove", "-y -q")

	args := c.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Error: No packages provided")
    os.Exit(1)
	}

	for _, arg := range args {
		cmd = cmd + " " + arg
	}

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
  app.Version = "0.1.0"

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
		cli.StringFlag{
	    Name: "conda, p",
	    Value: "/opt/anaconda/bin/conda",
	    Usage: "Conda path",
	  },
		cli.IntFlag{
	    Name: "timeout, t",
	    Value: 300,
	    Usage: "Timeout",
	  },
	}

	app.Commands = []cli.Command{
		{
			Name:     "list",
			Usage:    "List conda packages",
			Action:   list,
			Flags:    defaultFlags,
		},
		{
			Name:     "install",
			Usage:    "Install conda packages",
			Action:   install,
			Flags:    defaultFlags,
		},
		{
			Name:     "remove",
			Usage:    "Remove conda packages",
			Action:   remove,
			Flags:    defaultFlags,
		},
	}

  app.Run(os.Args)
}
