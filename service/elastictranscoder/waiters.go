// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package elastictranscoder

import (
	"github.com/aws/aws-sdk-go/private/waiter"
)

func (c *ElasticTranscoder) WaitUntilJobComplete(input *ReadJobInput) error {
	waiterCfg := waiter.Config{
		Operation:   "ReadJob",
		Delay:       30,
		MaxAttempts: 120,
		Acceptors: []waiter.WaitAcceptor{
			waiter.WaitAcceptor{
				State:    "success",
				Matcher:  "path",
				Argument: "Job.Status",
				Expected: "Complete",
			},
			waiter.WaitAcceptor{
				State:    "failure",
				Matcher:  "path",
				Argument: "Job.Status",
				Expected: "Canceled",
			},
			waiter.WaitAcceptor{
				State:    "failure",
				Matcher:  "path",
				Argument: "Job.Status",
				Expected: "Error",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}
