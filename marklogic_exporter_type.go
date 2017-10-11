package main

import (
	"time"
)

type Status struct {
	LocalClusterStatus struct {
		ID               string `json:"id"`
		Name             string `json:"name"`
		Version          string `json:"version"`
		EffectiveVersion int    `json:"effective-version"`
		Role             string `json:"role"`
		Meta             struct {
			URI         string    `json:"uri"`
			CurrentTime time.Time `json:"current-time"`
			ElapsedTime struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"elapsed-time"`
		} `json:"meta"`
		StatusRelations struct {
			HostsStatus struct {
				Uriref             string `json:"uriref"`
				Typeref            string `json:"typeref"`
				HostsStatusSummary struct {
					TotalHosts struct {
						Units string `json:"units"`
						Value int    `json:"value"`
					} `json:"total-hosts"`
					LoadProperties struct {
						TotalLoad struct {
							Units string  `json:"units"`
							Value float64 `json:"value"`
						} `json:"total-load"`
						LoadDetail struct {
							QueryReadLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"query-read-load"`
							JournalWriteLoad struct {
								Units string  `json:"units"`
								Value float64 `json:"value"`
							} `json:"journal-write-load"`
							SaveWriteLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"save-write-load"`
							MergeReadLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"merge-read-load"`
							MergeWriteLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"merge-write-load"`
							BackupReadLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"backup-read-load"`
							BackupWriteLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"backup-write-load"`
							RestoreReadLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"restore-read-load"`
							RestoreWriteLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"restore-write-load"`
							LargeReadLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"large-read-load"`
							LargeWriteLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"large-write-load"`
							ExternalBinaryReadLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"external-binary-read-load"`
							XdqpClientReceiveLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"xdqp-client-receive-load"`
							XdqpClientSendLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"xdqp-client-send-load"`
							XdqpServerReceiveLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"xdqp-server-receive-load"`
							XdqpServerSendLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"xdqp-server-send-load"`
							ForeignXdqpClientReceiveLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"foreign-xdqp-client-receive-load"`
							ForeignXdqpClientSendLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"foreign-xdqp-client-send-load"`
							ForeignXdqpServerReceiveLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"foreign-xdqp-server-receive-load"`
							ForeignXdqpServerSendLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"foreign-xdqp-server-send-load"`
							ReadLockWaitLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"read-lock-wait-load"`
							ReadLockHoldLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"read-lock-hold-load"`
							WriteLockWaitLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"write-lock-wait-load"`
							WriteLockHoldLoad struct {
								Units string  `json:"units"`
								Value float64 `json:"value"`
							} `json:"write-lock-hold-load"`
							DeadlockWaitLoad struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"deadlock-wait-load"`
						} `json:"load-detail"`
					} `json:"load-properties"`
					RateProperties struct {
						TotalRate struct {
							Units string  `json:"units"`
							Value float64 `json:"value"`
						} `json:"total-rate"`
						RateDetail struct {
							MemoryProcessSwapRate struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"memory-process-swap-rate"`
							MemorySystemPageinRate struct {
								Units string  `json:"units"`
								Value float64 `json:"value"`
							} `json:"memory-system-pagein-rate"`
							MemorySystemPageoutRate struct {
								Units string  `json:"units"`
								Value float64 `json:"value"`
							} `json:"memory-system-pageout-rate"`
							MemorySystemSwapinRate struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"memory-system-swapin-rate"`
							MemorySystemSwapoutRate struct {
								Units string `json:"units"`
								Value int    `json:"value"`
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
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"save-write-rate"`
							MergeReadRate struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"merge-read-rate"`
							MergeWriteRate struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"merge-write-rate"`
							BackupReadRate struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"backup-read-rate"`
							BackupWriteRate struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"backup-write-rate"`
							RestoreReadRate struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"restore-read-rate"`
							RestoreWriteRate struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"restore-write-rate"`
							LargeReadRate struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"large-read-rate"`
							LargeWriteRate struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"large-write-rate"`
							ExternalBinaryReadRate struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"external-binary-read-rate"`
							XdqpClientReceiveRate struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"xdqp-client-receive-rate"`
							XdqpClientSendRate struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"xdqp-client-send-rate"`
							XdqpServerReceiveRate struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"xdqp-server-receive-rate"`
							XdqpServerSendRate struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"xdqp-server-send-rate"`
							ForeignXdqpClientReceiveRate struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"foreign-xdqp-client-receive-rate"`
							ForeignXdqpClientSendRate struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"foreign-xdqp-client-send-rate"`
							ForeignXdqpServerReceiveRate struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"foreign-xdqp-server-receive-rate"`
							ForeignXdqpServerSendRate struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"foreign-xdqp-server-send-rate"`
							ReadLockRate struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"read-lock-rate"`
							WriteLockRate struct {
								Units string  `json:"units"`
								Value float64 `json:"value"`
							} `json:"write-lock-rate"`
							DeadlockRate struct {
								Units string `json:"units"`
								Value int    `json:"value"`
							} `json:"deadlock-rate"`
						} `json:"rate-detail"`
					} `json:"rate-properties"`
				} `json:"hosts-status-summary"`
			} `json:"hosts-status"`
			ServersStatus struct {
				Uriref               string `json:"uriref"`
				Typeref              string `json:"typeref"`
				ServersStatusSummary struct {
					RequestRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"request-rate"`
					ExpandedTreeCacheMissRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"expanded-tree-cache-miss-rate"`
					ExpandedTreeCacheHitRate struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"expanded-tree-cache-hit-rate"`
					RequestCount struct {
						Units string `json:"units"`
						Value int    `json:"value"`
					} `json:"request-count"`
				} `json:"servers-status-summary"`
			} `json:"servers-status"`
			ForestsStatus struct {
				Uriref               string `json:"uriref"`
				Typeref              string `json:"typeref"`
				ForestsStatusSummary struct {
					TotalForests struct {
						Units string `json:"units"`
						Value int    `json:"value"`
					} `json:"total-forests"`
					StateNotOpen struct {
						Units string `json:"units"`
						Value int    `json:"value"`
					} `json:"state-not-open"`
					MaxStandsPerForest struct {
						Units string `json:"units"`
						Value int    `json:"value"`
					} `json:"max-stands-per-forest"`
					MergeCount struct {
						Units string `json:"units"`
						Value int    `json:"value"`
					} `json:"merge-count"`
					BackupCount struct {
						Units string `json:"units"`
						Value int    `json:"value"`
					} `json:"backup-count"`
					RestoreCount struct {
						Units string `json:"units"`
						Value int    `json:"value"`
					} `json:"restore-count"`
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
							Units string `json:"units"`
							Value int    `json:"value"`
						} `json:"triple-cache-hit-rate"`
						TripleCacheMissRate struct {
							Units string `json:"units"`
							Value int    `json:"value"`
						} `json:"triple-cache-miss-rate"`
						TripleValueCacheHitRate struct {
							Units string `json:"units"`
							Value int    `json:"value"`
						} `json:"triple-value-cache-hit-rate"`
						TripleValueCacheMissRate struct {
							Units string `json:"units"`
							Value int    `json:"value"`
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
						CompressedTreeCacheRatio struct {
							Units string `json:"units"`
							Value string `json:"value"`
						} `json:"compressed-tree-cache-ratio"`
						LargeBinaryCacheHitRate struct {
							Units string `json:"units"`
							Value int    `json:"value"`
						} `json:"large-binary-cache-hit-rate"`
						LargeBinaryCacheMissRate struct {
							Units string `json:"units"`
							Value int    `json:"value"`
						} `json:"large-binary-cache-miss-rate"`
					} `json:"cache-properties"`
				} `json:"forests-status-summary"`
			} `json:"forests-status"`
			RequestsStatus struct {
				Uriref                string `json:"uriref"`
				Typeref               string `json:"typeref"`
				RequestsStatusSummary struct {
					MaxSeconds struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"max-seconds"`
					NinetiethPercentileSeconds struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"ninetieth-percentile-seconds"`
					MedianSeconds struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"median-seconds"`
					MeanSeconds struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"mean-seconds"`
					StandardDevSeconds struct {
						Units string `json:"units"`
						Value int    `json:"value"`
					} `json:"standard-dev-seconds"`
					MinSeconds struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"min-seconds"`
					TotalRequests struct {
						Units string `json:"units"`
						Value int    `json:"value"`
					} `json:"total-requests"`
					UpdateCount struct {
						Units string `json:"units"`
						Value int    `json:"value"`
					} `json:"update-count"`
					QueryCount struct {
						Units string `json:"units"`
						Value int    `json:"value"`
					} `json:"query-count"`
				} `json:"requests-status-summary"`
			} `json:"requests-status"`
			TransactionsStatus struct {
				Uriref                    string `json:"uriref"`
				Typeref                   string `json:"typeref"`
				TransactionsStatusSummary struct {
					MaxSeconds struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"max-seconds"`
					NinetiethPercentileSeconds struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"ninetieth-percentile-seconds"`
					MedianSeconds struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"median-seconds"`
					MeanSeconds struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"mean-seconds"`
					StandardDevSeconds struct {
						Units string `json:"units"`
						Value int    `json:"value"`
					} `json:"standard-dev-seconds"`
					MinSeconds struct {
						Units string  `json:"units"`
						Value float64 `json:"value"`
					} `json:"min-seconds"`
					TotalTransactions struct {
						Units string `json:"units"`
						Value int    `json:"value"`
					} `json:"total-transactions"`
				} `json:"transactions-status-summary"`
			} `json:"transactions-status"`
		} `json:"status-relations"`
		RelatedViews struct {
			RelatedView []struct {
				ViewType string `json:"view-type"`
				ViewName string `json:"view-name"`
				ViewURI  string `json:"view-uri"`
			} `json:"related-view"`
		} `json:"related-views"`
	} `json:"local-cluster-status"`
}