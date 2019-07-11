// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
)

{% for enum in enums %}

{{ enum.description }}
type {{ enum.name }} string

const (
{% for enum_member in enum.members %}
    {{ enum_member.description }}
    {{ enum_member.name }}{{ enum.name }} {{ enum.name }} = "{{ enum_member.name }}"
{%- endfor %}
)
{% endfor %}
{% for class in classes -%}

{{ class.description }}
type {{ class.name }} struct {
    common.Entity
{% for attr in class.attrs %}
    {{ attr.description }}
    {{ attr.name }}  {{ attr.type }}
{%- endfor %}
}

// UnmarshalJSON unmarshals a {{ class.name }} object from the raw JSON.
func ({{ class.name|lower }} *{{ class.name }}) UnmarshalJSON(b []byte) error {
    type temp {{ class.name }}
    var t struct {
        temp
    }

    err := json.Unmarshal(b, &t)
    if err != nil {
        return err
    }

    *{{ class.name|lower }} = {{ class.name }}(t.temp)

    // Extract the links to other entities for later

    return nil
}

{% if class.name == object_name %}
// Get{{ class.name }} will get a {{ class.name }} instance from the service.
func Get{{ class.name }}(c common.Client, uri string) (*{{ class.name }}, error) {
    resp, err := c.Get(uri)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var {{ class.name|lower }} {{ class.name }}
    err = json.NewDecoder(resp.Body).Decode(&{{ class.name|lower }})
    if err != nil {
        return nil, err
    }

    {{ class.name|lower }}.SetClient(c)
    return &{{ class.name|lower }}, nil
}

// ListReferenced{{ class.name }}s gets the collection of {{ class.name }} from
// a provided reference.
func ListReferenced{{ class.name }}s(c common.Client, link string) ([]*{{ class.name }}, error) {
    var result []*{{ class.name }}
    if link == "" {
        return result, nil
    }

    links, err := common.GetCollection(c, link)
    if err != nil {
        return result, err
    }

    for _, {{ class.name|lower }}Link := range links.ItemLinks {
        {{ class.name|lower }}, err := Get{{ class.name }}(c, {{ class.name|lower }}Link)
        if err != nil {
            return result, err
        }
        result = append(result, {{ class.name|lower }})
    }

    return result, nil
}

{% endif %}
{% endfor %}
