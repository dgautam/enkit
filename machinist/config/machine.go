package config

import (
	"path/filepath"
)

type Node struct {
	Name          string
	Tags          []string
	SSHPrincipals []string
	IpAddresses   []string

	RequireRoot bool

	// BUG(INFRA-2550): Machinist can unpack files/scripts/config onto the host
	// machine, but this is better managed out-of-band by another tool, such as
	// Ansible or Puppet. If this bool is set, perform the legacy unpacking
	// behavior.
	ModifyMachineConfig bool

	LibNssConfLocation string

	// Pam Location configs
	// "/etc/security/pam_script_acct"
	PamSecurityLocation string
	PamSSHDLocation     string
	// SSHD Configs
	AutoRestartSSHD           bool
	CaPublicKeyLocation       string
	HostKeyLocation           string
	SSHDConfigurationLocation string
	ReWriteConfigs            bool

	*Common
}


// HostCertificate will return the path of the HostCertificate based on the path set by HostKeyLocation
// for example /foo/bar.pem will output /foo/bar-cert.pub
func (nf *Node) HostCertificate() string {
	b := filepath.Base(nf.HostKeyLocation)
	bExt := filepath.Ext(b)
	rawName := b[0 : len(b)-len(bExt)]
	d := filepath.Dir(nf.HostKeyLocation)
	return filepath.Join(d, rawName+"-cert.pub")
}
