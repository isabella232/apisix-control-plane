/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package yml

import "github.com/apache/apisix-control-plane/pkg/mem"

func (d *Destination) ToMem() []mem.MemModel {
	result := make([]mem.MemModel, 0)
	for _, s := range d.Subsets {
		fullName := *d.Kind + seprator + *d.Name + seprator + *s.Name
		group := *d.Kind + seprator + *d.Name
		upstream := &mem.Upstream{
			Kind:     d.Kind,
			Name:     d.Name,
			FullName: &fullName,
			Host:     d.Host,
			Group:    &group,
			Weight:   s.Weight,
		}
		result = append(result, upstream)
	}
	return result
}