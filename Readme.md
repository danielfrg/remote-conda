# Remote-conda

Install conda packages in remote hosts

## Installation



## Usage

`remote-conda` is a basic CLI that has the basic conda commands: `install`, `remove`, `list` and follows the same API as conda with some extra options:

Host/IP (--host/-x). Can be used multiple times
User (--user/-u) is the username to ssh into the hosts
Private key (--pkey/-k) to ssh into the hosts
Pip Path (--conda/-p), useful to manage packages in multiple virtual environments. Default: `/opt/anaconda/bin/conda`

## Examples:

```
$ remote-conda install numpy -u ubuntu -k ~/.ssh/key.pem -x myhost1
$ remote-conda remove numpy -u ubuntu -k ~/.ssh/key.pem -x myhost1 -x myhost2

# Another conda location
$ remote-conda list -u ubuntu -k ~/.ssh/key.pem -x myhost1 -p /user/ubuntu/anaconda/bin/conda
```
