package db

type HostIPMap map[string]string

type Host struct {
	Hostname string `gorm:"primaryKey;not null"`
	IP       string `gorm:"unique;not null"`
}

func DiffHostIPMaps(newHostIPMap HostIPMap, existingHostIPMap HostIPMap) ([]Host, []Host, []string) {
	toAddHosts := make([]Host, 0)
	toUpdateHosts := make([]Host, 0)
	toDeleteHosts := make([]string, 0)

	for hostname, ip := range newHostIPMap {
		if _, ok := existingHostIPMap[hostname]; !ok {
			toAddHosts = append(toAddHosts, Host{Hostname: hostname, IP: ip})
		} else if existingHostIPMap[hostname] != ip {
			toUpdateHosts = append(toUpdateHosts, Host{Hostname: hostname, IP: ip})
		}
	}

	for hostname := range existingHostIPMap {
		if _, ok := newHostIPMap[hostname]; !ok {
			toDeleteHosts = append(toDeleteHosts, hostname)
		}
	}

	return toAddHosts, toUpdateHosts, toDeleteHosts
}
