/*Package exec is used for executing commands

Copyright © 2020 Addshore

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package exec

import (
	"os/exec"
)

/*Command passes through to exec.Command for running generic commands*/
func Command(name string, arg ...string) *exec.Cmd {
	return exec.Command(name, arg...)
}

/*DockerCompose executes the docker-compose command*/
func DockerCompose(command string, arg ...string) *exec.Cmd {
	arg = append([]string{command}, arg...)
	return exec.Command("docker-compose", arg...)
}