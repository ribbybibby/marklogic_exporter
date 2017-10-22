package main

import "time"

type hostList struct {
	HostDefaultList struct {
		Meta struct {
			URI         string    `json:"uri"`
			CurrentTime time.Time `json:"current-time"`
			ElapsedTime struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"elapsed-time"`
		} `json:"meta"`
		ListItems struct {
			ListCount struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"list-count"`
			ListItem []struct {
				Uriref  string `json:"uriref"`
				Roleref string `json:"roleref"`
				Idref   string `json:"idref"`
				Nameref string `json:"nameref"`
			} `json:"list-item"`
		} `json:"list-items"`
		RelatedViews struct {
			RelatedView []struct {
				ViewType string `json:"view-type"`
				ViewName string `json:"view-name"`
				ViewURI  string `json:"view-uri"`
			} `json:"related-view"`
		} `json:"related-views"`
	} `json:"host-default-list"`
}

type hostDetail struct {
	HostDefault struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Meta struct {
			URI         string    `json:"uri"`
			CurrentTime time.Time `json:"current-time"`
			ElapsedTime struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"elapsed-time"`
		} `json:"meta"`
		Relations struct {
			RelationGroup []struct {
				Typeref       string `json:"typeref"`
				RelationCount struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"relation-count,omitempty"`
				Relation []struct {
					Uriref  string `json:"uriref"`
					Idref   string `json:"idref"`
					Nameref string `json:"nameref"`
				} `json:"relation,omitempty"`
				Uriref string `json:"uriref,omitempty"`
			} `json:"relation-group"`
		} `json:"relations"`
		RelatedViews struct {
			RelatedView []struct {
				ViewType string `json:"view-type"`
				ViewName string `json:"view-name"`
				ViewURI  string `json:"view-uri"`
			} `json:"related-view"`
		} `json:"related-views"`
	} `json:"host-default"`
}

type hostStatus struct {
	HostStatus struct {
		ID                  string `json:"id"`
		Name                string `json:"name"`
		Version             string `json:"version"`
		EffectiveVersion    int    `json:"effective-version"`
		HostMode            string `json:"host-mode"`
		HostModeDescription string `json:"host-mode-description"`
		Meta                struct {
			URI         string    `json:"uri"`
			CurrentTime time.Time `json:"current-time"`
			ElapsedTime struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"elapsed-time"`
		} `json:"meta"`
		Relations struct {
			RelationGroup []struct {
				Uriref   string `json:"uriref,omitempty"`
				Typeref  string `json:"typeref"`
				Relation []struct {
					Uriref  string `json:"uriref"`
					Idref   string `json:"idref"`
					Nameref string `json:"nameref"`
				} `json:"relation,omitempty"`
			} `json:"relation-group"`
		} `json:"relations"`
		StatusProperties struct {
			Online struct {
				Units string `json:"units"`
				Value bool   `json:"value"`
			} `json:"online"`
			Secure struct {
				Units string `json:"units"`
				Value bool   `json:"value"`
			} `json:"secure"`
			CacheProperties struct {
				CacheDetail struct {
					CompressedTreeCachePartitions struct {
						CompressedTreeCachePartition []struct {
							PartitionSize     int     `json:"partition-size"`
							PartitionTable    float64 `json:"partition-table"`
							PartitionUsed     float64 `json:"partition-used"`
							PartitionFree     float64 `json:"partition-free"`
							PartitionOverhead float64 `json:"partition-overhead"`
						} `json:"compressed-tree-cache-partition"`
					} `json:"compressed-tree-cache-partitions"`
					ExpandedTreeCachePartitions struct {
						ExpandedTreeCachePartition []struct {
							PartitionSize     int     `json:"partition-size"`
							PartitionTable    float64 `json:"partition-table"`
							PartitionBusy     int     `json:"partition-busy"`
							PartitionUsed     float64 `json:"partition-used"`
							PartitionFree     float64 `json:"partition-free"`
							PartitionOverhead float64 `json:"partition-overhead"`
						} `json:"expanded-tree-cache-partition"`
					} `json:"expanded-tree-cache-partitions"`
					TripleCachePartitions struct {
						TripleCachePartition []struct {
							PartitionSize int `json:"partition-size"`
							PartitionBusy int `json:"partition-busy"`
							PartitionUsed int `json:"partition-used"`
							PartitionFree int `json:"partition-free"`
						} `json:"triple-cache-partition"`
					} `json:"triple-cache-partitions"`
					TripleValueCachePartitions struct {
						TripleValueCachePartition []struct {
							PartitionSize int `json:"partition-size"`
							PartitionBusy int `json:"partition-busy"`
							PartitionUsed int `json:"partition-used"`
							PartitionFree int `json:"partition-free"`
						} `json:"triple-value-cache-partition"`
					} `json:"triple-value-cache-partitions"`
				} `json:"cache-detail"`
			} `json:"cache-properties"`
			LoadProperties struct {
				TotalLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"total-load"`
				LoadDetail struct {
					QueryReadLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"query-read-load"`
					JournalWriteLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"journal-write-load"`
					SaveWriteLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"save-write-load"`
					MergeReadLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"merge-read-load"`
					MergeWriteLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"merge-write-load"`
					BackupReadLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"backup-read-load"`
					BackupWriteLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"backup-write-load"`
					RestoreReadLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"restore-read-load"`
					RestoreWriteLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"restore-write-load"`
					LargeReadLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"large-read-load"`
					LargeWriteLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"large-write-load"`
					ExternalBinaryReadLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"external-binary-read-load"`
					XdqpClientReceiveLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"xdqp-client-receive-load"`
					XdqpClientSendLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"xdqp-client-send-load"`
					XdqpServerReceiveLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"xdqp-server-receive-load"`
					XdqpServerSendLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"xdqp-server-send-load"`
					ForeignXdqpClientReceiveLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"foreign-xdqp-client-receive-load"`
					ForeignXdqpClientSendLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"foreign-xdqp-client-send-load"`
					ForeignXdqpServerReceiveLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"foreign-xdqp-server-receive-load"`
					ForeignXdqpServerSendLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"foreign-xdqp-server-send-load"`
					ReadLockWaitLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"read-lock-wait-load"`
					ReadLockHoldLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"read-lock-hold-load"`
					WriteLockWaitLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"write-lock-wait-load"`
					WriteLockHoldLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"write-lock-hold-load"`
					DeadlockWaitLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"deadlock-wait-load"`
				} `json:"load-detail"`
			} `json:"load-properties"`
			RateProperties struct {
				TotalRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"total-rate"`
				RateDetail struct {
					MemorySystemPageinRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"memory-system-pagein-rate"`
					MemorySystemPageoutRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"memory-system-pageout-rate"`
					MemorySystemSwapinRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"memory-system-swapin-rate"`
					MemorySystemSwapoutRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"memory-system-swapout-rate"`
					QueryReadRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"query-read-rate"`
					JournalWriteRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"journal-write-rate"`
					SaveWriteRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"save-write-rate"`
					MergeReadRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"merge-read-rate"`
					MergeWriteRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"merge-write-rate"`
					BackupReadRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"backup-read-rate"`
					BackupWriteRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"backup-write-rate"`
					RestoreReadRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"restore-read-rate"`
					RestoreWriteRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"restore-write-rate"`
					LargeReadRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"large-read-rate"`
					LargeWriteRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"large-write-rate"`
					ExternalBinaryReadRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"external-binary-read-rate"`
					XdqpClientReceiveRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"xdqp-client-receive-rate"`
					XdqpClientSendRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"xdqp-client-send-rate"`
					XdqpServerReceiveRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"xdqp-server-receive-rate"`
					XdqpServerSendRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"xdqp-server-send-rate"`
					ForeignXdqpClientReceiveRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"foreign-xdqp-client-receive-rate"`
					ForeignXdqpClientSendRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"foreign-xdqp-client-send-rate"`
					ForeignXdqpServerReceiveRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"foreign-xdqp-server-receive-rate"`
					ForeignXdqpServerSendRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"foreign-xdqp-server-send-rate"`
					ReadLockRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"read-lock-rate"`
					WriteLockRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"write-lock-rate"`
					DeadlockRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"deadlock-rate"`
				} `json:"rate-detail"`
			} `json:"rate-properties"`
			StatusDetail struct {
				Zone           string `json:"zone"`
				BindPort       int    `json:"bind-port"`
				ConnectPort    int    `json:"connect-port"`
				SslFipsEnabled struct {
					Units string `json:"units"`
					Value bool   `json:"value"`
				} `json:"ssl-fips-enabled"`
				ForeignBindPort    int `json:"foreign-bind-port"`
				ForeignConnectPort int `json:"foreign-connect-port"`
				BackgroundIoLimit  struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"background-io-limit"`
				MeteringEnabled struct {
					Units string `json:"units"`
					Value bool   `json:"value"`
				} `json:"metering-enabled"`
				MetersDatabase struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"meters-database"`
				PerformanceMeteringEnabled struct {
					Units string `json:"units"`
					Value bool   `json:"value"`
				} `json:"performance-metering-enabled"`
				PerformanceMeteringPeriod struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"performance-metering-period"`
				PerformanceMeteringRetainRaw struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"performance-metering-retain-raw"`
				PerformanceMeteringRetainHourly struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"performance-metering-retain-hourly"`
				PerformanceMeteringRetainDaily struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"performance-metering-retain-daily"`
				LastStartup struct {
					Units string    `json:"units"`
					Value time.Time `json:"value"`
				} `json:"last-startup"`
				Version          string `json:"version"`
				EffectiveVersion struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"effective-version"`
				SoftwareVersion struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"software-version"`
				OsVersion string `json:"os-version"`
				HostMode  struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"host-mode"`
				Architecture   string `json:"architecture"`
				Platform       string `json:"platform"`
				LicenseKeyCpus struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"license-key-cpus"`
				LicenseKeyCores struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"license-key-cores"`
				LicenseKeySize struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"license-key-size"`
				LicenseKeyOption []struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"license-key-option"`
				Edition struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"edition"`
				Environment struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"environment"`
				Cpus struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"cpus"`
				Cores struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"cores"`
				CoreThreads struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"core-threads"`
				TotalCPUStatUser      float64 `json:"total-cpu-stat-user"`
				TotalCPUStatNice      int     `json:"total-cpu-stat-nice"`
				TotalCPUStatSystem    float64 `json:"total-cpu-stat-system"`
				TotalCPUStatIdle      float64 `json:"total-cpu-stat-idle"`
				TotalCPUStatIowait    float64 `json:"total-cpu-stat-iowait"`
				TotalCPUStatIrq       int     `json:"total-cpu-stat-irq"`
				TotalCPUStatSoftirq   float64 `json:"total-cpu-stat-softirq"`
				TotalCPUStatSteal     float64 `json:"total-cpu-stat-steal"`
				TotalCPUStatGuest     int     `json:"total-cpu-stat-guest"`
				TotalCPUStatGuestNice int     `json:"total-cpu-stat-guest-nice"`
				MemoryProcessSize     struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"memory-process-size"`
				MemoryProcessRss struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"memory-process-rss"`
				MemoryProcessAnon struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"memory-process-anon"`
				MemoryProcessRssHwm struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"memory-process-rss-hwm"`
				MemoryProcessSwapSize struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"memory-process-swap-size"`
				MemorySystemPageinRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"memory-system-pagein-rate"`
				MemorySystemPageoutRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"memory-system-pageout-rate"`
				MemorySystemSwapinRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"memory-system-swapin-rate"`
				MemorySystemSwapoutRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"memory-system-swapout-rate"`
				MemorySize struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"memory-size"`
				HostSize struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"host-size"`
				HostLargeDataSize struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"host-large-data-size"`
				LogDeviceSpace struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"log-device-space"`
				DataDirSpace struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"data-dir-space"`
				BackupJobs struct {
					BackupJob []struct {
						JobID     string `json:"job-id"`
						Path      string `json:"path"`
						StartTime struct {
							Units string    `json:"units"`
							Value time.Time `json:"value"`
						} `json:"start-time"`
						Forests struct {
							Forest []struct {
								ForestID string `json:"forest-id"`
								Status   struct {
									Units string `json:"units"`
									Value string `json:"value"`
								} `json:"status"`
							} `json:"forest"`
						} `json:"forests"`
					} `json:"backup-job"`
				} `json:"backup-jobs"`
				QueryReadBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"query-read-bytes"`
				QueryReadTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"query-read-time"`
				QueryReadRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"query-read-rate"`
				QueryReadLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"query-read-load"`
				JournalWriteBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"journal-write-bytes"`
				JournalWriteTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"journal-write-time"`
				JournalWriteRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"journal-write-rate"`
				JournalWriteLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"journal-write-load"`
				SaveWriteBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"save-write-bytes"`
				SaveWriteTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"save-write-time"`
				SaveWriteRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"save-write-rate"`
				SaveWriteLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"save-write-load"`
				MergeReadBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"merge-read-bytes"`
				MergeReadTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"merge-read-time"`
				MergeReadRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"merge-read-rate"`
				MergeReadLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"merge-read-load"`
				MergeWriteBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"merge-write-bytes"`
				MergeWriteTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"merge-write-time"`
				MergeWriteRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"merge-write-rate"`
				MergeWriteLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"merge-write-load"`
				BackupReadBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"backup-read-bytes"`
				BackupReadTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"backup-read-time"`
				BackupReadRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"backup-read-rate"`
				BackupReadLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"backup-read-load"`
				BackupWriteBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"backup-write-bytes"`
				BackupWriteTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"backup-write-time"`
				BackupWriteRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"backup-write-rate"`
				BackupWriteLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"backup-write-load"`
				RestoreReadBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"restore-read-bytes"`
				RestoreReadTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"restore-read-time"`
				RestoreReadRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"restore-read-rate"`
				RestoreReadLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"restore-read-load"`
				RestoreWriteBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"restore-write-bytes"`
				RestoreWriteTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"restore-write-time"`
				RestoreWriteRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"restore-write-rate"`
				RestoreWriteLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"restore-write-load"`
				LargeReadBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"large-read-bytes"`
				LargeReadTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"large-read-time"`
				LargeReadRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"large-read-rate"`
				LargeReadLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"large-read-load"`
				LargeWriteBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"large-write-bytes"`
				LargeWriteTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"large-write-time"`
				LargeWriteRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"large-write-rate"`
				LargeWriteLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"large-write-load"`
				ExternalBinaryReadBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"external-binary-read-bytes"`
				ExternalBinaryReadTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"external-binary-read-time"`
				ExternalBinaryReadRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"external-binary-read-rate"`
				ExternalBinaryReadLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"external-binary-read-load"`
				WebDAVServerReceiveBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"webDAV-server-receive-bytes"`
				WebDAVServerReceiveTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"webDAV-server-receive-time"`
				WebDAVServerReceiveRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"webDAV-server-receive-rate"`
				WebDAVServerReceiveLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"webDAV-server-receive-load"`
				WebDAVServerSendBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"webDAV-server-send-bytes"`
				WebDAVServerSendTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"webDAV-server-send-time"`
				WebDAVServerSendRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"webDAV-server-send-rate"`
				WebDAVServerSendLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"webDAV-server-send-load"`
				HTTPServerReceiveBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"http-server-receive-bytes"`
				HTTPServerReceiveTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"http-server-receive-time"`
				HTTPServerReceiveRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"http-server-receive-rate"`
				HTTPServerReceiveLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"http-server-receive-load"`
				HTTPServerSendBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"http-server-send-bytes"`
				HTTPServerSendTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"http-server-send-time"`
				HTTPServerSendRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"http-server-send-rate"`
				HTTPServerSendLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"http-server-send-load"`
				XdbcServerReceiveBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"xdbc-server-receive-bytes"`
				XdbcServerReceiveTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"xdbc-server-receive-time"`
				XdbcServerReceiveRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"xdbc-server-receive-rate"`
				XdbcServerReceiveLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"xdbc-server-receive-load"`
				XdbcServerSendBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"xdbc-server-send-bytes"`
				XdbcServerSendTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"xdbc-server-send-time"`
				XdbcServerSendRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"xdbc-server-send-rate"`
				XdbcServerSendLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"xdbc-server-send-load"`
				OdbcServerReceiveBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"odbc-server-receive-bytes"`
				OdbcServerReceiveTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"odbc-server-receive-time"`
				OdbcServerReceiveRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"odbc-server-receive-rate"`
				OdbcServerReceiveLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"odbc-server-receive-load"`
				OdbcServerSendBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"odbc-server-send-bytes"`
				OdbcServerSendTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"odbc-server-send-time"`
				OdbcServerSendRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"odbc-server-send-rate"`
				OdbcServerSendLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"odbc-server-send-load"`
				XdqpClientReceiveBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"xdqp-client-receive-bytes"`
				XdqpClientReceiveTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"xdqp-client-receive-time"`
				XdqpClientReceiveRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"xdqp-client-receive-rate"`
				XdqpClientReceiveLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"xdqp-client-receive-load"`
				XdqpClientSendBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"xdqp-client-send-bytes"`
				XdqpClientSendTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"xdqp-client-send-time"`
				XdqpClientSendRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"xdqp-client-send-rate"`
				XdqpClientSendLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"xdqp-client-send-load"`
				XdqpServerReceiveBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"xdqp-server-receive-bytes"`
				XdqpServerReceiveTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"xdqp-server-receive-time"`
				XdqpServerReceiveRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"xdqp-server-receive-rate"`
				XdqpServerReceiveLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"xdqp-server-receive-load"`
				XdqpServerSendBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"xdqp-server-send-bytes"`
				XdqpServerSendTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"xdqp-server-send-time"`
				XdqpServerSendRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"xdqp-server-send-rate"`
				XdqpServerSendLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"xdqp-server-send-load"`
				ForeignXdqpClientReceiveBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"foreign-xdqp-client-receive-bytes"`
				ForeignXdqpClientReceiveTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"foreign-xdqp-client-receive-time"`
				ForeignXdqpClientReceiveRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"foreign-xdqp-client-receive-rate"`
				ForeignXdqpClientReceiveLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"foreign-xdqp-client-receive-load"`
				ForeignXdqpClientSendBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"foreign-xdqp-client-send-bytes"`
				ForeignXdqpClientSendTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"foreign-xdqp-client-send-time"`
				ForeignXdqpClientSendRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"foreign-xdqp-client-send-rate"`
				ForeignXdqpClientSendLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"foreign-xdqp-client-send-load"`
				ForeignXdqpServerReceiveBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"foreign-xdqp-server-receive-bytes"`
				ForeignXdqpServerReceiveTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"foreign-xdqp-server-receive-time"`
				ForeignXdqpServerReceiveRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"foreign-xdqp-server-receive-rate"`
				ForeignXdqpServerReceiveLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"foreign-xdqp-server-receive-load"`
				ForeignXdqpServerSendBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"foreign-xdqp-server-send-bytes"`
				ForeignXdqpServerSendTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"foreign-xdqp-server-send-time"`
				ForeignXdqpServerSendRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"foreign-xdqp-server-send-rate"`
				ForeignXdqpServerSendLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"foreign-xdqp-server-send-load"`
				ReadLockCount struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"read-lock-count"`
				ReadLockWaitTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"read-lock-wait-time"`
				ReadLockHoldTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"read-lock-hold-time"`
				ReadLockRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"read-lock-rate"`
				ReadLockWaitLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"read-lock-wait-load"`
				ReadLockHoldLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"read-lock-hold-load"`
				WriteLockCount struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"write-lock-count"`
				WriteLockWaitTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"write-lock-wait-time"`
				WriteLockHoldTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"write-lock-hold-time"`
				WriteLockRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"write-lock-rate"`
				WriteLockWaitLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"write-lock-wait-load"`
				WriteLockHoldLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"write-lock-hold-load"`
				DeadlockCount struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"deadlock-count"`
				DeadlockWaitTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"deadlock-wait-time"`
				DeadlockRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"deadlock-rate"`
				DeadlockWaitLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"deadlock-wait-load"`
			} `json:"status-detail"`
		} `json:"status-properties"`
		RelatedViews struct {
			RelatedView []struct {
				ViewType string `json:"view-type"`
				ViewName string `json:"view-name"`
				ViewURI  string `json:"view-uri"`
			} `json:"related-view"`
		} `json:"related-views"`
	} `json:"host-status"`
}

type groupList struct {
	GroupDefaultList struct {
		Meta struct {
			URI         string    `json:"uri"`
			CurrentTime time.Time `json:"current-time"`
			ElapsedTime struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"elapsed-time"`
		} `json:"meta"`
		ListItems struct {
			ListCount struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"list-count"`
			ListItem []struct {
				Uriref  string `json:"uriref"`
				Idref   string `json:"idref"`
				Nameref string `json:"nameref"`
			} `json:"list-item"`
		} `json:"list-items"`
		RelatedViews struct {
			RelatedView []struct {
				ViewType string `json:"view-type"`
				ViewName string `json:"view-name"`
				ViewURI  string `json:"view-uri"`
			} `json:"related-view"`
		} `json:"related-views"`
	} `json:"group-default-list"`
}

type groupDetail struct {
	GroupDefault struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Meta struct {
			URI         string    `json:"uri"`
			CurrentTime time.Time `json:"current-time"`
			ElapsedTime struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"elapsed-time"`
		} `json:"meta"`
		Relations struct {
			RelationGroup []struct {
				Typeref       string `json:"typeref"`
				RelationCount struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"relation-count"`
				Relation []struct {
					Uriref  string `json:"uriref"`
					Idref   string `json:"idref"`
					Nameref string `json:"nameref"`
				} `json:"relation"`
			} `json:"relation-group"`
		} `json:"relations"`
		RelatedViews struct {
			RelatedView []struct {
				ViewType string `json:"view-type"`
				ViewName string `json:"view-name"`
				ViewURI  string `json:"view-uri"`
			} `json:"related-view"`
		} `json:"related-views"`
	} `json:"group-default"`
}

type forestStatusSummary struct {
	ForestStatusList struct {
		Meta struct {
			URI         string    `json:"uri"`
			CurrentTime time.Time `json:"current-time"`
			ElapsedTime struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"elapsed-time"`
		} `json:"meta"`
		Relations struct {
			RelationGroup []struct {
				Typeref       string `json:"typeref"`
				RelationCount struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"relation-count"`
				Relation []struct {
					Uriref  string `json:"uriref"`
					Idref   string `json:"idref"`
					Nameref string `json:"nameref"`
				} `json:"relation"`
			} `json:"relation-group"`
		} `json:"relations"`
		StatusListSummary struct {
			TotalForests struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"total-forests"`
			StateNotOpen struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"state-not-open"`
			MaxStandsPerForest struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"max-stands-per-forest"`
			MergeCount struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"merge-count"`
			BackupCount struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"backup-count"`
			RestoreCount struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"restore-count"`
			MinCapacity struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"min-capacity"`
			LoadProperties struct {
				TotalLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"total-load"`
				LoadDetail struct {
					QueryReadLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"query-read-load"`
					JournalWriteLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"journal-write-load"`
					SaveWriteLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"save-write-load"`
					MergeReadLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"merge-read-load"`
					MergeWriteLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"merge-write-load"`
					BackupReadLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"backup-read-load"`
					BackupWriteLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"backup-write-load"`
					RestoreReadLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"restore-read-load"`
					RestoreWriteLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"restore-write-load"`
					LargeReadLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"large-read-load"`
					LargeWriteLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"large-write-load"`
					DatabaseReplicationSendLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"database-replication-send-load"`
					DatabaseReplicationReceiveLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"database-replication-receive-load"`
					ReadLockWaitLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"read-lock-wait-load"`
					ReadLockHoldLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"read-lock-hold-load"`
					WriteLockWaitLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"write-lock-wait-load"`
					WriteLockHoldLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"write-lock-hold-load"`
					DeadlockWaitLoad struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"deadlock-wait-load"`
				} `json:"load-detail"`
			} `json:"load-properties"`
			RateProperties struct {
				TotalRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"total-rate"`
				RateDetail struct {
					QueryReadRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"query-read-rate"`
					JournalWriteRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"journal-write-rate"`
					SaveWriteRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"save-write-rate"`
					MergeReadRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"merge-read-rate"`
					MergeWriteRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"merge-write-rate"`
					BackupReadRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"backup-read-rate"`
					BackupWriteRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"backup-write-rate"`
					RestoreReadRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"restore-read-rate"`
					RestoreWriteRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"restore-write-rate"`
					LargeReadRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"large-read-rate"`
					LargeWriteRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"large-write-rate"`
					DatabaseReplicationSendRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"database-replication-send-rate"`
					DatabaseReplicationReceiveRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"database-replication-receive-rate"`
					ReadLockRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"read-lock-rate"`
					WriteLockRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"write-lock-rate"`
					DeadlockRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"deadlock-rate"`
				} `json:"rate-detail"`
			} `json:"rate-properties"`
			CacheProperties struct {
				ListCacheHitRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"list-cache-hit-rate"`
				ListCacheMissRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"list-cache-miss-rate"`
				TripleCacheHitRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"triple-cache-hit-rate"`
				TripleCacheMissRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"triple-cache-miss-rate"`
				TripleValueCacheHitRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"triple-value-cache-hit-rate"`
				TripleValueCacheMissRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"triple-value-cache-miss-rate"`
				CompressedTreeCacheHitRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"compressed-tree-cache-hit-rate"`
				CompressedTreeCacheMissRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"compressed-tree-cache-miss-rate"`
				ListCacheRatio struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"list-cache-ratio"`
				TripleCacheRatio struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"triple-cache-ratio"`
				TripleValueCacheRatio struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"triple-value-cache-ratio"`
				CompressedTreeCacheRatio struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"compressed-tree-cache-ratio"`
				LargeBinaryCacheHitRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"large-binary-cache-hit-rate"`
				LargeBinaryCacheMissRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"large-binary-cache-miss-rate"`
			} `json:"cache-properties"`
		} `json:"status-list-summary"`
		StatusListItems struct {
			StatusListItem []struct {
				Uriref  string `json:"uriref"`
				Idref   string `json:"idref"`
				Nameref string `json:"nameref"`
			} `json:"status-list-item"`
		} `json:"status-list-items"`
		RelatedViews struct {
			RelatedView []struct {
				ViewType string `json:"view-type"`
				ViewName string `json:"view-name"`
				ViewURI  string `json:"view-uri"`
			} `json:"related-view"`
		} `json:"related-views"`
	} `json:"forest-status-list"`
}

type forestStatus struct {
	ForestStatus struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Meta struct {
			URI         string    `json:"uri"`
			CurrentTime time.Time `json:"current-time"`
			ElapsedTime struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"elapsed-time"`
		} `json:"meta"`
		Relations struct {
			RelationGroup []struct {
				Typeref  string `json:"typeref"`
				Relation []struct {
					Uriref  string `json:"uriref"`
					Idref   string `json:"idref"`
					Nameref string `json:"nameref"`
				} `json:"relation"`
			} `json:"relation-group"`
		} `json:"relations"`
		StatusProperties struct {
			State struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"state"`
			Enabled struct {
				Units string `json:"units"`
				Value bool   `json:"value"`
			} `json:"enabled"`
			Availability struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"availability"`
			Encryption struct {
				Units string `json:"units"`
				Value bool   `json:"value"`
			} `json:"encryption"`
			UpdatesAllowed struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"updates-allowed"`
			RebalancerEnable struct {
				Units string `json:"units"`
				Value bool   `json:"value"`
			} `json:"rebalancer-enable"`
			MasterForest             string `json:"master-forest"`
			CurrentMasterForest      string `json:"current-master-forest"`
			CurrentMasterPreciseTime struct {
				Units string    `json:"units"`
				Value time.Time `json:"value"`
			} `json:"current-master-precise-time"`
			CurrentMasterFsn struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"current-master-fsn"`
			CurrentForeignMasterCluster struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"current-foreign-master-cluster"`
			CurrentForeignMasterDatabase    string `json:"current-foreign-master-database"`
			CurrentForeignMasterForest      string `json:"current-foreign-master-forest"`
			CurrentForeignMasterPreciseTime struct {
				Units string    `json:"units"`
				Value time.Time `json:"value"`
			} `json:"current-foreign-master-precise-time"`
			CurrentForeignMasterFsn struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"current-foreign-master-fsn"`
			LastStateChange struct {
				Units string    `json:"units"`
				Value time.Time `json:"value"`
			} `json:"last-state-change"`
			NonblockingTimestamp struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"nonblocking-timestamp"`
			MaxQueryTimestamp struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"max-query-timestamp"`
			DataDir      string `json:"data-dir"`
			JournalsSize struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"journals-size"`
			LargeDataSize struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"large-data-size"`
			OrphanedBinaries struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"orphaned-binaries"`
			Stand []struct {
				StandID   []string `json:"stand-id"`
				Path      string   `json:"path"`
				StandKind struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"stand-kind"`
				IsFast struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"is-fast"`
				LabelVersion struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"label-version"`
				DiskSize struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"disk-size"`
				EncryptedDiskSize struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"encrypted-disk-size"`
				MemorySize struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"memory-size"`
				ListCacheHits struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"list-cache-hits"`
				ListCacheMisses struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"list-cache-misses"`
				ListCacheHitRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"list-cache-hit-rate"`
				ListCacheMissRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"list-cache-miss-rate"`
				CompressedTreeCacheHits struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"compressed-tree-cache-hits"`
				CompressedTreeCacheMisses struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"compressed-tree-cache-misses"`
				CompressedTreeCacheHitRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"compressed-tree-cache-hit-rate"`
				CompressedTreeCacheMissRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"compressed-tree-cache-miss-rate"`
				TripleCacheHits struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"triple-cache-hits"`
				TripleCacheMisses struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"triple-cache-misses"`
				TripleCacheHitRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"triple-cache-hit-rate"`
				TripleCacheMissRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"triple-cache-miss-rate"`
				TripleValueCacheHits struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"triple-value-cache-hits"`
				TripleValueCacheMisses struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"triple-value-cache-misses"`
				TripleValueCacheHitRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"triple-value-cache-hit-rate"`
				TripleValueCacheMissRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"triple-value-cache-miss-rate"`
			} `json:"stand"`
			ForestReserve struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"forest-reserve"`
			Rebalancing struct {
				Units string `json:"units"`
				Value bool   `json:"value"`
			} `json:"rebalancing"`
			Reindexing struct {
				Units string `json:"units"`
				Value bool   `json:"value"`
			} `json:"reindexing"`
			DeviceSpace struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"device-space"`
			LastBackup struct {
				Units string    `json:"units"`
				Value time.Time `json:"value"`
			} `json:"last-backup"`
			LastIncrBackup struct {
				Units string    `json:"units"`
				Value time.Time `json:"value"`
			} `json:"last-incr-backup"`
			TransactionJournalSize struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"transaction-journal-size"`
			TransactionJournalLimit struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"transaction-journal-limit"`
			QueryReadBytes struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"query-read-bytes"`
			QueryReadTime struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"query-read-time"`
			QueryReadRate struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"query-read-rate"`
			QueryReadLoad struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"query-read-load"`
			JournalWriteBytes struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"journal-write-bytes"`
			JournalWriteTime struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"journal-write-time"`
			JournalWriteRate struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"journal-write-rate"`
			JournalWriteLoad struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"journal-write-load"`
			SaveWriteBytes struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"save-write-bytes"`
			SaveWriteTime struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"save-write-time"`
			SaveWriteRate struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"save-write-rate"`
			SaveWriteLoad struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"save-write-load"`
			MergeReadBytes struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"merge-read-bytes"`
			MergeReadTime struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"merge-read-time"`
			MergeReadRate struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"merge-read-rate"`
			MergeReadLoad struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"merge-read-load"`
			MergeWriteBytes struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"merge-write-bytes"`
			MergeWriteTime struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"merge-write-time"`
			MergeWriteRate struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"merge-write-rate"`
			MergeWriteLoad struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"merge-write-load"`
			BackupReadBytes struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"backup-read-bytes"`
			BackupReadTime struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"backup-read-time"`
			BackupReadRate struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"backup-read-rate"`
			BackupReadLoad struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"backup-read-load"`
			BackupWriteBytes struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"backup-write-bytes"`
			BackupWriteTime struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"backup-write-time"`
			BackupWriteRate struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"backup-write-rate"`
			BackupWriteLoad struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"backup-write-load"`
			RestoreReadBytes struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"restore-read-bytes"`
			RestoreReadTime struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"restore-read-time"`
			RestoreReadRate struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"restore-read-rate"`
			RestoreReadLoad struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"restore-read-load"`
			RestoreWriteBytes struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"restore-write-bytes"`
			RestoreWriteTime struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"restore-write-time"`
			RestoreWriteRate struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"restore-write-rate"`
			RestoreWriteLoad struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"restore-write-load"`
			LargeReadBytes struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"large-read-bytes"`
			LargeReadTime struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"large-read-time"`
			LargeReadRate struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"large-read-rate"`
			LargeReadLoad struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"large-read-load"`
			LargeWriteBytes struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"large-write-bytes"`
			LargeWriteTime struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"large-write-time"`
			LargeWriteRate struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"large-write-rate"`
			LargeWriteLoad struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"large-write-load"`
			DatabaseReplicationReceiveBytes struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"database-replication-receive-bytes"`
			DatabaseReplicationReceiveTime struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"database-replication-receive-time"`
			DatabaseReplicationReceiveRate struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"database-replication-receive-rate"`
			DatabaseReplicationReceiveLoad struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"database-replication-receive-load"`
			DatabaseReplicationSendBytes struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"database-replication-send-bytes"`
			DatabaseReplicationSendTime struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"database-replication-send-time"`
			DatabaseReplicationSendRate struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"database-replication-send-rate"`
			DatabaseReplicationSendLoad struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"database-replication-send-load"`
			ReadLockCount struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"read-lock-count"`
			ReadLockWaitTime struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"read-lock-wait-time"`
			ReadLockHoldTime struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"read-lock-hold-time"`
			ReadLockRate struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"read-lock-rate"`
			ReadLockWaitLoad struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"read-lock-wait-load"`
			ReadLockHoldLoad struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"read-lock-hold-load"`
			WriteLockCount struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"write-lock-count"`
			WriteLockWaitTime struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"write-lock-wait-time"`
			WriteLockHoldTime struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"write-lock-hold-time"`
			WriteLockRate struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"write-lock-rate"`
			WriteLockWaitLoad struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"write-lock-wait-load"`
			WriteLockHoldLoad struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"write-lock-hold-load"`
			DeadlockCount struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"deadlock-count"`
			DeadlockWaitTime struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"deadlock-wait-time"`
			DeadlockRate struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"deadlock-rate"`
			DeadlockWaitLoad struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"deadlock-wait-load"`
			LargeBinaryCacheHits struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"large-binary-cache-hits"`
			LargeBinaryCacheMisses struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"large-binary-cache-misses"`
			LargeBinaryCacheHitRate struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"large-binary-cache-hit-rate"`
			LargeBinaryCacheMissRate struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"large-binary-cache-miss-rate"`
		} `json:"status-properties"`
		RelatedViews struct {
			RelatedView []struct {
				ViewType string `json:"view-type"`
				ViewName string `json:"view-name"`
				ViewURI  string `json:"view-uri"`
			} `json:"related-view"`
		} `json:"related-views"`
	} `json:"forest-status"`
}

type serverStatus struct {
	ServerStatus struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		ServerKind string `json:"server-kind"`
		Meta       struct {
			URI         string    `json:"uri"`
			CurrentTime time.Time `json:"current-time"`
			ElapsedTime struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"elapsed-time"`
		} `json:"meta"`
		Relations struct {
			RelationGroup []struct {
				Typeref       string `json:"typeref"`
				RelationCount struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"relation-count,omitempty"`
				Relation []struct {
					Uriref  string `json:"uriref"`
					Idref   string `json:"idref"`
					Nameref string `json:"nameref"`
				} `json:"relation,omitempty"`
				Uriref string `json:"uriref,omitempty"`
			} `json:"relation-group"`
		} `json:"relations"`
		StatusProperties struct {
			Enabled struct {
				Units string `json:"units"`
				Value bool   `json:"value"`
			} `json:"enabled"`
			Port             int    `json:"port"`
			Root             string `json:"root"`
			DisplayLastLogin struct {
				Units string `json:"units"`
				Value bool   `json:"value"`
			} `json:"display-last-login"`
			Backlog struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"backlog"`
			Threads struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"threads"`
			MaxThreads struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"max-threads"`
			RequestTimeout struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"request-timeout"`
			KeepAliveTimeout struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"keep-alive-timeout"`
			SessionTimeout struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"session-timeout"`
			StaticExpires struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"static-expires"`
			MaxTimeLimit struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"max-time-limit"`
			DefaultTimeLimit struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"default-time-limit"`
			MultiVersionConcurrencyControl struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"multi-version-concurrency-control"`
			Authentication struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"authentication"`
			DefaultUser            string `json:"default-user"`
			Privilege              int    `json:"privilege"`
			ConcurrentRequestLimit struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"concurrent-request-limit"`
			DefaultXqueryVersion struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"default-xquery-version"`
			OutputSgmlCharacterEntities struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"output-sgml-character-entities"`
			OutputEncoding   string `json:"output-encoding"`
			TotalRequestRate struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"total-request-rate"`
			TotalExpandedTreeCacheMissRate struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"total-expanded-tree-cache-miss-rate"`
			TotalExpandedTreeCacheHitRate struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"total-expanded-tree-cache-hit-rate"`
			TotalRequests struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"total-requests"`
			HostDetail []struct {
				RelationID    string `json:"relation-id"`
				RequestsCount struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"requests-count"`
				MaxInferenceSize struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"max-inference-size"`
				DefaultInferenceSize struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"default-inference-size"`
				DistributeTimestamps struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"distribute-timestamps"`
				RequestRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"request-rate"`
				ExpandedTreeCacheHits struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"expanded-tree-cache-hits"`
				ExpandedTreeCacheMisses struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"expanded-tree-cache-misses"`
				ExpandedTreeCacheHitRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"expanded-tree-cache-hit-rate"`
				ExpandedTreeCacheMissRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"expanded-tree-cache-miss-rate"`
				FsProgramCacheHits struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"fs-program-cache-hits"`
				FsProgramCacheMisses struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"fs-program-cache-misses"`
				DbProgramCacheHits struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"db-program-cache-hits"`
				DbProgramCacheMisses struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"db-program-cache-misses"`
				EnvProgramCacheHits struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"env-program-cache-hits"`
				EnvProgramCacheMisses struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"env-program-cache-misses"`
				FsMainModuleSeqCacheHits struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"fs-main-module-seq-cache-hits"`
				FsMainModuleSeqCacheMisses struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"fs-main-module-seq-cache-misses"`
				DbMainModuleSeqCacheHits struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"db-main-module-seq-cache-hits"`
				DbMainModuleSeqCacheMisses struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"db-main-module-seq-cache-misses"`
				FsLibModuleCacheHits struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"fs-lib-module-cache-hits"`
				FsLibModuleCacheMisses struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"fs-lib-module-cache-misses"`
				DbLibModuleCacheHits struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"db-lib-module-cache-hits"`
				DbLibModuleCacheMisses struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"db-lib-module-cache-misses"`
				RequestTime struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"request-time"`
				ServerReceiveBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"server-receive-bytes"`
				ServerReceiveTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"server-receive-time"`
				ServerReceiveRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"server-receive-rate"`
				ServerReceiveLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"server-receive-load"`
				ServerSendBytes struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"server-send-bytes"`
				ServerSendTime struct {
					Units string `json:"units"`
					Value string `json:"value"`
				} `json:"server-send-time"`
				ServerSendRate struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"server-send-rate"`
				ServerSendLoad struct {
					Units string  `json:"units"`
					Value float64 `json:"value"`
				} `json:"server-send-load"`
			} `json:"host-detail"`
		} `json:"status-properties"`
		RelatedViews struct {
			RelatedView []struct {
				ViewType string `json:"view-type"`
				ViewName string `json:"view-name"`
				ViewURI  string `json:"view-uri"`
			} `json:"related-view"`
		} `json:"related-views"`
	} `json:"server-status"`
}
