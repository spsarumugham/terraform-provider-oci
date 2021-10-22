// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v49/common"
)

// ExadataDatabaseStatisticsSummary Database details and statistics.
type ExadataDatabaseStatisticsSummary struct {
	ResourceDetails *DatabaseDetails `mandatory:"true" json:"resourceDetails"`

	CurrentStatistics *ExadataInsightResourceStatistics `mandatory:"true" json:"currentStatistics"`
}

func (m ExadataDatabaseStatisticsSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ExadataDatabaseStatisticsSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExadataDatabaseStatisticsSummary ExadataDatabaseStatisticsSummary
	s := struct {
		DiscriminatorParam string `json:"exadataResourceType"`
		MarshalTypeExadataDatabaseStatisticsSummary
	}{
		"DATABASE",
		(MarshalTypeExadataDatabaseStatisticsSummary)(m),
	}

	return json.Marshal(&s)
}
