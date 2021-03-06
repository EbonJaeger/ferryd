//
// Copyright © 2017-2019 Solus Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package jobs

import (
	"ferryd/core"
	"fmt"
	log "github.com/DataDrake/waterlog"
)

// TrimObsoleteJobHandler is responsible for indexing repositories and should only
// ever be used in sequential queues.
type TrimObsoleteJobHandler Job

// NewTrimObsoleteJob will return a job suitable for adding to the job processor
func NewTrimObsoleteJob(id string) *Job {
	return &Job{
		Type:    TrimObsolete,
		SrcRepo: id,
	}
}

// NewTrimObsoleteJobHandler will create a job handler for the input job and ensure it validates
func NewTrimObsoleteJobHandler(j *Job) (handler *TrimObsoleteJobHandler, err error) {
	if len(j.SrcRepo) == 0 {
		err = fmt.Errorf("job is missing a source repository")
		return
	}
	h := TrimObsoleteJobHandler(*j)
	handler = &h
	return
}

// Execute will try to remove any excessive packages marked as Obsolete
func (j *TrimObsoleteJobHandler) Execute(_ *JobStore, manager *core.Manager) error {
	if err := manager.TrimObsolete(j.SrcRepo); err != nil {
		return err
	}
	log.Infof("Trimmed obsoletes in repository '%s'\n", j.SrcRepo)
	return nil
}

// Describe returns a human readable description for this job
func (j *TrimObsoleteJobHandler) Describe() string {
	return fmt.Sprintf("Trim obsoletes from repository '%s'", j.SrcRepo)
}
