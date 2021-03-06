/*
   Copyright 2020 Docker Compose CLI authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package formatter

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/pkg/errors"

	"github.com/docker/compose-cli/errdefs"
)

// Print prints formatted lists in different formats
func Print(toJSON interface{}, format string, outWriter io.Writer, writerFn func(w io.Writer), headers ...string) error {
	switch strings.ToLower(format) {
	case PRETTY, "":
		return PrintPrettySection(outWriter, writerFn, headers...)
	case JSON:
		switch reflect.TypeOf(toJSON).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(toJSON)
			for i := 0; i < s.Len(); i++ {
				obj := s.Index(i).Interface()
				jsonLine, err := ToJSON(obj, "", "")
				if err != nil {
					return err
				}
				_, _ = fmt.Fprint(outWriter, jsonLine)
			}
		default:
			outJSON, err := ToStandardJSON(toJSON)
			if err != nil {
				return err
			}
			_, _ = fmt.Fprintln(outWriter, outJSON)
		}
	default:
		return errors.Wrapf(errdefs.ErrParsingFailed, "format value %q could not be parsed", format)
	}
	return nil
}
