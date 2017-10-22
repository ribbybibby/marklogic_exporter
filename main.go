package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
	digest "github.com/xinsnake/go-http-digest-auth-client"

	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/ini.v1"
)

const (
	namespace = "marklogic"
)

var (
	// Disk metrics
	diskRead = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "disk", "read_bytes"),
		"Placeholder",
		[]string{"type", "host", "database", "forest"}, nil,
	)
	diskWrite = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "disk", "write_bytes"),
		"Placeholder",
		[]string{"type", "host", "database", "forest"}, nil,
	)
	diskSize = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "disk", "size_bytes"),
		"Placeholder",
		[]string{"host", "database", "forest", "stand_id", "encrypted", "fast"}, nil,
	)
	diskForestReserveSize = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "disk", "forest_reserve_size_bytes"),
		"Placeholder",
		[]string{"host", "database", "forest"}, nil,
	)
	diskJournalsSize = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "disk", "journals_size_bytes"),
		"Placeholder",
		[]string{"host", "database", "forest"}, nil,
	)
	diskMinCapacity = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "disk", "min_capacity_percent"),
		"Placeholder",
		[]string{"host"}, nil,
	)

	// CPU metrics
	cpuUsage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "cpu", "usage_percent"),
		"Placeholder",
		[]string{"type", "host"}, nil,
	)

	// Memory metrics
	memorySize = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "memory", "size_mb"),
		"Placeholder",
		[]string{"host", "database", "forest", "stand"}, nil,
	)
	memoryFootprint = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "memory", "footprint_bytes"),
		"Placeholder",
		[]string{"type", "host"}, nil,
	)
	memoryPagesIn = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "memory", "pages_in_per_second"),
		"Placeholder",
		[]string{"host"}, nil,
	)
	memoryPagesOut = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "memory", "pages_out_per_second"),
		"Placeholder",
		[]string{"host"}, nil,
	)
	memorySwapIn = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "memory", "swap_in_per_second"),
		"Placeholder",
		[]string{"host"}, nil,
	)
	memorySwapOut = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "memory", "swap_out_per_second"),
		"Placeholder",
		[]string{"host"}, nil,
	)

	// Database metrics
	databaseListCacheHits = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "list_cache_hits"),
		"Placeholder",
		[]string{"host", "database", "forest", "stand_id"}, nil,
	)
	databaseListCacheMisses = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "list_cache_misses"),
		"Placeholder",
		[]string{"host", "database", "forest", "stand_id"}, nil,
	)
	databaseCompressedTreeCacheHits = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "compressed_tree_cache_hits"),
		"Placeholder",
		[]string{"host", "database", "forest", "stand_id"}, nil,
	)
	databaseCompressedTreeCacheMisses = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "compressed_tree_cache_misses"),
		"Placeholder",
		[]string{"host", "database", "forest", "stand_id"}, nil,
	)
	databaseTripleCacheHits = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "triple_cache_hits"),
		"Placeholder",
		[]string{"host", "database", "forest", "stand_id"}, nil,
	)
	databaseTripleCacheMisses = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "triple_cache_misses"),
		"Placeholder",
		[]string{"host", "database", "forest", "stand_id"}, nil,
	)
	databaseTripleValueCacheHits = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "triple_value_cache_hits"),
		"Placeholder",
		[]string{"host", "database", "forest", "stand_id"}, nil,
	)
	databaseTripleValueCacheMisses = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "triple_value_cache_misses"),
		"Placeholder",
		[]string{"host", "database", "forest", "stand_id"}, nil,
	)
	databaseReadLocks = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "read_locks"),
		"Placeholder",
		[]string{"host", "database", "forest"}, nil,
	)
	databaseReadLockWaitTime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "read_lock_wait_time_seconds"),
		"Placeholder",
		[]string{"host", "database", "forest"}, nil,
	)
	databaseReadLockHoldTime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "read_lock_hold_time_seconds"),
		"Placeholder",
		[]string{"host", "database", "forest"}, nil,
	)
	databaseWriteLocks = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "write_locks"),
		"Placeholder",
		[]string{"host", "database", "forest"}, nil,
	)
	databaseWriteLockWaitTime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "write_lock_wait_time_seconds"),
		"Placeholder",
		[]string{"host", "database", "forest"}, nil,
	)
	databaseWriteLockHoldTime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "write_lock_hold_time_seconds"),
		"Placeholder",
		[]string{"host", "database", "forest"}, nil,
	)
	databaseDeadLocks = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "deadlocks"),
		"Placeholder",
		[]string{"host", "database", "forest"}, nil,
	)
	databaseDeadLockWaitTime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "dead_lock_wait_time_seconds"),
		"Placeholder",
		[]string{"host", "database", "forest"}, nil,
	)
	databaseReplicationReceive = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "replication_receive_bytes"),
		"Placeholder",
		[]string{"host", "database", "forest"}, nil,
	)
	databaseReplicationSend = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "replication_send_bytes"),
		"Placeholder",
		[]string{"host", "database", "forest"}, nil,
	)

	// Server metrics
	serverRequestRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "server", "request_rate_per_second"),
		"Placeholder",
		[]string{"host", "database", "server", "group_id"}, nil,
	)
	serverReceive = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "server", "receive_bytes"),
		"Placeholder",
		[]string{"host", "database", "server", "group_id"}, nil,
	)
	serverSend = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "server", "send_bytes"),
		"Placeholder",
		[]string{"host", "database", "server", "group_id"}, nil,
	)
	serverExpandedTreeCacheHits = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "server", "expanded_tree_cache_hits"),
		"Placeholder",
		[]string{"host", "database", "server", "group_id"}, nil,
	)
	serverExpandedTreeCacheMisses = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "server", "expanded_tree_cache_misses"),
		"Placeholder",
		[]string{"host", "database", "server", "group_id"}, nil,
	)

	// Network metrics
	networkXDQPClientReceive = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "network", "xdqp_client_receive_bytes"),
		"Placeholder",
		[]string{"host"}, nil,
	)
	networkXDQPClientSend = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "network", "xdqp_client_send_bytes"),
		"Placeholder",
		[]string{"host"}, nil,
	)
	networkXDQPServerReceive = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "network", "xdqp_server_receive_bytes"),
		"Placeholder",
		[]string{"host"}, nil,
	)
	networkXDQPServerSend = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "network", "xdqp_server_send_bytes"),
		"Placeholder",
		[]string{"host"}, nil,
	)
	networkXDQPForeignClientReceive = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "network", "xdqp_foreign_client_receive_bytes"),
		"Placeholder",
		[]string{"host"}, nil,
	)
	networkXDQPForeignClientSend = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "network", "xdqp_foreign_client_send_bytes"),
		"Placeholder",
		[]string{"host"}, nil,
	)
	networkXDQPForeignServerReceive = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "network", "xdqp_foreign_server_receive_bytes"),
		"Placeholder",
		[]string{"host"}, nil,
	)
	networkXDQPForeignServerSend = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "network", "xdqp_foreign_server_send_bytes"),
		"Placeholder",
		[]string{"host"}, nil,
	)
)

// Exporter is our custom collector type
type Exporter struct {
	username string
	password string
	uri      string
}

// Describe implements the prometheus.Collector interface.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- diskRead
	ch <- diskWrite
	ch <- diskSize
	ch <- diskForestReserveSize
	ch <- diskJournalsSize
	ch <- diskMinCapacity
	ch <- cpuUsage
	ch <- memorySize
	ch <- memoryFootprint
	ch <- memoryPagesIn
	ch <- memoryPagesOut
	ch <- memorySwapIn
	ch <- memorySwapOut
	ch <- databaseListCacheHits
	ch <- databaseListCacheMisses
	ch <- databaseCompressedTreeCacheHits
	ch <- databaseCompressedTreeCacheMisses
	ch <- databaseTripleCacheHits
	ch <- databaseTripleCacheMisses
	ch <- databaseTripleValueCacheHits
	ch <- databaseTripleValueCacheMisses
	ch <- databaseReadLocks
	ch <- databaseReadLockWaitTime
	ch <- databaseReadLockHoldTime
	ch <- databaseWriteLocks
	ch <- databaseWriteLockWaitTime
	ch <- databaseWriteLockHoldTime
	ch <- databaseDeadLocks
	ch <- databaseDeadLockWaitTime
	ch <- databaseReplicationReceive
	ch <- databaseReplicationSend
	ch <- serverRequestRate
	ch <- serverReceive
	ch <- serverSend
	ch <- serverExpandedTreeCacheHits
	ch <- serverExpandedTreeCacheMisses
	ch <- networkXDQPClientReceive
	ch <- networkXDQPClientSend
	ch <- networkXDQPServerReceive
	ch <- networkXDQPServerSend
	ch <- networkXDQPForeignClientReceive
	ch <- networkXDQPForeignClientSend
	ch <- networkXDQPForeignServerReceive
	ch <- networkXDQPForeignServerSend
}

// Collect implements the prometheus.Collector interface.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	wg := sync.WaitGroup{}

	hosts, err := e.GetHosts()
	if err != nil {
		log.Errorln(err)
		return
	}

	groups, err := e.GetGroups()
	if err != nil {
		log.Errorln(err)
		return
	}

	forests, err := e.GetForests(hosts)
	if err != nil {
		log.Errorln(err)
		return
	}

	servers, err := e.GetServers(groups)
	if err != nil {
		log.Errorln(err)
		return
	}

	for _, hostName := range hosts {
		wg.Add(1)
		go func(hostName string) {
			err := e.GetHostMetrics(ch, hostName)
			if err != nil {
				log.Errorln(err)
				return
			}
			wg.Done()
		}(hostName)
	}

	for hostName, forestList := range forests {
		for _, forestName := range forestList {
			wg.Add(1)
			go func(forestName string, hostName string) {
				err := e.GetForestMetrics(ch, forestName, hostName)
				if err != nil {
					log.Errorln(err)
					return
				}
				wg.Done()
			}(forestName, hostName)
		}
	}

	for groupName, serverList := range servers {
		for _, serverName := range serverList {
			wg.Add(1)
			go func(serverName string, groupName string) {
				err := e.GetServerMetrics(ch, serverName, groupName)
				if err != nil {
					log.Errorln(err)
					return
				}
				wg.Done()
			}(serverName, groupName)
		}
	}

	wg.Wait()
}

// GetJSON unmarhsals the json returned from uri+path into a supplied interface
func (e *Exporter) GetJSON(path string, inf interface{}) error {
	var resp *http.Response
	var err error

	dr := digest.NewRequest(e.username, e.password, "GET", e.uri+path, "")
	if resp, err = dr.Execute(); err != nil {
		return err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorln(err)
		return err
	}

	err = json.Unmarshal(b, inf)

	return nil
}

// GetHosts returns a list of hostnames
func (e *Exporter) GetHosts() ([]string, error) {
	hosts := []string{}

	hl := hostList{}
	err := e.GetJSON("/manage/v2/hosts?format=json", &hl)
	if err != nil {
		log.Errorln(err)
		return hosts, err
	}

	for _, host := range hl.HostDefaultList.ListItems.ListItem {
		hosts = append(hosts, host.Nameref)
	}

	return hosts, nil
}

// GetGroups returns a list of groups
func (e *Exporter) GetGroups() ([]string, error) {
	groups := []string{}

	gl := groupList{}
	err := e.GetJSON("/manage/v2/groups?format=json", &gl)
	if err != nil {
		log.Errorln(err)
		return groups, err
	}

	for _, group := range gl.GroupDefaultList.ListItems.ListItem {
		groups = append(groups, group.Nameref)
	}

	return groups, nil

}

// GetServers returns a map of groups to servers
func (e *Exporter) GetServers(groups []string) (map[string][]string, error) {
	servers := map[string][]string{}

	for _, groupName := range groups {
		gd := groupDetail{}
		if err := e.GetJSON("/manage/v2/groups/"+groupName+"?format=json", &gd); err != nil {
			log.Errorln(err)
			return servers, err
		}
		for _, relation := range gd.GroupDefault.Relations.RelationGroup {
			if relation.Typeref == "servers" {
				for _, server := range relation.Relation {
					servers[groupName] = append(servers[groupName], server.Nameref)
				}
			}
		}
	}

	return servers, nil
}

// GetForests returns a map of hosts to forests
func (e *Exporter) GetForests(hosts []string) (map[string][]string, error) {

	forests := map[string][]string{}

	for _, hostName := range hosts {
		hd := hostDetail{}
		if err := e.GetJSON("/manage/v2/hosts/"+hostName+"?format=json", &hd); err != nil {
			log.Errorln(err)
			return forests, err
		}

		for _, relation := range hd.HostDefault.Relations.RelationGroup {
			if relation.Typeref == "forests" {
				for _, forest := range relation.Relation {
					forests[hostName] = append(forests[hostName], forest.Nameref)
				}
			}
		}
	}

	return forests, nil
}

// GetHostMetrics gathers metrics from v2/hosts/$host?view=status
func (e *Exporter) GetHostMetrics(ch chan<- prometheus.Metric, hostName string) error {
	hs := hostStatus{}
	if err := e.GetJSON("/manage/v2/hosts/"+hostName+"?view=status&format=json", &hs); err != nil {
		log.Errorln(err)
		return err
	}

	fss := forestStatusSummary{}
	if err := e.GetJSON("/manage/v2/forests?view=status&host-id="+hostName+"&format=json", &fss); err != nil {
		log.Errorln(err)
		return err
	}

	// Disk
	minCapacityValue, _ := strconv.ParseFloat(fss.ForestStatusList.StatusListSummary.MinCapacity.Value, 64)
	ch <- prometheus.MustNewConstMetric(
		diskMinCapacity, prometheus.GaugeValue,
		minCapacityValue,
		hostName,
	)

	// Memory
	ch <- prometheus.MustNewConstMetric(
		memoryFootprint, prometheus.GaugeValue,
		hs.HostStatus.StatusProperties.StatusDetail.MemoryProcessAnon.Value*1048576,
		"anon", hostName,
	)
	ch <- prometheus.MustNewConstMetric(
		memoryFootprint, prometheus.GaugeValue,
		hs.HostStatus.StatusProperties.StatusDetail.MemoryProcessRss.Value*1048576,
		"rss", hostName,
	)
	ch <- prometheus.MustNewConstMetric(
		memoryFootprint, prometheus.GaugeValue,
		hs.HostStatus.StatusProperties.StatusDetail.MemoryProcessRssHwm.Value*1048576,
		"rsshwm", hostName,
	)
	ch <- prometheus.MustNewConstMetric(
		memoryPagesIn, prometheus.GaugeValue,
		hs.HostStatus.StatusProperties.RateProperties.RateDetail.MemorySystemPageinRate.Value,
		hostName,
	)
	ch <- prometheus.MustNewConstMetric(
		memoryPagesOut, prometheus.GaugeValue,
		hs.HostStatus.StatusProperties.RateProperties.RateDetail.MemorySystemPageoutRate.Value,
		hostName,
	)
	ch <- prometheus.MustNewConstMetric(
		memorySwapIn, prometheus.GaugeValue,
		hs.HostStatus.StatusProperties.RateProperties.RateDetail.MemorySystemSwapinRate.Value,
		hostName,
	)
	ch <- prometheus.MustNewConstMetric(
		memorySwapOut, prometheus.GaugeValue,
		hs.HostStatus.StatusProperties.RateProperties.RateDetail.MemorySystemSwapoutRate.Value,
		hostName,
	)

	//Network
	ch <- prometheus.MustNewConstMetric(
		networkXDQPClientReceive, prometheus.GaugeValue,
		hs.HostStatus.StatusProperties.StatusDetail.XdqpClientReceiveBytes.Value,
		hostName,
	)
	ch <- prometheus.MustNewConstMetric(
		networkXDQPClientSend, prometheus.GaugeValue,
		hs.HostStatus.StatusProperties.StatusDetail.XdqpClientSendBytes.Value,
		hostName,
	)
	ch <- prometheus.MustNewConstMetric(
		networkXDQPServerReceive, prometheus.GaugeValue,
		hs.HostStatus.StatusProperties.StatusDetail.XdqpServerReceiveBytes.Value,
		hostName,
	)
	ch <- prometheus.MustNewConstMetric(
		networkXDQPServerSend, prometheus.GaugeValue,
		hs.HostStatus.StatusProperties.StatusDetail.XdqpServerSendBytes.Value,
		hostName,
	)
	ch <- prometheus.MustNewConstMetric(
		networkXDQPForeignClientReceive, prometheus.GaugeValue,
		hs.HostStatus.StatusProperties.StatusDetail.ForeignXdqpClientReceiveBytes.Value,
		hostName,
	)
	ch <- prometheus.MustNewConstMetric(
		networkXDQPForeignClientSend, prometheus.GaugeValue,
		hs.HostStatus.StatusProperties.StatusDetail.ForeignXdqpClientSendBytes.Value,
		hostName,
	)
	ch <- prometheus.MustNewConstMetric(
		networkXDQPForeignServerReceive, prometheus.GaugeValue,
		hs.HostStatus.StatusProperties.StatusDetail.ForeignXdqpServerReceiveBytes.Value,
		hostName,
	)
	ch <- prometheus.MustNewConstMetric(
		networkXDQPForeignServerSend, prometheus.GaugeValue,
		hs.HostStatus.StatusProperties.StatusDetail.ForeignXdqpServerSendBytes.Value,
		hostName,
	)
	return nil
}

func (e *Exporter) GetServerMetrics(ch chan<- prometheus.Metric, serverName string, groupName string) error {
	var databaseName string

	ss := serverStatus{}

	err := e.GetJSON("/manage/v2/servers/"+serverName+"view=status&group-id="+groupName+"&format=json", &ss)
	if err != nil {
		log.Errorln(err)
		return err
	}

	for _, relation := range ss.ServerStatus.Relations.RelationGroup {
		if relation.Typeref == "databases" {
			databaseName = relation.Relation[0].Nameref
		}
	}

	for _, host := range ss.ServerStatus.StatusProperties.HostDetail {
		ch <- prometheus.MustNewConstMetric(
			serverRequestRate, prometheus.GaugeValue,
			host.RequestRate.Value,
			host.RelationID, databaseName, serverName, groupName,
		)
		ch <- prometheus.MustNewConstMetric(
			serverReceive, prometheus.GaugeValue,
			host.ServerReceiveBytes.Value,
			host.RelationID, databaseName, serverName, groupName,
		)
		ch <- prometheus.MustNewConstMetric(
			serverSend, prometheus.GaugeValue,
			host.ServerSendBytes.Value,
			host.RelationID, databaseName, serverName, groupName,
		)
		ch <- prometheus.MustNewConstMetric(
			serverExpandedTreeCacheHits, prometheus.CounterValue,
			host.ExpandedTreeCacheHits.Value,
			host.RelationID, databaseName, serverName, groupName,
		)
		ch <- prometheus.MustNewConstMetric(
			serverExpandedTreeCacheMisses, prometheus.CounterValue,
			host.ExpandedTreeCacheMisses.Value,
			host.RelationID, databaseName, serverName, groupName,
		)
	}

	return nil
}

// GetForestMetrics gathers metrics from v2/forests/$forest?view=status
func (e *Exporter) GetForestMetrics(ch chan<- prometheus.Metric, forestName string, hostName string) error {
	var databaseName string

	fs := forestStatus{}
	err := e.GetJSON("/manage/v2/forests/"+forestName+"?view=status&format=json", &fs)
	if err != nil {
		log.Errorln(err)
		return err
	}

	for _, relation := range fs.ForestStatus.Relations.RelationGroup {
		if relation.Typeref == "databases" {
			databaseName = relation.Relation[0].Nameref
			break
		}
	}

	// Disk read
	ch <- prometheus.MustNewConstMetric(
		diskRead, prometheus.CounterValue,
		fs.ForestStatus.StatusProperties.QueryReadBytes.Value,
		"query", hostName, databaseName, forestName,
	)
	ch <- prometheus.MustNewConstMetric(
		diskRead, prometheus.CounterValue,
		fs.ForestStatus.StatusProperties.MergeReadBytes.Value,
		"merge", hostName, databaseName, forestName,
	)
	ch <- prometheus.MustNewConstMetric(
		diskRead, prometheus.CounterValue,
		fs.ForestStatus.StatusProperties.BackupReadBytes.Value,
		"backup", hostName, databaseName, forestName,
	)
	ch <- prometheus.MustNewConstMetric(
		diskRead, prometheus.CounterValue,
		fs.ForestStatus.StatusProperties.RestoreReadBytes.Value,
		"restore", hostName, databaseName, forestName,
	)
	ch <- prometheus.MustNewConstMetric(
		diskRead, prometheus.CounterValue,
		fs.ForestStatus.StatusProperties.LargeReadBytes.Value,
		"large", hostName, databaseName, forestName,
	)
	// Disk write
	ch <- prometheus.MustNewConstMetric(
		diskWrite, prometheus.CounterValue,
		fs.ForestStatus.StatusProperties.JournalWriteBytes.Value,
		"journal", hostName, databaseName, forestName,
	)
	ch <- prometheus.MustNewConstMetric(
		diskWrite, prometheus.CounterValue,
		fs.ForestStatus.StatusProperties.SaveWriteBytes.Value,
		"save", hostName, databaseName, forestName,
	)
	ch <- prometheus.MustNewConstMetric(
		diskWrite, prometheus.CounterValue,
		fs.ForestStatus.StatusProperties.MergeWriteBytes.Value,
		"merge", hostName, databaseName, forestName,
	)
	ch <- prometheus.MustNewConstMetric(
		diskWrite, prometheus.CounterValue,
		fs.ForestStatus.StatusProperties.BackupWriteBytes.Value,
		"backup", hostName, databaseName, forestName,
	)
	ch <- prometheus.MustNewConstMetric(
		diskWrite, prometheus.CounterValue,
		fs.ForestStatus.StatusProperties.RestoreWriteBytes.Value,
		"restore", hostName, databaseName, forestName,
	)
	ch <- prometheus.MustNewConstMetric(
		diskWrite, prometheus.CounterValue,
		fs.ForestStatus.StatusProperties.LargeWriteBytes.Value,
		"large", hostName, databaseName, forestName,
	)
	// Disk journals size
	ch <- prometheus.MustNewConstMetric(
		diskJournalsSize, prometheus.GaugeValue,
		fs.ForestStatus.StatusProperties.JournalsSize.Value*1048576,
		hostName, databaseName, forestName,
	)
	// Disk forest reserve size
	ch <- prometheus.MustNewConstMetric(
		diskForestReserveSize, prometheus.GaugeValue,
		fs.ForestStatus.StatusProperties.ForestReserve.Value*1048576,
		hostName, databaseName, forestName,
	)

	// Database metrics
	ch <- prometheus.MustNewConstMetric(
		databaseReadLocks, prometheus.GaugeValue,
		fs.ForestStatus.StatusProperties.ReadLockCount.Value,
		hostName, databaseName, forestName,
	)
	readLockWaitTimeValue, _ := strconv.ParseFloat(fs.ForestStatus.StatusProperties.ReadLockWaitTime.Value, 64)
	ch <- prometheus.MustNewConstMetric(
		databaseReadLockWaitTime, prometheus.GaugeValue,
		readLockWaitTimeValue,
		hostName, databaseName, forestName,
	)
	readLockHoldTimeValue, _ := strconv.ParseFloat(fs.ForestStatus.StatusProperties.ReadLockHoldTime.Value, 64)
	ch <- prometheus.MustNewConstMetric(
		databaseReadLockHoldTime, prometheus.GaugeValue,
		readLockHoldTimeValue,
		hostName, databaseName, forestName,
	)
	ch <- prometheus.MustNewConstMetric(
		databaseWriteLocks, prometheus.GaugeValue,
		fs.ForestStatus.StatusProperties.WriteLockCount.Value,
		hostName, databaseName, forestName,
	)
	writeLockWaitTimeValue, _ := strconv.ParseFloat(fs.ForestStatus.StatusProperties.WriteLockWaitTime.Value, 64)
	ch <- prometheus.MustNewConstMetric(
		databaseWriteLockWaitTime, prometheus.GaugeValue,
		writeLockWaitTimeValue,
		hostName, databaseName, forestName,
	)
	writeLockHoldTimeValue, _ := strconv.ParseFloat(fs.ForestStatus.StatusProperties.WriteLockHoldTime.Value, 64)
	ch <- prometheus.MustNewConstMetric(
		databaseWriteLockHoldTime, prometheus.GaugeValue,
		writeLockHoldTimeValue,
		hostName, databaseName, forestName,
	)
	ch <- prometheus.MustNewConstMetric(
		databaseDeadLocks, prometheus.GaugeValue,
		fs.ForestStatus.StatusProperties.DeadlockCount.Value,
		hostName, databaseName, forestName,
	)
	deadLockWaitTimeValue, _ := strconv.ParseFloat(fs.ForestStatus.StatusProperties.DeadlockWaitTime.Value, 64)
	ch <- prometheus.MustNewConstMetric(
		databaseDeadLockWaitTime, prometheus.GaugeValue,
		deadLockWaitTimeValue,
		hostName, databaseName, forestName,
	)
	ch <- prometheus.MustNewConstMetric(
		databaseReplicationReceive, prometheus.GaugeValue,
		fs.ForestStatus.StatusProperties.DatabaseReplicationReceiveBytes.Value,
		hostName, databaseName, forestName,
	)
	ch <- prometheus.MustNewConstMetric(
		databaseReplicationSend, prometheus.GaugeValue,
		fs.ForestStatus.StatusProperties.DatabaseReplicationSendBytes.Value,
		hostName, databaseName, forestName,
	)

	for _, stand := range fs.ForestStatus.StatusProperties.Stand {
		standID := stand.StandID[0]

		// Disk metrics
		ch <- prometheus.MustNewConstMetric(
			diskSize, prometheus.GaugeValue,
			stand.DiskSize.Value*1048576,
			hostName, databaseName, forestName, standID, "false", stand.IsFast.Value,
		)
		encryptedDiskSizeValue, _ := strconv.ParseFloat(stand.EncryptedDiskSize.Value, 64)
		ch <- prometheus.MustNewConstMetric(diskSize, prometheus.GaugeValue,
			encryptedDiskSizeValue*1048576,
			hostName, databaseName, forestName, standID, "true", stand.IsFast.Value,
		)

		// Database metrics
		listCacheHitsValue, _ := strconv.ParseFloat(stand.ListCacheHits.Value, 64)
		ch <- prometheus.MustNewConstMetric(
			databaseListCacheHits, prometheus.CounterValue,
			listCacheHitsValue,
			hostName, databaseName, forestName, standID,
		)
		listCacheMissesValue, _ := strconv.ParseFloat(stand.ListCacheMisses.Value, 64)
		ch <- prometheus.MustNewConstMetric(
			databaseListCacheMisses, prometheus.CounterValue,
			listCacheMissesValue,
			hostName, databaseName, forestName, standID,
		)
		compressedTreeCacheHitsValue, _ := strconv.ParseFloat(stand.CompressedTreeCacheHits.Value, 64)
		ch <- prometheus.MustNewConstMetric(
			databaseCompressedTreeCacheHits, prometheus.CounterValue,
			compressedTreeCacheHitsValue,
			hostName, databaseName, forestName, standID,
		)
		compressedTreeCacheMissesValue, _ := strconv.ParseFloat(stand.CompressedTreeCacheMisses.Value, 64)
		ch <- prometheus.MustNewConstMetric(
			databaseCompressedTreeCacheMisses, prometheus.CounterValue,
			compressedTreeCacheMissesValue,
			hostName, databaseName, forestName, standID,
		)
		tripleCacheHitsValue, _ := strconv.ParseFloat(stand.TripleCacheHits.Value, 64)
		ch <- prometheus.MustNewConstMetric(
			databaseTripleCacheHits, prometheus.CounterValue,
			tripleCacheHitsValue,
			hostName, databaseName, forestName, standID,
		)
		tripleCacheMissesValue, _ := strconv.ParseFloat(stand.TripleCacheMisses.Value, 64)
		ch <- prometheus.MustNewConstMetric(
			databaseTripleCacheMisses, prometheus.CounterValue,
			tripleCacheMissesValue,
			hostName, databaseName, forestName, standID,
		)
		tripleValueCacheHitsValue, _ := strconv.ParseFloat(stand.TripleValueCacheHits.Value, 64)
		ch <- prometheus.MustNewConstMetric(
			databaseTripleValueCacheHits, prometheus.CounterValue,
			tripleValueCacheHitsValue,
			hostName, databaseName, forestName, standID,
		)
		tripleValueCacheMissesValue, _ := strconv.ParseFloat(stand.TripleValueCacheMisses.Value, 64)
		ch <- prometheus.MustNewConstMetric(
			databaseTripleValueCacheMisses, prometheus.CounterValue,
			tripleValueCacheMissesValue,
			hostName, databaseName, forestName, standID,
		)
	}

	return nil
}
func main() {
	var (
		listenAddress = kingpin.Flag("web.listen-address", "Address to listen on for web interface and telemetry.").Default(":9307").String()
		metricsPath   = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").Default("/metrics").String()
		iniFile       = kingpin.Flag("config.ini", "Configuration file").Default("./marklogic.ini").String()
		uri           = kingpin.Flag("marklogic.uri", "HTTP API address of a MarkLogic server").Default("http://localhost:8002").String()
	)

	log.AddFlags(kingpin.CommandLine)
	kingpin.Version(version.Print(namespace + "_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	cfg, err := ini.InsensitiveLoad(*iniFile)
	if err != nil {
		log.Fatal("Issue when loading config from " + *iniFile)
	}

	log.Infoln("Starting "+namespace+"_exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())

	prometheus.MustRegister(&Exporter{
		username: cfg.Section("auth").Key("username").String(),
		password: cfg.Section("auth").Key("password").String(),
		uri:      *uri,
	})

	http.Handle(*metricsPath, prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
						 <head><title>MarkLogic Exporter</title></head>
						 <body>
						 <h1>MarkLogic Exporter</h1>
						 <p><a href='` + *metricsPath + `'>Metrics</a></p>
						 </body>
						 </html>`))
	})

	log.Infoln("Listening on", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
