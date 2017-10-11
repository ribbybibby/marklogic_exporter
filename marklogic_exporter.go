package main

import (
	"encoding/json"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"

	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/ini.v1"
	digest "github.com/xinsnake/go-http-digest-auth-client"

)

const (
	namespace = "marklogic"
)

var (
	up = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "up"),
		"Was the last query of MarkLogic successful?",
		nil, nil,
	)
	HostsTotalHosts = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "hosts", "total_hosts"),
		"Number of hosts in this cluster.",
		nil,nil,
	)
	DatabaseReadLockWaitLoad = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "read_lock_wait_load"),
		"Read lock wait load",
		nil,nil,
	)
	DatabaseReadLockHoldLoad = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "read_lock_hold_load"),
		"Read lock hold load",
		nil,nil,
	)
	DatabaseWriteLockWaitLoad = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "write_lock_wait_load"),
		"Write lock wait load",
		nil,nil,
	)
	DatabaseWriteLockHoldLoad = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "write_lock_hold_load"),
		"Write lock hold load",
		nil,nil,
	)
	DatabaseDeadlockWaitLoad = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "deadlock_wait_load"),
		"Deadlock wait load",
		nil,nil,
	)
	DatabaseReadLockRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "read_lock_rate"),
		"The number of read locks set per second on each database.",
		nil,nil,
	)
	DatabaseWriteLockRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "write_lock_rate"),
		"The number of write locks set per second on each database.",
		nil,nil,
	)
	DatabaseDeadlockRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "database", "deadlock_rate"),
		"The number of deadlocks per second on each database",
		nil,nil,
	)
	MemoryProcessSwapRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "memory", "process_swap_rate"),
		"The overall swap rate (from Linux /proc/vmstat for the cluster in pages/sec",
		nil,nil,
	)
	MemorySystemPageinRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "memory", "system_pagein_rate"),
		"The page-in rate (from Linux /proc/vmstat) for the cluster in pages/sec.",
		nil,nil,
	)
	MemorySystemPageoutRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "memory", "system_pageout_rate"),
		"The page-out rate (from Linux /proc/vmstat) for the cluster in pages/sec.",
		nil,nil,
	)
	MemorySystemSwapinRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "memory", "system_swapin_rate"),
		"The swap-in rate (from Linux /proc/vmstat) for the cluster in pages/sec.",
		nil,nil,
	)
	MemorySystemSwapoutRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "memory", "system_swapout_rate"),
		"The swap-out rate (from Linux /proc/vmstat) for the cluster in pages/sec.",
		nil,nil,
	)
	DiskQueryReadRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "disk", "query_read_rate"),
		"The moving average of reading query data from disk",
		nil,nil,
	)
	DiskJournalWriteRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "disk", "journal_write_rate"),
		"The moving average of data writes to the journal.",
		nil,nil,
	)
	DiskSaveWriteRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "disk", "save_write_rate"),
		"The moving average of data writes to in-memory stands.",
		nil,nil,
	)
	DiskMergeReadRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "disk", "merge_read_rate"),
		"The moving average of reading merge data from disk.",
		nil,nil,
	)
	DiskMergeWriteRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "disk", "merge_write_rate"),
		"The moving average of writing data for merges.",
		nil,nil,
	)
	DiskBackupReadRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "disk", "backup_read_rate"),
		"The moving average of reading backup data from disk.",
		nil,nil,
	)
	DiskBackupWriteRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "disk", "backup_write_rate"),
		"The moving average of writing backup data to disk.",
		nil,nil,
	)
	DiskRestoreReadRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "disk", "restore_read_rate"),
		"The moving average of reading restore data from disk.",
		nil,nil,
	)
	DiskRestoreWriteRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "disk", "restore_write_rate"),
		"The moving average of writing restore data to disk.",
		nil,nil,
	)
	DiskLargeReadRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "disk", "large_read_rate"),
		"The moving average of reading large documents from disk.",
		nil,nil,
	)
	DiskLargeWriteRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "disk", "large_write_rate"),
		"The moving average of writing data for large documents to disk.",
		nil,nil,
	)
	DiskExternalBinaryReadRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "disk", "external_binary_read_rate"),
		"External binary read rate",
		nil,nil,
	)
	XDQPClientReceiveRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "xdqp", "client_receive_rate"),
		"XDQP client receive rate",
		nil,nil,
	)
	XDQPClientSendRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "xdqp", "client_send_rate"),
		"XDQP client send rate",
		nil,nil,
	)
	XDQPServerReceiveRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "xdqp", "server_receive_rate"),
		"XDQP server receive rate",
		nil,nil,
	)
	XDQPServerSendRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "xdqp", "server_send_rate"),
		"XDQP server send rate",
		nil,nil,
	)
	XDQPForeignClientReceiveRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "xdqp", "foreign_client_receive_rate"),
		"XDQP foreign client receive rate",
		nil,nil,
	)
	XDQPForeignClientSendRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "xdqp", "foreign_client_send_rate"),
		"XDQP foreign client send rate",
		nil,nil,
	)
	XDQPForeignServerReceiveRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "xdqp", "foreign_server_receive_rate"),
		"XDQP foreign server receive rate",
		nil,nil,
	)
	XDQPForeignServerSendRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "xdqp", "foreign_server_send_rate"),
		"XDQP foreign server send rate",
		nil,nil,
	)
	ServerRequestRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "server", "request_rate"),
		"The total number of queries being processed per second, across all of the App Servers.",
		nil,nil,
	)
	ServerExpandedTreeCacheMissRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "server", "expanded_tree_cache_miss_rate"),
		"The number of times per second that queries couldn't use the expanded tree cache.",
		nil,nil,
	)
	ServerExpandedTreeCacheHitRate = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "server", "expanded_tree_cache_hit_rate"),
		"The number of times per second that queries could use the expanded tree cache.",
		nil,nil,
	)
	ServerRequestCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "server", "request_count"),
		"Server request count",
		nil,nil,
	)
	RequestsTotalRequests = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "requests", "total"),
		"Total requests",
		nil,nil,
	)
	RequestsUpdateCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "requests", "update_count"),
		"Number of updates ",
		nil,nil,
	)
	RequestsQueryCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "requests", "query_count"),
		"Query count",
		nil,nil,
	)
	TransactionsTotalTransactions = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "transactions", "total"),
		"Total transactions",
		nil,nil,
	)
	ForestsTotalForests = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "forests", "total"),
		"Total number of forests",
		nil,nil,
	)
)

type Exporter struct {
	password string
	uri		 string
	username string
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- up
	ch <- HostsTotalHosts
	ch <- DatabaseReadLockWaitLoad
	ch <- DatabaseReadLockHoldLoad
	ch <- DatabaseWriteLockWaitLoad
	ch <- DatabaseWriteLockHoldLoad
	ch <- DatabaseDeadlockWaitLoad
	ch <- DatabaseReadLockRate
	ch <- DatabaseWriteLockRate
	ch <- DatabaseDeadlockRate
	ch <- MemoryProcessSwapRate
	ch <- MemorySystemPageinRate
	ch <- MemorySystemPageoutRate
	ch <- MemorySystemSwapinRate
	ch <- MemorySystemSwapoutRate
	ch <- DiskQueryReadRate
	ch <- DiskJournalWriteRate
	ch <- DiskSaveWriteRate
	ch <- DiskMergeReadRate
	ch <- DiskMergeWriteRate
	ch <- DiskBackupReadRate
	ch <- DiskBackupWriteRate
	ch <- DiskRestoreReadRate
	ch <- DiskRestoreWriteRate
	ch <- DiskLargeReadRate
	ch <- DiskLargeWriteRate
	ch <- DiskExternalBinaryReadRate
	ch <- XDQPClientReceiveRate
	ch <- XDQPClientSendRate
	ch <- XDQPServerReceiveRate
	ch <- XDQPServerSendRate
	ch <- XDQPForeignClientReceiveRate
	ch <- XDQPForeignClientSendRate
	ch <- XDQPForeignServerReceiveRate
	ch <- XDQPForeignServerSendRate
	ch <- ServerRequestRate
	ch <- ServerExpandedTreeCacheMissRate
	ch <- ServerExpandedTreeCacheHitRate
	ch <- ServerRequestCount
	ch <- RequestsTotalRequests
	ch <- RequestsUpdateCount
	ch <- RequestsQueryCount
	ch <- TransactionsTotalTransactions
	ch <- ForestsTotalForests
}

func (e *Exporter) GetStatus() (Status, error) {

	var resp *http.Response
	var err error

	s := Status{}

	dr := digest.NewRequest(e.username, e.password, "GET", e.uri + "/manage/v2?view=status&format=json", "")
	if resp, err = dr.Execute(); err != nil {
		return s, err
	}

	defer resp.Body.Close()
	
	json.NewDecoder(resp.Body).Decode(&s)	

	return s, nil
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {

	s, err := e.GetStatus()
	if err != nil {
		log.Errorln(err)
		ch <- prometheus.MustNewConstMetric(
			up, prometheus.GaugeValue, 0,
		)
		return
	}

	// up
	ch <- prometheus.MustNewConstMetric(
		up, prometheus.GaugeValue, 1,
	)

	// Hosts Status
	ch <- prometheus.MustNewConstMetric(
		HostsTotalHosts, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.TotalHosts.Value),
	)

	// Database Performance
	ch <- prometheus.MustNewConstMetric(
		DatabaseReadLockWaitLoad, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.LoadProperties.LoadDetail.ReadLockWaitLoad.Value),
	)
	ch <- prometheus.MustNewConstMetric(
		DatabaseReadLockHoldLoad, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.LoadProperties.LoadDetail.ReadLockHoldLoad.Value),
	)
	ch <- prometheus.MustNewConstMetric(
		DatabaseWriteLockWaitLoad, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.LoadProperties.LoadDetail.WriteLockWaitLoad.Value),
	)
	ch <- prometheus.MustNewConstMetric(
		DatabaseWriteLockHoldLoad, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.LoadProperties.LoadDetail.WriteLockHoldLoad.Value),
	)
	ch <- prometheus.MustNewConstMetric(
		DatabaseDeadlockWaitLoad, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.LoadProperties.LoadDetail.DeadlockWaitLoad.Value),
	)
	ch <- prometheus.MustNewConstMetric(
		DatabaseReadLockRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.ReadLockRate.Value),
	)
	ch <- prometheus.MustNewConstMetric(
		DatabaseWriteLockRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.WriteLockRate.Value),
	)
	ch <- prometheus.MustNewConstMetric(
		DatabaseDeadlockRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.DeadlockRate.Value),
	)

	// Memory
	ch <- prometheus.MustNewConstMetric(
		MemoryProcessSwapRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.MemoryProcessSwapRate.Value * 1048576),
	)
	ch <- prometheus.MustNewConstMetric(
		MemorySystemPageinRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.MemorySystemPageinRate.Value * 1048576),
	)
	ch <- prometheus.MustNewConstMetric(
		MemorySystemPageoutRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.MemorySystemPageoutRate.Value * 1048576),
	)
	ch <- prometheus.MustNewConstMetric(
		MemorySystemSwapinRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.MemorySystemSwapinRate.Value * 1048576),
	)
	ch <- prometheus.MustNewConstMetric(
		MemorySystemSwapoutRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.MemorySystemSwapoutRate.Value * 1048576),
	)

	// Disk
	ch <- prometheus.MustNewConstMetric(
		DiskQueryReadRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.QueryReadRate.Value * 1048576),
	)
	ch <- prometheus.MustNewConstMetric(
		DiskJournalWriteRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.JournalWriteRate.Value * 1048576),
	)	
	ch <- prometheus.MustNewConstMetric(
		DiskSaveWriteRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.SaveWriteRate.Value * 1048576),
	)	
	ch <- prometheus.MustNewConstMetric(
		DiskMergeReadRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.MergeReadRate.Value * 1048576),
	)	
	ch <- prometheus.MustNewConstMetric(
		DiskMergeWriteRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.MergeWriteRate.Value * 1048576),
	)	
	ch <- prometheus.MustNewConstMetric(
		DiskBackupReadRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.BackupReadRate.Value * 1048576),
	)	
	ch <- prometheus.MustNewConstMetric(
		DiskBackupWriteRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.BackupWriteRate.Value * 1048576),
	)	
	ch <- prometheus.MustNewConstMetric(
		DiskRestoreReadRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.RestoreReadRate.Value * 1048576),
	)	
	ch <- prometheus.MustNewConstMetric(
		DiskRestoreWriteRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.RestoreWriteRate.Value * 1048576),
	)	
	ch <- prometheus.MustNewConstMetric(
		DiskLargeReadRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.LargeReadRate.Value * 1048576),
	)	
	ch <- prometheus.MustNewConstMetric(
		DiskLargeWriteRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.LargeWriteRate.Value * 1048576),
	)	
	ch <- prometheus.MustNewConstMetric(
		DiskExternalBinaryReadRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.ExternalBinaryReadRate.Value * 1048576),
	)

	// XDQP (Network)
	ch <- prometheus.MustNewConstMetric(
		XDQPClientReceiveRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.XdqpClientReceiveRate.Value * 1048576),
	)
	ch <- prometheus.MustNewConstMetric(
		XDQPClientSendRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.XdqpClientSendRate.Value * 1048576),
	)
	ch <- prometheus.MustNewConstMetric(
		XDQPServerReceiveRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.XdqpServerReceiveRate.Value * 1048576),
	)
	ch <- prometheus.MustNewConstMetric(
		XDQPServerSendRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.XdqpServerSendRate.Value * 1048576),
	)
	ch <- prometheus.MustNewConstMetric(
		XDQPForeignClientReceiveRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.ForeignXdqpClientReceiveRate.Value * 1048576),
	)
	ch <- prometheus.MustNewConstMetric(
		XDQPForeignClientSendRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.ForeignXdqpClientSendRate.Value * 1048576),
	)
	ch <- prometheus.MustNewConstMetric(
		XDQPForeignServerReceiveRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.ForeignXdqpServerReceiveRate.Value * 1048576),
	)
	ch <- prometheus.MustNewConstMetric(
		XDQPForeignServerSendRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.HostsStatus.HostsStatusSummary.RateProperties.RateDetail.ForeignXdqpServerSendRate.Value * 1048576),
	)

	// Server Performance
	ch <- prometheus.MustNewConstMetric(
		ServerRequestRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.ServersStatus.ServersStatusSummary.RequestRate.Value),
	)	
	ch <- prometheus.MustNewConstMetric(
		ServerExpandedTreeCacheMissRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.ServersStatus.ServersStatusSummary.ExpandedTreeCacheMissRate.Value),
	)
	ch <- prometheus.MustNewConstMetric(
		ServerExpandedTreeCacheHitRate, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.ServersStatus.ServersStatusSummary.ExpandedTreeCacheHitRate.Value),
	)
	ch <- prometheus.MustNewConstMetric(
		ServerRequestCount, prometheus.CounterValue, float64(s.LocalClusterStatus.StatusRelations.ServersStatus.ServersStatusSummary.RequestCount.Value),
	)

	// Requests
	ch <- prometheus.MustNewConstMetric(
		RequestsTotalRequests, prometheus.CounterValue, float64(s.LocalClusterStatus.StatusRelations.RequestsStatus.RequestsStatusSummary.TotalRequests.Value),
	)
	ch <- prometheus.MustNewConstMetric(
		RequestsQueryCount, prometheus.CounterValue, float64(s.LocalClusterStatus.StatusRelations.RequestsStatus.RequestsStatusSummary.UpdateCount.Value),
	)
	ch <- prometheus.MustNewConstMetric(
		RequestsUpdateCount, prometheus.CounterValue, float64(s.LocalClusterStatus.StatusRelations.RequestsStatus.RequestsStatusSummary.QueryCount.Value),
	)

	// Transactions
	ch <- prometheus.MustNewConstMetric(
		TransactionsTotalTransactions, prometheus.CounterValue, float64(s.LocalClusterStatus.StatusRelations.TransactionsStatus.TransactionsStatusSummary.TotalTransactions.Value),
	)

	// Forests
	ch <- prometheus.MustNewConstMetric(
		ForestsTotalForests, prometheus.GaugeValue, float64(s.LocalClusterStatus.StatusRelations.ForestsStatus.ForestsStatusSummary.TotalForests.Value),
	)	
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

	prometheus.MustRegister(&Exporter{
		username: cfg.Section("auth").Key("username").String(),
		password: cfg.Section("auth").Key("password").String(),
		uri: *uri,
	})

	log.Infoln("Starting " + namespace + "_exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())

	http.Handle(*metricsPath, prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
						 <head><title>MarkLogic Exporter</title></head>
						 <body>
						 <h1>MarkLogic Exporter</h1>
						 <p><a href='`+ *metricsPath + `'>Metrics</a></p>
						 </body>
						 </html>`))		
	})

	log.Infoln("Listening on", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}